# [Workspaces](https://golang.google.cn/blog/get-familiar-with-workspaces)

- [Support for go 1.18 "workspace" mode?](https://github.com/bazelbuild/bazel-gazelle/issues/1232)
- While working on multiple modules in the same repository, the go.work file defines the workspace instead of using replace directives in each module’s go.mod file.
    1. Create separate directories for different dependency needs.
    2. Run go work init in each of your workspace directories.
    3. Add the dependencies you want within each directory via go work use [path-to-dependency].
    4. To replace existing dependencies in your modules' go.mod files use go work replace [path-to-module].
    5. Run go run [path-to-your-module] in each workspace directory to use the dependencies specified by its go.work file.
