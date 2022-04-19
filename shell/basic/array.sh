#!/bin/bash

arr=(xx xyz "hello world" 1234 20.0 "你好")

echo "数组arr的内容: ${arr[@]}"
echo "数组arr的内容: ${arr[*]}"

echo "数组arr的长度: ${#arr[@]}"
echo "数组arr的长度: ${#arr[*]}"

# 常用的数组遍历方式

# 方式一
end_len=${#arr[@]}
end_index=`expr $end_len - 1`

for i in `seq 0 $end_index`
do
    pos=`expr $i + 1`
    echo "arr中第${pos}位置的元素为:${arr[$i]}"
done

# 方式二
count=1
for ele in "${arr[@]}"
do
    echo "arr中第${count}位置的元素为:$ele"
    count=`expr $count + 1`
done