#!/bin/bash
# redis 服务的状态检测脚本

# 定制本地变量
status_arg="$1"

# 定制状态获取函数
redis_status(){
	#接收函数调用的参数
	status="$1"
	redis_status_value=$(/usr/bin/redis-cli -h 127.0.0.1 -p 6379 info | grep -w $status | cut -d ':' -f2)
	echo "${redis_status_value}"
}

# 调用函数
if [ "$#" == "1" ]
then
	redis_status $status_arg
else
	echo "Usage: /bin/bash $0 status_arg"
fi