load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary")

package(default_visibility = ["//visibility:public"])

genrule(
  name = "bot_api_docs",
  outs = ["api.html"],
  cmd = "curl https://core.telegram.org/bots/api > \"$@\"",
  visibility = ["//visibility:private"],
)

genrule(
    name = "telegram_types",
    srcs = [":bot_api_docs"],
    outs = ["telegram_types.go"],
    cmd = "python $(location fetch_types.py) $(location //:bot_api_docs) > \"$@\"",
    tools = ["fetch_types.py"],
    visibility = ["//visibility:private"],
)

go_library(
    name = "telegram_bot",
    srcs = [
        "bot.go",
        ":telegram_types",
    ],
    importpath = "github.com/lanseg/tgbot",
)

go_binary(
    name = "main",
    srcs = [
        "main.go",
    ],
    deps = [
      ":telegram_bot",
    ],
)
