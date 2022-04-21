# centos7安装python3.8.3以及pip私有源

## 安装python3.8.3

centos7默认python版本为2.7.5
为AI环境独立编译安装一套python3.8.3

### step 1: 安装依赖包

yum install zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel readline-devel tk-devel gcc make libffi-devel

### step 2: 下载源码包

https://www.python.org/ftp/python/3.8.3/Python-3.8.3.tgz

### step 3: 解压安装

tar -zxvf Python-3.8.3.tgz  

cd Python-3.8.3

./configure prefix=/usr/local/Python3.8.3

make && make install

### step 4: 建立软连接

添加python3的软链接
ln -s /usr/local/Python3.8.3/bin/python3.8 /usr/bin/python3 

添加 pip3 的软链接
ln -s /usr/local/Python3.8.3/bin/pip3.8 /usr/bin/pip3