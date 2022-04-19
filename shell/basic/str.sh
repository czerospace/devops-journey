#!/bin/bash

# 字符串长度
var1=abcdefghijklmnabcdefg
echo "获取变量var1的长度:${#var1}"

var2="中文长度"
echo "中文字符长度=${#var2}"

# 字符串切片
echo "去掉var1前面两个字符: ${var1:2}"

echo "从第3个字符开始，连续取三个字符: ${var1:2:3}"

# 删除某个子串 # ## % %% 删除匹配子串
# # 和 ## 是从左开始删除匹配
# % 和 %% 是从右开始删除匹配

echo "从左边删除，最小匹配: a*d,结果为：${var1#a*d}" 
#结果:efghijklmnabcdefg
echo "从左边删除，最大匹配: a*d,结果为：${var1##a*d}" 
#结果:efg

echo "从右边删除，最小匹配: a*d,结果为: ${var1%a*g}"
#结果为: abcdefghijklmn
echo "从右边删除，最大匹配: a*d,结果为: ${var1%%a*g}"
#结果为: 

echo "原始字符串为: $var1"
echo "替换abc字符串为1234，只替换第一个匹配结果 ${var1/abc/1234}" 
echo "替换abc字符串为1234，替换全部匹配结果 ${var1//abc/1234}" 

echo "从头开始匹配替换，结果: ${var1/#abc/你好}"
echo "从尾开始匹配替换，结果: ${var1/%efg/你好}"


# 直接给变量 var3 赋值
echo "变量var3的值: ${var3-xyz}"
echo "变量var3的值: ${var3-abc}"

var3=  # var3 赋空值

# 如果变量值为空，则不会生效赋值
echo "使用短横线默认值，输出结果: ${var3-xyz}"

# 如果变量值为空，则生效赋值
echo "使用短横线默认值，输出结果: ${var3:-xyz}"

var4=123

# 只要变量有赋值（空值或其它值） 就会生效赋值
echo "变量var4的值：${var4+xyz}"

# 只有变量赋值不为空，才会生效赋值
echo "变量var4的值：${var4:+abc}"


# 问号后面跟错误信息，如果var5没有定义，则报错推出，并提示后面定义的错误信息 
#echo "变量var5的值：${var5?"请输入var5的值"}"

var6=123

# 如果 var6 值不为空，则输出 var6 的值
# 如果 var6 值为空，则输出提示
echo "变量var6的值：${var6:?"请输入var6的值"}"