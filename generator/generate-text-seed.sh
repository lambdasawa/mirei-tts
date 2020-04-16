#!/bin/bash

set -xe

pushd "$(pwd)"

direnv allow .
direnv exec . go run cmd/generate-text/main.go

popd
