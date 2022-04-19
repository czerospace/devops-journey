#!/bin/bash

sum=0
i=0

while [ $i -le 100 ]
do
    let "sum += $i"
    let "i++"
done
echo "1+2+..+100=$sum"


echo "while遍历文件"
while read line
do
    echo "本行内容: $line"
done < test_while.txt

#按列遍历
echo
while read ch en
do
    echo "本行中文: $ch,本行英文:$en"
done < test_while.txt

echo

cat test_while.txt|while read line
do
     echo "本行内容: $line"
done

echo "########分割##########"
# until语法
sum=0
i=0
until [ $i -gt 100 ]
do
    let "sum += $i"
    let "i++"
done
echo "1+2+..+100=$sum"

echo "使用until语法遍历文件"
until ! read ch en
do
    echo "本行中文: $ch,本行英文:$en"
done < test_while.txt