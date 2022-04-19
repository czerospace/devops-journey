#!/bin/bash

# [] 字符串判断 < > = !=
# [] 数值判断 -gt -lt -ge -le -eq

x="abc"
if [ $x = "abc" ];then
    echo "x = abc"
else
    echo "x != abc"
fi

y=123
if [ $y -eq 123 ];then
    echo "y = 123"
else
    echo "y != 123"
fi


# -e "文件路径" 判断文件是否存在；
# -f 是否为普通文件；
# -d  目录是否存在；
if [ -d "etc" ];then
    echo "etc/存在"
elif [ ! -d "etc" ];then
    echo "etc不存在"
fi


# -z "变量" 如果变量为0，则条件成立;
DATA_DIR=$1
if [ -z "$DATA_DIR" ];then
    echo "DATA_DIR不存在"
    exit 1
fi

# -n "变量" 如果变量不为0，则条件成立；
if [ -n "$DATA_DIR" ];then
    echo "DATA_DIR=$DATA_DIR"
fi


# [[]]

# linux命令中 [[ -e "/etc" ]] && echo "/etc存在"  相当于下面 if
if [[ -e "/etc" ]];then
    echo "/etc路径存在"
fi

# if 中的条件 与 或
&& -a
|| -o