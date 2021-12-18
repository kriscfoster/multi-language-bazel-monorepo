# multi-language-bazel-monorepo

![getting started with bazel](https://user-images.githubusercontent.com/17026751/146653297-18db0708-f9e4-4bb3-ba2f-469be7774e25.png)

I wanted to try to set up a multi-language monorepo using [bazel](https://bazel.build/) for experimentation purposes. You can follow the journey on YouTube.

1. [Bazel Tutorial: (Part 1) Getting Started, Bazelisk & our First Build Targets](https://youtu.be/BZYj6yfA6Bs)
   1. Bazel overview.
   1. Installing & using [bazelisk](https://github.com/bazelbuild/bazelisk).
   1. `WORKSPACE.bazel`, `BUILD.bazel` & creating our first build targets.
1. [Bazel Tutorial: (Part 2) Python (Creating Build & Test Targets)]()
   1. Visual Studio Code Extension.
   1. [rules_python](https://github.com/bazelbuild/rules_python).
   1. Creating simple python application without any external dependencies. 
   1. Bringing dependencies into the mix by creating a Flask application.
   1. Test targets using `py_test` rule.
   1. TODO: need to look into virtual environment stuff, python version used.

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
