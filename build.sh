#!/bin/bash

source ~/.nvm/nvm.sh

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


mkdir -p build/darwin/intel
mkdir -p build/darwin/arm
mkdir -p build/windows
mkdir -p build/linux

env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build --trimpath -o build/darwin/intel/${PROG_NAME} -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
env CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build --trimpath -o build/darwin/arm/${PROG_NAME} -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
#lipo -create -output build/darwin/$PROG_NAME build/darwin/intel/${PROG_NAME} build/darwin/arm/${PROG_NAME}

env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build --trimpath -o build/linux/$PROG_NAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"

# env CGO_ENABLED=1 GOOS=windows GOARCH=386 go build --trimpath -o build/windows/$PROG_NAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
env CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" /root/go/bin/go build  --trimpath  -o build/windows/$PROG_NAME -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
cd build/windows
jar cvf taal-client.zip ./taal-client.exe
cd ../..

# # env GOOS=windows GOARCH=386 go build --trimpath -o build/windows/$PROG_NAME.exe -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
# ssh root@138.201.123.88 << EOL
#   cd taal-client
#   git reset --hard
#   git pull ordishs master

#   rm -f taal-client
#   rm -f taal-client.exe

#   cd console
#   npm i
#   npm run build
#   cd ..

#   CGO_ENABLED=1 /root/go/bin/go build --trimpath -ldflags="-s -w -X main.commit=${GIT_COMMIT}"

#   CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="i686-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp" /root/go/bin/go build  --trimpath -ldflags="-s -w -X main.commit=${GIT_COMMIT}"
# EOL

# scp root@138.201.123.88:./taal-client/taal-client build/linux/
# scp root@138.201.123.88:./taal-client/taal-client.exe build/windows/
# cd build/windows
# jar cvf taal-client.zip ./taal-client.exe
# cd ../..

if [[ "$?" == "0" ]]; then
  echo $GIT_COMMIT > build/commit.dat
  echo "${PROG_NAME}: Built $FILENAME"

  cd build
  tar cvfz ../$FILENAME.tar.gz *  
  echo "${PROG_NAME}: Artifact $FILENAME.tar.gz"
else
  echo "${PROG_NAME}: Build FAILED"
fi
