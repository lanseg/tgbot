#!/usr/bin/env python3
"""Generating telegram bot api from the API docs.

API definition taken from here: https://core.telegram.org/bots/api
"""

import api_parser
import sys
from format import golang


if __name__ == "__main__":
    with open(sys.argv[1], encoding="utf-8") as api:
        docParser = api_parser.Parser()
        docParser.feed(api.read())
        print(golang.format(docParser.tokens))
