#!/bin/bash

source ~/.nvm/nvm.sh

for i in "$@"
do
case $i in
    --force*)
    FORCE="true"
    shift # past argument=value
    ;;
    --build=*)
    BUILD="${i#*=}"
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

#env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build --trimpath -o build/darwin/intel/${PROG_NAME} -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
#env CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build --trimpath -o  -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
#lipo -create -output build/darwin/$PROG_NAME build/darwin/intel/${PROG_NAME} build/darwin/arm/${PROG_NAME}
ssh taal@www.masa.gi << EOL
  cd taal-client
  git reset --hard
  git pull

  rm -rf build
  mkdir -p build/darwin/intel
  mkdir -p build/darwin/arm

  cd console
  PATH=$PATH:/opt/homebrew/bin npm i
  PATH=$PATH:/opt/homebrew/bin npm run build
  cd ..

  env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 /opt/homebrew/bin/go build --trimpath  -o build/${PROG_NAME}_amd64 -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
  env CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 /opt/homebrew/bin/go build --trimpath  -o build/${PROG_NAME}_arm64 -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
  lipo -create -output build/$PROG_NAME build/${PROG_NAME}_amd64 build/${PROG_NAME}_arm64
EOL

scp -r taal@www.masa.gi:./taal-client/build/* ./build/darwin 

env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build --trimpath -o build/linux/$PROG_NAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"

env CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" go build  --trimpath  -o build/windows/${PROG_NAME}.exe -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
cd build/windows
jar cvf taal-client.zip ./taal-client.exe
cd ../..

if [[ $BUILD == "" ]]; then
  FILENAME=${PROG_NAME}_${GIT_COMMIT}
else
  FILENAME=${PROG_NAME}_${GIT_COMMIT}_${BUILD}
fi

if [[ "$?" == "0" ]]; then
  echo $GIT_COMMIT > build/commit.dat
  
  cd build
  tar cvfz ../$FILENAME.tar.gz ./*
  echo "${PROG_NAME}: Built $FILENAME"
else
  echo "${PROG_NAME}: Build FAILED"
fi
