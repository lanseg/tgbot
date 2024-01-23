#!/usr/bin/env python3
"""Parser for the telegram bot api docs.

API definition taken from here: https://core.telegram.org/bots/api
"""

from dataclasses import dataclass
from enum import Enum
from html import parser

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

# An override for some of the methods with an unclear return type
RETURN_TYPES: dict[str, str] = {
    "banChatMember": "bool",
    "close": "string",  # True or error 429
    "copyMessage": "MessageId",
    "createInvoiceLink": "string",
    "getGameHighScores": "GameHighScore",
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


class State(Enum):
    """Which part of the html doc are we parsing now."""

    DOCUMENT = 1
    SECTION = 2
    SUBSECTION = 3
    TABLE = 4
    TABLE_ROW = 5
    TABLE_CELL = 6


class Parser(parser.HTMLParser):
    """Parses Telegram's bot api doc into a set of tokens."""

    def __init__(self):
        super().__init__()
        self.tokens = []
        self.state = State.DOCUMENT
        self.section = ""
        self.subsection = ""

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
