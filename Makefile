SHELL := /bin/bash

.PHONY: install fmt vet test devchecknotest devcheck importfmtlint

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

devchecknotest: install importfmtlint fmt vet

devcheck: devchecknotest test

importfmtlint:
	go run github.com/pavius/impi/cmd/impi --local . --scheme stdThirdPartyLocal ./...