# multi-language-bazel-monorepo

I wanted to try to set up a multi-language monorepo using [bazel](https://bazel.build/) for experimentation purposes. You can follow the journey on YouTube.

1. `bazelisk`, `WORKSPACE` & first `BUILD` file - *TODO: link to video*.

## setup instructions

### prerequisites

- [bazelisk installed](https://github.com/bazelbuild/bazelisk) (`brew install bazelisk` on mac) - bazel launcher that also manages the bazel installation & version using the [.bazelversion](./.bazelversion) file.

### running all builds

`bazel build //...`

### running all tests

`bazel test //...`
