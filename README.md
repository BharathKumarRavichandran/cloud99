# cloud99

## Prerequisites
- Install [Go 1.13.5](https://golang.org/dl/#go1.13.5)
- MySQL

## Check prerequisites
- Check the go version installed.
```
go version
```

## Build instructions

- Download the repository and `cd` into it.
```
go get github.com/BharathKumarRavichandran/cloud99
cd $GOPATH/src/github.com/BharathKumarRavichandran/cloud99
```
- Install dependencies - `./scripts/install_dep.sh`
- Run `cp config.json.example config.json`
- Fill in the database credentials in the `Dev` section of **config.json**.
- Run `go run main.go`
