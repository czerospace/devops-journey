#!/bin/bash

# 不带参数的函数
xyz(){
    echo "hello world!"
}

# 如何向定义的函数传参呢？位置参数
xyz_with_params(){
    echo "函数传入参数个数为: $#"
    for i in `seq 1 $#`
    do
        echo "第$i个位置参数:$1"
        shift
        # shift 的意思是移走前面一个参数，将后面一个参数补到前面一个参数的位置
    done
}

xyz
xyz_with_params abc xyz 123

# 函数里的返回值
xyz_with_return(){
    echo "函数的返回值为100"
    return 100
}

xyz_with_return
# 使用 $? 输出上一次的执行结果,但是结果值在0-255之间，超过255会取余
echo "函数的返回值:$?"

# $? 返回 0 为 true，返回 非0 为 false

if xyz_with_return
then 
    echo "成功"
else
    echo "失败"
fi
