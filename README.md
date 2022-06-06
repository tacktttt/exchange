# exchange

## Install

```sh
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

## Generate API code from openapi.yml

```sh
oapi-codegen -package openapi openapi.yml > ./presentation/openapi/openapi.gen.go
```
