#!/bin/bash

# 0 stdin 标准输入
# 1 stdout 标准输出
# 2 stderr 标准错误输出

echo -e "hello\nworld\n你好" > /tmp/file.in

exec 4</tmp/file.in
exec 5>/tmp/file.out

while read line
do
    echo "文件内容: $line"
done <&4

# exec M>&N M和N都是文件描述符
# > 等价于 1>
# 将标准输出重定向到 &5
exec >&5

echo "hello" #输出到stdout

exec 2>&5

ls xxxxx # stderr 重定向到 &5


# 关闭文件描述符
exec 4<&-
exec 5>&-

echo "hello"
