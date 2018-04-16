#!/bin/sh

set -euo pipefail

# function that ensures a variable is set and prints nice error message.
ensure() {
  if [ -z "$1" ]; then
    echo "$2 must be set"
    exit 1
  fi
}

ensure "${PKG:-}" "PKG"
ensure "${OS:-}" "OS"
ensure "${ARCH:-}" "ARCH"
ensure "${VERSION:-}" "VERSION"

export CGO_ENABLED=0
export GOARCH="${ARCH}"
export GOOS="${OS}"

BUILD_DATE=$(date "+%FT%H:%M:%S %Z")

go install \
  -v \
  -installsuffix "static" \
  -ldflags "-extldflags -static -X ${PKG}/pkg/version.Version=${VERSION} -X \"${PKG}/pkg/version.BuildDate=${BUILD_DATE}\"" \
  ./...