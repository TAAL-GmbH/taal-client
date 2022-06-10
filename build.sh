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

cd $(dirname $BASH_SOURCE)

PROG_NAME="taal-client"


if [ -z "$(git status --porcelain)" ]; then
  # Working directory clean
  GIT_COMMIT=$(git rev-parse HEAD)
  LAST_BUILD=$(cat build/commit.dat)
  if [[ $GIT_COMMIT == $LAST_BUILD ]]; then
    echo "$PROG_NAME: Nothing new."
    exit 1
  fi
elif [[ "$FORCE" == "true" ]]; then
  echo "Force build requested."
  GIT_COMMIT=$(date +%Y-%m-%d-%H-%M-%S)
else
  echo "${PROG_NAME}: Project must be clean before you can build"
  exit 1
fi

rm -rf build

echo "${PROG_NAME}: Building..."

cd console
npm i
npm run build
cd ..


mkdir -p build/darwin
mkdir -p build/windows
mkdir -p build/linux

env GOOS=darwin GOARCH=amd64 go build -o build/darwin/${PROG_NAME}_intel -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
env GOOS=darwin GOARCH=arm64 go build -o build/darwin/${PROG_NAME}_arm -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
lipo -create -output build/darwin/$PROG_NAME build/darwin/${PROG_NAME}_intel build/darwin/${PROG_NAME}_arm

env GOOS=linux GOARCH=amd64 go build -o build/linux/$PROG_NAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
env GOOS=windows GOARCH=386 go build -o build/windows/$PROG_NAME.exe -ldflags="-s -w -X main.commit=${GIT_COMMIT}"

if [[ "$?" == "0" ]]; then
  echo $GIT_COMMIT > build/commit.dat
  echo "${PROG_NAME}: Built $PROG_NAME"
else
  echo "${PROG_NAME}: Build FAILED"
fi
