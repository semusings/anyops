#!/bin/bash
#
# AnyOps installer
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/bhuwanupadhyay/anyops/master/install.sh | bash

# When releasing AnyOps, the releaser should update this version number
# AFTER they upload new binaries.
VERSION="1.0.0-alpha"

set -e

function copy_binary() {
  if [[ ":$PATH:" == *":$HOME/.local/bin:"* ]]; then
    if [ ! -d "$HOME/.local/bin" ]; then
      mkdir -p "$HOME/.local/bin"
    fi
    mv anyops "$HOME/.local/bin/anyops"
  else
    echo "Installing Tilt to /usr/local/bin which is write protected"
    echo "If you'd prefer to install Anyops without sudo permissions, add \$HOME/.local/bin to your \$PATH and rerun the installer"
    sudo mv anyops /usr/local/bin/anyops
  fi
}

function install_anyops() {
  if [[ "$OSTYPE" == "linux"* ]]; then
    case $(uname -m) in
    aarch64) ARCH=arm64 ;;
    armv7l) ARCH=arm ;;
    *) ARCH=$(uname -m) ;;
    esac
    set -x
    curl -fsSL https://github.com/bhuwanupadhyay/anyops/releases/download/v$VERSION/anyops-v$VERSION-linux-$ARCH.tar.gz | tar -xzv anyops
    copy_binary
  elif [[ "$OSTYPE" == "darwin"* ]]; then
    ARCH=$(uname -m)
    set -x
    curl -fsSL https://github.com/bhuwanupadhyay/anyops/releases/download/v$VERSION/anyops-v$VERSION-darwin-$ARCH.tar.gz | tar -xzv anyops
    copy_binary
  else
    set +x
    echo "The AnyOps installer does not work for your platform: $OSTYPE"
    echo "For other installation options, download from the following page:"
    echo "https://github.com/bhuwanupadhyay/anyops/releases/"
    echo "Thank you!"
    exit 1
  fi

  set +x
}

# so that we can skip installation in CI and just test the version check
if [[ -z $NO_INSTALL ]]; then
  install_anyops
fi
