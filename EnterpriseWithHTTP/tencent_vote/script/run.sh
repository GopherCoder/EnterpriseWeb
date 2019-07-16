#!/usr/bin/env bash

echo "Start Web Server..."

while read line
do
    l=`echo ${line} | awk '{print $0}'`
    echo ${l}
done < ../main.go
go run ../main.go




