# ecs-gen/

This is a cli tool to generate AWS Fargate api infrastructure which are provisioned by terraform.

## Description

## Usage

### ecs-gen init

### ecs-gen api

## Install

To install, use `go get`:

```bash
$ go get -d github.com/keng000/ecs-gen/
$ cd $GOPATH/src/github.com/keng000/ecs-gen/
$ go install
```

## Contribution

1. Fork ([https://github.com/keng000/ecs-gen//fork](https://github.com/keng000/ecs-gen/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[keng000](https://github.com/keng000)


### Deps

- go-assets-builder
- go-assets

when you update the tpml file, then you need to append into assets.
run below.

```
cd skeleton
go-assets-builder -p skeleton resource/ > assets.go
```