#!/bin/bash

# ---------------------------------
# Install middlewares.
# ---------------------------------

# Install golang
curl -LO https://go.dev/dl/go1.20.3.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz
mkdir $HOME/go
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# ---------------------------------
# Install load-params.service
# ---------------------------------
INITENV_NAME="load-params"
INITENV_SHELL="/etc/rc.d/${INITENV_NAME}.sh"
INITENV_SERVICE="/etc/systemd/system/${INITENV_NAME}.service"

# Install initialize environments shell script.
rm -rf "${INITENV_SHELL}" "${INITENV_SERVICE}"
cp "${INITENV_SHELL##*/}" "${INITENV_SHELL}"
chmod +x "${INITENV_SHELL}"

# Install load systems parameter service.
cp "${INITENV_SERVICE##*/}" "${INITENV_SERVICE}"
chmod +x "${INITENV_SERVICE}"
systemctl daemon-reload
systemctl enable ${INITENV_NAME}
systemctl start ${INITENV_NAME}


# ---------------------------------
# Install Sayka.service
# ---------------------------------
APP_NAME="Sayka"
APP_SERVICE="/etc/systemd/system/${APP_NAME}.service"

# Install application service
rm -rf "${APP_SERVICE}"
cp "${APP_SERVICE##*/}" "${APP_SERVICE}"
chmod +x "${APP_SERVICE}"
systemctl daemon-reload
