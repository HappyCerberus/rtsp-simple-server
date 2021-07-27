load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_gookit_color",
    importpath = "github.com/gookit/color",
    sum = "h1:tXy44JFSFkKnELV6WaMo/lLfu/meqITX3iAV52do7lk=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_aler9_gortsplib",
    importpath = "github.com/aler9/gortsplib",
    sum = "h1:PPTqxgdDmDBQcDziEuLqS4VzmMTp5NSd7b3WZqQCtR4=",
    version = "v0.0.0-20210626112538-649c63cf5b62",
)

go_repository(
    name = "com_github_pion_rtp",
    importpath = "github.com/pion/rtp",
    sum = "h1:iGBerLX6JiDjB9NXuaPzHyxHFG9JsIEdgwTC0lp5n/U=",
    version = "v1.6.2",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "com_github_notedit_rtmp",
    importpath = "github.com/notedit/rtmp",
    replace = "github.com/aler9/rtmp",
    sum = "h1:95mXJ5fUCYpBRdSOnLAQAdJHHKxxxJrVCiaqDi965YQ=",
    version = "v0.0.0-20210403095203-3be4a5535927",
)

go_repository(
    name = "com_github_asticode_go_astits",
    importpath = "github.com/asticode/go-astits",
    sum = "h1:69cilL0/7uwsxdGNQgwmVBu6JP0aMicXSm91ukJDjgQ=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:hDPOHmpOpP40lSULcqw7IrRb/u7w6RpDC9399XyoNd0=",
    version = "v1.6.1",
)

go_repository(
    name = "com_github_xo_terminfo",
    importpath = "github.com/xo/terminfo",
    sum = "h1:QldyIu/L63oPpyvQmHgvgickp1Yw510KJOqX7H24mg8=",
    version = "v0.0.0-20210125001918-ca9a967f8778",
)

go_repository(
    name = "com_github_pion_rtcp",
    importpath = "github.com/pion/rtcp",
    sum = "h1:NT3H5LkUGgaEapvp0HGik+a+CpflRF7KTD7H+o7OWIM=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:UQZhZ2O0vMHr2cI+DC1Mbh0TJxzA3RcLoMsFw+aXw7E=",
    version = "v0.0.0-20190924025748-f65c72e2690d",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=",
    version = "v3.0.0-20200313102051-9f266ea9e77c",
)

go_repository(
    name = "com_github_icza_bitio",
    importpath = "github.com/icza/bitio",
    sum = "h1:squ/m1SHyFeCA6+6Gyol1AxV9nmPPlJFT8c2vKdj3U8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_pion_sdp",
    importpath = "github.com/pion/sdp",
    sum = "h1:21lpgEILHyolpsIrbCBagZaAPj4o057cFjzaFebkVOs=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_pion_randutil",
    importpath = "github.com/pion/randutil",
    sum = "h1:CFG1UdESneORglEsnimhUjf33Rwjubwj6xfiOXBa3mA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_asticode_go_astikit",
    importpath = "github.com/asticode/go-astikit",
    sum = "h1:+7N+J4E4lWx2QOkRdOf6DafWJMv6O4RRfgClwQokrH8=",
    version = "v0.20.0",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:hsms1Qyu0jgnwNXIxa+/V/PDsU6CfLf6CNO8H7IWoS4=",
    version = "v1.4.9",
)

go_repository(
    name = "com_github_icza_mighty",
    importpath = "github.com/icza/mighty",
    sum = "h1:8UsGZ2rr2ksmEru6lToqnXgA8Mz1DP11X4zSJ159C3k=",
    version = "v0.0.0-20180919140131-cfd07d671de6",
)

go_repository(
    name = "com_github_kballard_go_shellquote",
    importpath = "github.com/kballard/go-shellquote",
    sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
    version = "v0.0.0-20180428030007-95032a82bc51",
)

go_repository(
    name = "com_github_pion_sdp_v3",
    importpath = "github.com/pion/sdp/v3",
    sum = "h1:UNnSPVaMM+Pdu/mR9UvAyyo6zkdYbKeuOooCwZvTl/g=",
    version = "v3.0.2",
)

go_repository(
    name = "com_github_pkg_profile",
    importpath = "github.com/pkg/profile",
    sum = "h1:uCmaf4vVbWAOZz36k1hrQD7ijGRzLwaME8Am/7a4jZI=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=",
    version = "v2.2.8",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:DN0cp81fZ3njFcrLCytUHRSUkqBjfTo4Tx9RJTWs0EY=",
    version = "v0.0.0-20201221181555-eec23a3978ad",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:k+E048sYJHyVnsr1GDrRZWQ32D2C7lWs9JRc0bel53A=",
    version = "v0.0.0-20210610132358-84b48f89b13b",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:b3NXsE2LusjYGGjL5bxEVZZORm/YEFFrWFjR8eFrw/c=",
    version = "v0.0.0-20210423082822-04245dca01da",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:aRYxNxv6iGQlyVaZmk6ZgYEDa+Jg18DxebPSrd6bg1M=",
    version = "v0.3.6",
)


go_rules_dependencies()

go_register_toolchains(version = "1.16.5")

gazelle_dependencies()
