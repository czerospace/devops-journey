#!/bin/bash

var1="a b  c   d"
echo $var1
echo "$var1"

# echo 常用选项
# -n -e

echo "hello"     # 默认输出的最后会有一个换行符
echo -n "hello"  # 不要加默认的换行符

echo "换行"

#echo -n "please input your name:"
#read name
#echo "hello,$name"

echo "hello\nwor\tld"
echo -e "hello\nwor\tld"  # -e 会解释 "" 中的转义符

# echo 输出带颜色的字符 \033[33mxxxxxx\033[0m
echo -e "\033[33mhelloniko\033[0m"