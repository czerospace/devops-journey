#!/bin/bash
#当使用exec打开文件后，read命令每次都会将文件指针移动到文件的下一行读取，直到文件末尾，利用这个可以实现处理文件内容
seq 5 > /tmp/tmp.log

exec < /tmp/tmp.log
while read line
do
    echo $line
done

echo ok
rm -f /tmp/tmp.log