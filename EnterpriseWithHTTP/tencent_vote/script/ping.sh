#!/usr/bin/env bash

<<EOF
1, http command
2, ping web server
EOF

results=`http localhost:7201/ping`

echo ${results}
echo ""

function GetCode() {
    echo ${results:8:3}
}

function GetData() {
    echo ${results:20:4}
}

command=${1}
case ${command} in
code) GetCode;;
data) GetData;;
esac;


Search="https://localhost:7201/v1/api/votes"

function GetPartOne() {
    echo ${Search#*/}
    echo ${Search##*/}
    echo ${Search%/*}
    echo ${Search%%/*}
    echo ${Search} | cut -f 1,2 -d '/'
}
GetPartOne