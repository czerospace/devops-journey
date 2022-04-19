#!/bin/bash

arr1=(a b c "hello,  world" d)
echo "遍历数组arr1,不带引号"
for a in ${arr1[@]}
do
    echo "遍历数组元素:$a"
done

echo

echo "遍历数组arr1,带引号"
for a in "${arr1[@]}"
do
    echo "遍历数组元素: $a"
done


echo "遍历数组arr1,用 * 不带引号"
for a in ${arr1[*]}
do
    echo "遍历数组元素:$a"
done

echo "遍历数组arr1,用 * 带引号"
for a in "${arr1[*]}"
do
    echo "遍历数组元素:$a"
done

#改变内置的分隔符
OLD_IFS=$IFS
IFS=$'\n'

for line in `cat test_for.txt`
do
    echo "循环变量:$line"
done

IFS=$OLD_IFS
