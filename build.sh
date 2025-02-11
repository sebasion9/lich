#!/bin/bash

# LICH
# author: sebasion9
# config to compile for ARM
# this build aims to run on rasp zero 2 w

export CGO_ENABLED=1
export GOOS=linux
export GOARCH=arm64
export CC=arm-linux-gnueabihf-gcc 
export CXX=arm-linux-gnueabihf-g++
export GOARM=7


HOST="raspserver"
SRC="main.go"
OUT_NAME="lich"
REMOTE_USER="sion"
REMOTE_DIR="/home/${REMOTE_USER}/${OUT_NAME}"
ARGS="-o"

echo "BUILDING ${OUT_NAME}"
go build $ARGS $OUT_NAME $SRC

if [ $? -ne 0 ]
then
    echo "BUILD EXITED WITH ${$?}"
    echo "EXITING SCRIPT WITH 1"
    exit 1
fi


echo "EXPORTING ${OUT_NAME} TO $REMOTE_USER:$HOST"
scp $OUT_NAME $REMOTE_USER@$HOST:$REMOTE_DIR


if [ $? -ne 0 ]
then
    echo "EXPORT EXITED WITH ${$?}"
    echo "EXITING SCRIPT WITH 1"
    exit 1
fi

echo "CLEANING BUILD"
go clean



