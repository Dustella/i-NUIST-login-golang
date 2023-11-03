#!/bin/bash

PLATFORMS="darwin/amd64 darwin/arm64" # amd64 only as of go1.5
PLATFORMS="$PLATFORMS windows/amd64 windows/386 windows/arm64" # arm compilation not available for Windows
PLATFORMS="$PLATFORMS linux/amd64 linux/386 linux/arm64"
PLATFORMS="$PLATFORMS linux/mips64 linux/mips64le" # experimental in go1.6
PLATFORMS_ARM="linux"

##############################################################
# Shouldn't really need to modify anything below this line.  #
##############################################################

type setopt >/dev/null 2>&1

SCRIPT_NAME=`basename "$0"`
FAILURES=""
SOURCE_FILE=`echo $@ | sed 's/\.go//'`
CURRENT_DIRECTORY=${PWD##*/}
OUTPUT="build/go-NUIST-login"

for PLATFORM in $PLATFORMS; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_FILENAME="${OUTPUT}-${GOOS}-${GOARCH}"
  if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
  CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILENAME} $@"
  echo "${CMD}"
  eval $CMD || FAILURES="${FAILURES} ${PLATFORM}"
done

# old ARM builds
for GOOS in $PLATFORMS_ARM; do
  GOARCH="arm"
  # build for each ARM version
  for GOARM in 7; do
    BIN_FILENAME="${OUTPUT}-${GOOS}-${GOARCH}${GOARM}"
    CMD="GOARM=${GOARM} GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILENAME} $@"
    echo "${CMD}"
    eval "${CMD}" || FAILURES="${FAILURES} ${GOOS}/${GOARCH}${GOARM}" 
  done
done

# eval errors
if [[ "${FAILURES}" != "" ]]; then
  echo ""
  echo "${SCRIPT_NAME} failed on: ${FAILURES}"
  exit 1
fi
