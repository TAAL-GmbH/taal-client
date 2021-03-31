#!/bin/bash

set -e


decho(){
    1>&2 echo $@
}

if [ -z ${LISTEN+x} ]; then
    decho "set LISTEN env variable"
    exit 1;
fi

if [ -z ${MAPI_URL+x} ]; then
    decho "set MAPI_URL env variable"
    exit 1;
fi

if [ -z ${API_KEY+x} ]; then
    decho "set API_KEY as environmental variable"
    exit 1;
fi

decho "registering API_KEY"
taal-client register $API_KEY


exec taal-client start

