#!/bin/bash

# LICH
# author: sebasion9
# config to compile for ARM
# this build aims to run on rasp zero 2 w

export CGO_ENABLED=1
export GOOS=linux
export GOARCH=arm
#export GOARCH=arm64
export GOHOSTARCH=x86
export CC=arm-linux-gnueabihf-gcc 
export CXX=arm-linux-gnueabihf-g++
export GOARM=7


HOST="raspserver"
SRC="main.go"
OUT_NAME="lich"
REMOTE_USER="sion"
REMOTE_DIR="/home/${REMOTE_USER}/${OUT_NAME}"
ARGS="-o"

echo "[BUILD] started building ${OUT_NAME} ..."
go build -v $ARGS $OUT_NAME $SRC

if [ $? -ne 0 ]
then
    echo "[BUILD] build failed"
    echo "[BUILD] exited with ${$?}"
    echo "[ERR] exiting script with 1"
    exit 1
fi

echo "[BUILD] build succeded"


echo "[SCP] exporting ${OUT_NAME} to $REMOTE_USER:$HOST"
scp $OUT_NAME $REMOTE_USER@$HOST:$REMOTE_DIR


if [ $? -ne 0 ]
then
    echo "[SCP] exited with ${$?}"
    echo "[ERR] exiting script with 1"
    exit 1
fi

echo "[SCP] copy succeded"

echo "[INFO] cleaning build"

go clean -v
if [ $? -ne 0 ]
then
    echo "[BUILD] clean failed"
    echo "[BUILD] exited with ${$?}"
    echo "[ERR] exiting script with 1"
    exit 1
fi

echo "[INFO] script succeded"
echo "[INFO] exiting script with 0"
exit 0



