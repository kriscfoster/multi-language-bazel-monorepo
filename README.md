# multi-language-bazel-monorepo

I wanted to try to set up a multi-language monorepo using bazel for experimentation purposes. You can follow the journey on YouTube.

1. bazelisk, WORKSPACE & first BUILD file - TODO: link to video.

## setup instructions

### Prerequisites

- bazelisk installed (`brew install bazelisk` on mac) - this manages the bazel installation & version using the [.bazelversion](./bazelversion) file.

### Running all builds

`bazel build //...`

### Running all tests

`bazel test //...`
