#!/usr/bin/env bash
:<<EOF
1. project script
2. run
3. build
EOF
echo "Start Web Server..."

while read line
do
    l=`echo ${line} | awk '{print $0}'`
    echo ${l}
done < ../main.go

echo "Read File Done..."

target="tencent_vote"
readonly target

function build() {
    echo "go build -o ${target} ../main.go"
    `go build -o ${target} ../main.go`
}

function runWithGo() {
    echo "go run ../main.go"
    `go run ../main.go`
}

function runWithBinary(){
    echo "execute binary file"
    if [[ ! -e ${target} ]];then
        build
        ./${target}
    else
        ./${target}
    fi
}


function Remove() {
    if [[ -e ${target} ]]; then
        rm ${target}
    else
        echo "Success!"
    fi
}

function Commands() {
    commands=("run" "build" "start" "remove" "list")
    echo ${commands[@]}
    echo ${commands[*]}
    echo ${#commands[*]}
    echo ${#commands[@]}
    for i in ${commands[@]}
    do
        echo ${i}
    done
}

command=${1}
case ${command} in
run) runWithGo;;
build) build;;
start) runWithBinary;;
remove) Remove;;
list) Commands;;
esac
