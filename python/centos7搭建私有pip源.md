# 安装pip私有源
devpi工具
## step 1: 安装supervisor
yum install supervisor

## step 2: 安装devpi
pip3 install devpi-server devpi-web devpi-client

## step 3: 配置devpi
创建存放库的目录
mkdir -p /r2/devpidata

第一次启动初始化
devpi-init --serverdir=/r2/devpidata --no-root-pypi

创建一些配置文件
devpi-gen-config --serverdir=/r2/devpidata --host=0.0.0.0 --port=3141  --request-timeout=60 --proxy-timeout=60

当前目录下会生成一个gen-config目录

mkdir /etc/devpi
mv gen-config /etc/devpi/

## step 4: 启动devpi
supervisorctl -c /etc/devpi/gen-config/supervisord.conf shutdown
supervisord -c /etc/devpi/gen-config/supervisord.conf

## step 5: 创建用户

连接server，并创建一个用户，登录
devpi use http://0.0.0.0:3141 -l

创建用户 pypi
devpi user -c pypi password=123
#登录
devpi login pypi --password=123

## step 5: 创建索引
创建索引 simple，并设置镜像
可以配置国内源
注意，有一些源在一些机器上不能访问，可以换别的国内源尝试，比如豆瓣，清华源等等
devpi index -c simple type=mirror mirror_url=http://mirrors.aliyun.com/pypi/simple mirror_web_url_fmt=http://mirrors.aliyun.com/pypi/simple/{name}/

## 服务端安装按完毕
去客户端测试
创建pip配置
mkdir ~/.pip/

vim ~/.pip/pip.conf

```shell
[global]
index-url = http://pypi.ai.ipanel.cn:3141/pypi/simple
timeout = 6000
trusted-host=pypi.ai.ipanel.cn
```

配置dns或者客户端服务器hosts将pypi.ai.ipanel.cn指向 devpi所在的服务器即可
pip3 install  测试