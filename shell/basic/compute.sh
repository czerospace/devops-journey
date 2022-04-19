#!/bin/bash

# expr

# 加
a=`expr 3 + 2`
echo $a

b=`expr $a + 10`
echo "$a + 10 = $b"

# 乘  * 符号需要转义
c=`expr $a \* 10`
echo "$a * 10 = $c"

# 除 向下取整
d=`expr $c / 3`
echo "$c / 3 = $d"

# 取余
e=`expr $c % 3`
echo "$c % 3 = $e"

# let

let "a += 10"
echo "a + 10 =$a"

let "a++"
echo "a++后,a=$a"

let "b=$a+$b+$c"
echo "计算后b的结果:$b"

(( b++ ))
echo "b自增后的值:$b"

# 浮点计算用 bc 命令，使用 scale 控制精度