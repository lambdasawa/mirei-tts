#!/bin/bash

set -xeu

pushd "$(pwd)"

cd cmd/generate-text

direnv exec . go run main.go

popd

mv cmd/generate-text/text-seed.json ./

rm -rf tmp/dic
