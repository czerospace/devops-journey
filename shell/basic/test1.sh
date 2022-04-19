#!/bin/bash

set -e // -e 如果报错不会继续执行后面的内容，+e 如果报错会继续执行后面的内容

DEPLOY_DIR=test
USE_FILE=test.txt

if [ $# -ge 1 ]
then
    DEPLOY_DIR=$1
fi

if [ $# -eq 2 ]
then
    USE_FILE=$2
fi

echo $DEPLOY_DIR
echo $USE_FILE