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


rm -rf build

echo "${PROG_NAME}: Building..."

mkdir -p build/darwin
mkdir -p build/windows
mkdir -p build/linux
mkdir -p build/raspian

if [[ $BUILD == "" ]]; then
  FILENAME=${PROG_NAME}_${GIT_COMMIT}
else
  FILENAME=${PROG_NAME}_${GIT_COMMIT}_${BUILD}
fi

env GOOS=linux GOARCH=amd64 go get -d -v

#env GOOS=darwin GOARCH=amd64 go build -o build/darwin/$FILENAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o build/linux/$FILENAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
#env GOOS=linux GOARCH=arm go build -o build/raspian/$FILENAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
#env GOOS=windows GOARCH=386 go build -o build/windows/$FILENAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"



if [[ "$?" == "0" ]]; then
  echo $GIT_COMMIT > build/commit.dat
  echo "${PROG_NAME}: Built $FILENAME"

 # cp -r assets build/darwin/
  cp -r assets build/linux/
 # cp -r assets build/raspian/
 # cp -r assets build/windows/
else
  echo "${PROG_NAME}: Build FAILED"
fi

cp build/linux/$FILENAME /tmp/taal-client

chmod +x /tmp/taal-client