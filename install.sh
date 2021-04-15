#!/bin/bash
########################################################################################################################
#
# This script is used to Install a Ping Identity DevOps tool and bash aliases
#
# ------------
# Installs pi               into     .
# Installs bash_profile.pi  into     .
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
cp etc/bash_profile.pi "$INSTALL_DIR/."

echo "
################################################################################
                     Welcome to Ping Identity pi CLI!

  You have just downloaded:

    ${INSTALL_DIR}/pi
    ${INSTALL_DIR}/bash_profile.pi

  It is recommended to:
    1. copy your 'pi' to a location in your PATH (i.e. ~/bin or /usr/local/bin)
    2. source the 'bash_profile.pi' in your shell to get all the goodness
       from Ping Identity DevOps!
    3. Ensure you have additional utilities used by pi:
          jq
          docker
          docker-compose
          openssl
          base64
          kustomize
          kubectl
          kubectx (includes kubens)
          envsubst

  Example:
    sudo mv ${INSTALL_DIR}/pi /usr/local/bin/.
    sudo mv ${INSTALL_DIR}/bash_profile.pi /usr/local/etc/.
    echo 'source /usr/local/etc/bash_profile.pi' >> ~/.bash_profile
    echo 'sourcePingIdentityFiles' >> ~/.bash_profile
    . ~/.bash_profile
    pi

  For more information please visit us at:

     http://devops.pingidentity.com

################################################################################
"
