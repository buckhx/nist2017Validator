#!/bin/sh

set -e -x

export GOPATH=$PWD

apk add --no-cache git mercurial curl \
 &&  apk add --no-cache bash gawk sed grep bc coreutils openssl \
 && curl -L -s https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 -o $GOPATH/bin/dep \
 && chmod +x $GOPATH/bin/dep

ls $GOPATH/bin/dep

mkdir -p src/github.com/aminasian/
cp -R ./nist-source src/github.com/aminasian/.

$GOPATH/bin/dep ensure

go test -cover ./... | tee test_coverage_results.txt
sed -i -e 's/^/     /' test_coverage_results.txt

mv test_coverage_results.txt $GOPATH/test-output/.