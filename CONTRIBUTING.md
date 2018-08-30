# Contributing

## Vendoring a new dependency

The project currently checks in all vendored dependencies.
Our vendoring tool of choice at present is
[dep](https://github.com/golang/dep).

Adding a dependency is relatively straightforward:

```go
  dep ensure -add github.com/some-user/some-dep
```

## Running the tests

You'll need ginkgo! You can install it
[here](https://onsi.github.io/ginkgo/#getting-ginkgo).

Then you can run the tests like this:

```bash
  ginkgo -r -race -p .
```
