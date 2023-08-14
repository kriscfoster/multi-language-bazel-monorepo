# multi-language-bazel-monorepo

![Getting Started with Bazel](https://user-images.githubusercontent.com/17026751/147872728-cba68987-9a5c-4cb6-9777-47c51efc4c75.png)

I wanted to try to set up a multi-language monorepo using [bazel](https://bazel.build/) for experimentation purposes. You can follow the journey on YouTube.

1. [Bazel Tutorial: (Part 1) Getting Started, Bazelisk & our First Build Targets](https://youtu.be/BZYj6yfA6Bs).
2. [Bazel Tutorial: How to Build, Run, Test & Query (deps, rdeps, tags)](https://youtu.be/vZnXXx4Oh7c).
3. [Bazel Tutorial: Python targets with py_library, py_test, py_binary (internal & external deps)](https://youtu.be/8P3m1-U7v0k).
4. [GitHub Actions for Bazel Monorepo - Building & Testing (CI)](https://youtu.be/qiZXFdd8OPo).
5. [Bazel & Go Tutorial: Targets with go_library, go_test & go_binary (internal & external deps)](https://youtu.be/DB_kWimE2bw).
6. [Bazel & NodeJS Tutorial: library, test & nodejs_binary (internal & external deps)](https://youtu.be/lmWjRhFhvSc).
7. [Bazel & Docker Tutorial: Building container images with bazel (local & remote registry)](https://youtu.be/hLD6vKl4Txc).
8. [Deploying from a Bazel Monorepo to Heroku](https://youtu.be/AHvON-xl_Ds).
9. [Using Gazelle to Improve Multi-Language Bazel Monorepo](https://youtu.be/MUP35hfK0q4).
10. [Bazel & Java Tutorial: java_library & java_binary (internal & external dependencies)](https://youtu.be/HPTzVHOcins).
11. [Bazel & Docker: Using Custom Base Images](https://youtu.be/thYPUrhA82A).

## Prerequisites

- [bazelisk installed](https://github.com/bazelbuild/bazelisk) (`brew install bazelisk` on mac) - bazel launcher that also manages the bazel installation & version using the [.bazelversion](./.bazelversion) file.

## Useful Commands

### Updating npm dependencies from package.json

- `bazel run -- @pnpm//:pnpm i --dir $PWD`

### Build all targets

- `bazel build //...`

### Test all test targets

- `bazel test //...`

### Sync Go Dependencies with `go.mod`

- `bazel run //:gazelle-update-repos`

### Format BUILD.bazel files (completely generates BUILd.bazel files for go projects)

- `bazel run //:gazelle`

### Run Python web app

- `bazel run //projects/python_web`
- http://localhost:5000

### Run Go web app

- `bazel run //projects/go_web`
- http://localhost:8080

### Run NodeJS web app

- `bazel run //projects/node_web`
- http://localhost:8080

### Run TypeScript app

- `bazel run //projects/ts_app`

### Build & Run NodeJS web app docker image

```
➜ bazel run projects/node_web:oci_tarball
...
Loaded image: projects/node_web:oci_tarball
...
➜ docker run -p 8080:8080 projects/node_web:oci_tarball
...
listening on port 8080
```

### Build & Run Go web app docker image
```
➜ bazel run projects/go_web:oci_tarball
...
Loaded image: projects/go_web:oci_tarball
➜
➜ docker run -p 8080:8080 projects/go_web:oci_tarball
2022/05/30 20:35:51 Going to listen on port 8080
```

### Deploying Go web app to Heroku

```
➜ bazel** run projects/go_web:bazoku-deployment --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
...
```
