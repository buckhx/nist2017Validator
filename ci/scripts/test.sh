#!/bin/bash

set -e -x

export GOPATH=$PWD

mkdir -p src/github.com/aminasian/
cp -R ./nist-source src/github.com/aminasian/.
go get -u github.com/golang/dep/cmd/dep
dep ensure
go test -cover ./... | tee test_coverage_results.txt
sed -i -e 's/^/     /' test_coverage_results.txt

mv test_coverage_results.txt $GOPATH/test-output/.