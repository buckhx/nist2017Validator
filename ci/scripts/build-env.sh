#!/bin/bash

set -e -x

output_directory = build-env

cat << EOF >> "${output_directory}/Dockerfile"

	  FROM golang:1.8-alpine

      RUN  apk add --no-cache git mercurial curl \
        &&  apk add --no-cache bash gawk sed grep bc coreutils openssl \
        && curl -L -s https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 -o $GOPATH/bin/dep \
        && chmod +x $GOPATH/bin/dep

EOF

cp -R ./nist-source "${output_directory}/nist-source"
