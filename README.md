# tgbot

Script that generates telegram bot directly from the [official bot API docs](https://core.telegram.org/bots/api).

## How it works?

The docs are quite simple, each method or structure definition consist of three parts:
1. &lt;H4&gt; Header: structs starts with uppercase letter, methods starts with a lowercase letter
2. Description
3. Table with methods or structure fields: name, type and description

Python script parses the document into a list of tokens and then converts the tokens into Golang (for now)
code. Struct is a simple golang struct and method is a function that accepts a FunctionNameRequest variable
and returns (FunctionNameResponse, error).

## Status

### What is working now

* Generating json-compatible golang types from the Telegram api docs
* Basic code to send queries and receive responses via htto

### What I am planning to add

* Generate complete methods so the most of the bot api would be available immediately

## Usage example for plain go, pre-generated types

A simple program that gets the latest updates available to your telegram bot

```go
package main

import (
        "encoding/json"
        "fmt"
        "os"

        "github.com/lanseg/tgbot"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Usage: ./main <bot_api_key>")
                os.Exit(1)
        }
        apiKey := os.Args[1]
        bot := tgbot.NewBot(apiKey)
        result, err := bot.GetUpdates(&tgbot.GetUpdatesRequest{})
        if err != nil {
                fmt.Printf("Could not perform request: %s\n", err)
                os.Exit(1)
        }

        asJson, _ := json.MarshalIndent(result, "", "    ")
        fmt.Printf("Got an update: %s\n", asJson)
}
```

## Usage example (bazel), code is generated on build

Add github repo to your workspace
```Bazel
workspace(name = "com_github_lanseg_tgbot_demo")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "tgbot",
    branch = "main",
    remote = "https://github.com/lanseg/tgbot",
)

git_repository(
  name = "io_bazel_rules_go",
  branch = "master",
  remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains(version = "1.19.1")

```

And this targets to the BUILD file:

```Bazel
load("@io_bazel_rules_go//go:def.bzl", "go_binary")

package(default_visibility = ["//visibility:public"])

go_binary(
  name = "main",
  srcs = ["main.go"],
  deps = ["@tgbot//:telegram_bot"]
)
```

And run it with ```bash bazel run //:main```
