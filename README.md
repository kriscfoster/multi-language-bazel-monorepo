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

### Run TypeScript app

- `bazel run //projects/ts_app`

### Run React app

- `bazel run //projects/react_app:start`

### Build & Run NodeJS web app docker image
```
➜ bazel run projects/node_web:node_web_image --@io_bazel_rules_docker//transitions:enable=yes -- --norun
...
INFO: Build completed successfully, 1 total action
Loaded image ID: sha256:XXX
Tagging YYY as bazel/projects/node_web:node_web_image
➜
➜ docker run -p 8080:8080 bazel/projects/node_web:node_web_image
listening on port 8080
```

### Build & Run Go web app docker image
```
➜ bazel run projects/go_web:go_web_image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 -- --norun
...
INFO: Build completed successfully, 1 total action
Loaded image ID: sha256:XXX
Tagging YYY as bazel/projects/go_web:go_web_image
➜
➜ docker run -p 8000:8000 bazel/projects/go_web:go_web_image
2022/05/30 20:35:51 Going to listen on port 8000
```

### Build & Run Python web app docker image
```
➜ bazel run projects/python_web:python_web_image -- --norun
...
INFO: Build completed successfully, 1 total action
Loaded image ID: sha256:XXX
Tagging YYY as bazel/projects/python_web:python_web_image
➜
➜ docker run -p 5000:5000 bazel/projects/python_web:python_web_image
...
* Running on http://127.0.0.1:5000
```

### Publishing Python web app docker image
```
➜ bazel run projects/python_web:publish         
...
INFO: Build completed successfully, 1 total action
2022/06/24 20:13:33 Successfully pushed Docker image to registry.hub.docker.com/krisfoster96/monorepo-python-web:1 - registry.hub.docker.com/krisfoster96/monorepo-python-web@sha256:024bcf5dd677d6fbce32fcf9d09329f4c80931cc12c90965bb397af1f497bf39
```

### Deploying Go web app to Heroku

```
➜ bazel** run projects/go_web:bazoku-deployment --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
...
```
