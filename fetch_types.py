#!/usr/bin/env python3
"""Generating golang telegram bot api from the API docs.

API definition taken from here: https://core.telegram.org/bots/api
"""

from dataclasses import dataclass
from enum import Enum
from html import parser
from urllib import request

# Whole sections that should be skipped because they contain no actionable
# types or definitions or are not supported by the generator at the moment.
SKIP_SECTIONS: set[str] = set(
    [
        "Recent changes",
        "Authorizing your bot",
        "Do I need a Local Bot API Server",
        "Making requests",
        "Using a Local Bot API Server",
        "Accent colors",
        "Profile accent colors",
        "Determining list of commands",
        "Sending files",
        "Inline mode objects",
        "Inline mode methods",
        "Formatting options",
    ]
)

SKIP_METHODS: set[str] = set(["sendMediaGroup"])

# Mapping from some primitive type name to its go equivalent
PRIMITIVE_TYPES: dict[str, str] = {
    "byte": "byte",
    "Integer": "int64",
    "Float": "float32",
    "Float number": "float32",
    "String": "string",
    "Boolean": "bool",
    "True": "bool",
    "False": "bool",
    "Integer or String": "string",
}

# An aggregate type that can be one of the following
ONEOF_TYPES: dict[str, str] = {
    "MessageOrigin": [
        "MessageOriginUser",
        "MessageOriginHiddenUser",
        "MessageOriginChat",
        "MessageOriginChannel",
    ],
    "ReactionType": [
        "ReactionTypeEmoji",
        "ReactionTypeCustomEmoji",
    ],
}


@dataclass
class Param:
    """A parameter: struct field or method argument."""

    name: str
    typeName: str
    required: bool = False
    description: str = ""


@dataclass
class Token:
    """A token: either struct or method definition."""

    name: str
    description: str
    params: list[Param]


# Utility code
def formatWord(word: str) -> str:
    if word == "":
        return ""
    if word.lower() in ["id", "url", "ip"]:
        return word.upper()
    return word[0].upper() + word[1:]


def toCamelCase(usstr: str, capFirst: str = True) -> str:
    result = "".join(map(formatWord, usstr.split("_")))
    if not capFirst and result.upper() != result:
        return result[0].lower() + result[1:]
    return result


def formatComment(comment: str, offset: int, maxWidth: int = 95) -> str:
    prefix = " " * offset + "// "
    currentLines = list(reversed(comment.split("\n")))
    commentLines = []

    while currentLines:
        line = currentLines.pop()
        if len(line) <= maxWidth:
            commentLines.append(line)
            continue
        splitAt = maxWidth
        while splitAt >= 0 and not line[splitAt].isspace():
            splitAt -= 1
        commentLines.append(line[:splitAt].strip())
        currentLines.append(line[splitAt:].strip())
    return prefix + ("\n" + prefix).join(commentLines)


def formatType(tgType: str) -> str:
    if " or " in tgType:
        return "interface{}"
    maybeArray = tgType.split("Array of ")
    tgType = PRIMITIVE_TYPES.get(maybeArray[-1], toCamelCase(maybeArray[-1]))
    if tgType not in PRIMITIVE_TYPES.values():
        tgType = "*" + tgType
    return "[]" * (len(maybeArray) - 1) + tgType


class State(Enum):
    DOCUMENT = 1
    SECTION = 2
    SUBSECTION = 3
    TABLE = 4
    TABLE_ROW = 5
    TABLE_CELL = 6


class Parser(parser.HTMLParser):
    def __init__(self):
        super().__init__()
        self.tokens = []
        self.state = State.DOCUMENT
        self.section = ""
        self.subsection = ""

        self.currentToken = None
        self.currentParam = None
        self.description = []
        self.table = []
        self.tableRow = []
        self.tableCell = ""

    def handle_starttag(self, tag, attrs):
        self.state = {
            "h3": State.SECTION,
            "h4": State.SUBSECTION,
            "tr": State.TABLE_ROW,
            "td": State.TABLE_CELL,
            "table": State.TABLE,
        }.get(tag, self.state)
        if (
            self.state == State.SUBSECTION
            and self.subsection != ""
            and self.section not in SKIP_SECTIONS
            and self.subsection not in SKIP_SECTIONS
        ):
            self.tokens.append(
                Token(self.subsection, " ".join(self.description), self.table)
            )
            self.subsection = ""
            self.description = []
            self.table = []

    def handle_endtag(self, tag):
        oldstate = self.state
        self.state = (
            {
                State.SECTION: {"h3": State.DOCUMENT},
                State.SUBSECTION: {"h4": State.DOCUMENT},
                State.TABLE_ROW: {"tr": State.TABLE},
                State.TABLE_CELL: {"td": State.TABLE_ROW},
                State.TABLE: {"table": State.DOCUMENT},
            }
            .get(self.state, {})
            .get(tag, self.state)
        )

        if self.section in SKIP_SECTIONS or self.subsection in SKIP_SECTIONS:
            return

        if oldstate == State.TABLE_CELL and self.state == State.TABLE_ROW:
            self.tableRow.append(self.tableCell)
            self.tableCell = ""
        elif (
            oldstate == State.TABLE_ROW and self.state == State.TABLE and self.tableRow
        ):
            row = self.tableRow
            if len(row) == 3:
                self.table.append(Param(row[0], row[1], "", row[2]))
            elif len(self.tableRow) == 4:
                self.table.append(Param(row[0], row[1], row[2], row[3]))
            self.tableRow = []
        elif oldstate == State.TABLE and self.state == State.DOCUMENT:
            self.tokens.append(
                Token(self.subsection, " ".join(self.description), self.table)
            )
            self.subsection = ""
            self.description = []
            self.table = []

    def handle_data(self, data):
        if self.state == State.SECTION:
            self.section = data
        elif self.state == State.SUBSECTION:
            self.subsection = data

        if self.section in SKIP_SECTIONS or self.subsection in SKIP_SECTIONS:
            return

        if self.state == State.TABLE_CELL:
            self.tableCell += data
        elif self.state == State.DOCUMENT:
            self.description.append(data.strip())


def formatStruct(token: Token) -> str:
    result = [
        formatComment(token.description, 0),
        f"type {toCamelCase(token.name)} struct {{",
    ]
    for param in token.params:
        result.append("")
        result.append(formatComment(param.description, 2))
        result.append(
            f'  {toCamelCase(param.name)} {formatType(param.typeName)} `json:"{param.name}"`'
        )
    result.append("}\n")
    return "\n".join(result)


def getResultType(token: Token, allTypes: dict[str, str]) -> list[str]:
    maybeTypes = map(
        lambda x: x.split(),
        filter(
            lambda x: x.find("is returned") != -1 or x.find("returns an ") != -1,
            token.description.lower().split("."),
        ),
    )
    return list(
        {allTypes[t].name for mt in maybeTypes for t in mt if t in allTypes.keys()}
    )


def formatRequestResponse(token: Token, allTypes: dict[str, str]) -> str:
    ret = getResultType(token, allTypes)
    methodResult = [
        Param("raw", "Array of byte", False, "Raw response from the server")
    ]
    if len(ret) == 1:
        methodResult.append(
            Param("result", ret[0], False, "Decoded response from the server")
        )

    return "\n".join(
        [
            formatStruct(
                Token(
                    token.name + "Request",
                    f"Request for API call '{token.name}'",
                    token.params,
                )
            ),
            formatStruct(
                Token(
                    token.name + "Response",
                    f"Response for API call '{token.name}'",
                    methodResult,
                )
            ),
        ]
    )


def formatMethod(token: Token) -> str:
    methodName = toCamelCase(token.name)
    return "\n".join(
        [
            formatComment(token.description, 2),
            f"  {methodName}(request *{methodName}Request) (*{methodName}Response, error)",
        ]
    )


if __name__ == "__main__":
    with request.urlopen("https://core.telegram.org/bots/api") as response:
        data = response.read().decode("utf-8")
        # with open("api") as data:
        parser = Parser()
        parser.feed(data)
        tokenByName = {}
        structNames = {}

        print("package telegram")
        for t in parser.tokens:
            tokenByName[t.name] = t
            if t.name[0].islower() or t.name in ONEOF_TYPES:
                continue
            structNames[t.name.lower()] = t
            print(formatStruct(t))

        print("// Oneof type fields are merged into one")
        for k, ts in ONEOF_TYPES.items():
            names = []
            fields = {}
            for oneof in [tokenByName[t] for t in ts]:
                names.append(oneof.name)
                for param in oneof.params:
                    fields[param.name] = Param(
                        param.name, param.typeName, "", param.description
                    )
            print(
                formatStruct(
                    Token(k, f'Merged fields of {", ".join(names)}', fields.values())
                )
            )

        print("// Bot request and response types")
        for t in parser.tokens:
            if t.name[0].isupper() or t.name in SKIP_METHODS:
                continue
            print(formatRequestResponse(t, structNames))
            print()

        print("// Bot interface")
        print("type Telegram interface {")
        for t in parser.tokens:
            if t.name[0].isupper() or t.name in SKIP_METHODS:
                continue
            print(formatMethod(t))
            print()
        print("}")
