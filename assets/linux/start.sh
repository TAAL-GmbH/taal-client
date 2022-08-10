#!/bin/bash

cd $(dirname "$0")

BASENAME=taal-client

if [ -z "$BASENAME" ]; then
    echo "$BASENAME not found"
    exit 1
fi

if [ ! -f /etc/systemd/system/$BASENAME.service ]; then
    cat << EOB > $BASENAME.service
[Unit]
Description=$BASENAME
After=network.target

[Service]
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
User=$(whoami)
ExecStart=$(pwd)/$BASENAME
WorkingDirectory=$(pwd)
Type=simple
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target

EOB

sudo cp $BASENAME.service /etc/systemd/system/
sudo systemctl daemon-reload

fi

sudo systemctl start $BASENAME
sudo systemctl enable $BASENAME
