SHELL := /bin/bash

.PHONY: install fmt vet test devchecknotest devcheck golangcilint importfmtlint

default: install

install:
	go mod tidy
	go install .

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -parallel=4 ./...

devchecknotest: install golangcilint importfmtlint

devcheck: devchecknotest test

golangcilint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 10m ./...

importfmtlint:
	go run github.com/pavius/impi/cmd/impi --local . --scheme stdThirdPartyLocal ./...