#!/bin/bash

for i in "$@"
do
case $i in
    --force*)
    FORCE="true"
    shift # past argument=value
    ;;

    *)
      # unknown option: ignore
    ;;
esac
done

# cd $(dirname $BASH_SOURCE)

PROG_NAME="taal-client"


GIT_COMMIT=$(date +%Y-%m-%d-%H-%M-%S)

cd console
npm run build
cd ..

rm -rf build

echo "${PROG_NAME}: Building..."

mkdir -p build/linux

if [[ $BUILD == "" ]]; then
  FILENAME=${PROG_NAME}_${GIT_COMMIT}
else
  FILENAME=${PROG_NAME}_${GIT_COMMIT}_${BUILD}
fi

env GOOS=linux GOARCH=amd64 go get -d -v

env GOOS=linux GOARCH=amd64 go build -o build/linux/$FILENAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"

if [[ "$?" == "0" ]]; then
  echo $GIT_COMMIT > build/commit.dat
  echo "${PROG_NAME}: Built $FILENAME"
else
  echo "${PROG_NAME}: Build FAILED"
fi

cp build/linux/$FILENAME /tmp/taal-client

chmod +x /tmp/taal-client