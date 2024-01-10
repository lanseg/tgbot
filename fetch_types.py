#!/usr/bin/env python3
"""Generating golang telegram bot api from the API docs.

API definition taken from here: https://core.telegram.org/bots/api
"""

import api_parser
import format_golang


if __name__ == "__main__":
    with open("api", encoding="utf-8") as api:
        docParser = api_parser.Parser()
        docParser.feed(api.read())
        tokenByName = {}
        structNames = {}

        print("package telegram")
        for t in docParser.tokens:
            tokenByName[t.name] = t
            if t.name[0].islower() or t.name in api_parser.ONEOF_TYPES:
                continue
            structNames[t.name.lower()] = t
            print(format_golang.formatStruct(t))

        print("// Oneof type fields are merged into one")
        for k, ts in api_parser.ONEOF_TYPES.items():
            names = []
            fields = {}
            for oneof in [tokenByName[t] for t in ts]:
                names.append(oneof.name)
                for param in oneof.params:
                    fields[param.name] = api_parser.Param(
                        param.name, param.typeName, "", param.description
                    )
            print(
                format_golang.formatStruct(
                    api_parser.Token(
                        k, f'Merged fields of {", ".join(names)}', fields.values()
                    )
                )
            )

        print("// Bot request and response types")
        for t in docParser.tokens:
            if t.name[0].isupper() or t.name in api_parser.SKIP_METHODS:
                continue
            print(format_golang.formatRequestResponse(t, structNames))
            print()

        print("// Bot interface")
        print("type TelegramBot interface {")
        for t in docParser.tokens:
            if t.name[0].isupper() or t.name in api_parser.SKIP_METHODS:
                continue
            print(format_golang.formatMethod(t))
            print()
        print("}")
