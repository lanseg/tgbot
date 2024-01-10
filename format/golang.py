#!/usr/bin/env python3
"""Generating golang telegram bot from the parsed api.

API definition taken from here: https://core.telegram.org/bots/api
"""

import api_parser

# Mapping from some primitive type name to its go equivalent
PRIMITIVE_TYPES: dict[str, str] = {
    "byte": "byte",
    "int": "int",
    "Integer": "int64",
    "Float": "float32",
    "Float number": "float32",
    "String": "string",
    "string": "string",
    "Boolean": "bool",
    "bool": "bool",
    "True": "bool",
    "true": "bool",
    "False": "bool",
    "false": "bool",
    "error": "string",
    "Integer or String": "string",
}

# An aggregate type that can be one of the following
ONEOF_TYPES: dict[str, list[str]] = {
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


def formatStruct(token: api_parser.Token) -> str:
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


def formatType(tgType: str) -> str:
    if " or " in tgType:
        return "interface{}"
    maybeArray = tgType.split("Array of ")
    tgType = PRIMITIVE_TYPES.get(maybeArray[-1], toCamelCase(maybeArray[-1]))
    if tgType not in PRIMITIVE_TYPES.values():
        tgType = "*" + tgType
    return "[]" * (len(maybeArray) - 1) + tgType


def getResultType(token: api_parser.Token, allTypes: dict[str, str]) -> list[str]:
    maybeTypes = map(
        lambda x: x.split(),
        filter(
            lambda x: x.find("is returned") != -1 or x.find("returns an ") != -1,
            token.description.lower().split("."),
        ),
    )
    result = []
    for maybeType in {
        allTypes[t].name for mt in maybeTypes for t in mt if t in allTypes.keys()
    }:
        if token.description.find(f"rray of {maybeType}") != -1:
            result.append(f"Array of {maybeType}")
        else:
            result.append(maybeType)
    return result


def formatRequestResponse(token: api_parser.Token, allTypes: dict[str, str]) -> str:
    ret = getResultType(token, allTypes)
    methodResult = [
        api_parser.Param("raw", "Array of byte", False, "Raw response from the server")
    ]
    if len(ret) == 1:
        methodResult.append(
            api_parser.Param(
                "result", ret[0], False, "Decoded response from the server"
            )
        )

    return "\n".join(
        [
            formatStruct(
                api_parser.Token(
                    token.name + "Request",
                    f"Request for API call '{token.name}'",
                    token.params,
                )
            ),
            formatStruct(
                api_parser.Token(
                    token.name + "Response",
                    f"Response for API call '{token.name}'",
                    methodResult,
                )
            ),
        ]
    )


def formatMethod(token: api_parser.Token) -> str:
    methodName = toCamelCase(token.name)
    return "\n".join(
        [
            formatComment(token.description, 2),
            f"  {methodName}(request *{methodName}Request) (*{methodName}Response, error)",
        ]
    )


def format(tokens: list[api_parser.Token]) -> str:
    tokenByName = {}
    structNames = {}
    result = ["package telegram"]
    for t in tokens:
        tokenByName[t.name] = t
        if t.name[0].islower() or t.name in ONEOF_TYPES:
            continue
        structNames[t.name.lower()] = t
        result.append(formatStruct(t))

    result.append("// Oneof type fields are merged into one")
    for k, ts in ONEOF_TYPES.items():
        names = []
        fields = {}
        for oneof in [tokenByName[t] for t in ts]:
            names.append(oneof.name)
            for param in oneof.params:
                fields[param.name] = api_parser.Param(
                    param.name, param.typeName, "", param.description
                )
        result.append(
            formatStruct(
                api_parser.Token(
                    k, f'Merged fields of {", ".join(names)}', fields.values()
                )
            )
        )

    result.append("// Bot request and response types")
    for t in tokens:
        if t.name[0].isupper() or t.name in api_parser.SKIP_METHODS:
            continue
        result.append(formatRequestResponse(t, structNames))
        result.append("")

    result.append("// Bot interface")
    result.append("type TelegramBot interface {")
    for t in tokens:
        if t.name[0].isupper() or t.name in api_parser.SKIP_METHODS:
            continue
        result.append(formatMethod(t))
        result.append("")
    result.append("}")
    return "\n".join(result)
