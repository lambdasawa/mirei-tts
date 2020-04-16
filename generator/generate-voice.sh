#!/bin/bash

set -xe

mkdir -p data/voice

direnv allow .
direnv exec . go run cmd/trim/main.go --config generator/trim.yml
