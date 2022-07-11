#!/bin/bash

cd $(dirname $BASH_SOURCE)

DIR=./

for i in "$@"
do
case $1 in
    --dir=*)
    DIR="${i#*=}"
    shift # past argument=value
    ;;

    --woc_env=*)
    shift # past argument=value
    ;;

    *)
    SERVERS="$SERVERS $1"
    shift
    ;;
esac
done

echo $SERVERS

PROG_NAME=taal-client
cd $DIR

BASENAME=taal-client
LATEST=$(ls -rtd ${BASENAME}_* | tail -1)
echo $LATEST

if [[ "$LATEST" == "" ]]; then
  exit 1
fi

for var in $SERVERS
do
  echo Deploying $LATEST to $var
  ssh $var mkdir -p ./taal-client
  scp -r $LATEST $var:./taal-client/
  ssh $var <<EOL
cd taal-client
mkdir ${LATEST%.tar.gz}
cd ${LATEST%.tar.gz}
tar xvfz ../$LATEST
cd ../..
rm -f public_taal-client
ln -s ./taal-client/${LATEST%.tar.gz} public_taal-client
EOL
done
