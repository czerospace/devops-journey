简单的根据 Failed 关键字过滤出登陆失败的源 IP，写入到 /etc/hosts.deny 中

```shell
#!/bin/bash
cat /var/log/secure|awk '/Failed/{print $(NF-3)}'|sort|uniq -c|awk '{print $2"="$1;}' > ip.txt
for i in `cat ip.txt`
do
  IP=`echo $i |awk -F= '{print $1}'`
  NUM=`echo $i|awk -F= '{print $2}'`
  if [ ${NUM} -gt 2 ]; then
    ipExists=`grep $IP /etc/hosts.deny |grep -v grep |wc -l`
    if [ ${ipExists} -lt 1 ];then
      echo "sshd:$IP:deny" >> /etc/hosts.deny
    fi
  fi
done
```

