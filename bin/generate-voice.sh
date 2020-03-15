#!/bin/bash

set -xeu

pushd "$(pwd)"

cd cmd/trim

go run main.go

popd

mkdir -p ./voice
mv cmd/trim/trim/* ./voice/

rm -rf tmp/dic
