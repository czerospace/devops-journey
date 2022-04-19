#!/bin/bash
# 使用 ping 命令检测主机是否存在

# 建立一个管道
pipe_path=/tmp/pipe1
mkfifo ${pipe_path}
exec 6<>${pipe_path}

max_processes=10
if [ $# -eq 1 ];then
    max_processes=$1
fi

for i in `seq 1 $max_processes`
do
    echo "hello" >&6
done

# 多进程
check_host(){
    host=${host_prefix}.$1
    if ping $host -c1 -w1 > /dev/null 2>/dev/null;then
        echo "$host可达"
    else
        echo "$host不可达"
    fi
    echo "hello" >&6
}

host_prefix=192.168.51

for i in `seq 2 254`
do
    read -u6 name
    check_host $i &
done

wait

rm -f $pipe_path