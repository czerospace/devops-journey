# CentOS7安装MySQL-5.7.38

## 一、下载

```html
官网下载社区版本
https://www.mysql.com/downloads/
```

```shell
wget https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-5.7.38-1.el7.x86_64.rpm-bundle.tar
```

## 二、安装

```shell
#解压tar包
tar xf mysql-5.7.38-1.el7.x86_64.rpm-bundle.tar
#卸载mariadb
yum remove mariadb*
#yum安装rpm包，自动解决依赖问题
yum install *.rpm
```

## 三、启动mysql

```shell
#按需修改配置文件
/etc/my.cnf
#启动mysqld
systemctl start mysqld
```

## 四、登陆mysql

```shell
#获取初始密码
grep password /var/log/mysqld.log
#使用初始密码登陆
mysql -uroot -p
```

## 五、修改root密码

```shell
alter user 'root'@'localhost' identified by 'xxx';
```