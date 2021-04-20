#!/bin/bash
########################################################################################################################
#
# This script is used to Install a Ping Identity pi CLI
#
# ------------
# Installs pi               into     .
# ------------

INSTALL_DIR=$(pwd)

if [ -f "$INSTALL_DIR/pi" ]; then
  echo "pi already installed in ${INSTALL_DIR}"
  echo "Please remove or move to reinstall"
  exit 1
fi

TMP_DIR=$(mktemp -d)
if [[ ! "$TMP_DIR" || ! -d "$TMP_DIR" ]]; then
  echo "Could not create temp dir."
  exit 1
fi

function cleanup {
  rm -rf "$TMP_DIR"
}

trap cleanup EXIT

cd "$TMP_DIR" >& /dev/null || echo "Unable to chanage to temporary directory" || exit 1

curl -s https://raw.githubusercontent.com/pingidentity/homebrew-devops/master/Formula/pi.rb |\
  grep url |\
  cut -d '"' -f 2 |\
  xargs curl -s -O -L

tar xzf ./*.tar.gz

cd pi-* || echo "Unable to chanage to pi-*" || exit 1

cp pi                  "$INSTALL_DIR/."

echo "
################################################################################
                     Welcome to Ping Identity pi CLI!

  You have just downloaded:

    ${INSTALL_DIR}/pi

  It is recommended to:
    1. copy your 'pi' to a location in your PATH (i.e. ~/bin or /usr/local/bin)
    2. Recommended additional utilities:
          base64 (used by pi)
          docker
          docker-compose
          envsubst
          helm
          jq (used by pi)
          jwt (used by pi)
          k9s
          kubectl
          kubectx (includes kubens)
          openssl (used by pi)

  Example:
    sudo mv ${INSTALL_DIR}/pi /usr/local/bin/.
    pi config
################################################################################
"
