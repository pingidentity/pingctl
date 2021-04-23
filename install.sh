#!/bin/bash
########################################################################################################################
#
# This script is used to Install pingctl, a Ping Identity CLI
#
# ------------
# Installs pingctl               into     .
# ------------

INSTALL_DIR=$(pwd)

if [ -f "$INSTALL_DIR/pingctl" ]; then
  echo "pingctl already installed in ${INSTALL_DIR}"
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

curl -s https://raw.githubusercontent.com/pingidentity/homebrew-devops/master/Formula/pingctl.rb |\
  grep url |\
  cut -d '"' -f 2 |\
  xargs curl -s -O -L

tar xzf ./*.tar.gz

cd pingctl-* || echo "Unable to chanage to pingctl-*" || exit 1

cp pingctl                  "$INSTALL_DIR/."

echo "
################################################################################
                     Welcome to Ping Identity pingctl CLI!

  You have just downloaded:

    ${INSTALL_DIR}/pingctl

  It is recommended to:
    1. copy your 'pingctl' to a location in your PATH (i.e. ~/bin or /usr/local/bin)
    2. Recommended additional utilities:
          base64 (used by pingctl)
          docker
          docker-compose
          envsubst
          helm
          jq (used by pingctl)
          jwt (used by pingctl)
          k9s
          kubectl
          kubectx (includes kubens)
          openssl (used by pingctl)

  Example:
    sudo mv ${INSTALL_DIR}/pingctl /usr/local/bin/.
    pingctl config
################################################################################
"
