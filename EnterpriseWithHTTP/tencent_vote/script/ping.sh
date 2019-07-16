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
    echo ${Search} | cut -f 1,2 -d ${1}
}
GetPartOne '/'


function Sum() {
    sum=0
    echo `expr 100 \* ${1}`
    for ((i=0;i<${1};i++));
    do
        let sum=sum+i
    done
    for i in `seq 1 2 ${2}`;
    do
        echo ${i}
    done
    echo `seq -s " | " ${2}`
    return ${sum}
}

Sum 10 12
echo $?

function Echo() {
    if [[ -e users ]];then
        `touch users`
    fi
    use=$(who > users)
    echo $(cat users)
}

Echo

echo $?

function Remove() {
    if test  -f users ; then
        `rm users`
    fi
}

Remove

function List() {
    l=("go" "python" "java")
    echo ${l[@]}
    echo `type ls`

    if [[ $[num1] -eq $[num2] ]]; then
        echo ""
    fi

}
List