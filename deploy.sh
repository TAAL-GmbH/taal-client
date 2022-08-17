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
      # unknown option: ignore
    ;;
  esac
done

echo $SERVERS

PROG_NAME=$(awk -F'"' '/^const progname =/ {print $2}' main.go)

cd $DIR

LATEST=$(ls -rt ${PROG_NAME}_*.tar.gz | tail -1)
echo $LATEST

if [[ "$LATEST" == "" ]]; then
  exit 1
fi

for var in $SERVERS
do
  echo Deploying ${PROG_NAME} to $var
  
  ssh $var mkdir -p ./${PROG_NAME}
  if [[ $? -ne 0 ]]; then
    exit 1
  fi

  scp $LATEST $var:./${PROG_NAME}/
  if [[ $? -ne 0 ]]; then
    exit 1
  fi

  ssh $var << EOL
cd ${PROG_NAME}
mkdir ${LATEST%.tar.gz}
cd ${LATEST%.tar.gz}
tar xvfz ../$LATEST
cd ../..
mkdir -p public_cdn
rm -f public_cdn/${PROG_NAME}
ln -s ./${PROG_NAME}/${LATEST%.tar.gz} public_cdn/${PROG_NAME}
EOL

done

if [[ $? -ne 0 ]]; then
    exit 1
fi

exit 0
