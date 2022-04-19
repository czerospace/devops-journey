#!/bin/bash


# 通过进程检测 nginx
check_nginx(){
    ps -ef|grep -v grep|grep "/usr/sbin/nginx" > /dev/null
    return $?
}

while ls > /dev/null
do
    sleep 2
    if check_nginx;then
        echo "nginx存在，休息2s"
        continue
    fi
    echo "启动 nginx..."
    systemctl start nginx
    sleep 2
    if check_nginx;then
        echo "启动nginx成功"
        continue
    fi
    echo "nginx 无法启动，请检查配置..."
    break
done