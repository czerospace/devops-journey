#!/bin/bash
# 定制 TCP 连接状态效果

# 定制本地变量
user_status=$1

tcp_status(){
    TCP_STAT=$1
    TCP_STAT_VALUE=$(ss -ant | awk 'NR >1 {++s[$1]} END {for(k in s) print k,s[k]}'| grep $TCP_STAT |cut -d ' ' -f2)
    # 为了保证检测的数据都有内容，不存在的状态设置为0值
    if [ -z $TCP_STAT_VALUE ];then
        TCP_STAT_VALUE=0
    fi
    echo $TCP_STAT_VALUE
}

tcp_status $user_status