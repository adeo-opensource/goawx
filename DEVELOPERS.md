# How to develop

## Build locally the library

### Needed

* Go 1.14
* Python 3
* [pre-commit](https://pre-commit.com/#install)

### Prerequisites

* Install pre-commit

```shell
pre-commit install
```

[//]: # (### Run tests and ensure they're all passing)

[//]: # ()
[//]: # (```shell)

[//]: # (go test ./test -count=1)

[//]: # (```)

### Push your code in the repo

```shell
git add .
git commit -m "feat: my new feature"
git push
```

### Test it in your code

```shell
go get github.com/adeo-opensource/goawx@my-branch-feature

go mod vendor
```
