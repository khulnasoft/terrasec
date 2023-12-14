#!/bin/bash
set -ex

if [[ -z "${TERRASEC_VERSION}" ]]; then
  TERRASEC_VERSION=${DEFAULT_TERRASEC_VERSION}
fi

VERSION=${TERRASEC_VERSION}

curl -LOs https://github.com/khulnasoft/terrasec/releases/download/v${VERSION}/terrasec_${VERSION}_Linux_x86_64.tar.gz
mkdir /usr/local/bin/terrasec_${VERSION}
tar -C  /usr/local/bin/terrasec_${VERSION} -xzf terrasec_${VERSION}_Linux_x86_64.tar.gz

mv /usr/local/bin/terrasec_${VERSION}/terrasec /usr/local/bin/terrasec

rm terrasec_${VERSION}_Linux_x86_64.tar.gz
rm -rf /usr/local/bin/terrasec_${VERSION}/
