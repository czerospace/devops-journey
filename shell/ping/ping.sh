#!/bin/bash

# 单进程

host_prefix=192.168.51
for i in `seq 2 254`
do
    host=${host_prefix}.$i
    if ping $host -c1 -w1 > /dev/null 2>/dev/null;then
        echo "$host可达"
    else
        echo "$host不可达"
    fi
done