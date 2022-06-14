#!/bin/zsh

NUMBER=$1
RANK=$2

mkdir -p ./ABC/${NUMBER}/${RANK}
cd ./ABC/${NUMBER}/${RANK}
touch mian.go && go mod init main
