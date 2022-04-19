#!/bin/bash

var1=hello
var1_xyz=abc

echo "$var1_xyz"    # abc
echo "${var1}_xyz"  # hello_xyz

# declare -i 指定整型 -r 只读变量 -a 指定数组 -f 指定函数名
declare -i var2
var2=xyz    # 赋值语句失败，var2 默认值为 0
echo "\$var2=$var2"

var2=234    # 赋值成功
echo "\$var2=$var2"

declare -r var3=abc # 只读变量,不会往下运行var3=xyz
#var3=xyz
echo "\$var3=$var3"

declare |grep ^var1

# declare 同义命令 set
unset var1  #清除变量var1
echo "\$var1=$var1"

# `` 反引号 $() 获取里面的Linux命令执行结果
LOCAL_IP=`cat /etc/hosts|grep \`hostname\`|awk '{print $1}'`
LOCAL_IP_OTHER=$(cat /etc/hosts|grep `hostname`|awk '{print $1}')

echo "\$LOCAL_IP=$LOCAL_IP"
echo "\$LOCAL_IP_OTHER=$LOCAL_IP_OTHER"