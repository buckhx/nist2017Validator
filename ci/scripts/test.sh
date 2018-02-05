#!/bin/bash

set -e -x

export GOPATH=$PWD

sudo apt-get install unzip

curl -fsSL -o /tmp/dep.zip \
https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-amd64.zip \
&& unzip /tmp/dep.zip -d /bin \
&& chmod +x /bin/dep

echo 'current directory contents'
echo $PWD
ls -lat

ls /bin


mkdir -p src/github.com/aminasian/
cp -R ./nist-source src/github.com/aminasian/.

$GOPATH/bin/dep ensure
go test -cover ./... | tee test_coverage_results.txt
sed -i -e 's/^/     /' test_coverage_results.txt

mv test_coverage_results.txt $GOPATH/test-output/.