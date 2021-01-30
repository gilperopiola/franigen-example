# Franigen-example

## Setup

### Go

[Install go 1.14](https://golang.org/dl/).

### Pre-commit

You should have [pre-commit](https://pre-commit.com/) installed locally.

Then run the following to install it in the repository.

```shell script
pre-commit install
```

It will pick up the configuration from `.pre-commit-config.yaml` on its own.

### Environment variables

[Install git-secret](https://git-secret.io/installation), get yourself a [gpg](https://gnupg.org/)
key if you don't have one, and have a member of the team add you to the list of
people who know the secret.

Then get the .env files with:

```shell script
git secret reveal
```

### Dependencies

```shell script
go mod download
go mod verify
```

## Run

Set the environment variables and then:

```shell script
go build . && go run franigen-example
```
