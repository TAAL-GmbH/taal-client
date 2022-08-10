#!/bin/bash

cd $(dirname "$0")

BASENAME=taal-client

if [ -z "$BASENAME" ]; then
    echo "$BASENAME not found"
    exit 1
fi

sudo systemctl stop $BASENAME
sudo systemctl disable $BASENAME

