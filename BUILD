load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

exports_files(["fetch_types.py"])

genrule(
    name = "telegram_types",
    srcs = ["api"],
    outs = ["telegram_types.go"],
    cmd = "python $(location fetch_types.py) > \"$@\"",
    tools = ["fetch_types.py"],
    visibility = ["//visibility:private"],
)

go_library(
    name = "telegram_bot",
    srcs = [
        ":telegram_types",
    ],
    importpath = "github.com/lanseg/tgbot",
    deps = [
    ],
)
