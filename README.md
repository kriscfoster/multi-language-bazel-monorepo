# multi-language-bazel-monorepo

![Getting Started with Bazel](https://user-images.githubusercontent.com/17026751/147872728-cba68987-9a5c-4cb6-9777-47c51efc4c75.png)

I wanted to try to set up a multi-language monorepo using [bazel](https://bazel.build/) for experimentation purposes. You can follow the journey on YouTube.

1. [Bazel Tutorial: (Part 1) Getting Started, Bazelisk & our First Build Targets](https://youtu.be/BZYj6yfA6Bs)
   1. Bazel overview.
   1. Installing & using [bazelisk](https://github.com/bazelbuild/bazelisk).
   1. `WORKSPACE.bazel`, `BUILD.bazel` & creating our first build targets.
1. [Bazel Tutorial: Python targets with py_library, py_test, py_binary (internal & external deps)](https://youtu.be/8P3m1-U7v0k)
   1. Enabline Visual Studio Code Extension.
   1. [rules_python](https://github.com/bazelbuild/rules_python).
   1. Creating simple python library without any external dependencies.
      1. `py_library` target.
      1. `py_test` target.
   1. Flask application.
      1. `py_binary` target.
      1. Using third party dependencies.
      1. Using dependencies from the monorepo.
   1. Discuss python version, hermetic etc.
1. [GitHub Actions for Bazel Monorepo - Building & Testing (CI)](https://youtu.be/qiZXFdd8OPo).
1. [Bazel & Go Tutorial: Targets with go_library, go_test & go_binary (internal & external deps)](https://youtu.be/DB_kWimE2bw).
1. [Bazel & NodeJS Tutorial: library, test & nodejs_binary (internal & external deps)](https://youtu.be/lmWjRhFhvSc).

## Prerequisites

- [bazelisk installed](https://github.com/bazelbuild/bazelisk) (`brew install bazelisk` on mac) - bazel launcher that also manages the bazel installation & version using the [.bazelversion](./.bazelversion) file.

## Useful Commands

### Build all targets

- `bazel build //...`

### Run all test targets

- `bazel test //...`

### Run Python web app

- `bazel run //projects/python_web`
- http://localhost:5000

### Run Go web app

- `bazel run //projects/go_web`
- http://localhost:8000

### Run NodeJS web app

- `bazel run //projects/node_web`
- http://localhost:8080
