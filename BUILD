load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/aler9/rtsp-simple-server
# gazelle:go_naming_convention import_alias
gazelle(name = "gazelle")

go_library(
    name = "rtsp-simple-server_lib",
    srcs = ["main.go"],
    importpath = "github.com/aler9/rtsp-simple-server",
    visibility = ["//visibility:private"],
    deps = [
        "//internal/conf",
        "//internal/confwatcher",
        "//internal/hlsserver",
        "//internal/logger",
        "//internal/metrics",
        "//internal/pathman",
        "//internal/pprof",
        "//internal/rlimit",
        "//internal/rtmpserver",
        "//internal/rtspserver",
        "//internal/stats",
        "@in_gopkg_alecthomas_kingpin_v2//:go_default_library",
    ],
)

go_binary(
    name = "rtsp-simple-server",
    embed = [":rtsp-simple-server_lib"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "ingress",
    srcs = [ "ingress.go" ],
    deps = [
        "//internal/conf",
        "//internal/confwatcher",
        "//internal/hlsserver",
        "//internal/logger",
        "//internal/metrics",
        "//internal/pathman",
        "//internal/pprof",
        "//internal/rlimit",
        "//internal/rtmpserver",
        "//internal/rtspserver",
        "//internal/stats",
        "@com_github_aler9_gortsplib//pkg/headers:go_default_library",
        "@in_gopkg_alecthomas_kingpin_v2//:go_default_library",
    ],
    visibility = ["//visibility:public"],
)

# TODO: this test is directly executing docker commands, which is not something
#       you can do in a bazel sandbox. Needs fixing, commented out for now.
#go_test(
#    name = "rtsp-simple-server_test",
#    srcs = [
#        "main_hlsreader_test.go",
#        "main_rtmpreadpub_test.go",
#        "main_rtmpsource_test.go",
#        "main_rtspreadpub_test.go",
#        "main_rtspsource_test.go",
#        "main_test.go",
#    ],
#    embed = [":rtsp-simple-server_lib"],
#    deps = [
#        "@com_github_aler9_gortsplib//:go_default_library",
#        "@com_github_aler9_gortsplib//pkg/auth:go_default_library",
#        "@com_github_aler9_gortsplib//pkg/base:go_default_library",
#        "@com_github_aler9_gortsplib//pkg/headers:go_default_library",
#        "@com_github_stretchr_testify//require:go_default_library",
#    ],
#)
