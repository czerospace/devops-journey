#!/bin/bash

# 建立一个管道
pipe_path=/tmp/pipe1
# 如果 pipe1 不存在就建立
[ -e ${pipe_path} ] || mkfifo $pipe_path

# 创建可输出输出的描述符7
exec 7<>${pipe_path}

for i in `seq 1 5`
do
    echo "向管道写入hello$i"
    echo "hello$i" >&7
    read -u7 name
    echo "从管道中读取内容: $name"
done