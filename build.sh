#!/bin/bash

# LICH
# author: sebasion9
# config to compile for ARM
# this build aims to run on rasp zero 2 w

# check if build & run
RUN=false
if [ "$1" = "run" ] || [ "$1" = "RUN" ]; then
    RUN=true
fi

# env variables, idk why this way
declare -A ENVS
ENVS["CGO_ENABLED"]=1
ENVS["GOOS"]=linux
ENVS["GOARCH"]=arm
ENVS["GOHOSTARCH"]=x86
ENVS["CC"]=arm-linux-gnueabihf-gcc 
ENVS["CXX"]=arm-linux-gnueabihf-g++
ENVS["GOARM"]=7

for key in ${!ENVS[@]}; do
    export ${key}=${ENVS[${key}]}
done



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

# run
if [ "$RUN" = "true" ];then
    echo "[RUN] entering build & run mode"
    echo "[RUN] opening ssh session"
    ssh -t -Y $REMOTE_USER@$HOST "$REMOTE_DIR/$OUT_NAME"
    echo "[RUN] exiting ssh"
else
    echo "[RUN] build only mode"
fi

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



