#!/bin/bash

case "$1" in
    dev) 
        echo "Running in test mode "
        go build -o questionbankgenerator && ./questionbankgenerator dev

        ;;

    dep) 
        echo "running in dev mode "
        go build -o questionbankgenerator && ./questionbankgenerator dep
        ;;
    
    *)
        echo "Usage : $0 {dev|dep}"
        exit 1

esac


