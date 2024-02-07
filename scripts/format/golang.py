#!/usr/bin/env python3
"""Generating golang telegram bot from the parsed api.

API definition taken from here: https://core.telegram.org/bots/api
"""

import textwrap
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
    """Format a word to fit a golang convention: first capital, known words to uppercase."""

    if word == "":
        return ""
    if word.lower() in ["id", "url", "ip"]:
        return word.upper()
    return word[0].upper() + word[1:]


def toCamelCase(usstr: str, capFirst: str = True) -> str:
    """Converts snake case (underscores) word to camel case."""

    result = "".join(map(formatWord, usstr.split("_")))
    if not capFirst and result.upper() != result:
        return result[0].lower() + result[1:]
    return result


def formatComment(comment: str, offset: int = 0, maxWidth: int = 95) -> str:
    """Turns a text into a golang comment block with intendation and no word breaking."""
    indent = " " * offset + "// "
    return textwrap.fill(
        comment.strip(), maxWidth, initial_indent=indent, subsequent_indent=indent
    )


def formatStruct(token: api_parser.Token) -> str:
    """Formats token as a golang struct definition."""

    result = []
    for param in token.params:
        result.extend(
            [
                formatComment(param.description, 2),
                f'  {toCamelCase(param.name)} {formatType(param.typeName)} `json:"{param.name},omitempty"`\n',
            ]
        )

    return "\n".join(
        [formatComment(token.description), f"type {toCamelCase(token.name)} struct {{"]
        + result
        + ["}"]
    )


def formatType(tgType: str) -> str:
    """Formats type name as a golang type definition, replacing unknown types with interface{}."""

    if " or " in tgType and tgType not in PRIMITIVE_TYPES:
        return "interface{}"
    maybeArray = tgType.split("Array of ")
    tgType = PRIMITIVE_TYPES.get(maybeArray[-1], toCamelCase(maybeArray[-1]))
    if tgType not in PRIMITIVE_TYPES.values():
        tgType = "*" + tgType
    return "[]" * (len(maybeArray) - 1) + tgType


def getResultType(token: api_parser.Token, allTypes: dict[str, str]) -> list[str]:
    """Tries to get information about the method result from a method description."""

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
    """Generates request and response structs for the method token."""

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


def formatMethod(token: api_parser.Token, allTypes: dict[str, str]) -> str:
    """Formats token as a golang method (member of a TelegramApi struct)."""
    name = toCamelCase(token.name)
    maybeReturnType = getResultType(token, allTypes)
    result = formatComment(token.description)
    if len(maybeReturnType) != 1:
        return result + textwrap.dedent(
            f"""
        func (a *TelegramApi) {name}(request *{name}Request) (*{name}Response, error) {{
          _, err := queryAndUnmarshal[interface{{}}](a.bot, \"{name}\", request)
          if err != nil {{
              return nil, err
          }}
          return &{name}Response {{ }}, nil
        }}"""
        )

    returnType = formatType(maybeReturnType[0])
    return result + textwrap.dedent(
        f"""
          func (a *TelegramApi) {name}(request *{name}Request) (*{name}Response, error) {{
              apiResponse, err := queryAndUnmarshal[{returnType}](a.bot, \"{name}\", request)
              if err != nil {{
                  return nil, err
              }}
              return &{name}Response {{ Result: apiResponse.Result }}, nil
          }}"""
    )


def formatTokens(tokens: list[api_parser.Token]) -> str:
    """Formats all tokens (types, methods, etc) to a golang file."""

    tokenByName = {}
    structNames = {}
    result = [
        "// Telegram bot API classes and enpoint",
        "package tgbot",
    ]
    for tok in tokens:
        tokenByName[tok.name] = tok
        if tok.name[0].islower() or tok.name in ONEOF_TYPES:
            continue
        structNames[tok.name.lower()] = tok
        result.append(formatStruct(tok))

    result.append("// Oneof type fields are merged into one")
    for typeName, memberTypes in ONEOF_TYPES.items():
        names = []
        fields = {}
        for oneof in [tokenByName[memberType] for memberType in memberTypes]:
            names.append(oneof.name)
            for param in oneof.params:
                fields[param.name] = api_parser.Param(
                    param.name, param.typeName, "", param.description
                )
        result.append(
            formatStruct(
                api_parser.Token(
                    typeName, f'Merged fields of {", ".join(names)}', fields.values()
                )
            )
        )

    result.append("// Bot request and response types")
    for tok in tokens:
        if tok.name[0].isupper() or tok.name in api_parser.SKIP_METHODS:
            continue
        result.append(formatRequestResponse(tok, structNames))
        result.append("")

    result.append(
        textwrap.dedent(
            """
      // Bot interface
      type TelegramApi struct {
        bot TelegramBot
      }
      
      func NewTelegramApi (bot TelegramBot) *TelegramApi {
          return &TelegramApi { bot: bot }
      }
      """
        )
    )
    for tok in tokens:
        if tok.name[0].isupper() or tok.name in api_parser.SKIP_METHODS:
            continue
        result.append(formatMethod(tok, structNames))
        result.append("")
    return "\n".join(result)
