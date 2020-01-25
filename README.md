# ecs-gen

This is a cli tool to generate AWS Fargate api infrastructure which are provisioned by terraform.

## Description

## Usage

```
NAME:
   ecs-gen

USAGE:
   ecs-gen [global options] command [command options] [arguments...]

AUTHOR:
   keng000

COMMANDS:
   init     Define basic infrastructure. e.g. VPC, Subnet, SG, TG...
   api      Generate target groups and ecr repos with auto scale setting
   deploy   Generate region derectory for module deploy
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### ecs-gen init

```
NAME:
   ecs-gen init - Define basic infrastructure. e.g. VPC, Subnet, SG, TG...

USAGE:
   ecs-gen init [PROJECT]
```

### ecs-gen api

```
NAME:
   ecs-gen api - Generate target groups and ecr repos with auto scale setting

USAGE:
   ecs-gen api [APIs...]
```

### ecs-gen deploy

```
NAME:
   ecs-gen deploy - Generate region derectory for module deploy

USAGE:
   ecs-gen deploy [command options] [REGIONs...]

OPTIONS:
   --list, -l  List available region list
```

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
$ go generate ./...
```
