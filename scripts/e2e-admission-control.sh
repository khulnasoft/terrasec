#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

export TERRASEC_BIN_PATH=${PWD}/bin/terrasec

go test -p 1 -v ./test/e2e/validatingwebhook/...