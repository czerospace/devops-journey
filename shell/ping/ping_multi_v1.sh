#!/bin/bash

# 多进程
check_host(){
    host=${host_prefix}.$1
    if ping $host -c1 -w1 > /dev/null 2>/dev/null;then
        echo "$host可达"
    else
        echo "$host不可达"
    fi
}

host_prefix=192.168.51

for i in `seq 2 254`
do
    check_host $i &
done

wait