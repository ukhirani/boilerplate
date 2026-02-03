#!/usr/bin/env bash
set -e

echo "installing bp . . ."


OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
x86_64) ARCH=amd64 ;;
aarch64 | arm64) ARCH=arm64 ;;
*)
  echo "Unsupported arch: $ARCH"
  exit 1
  ;;
esac

URL="https://github.com/ukhirani/boilerplate/releases/latest/download/bp_${OS}_${ARCH}.tar.gz"

curl -L "$URL" | tar -xz
sudo mv bp /usr/local/bin/bp 2>/dev/null
hash -r 2>/dev/null || true

echo "bp installed successfully . . ."
