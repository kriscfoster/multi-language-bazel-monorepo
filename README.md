# multi-language-bazel-monorepo

![getting started with bazel](https://user-images.githubusercontent.com/17026751/146653297-18db0708-f9e4-4bb3-ba2f-469be7774e25.png)

I wanted to try to set up a multi-language monorepo using [bazel](https://bazel.build/) for experimentation purposes. You can follow the journey on YouTube.

1. [Bazel Tutorial: (Part 1) Getting Started, Bazelisk & our First Build Targets](https://youtu.be/BZYj6yfA6Bs)

    a. Bazel overview.

    b. Installing & using [bazelisk](https://github.com/bazelbuild/bazelisk).

    c. `WORKSPACE.bazel`, `BUILD.bazel` & creating our first build targets.

2. [Bazel Tutorial: (Part 2) Python (Creating Build & Test Targets)]()

    a. Visual Studio Code Extension.

    b. [rules_python](https://github.com/bazelbuild/rules_python).

    c. Creating simple python application without any external dependencies. 

    d. Bringing dependencies into the mix by creating a Flask application.

    e. Test targets using `py_test` rule.

## setup instructions

### prerequisites

- [bazelisk installed](https://github.com/bazelbuild/bazelisk) (`brew install bazelisk` on mac) - bazel launcher that also manages the bazel installation & version using the [.bazelversion](./.bazelversion) file.

### building all targets

`bazel build //...`

### running all test targets

`bazel test //...`

### python application

- `bazel test //projects/python/...`
- `bazel run //projects/python:main`
