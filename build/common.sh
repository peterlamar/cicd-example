#!/bin/bash

# This will canonicalize the path
ROOT=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)
cd $ROOT

function run_tests() {
  go test
}

function build_bin() {
  go build
}
