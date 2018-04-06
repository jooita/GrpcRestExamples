#!/bin/sh

function go_grpc_server {
	go build -o grpc-server/go/EchoServer grpc-server/go/EchoServer.go
	./grpc-server/go/EchoServer & export PID=$!
	printf "go_grpc_server pid: ${PID}\n"
}

function python_grpc_server {
	PYTHONPATH=$PYTHONPATH:$PWD/../echopb python grpc-server/python/EchoServer.py & export PID=$!
	printf "python_grpc_server pid: ${PID}\n"
}

function reverse_proxy_server {
	go build -o reverse-proxy-server/main reverse-proxy-server/main.go
	./reverse-proxy-server/main & export PID=$!
	printf "reverse_proxy_server pid: ${PID}\n"
}

if [[ $# -eq 0 ]] ; then                                                          
	go_grpc_server
	reverse_proxy_server
    exit 1                                                                        
fi

while [ "$1" != "" ]; do
    case $1 in
        python)
				python_grpc_server
				reverse_proxy_server
            ;;
        go)
				go_grpc_server
				reverse_proxy_server
            ;;
        *)
            ;;
    esac
    shift
done
