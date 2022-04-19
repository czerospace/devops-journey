# CentOS7部署私有镜像仓库Harbor
**因为harbor采用docker-compose方式部署在docker中，所以用harbor建立私有镜像仓库需要先部署docker、docker-compose**

## 一、部署docker
### step 1: 安装必要的一些系统工具
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
### Step 2: 添加软件源信息
sudo yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
### Step 3：修改源为阿里云
sudo sed -i 's+download.docker.com+mirrors.aliyun.com/docker-ce+' /etc/yum.repos.d/docker-ce.repo

sudo yum makecache fast

### Step 4: 查找docekr版本
yum list docker-ce.x86_64 --showduplicates | sort -r|grep 20.10.2

### Step 5: 安装Docker version 20.10.2

sudo yum install docker-ce-20.10.2
### Step 6: 开启Docker服务
sudo systemctl  start docker

### Step 7: 检查Docker服务
docker info|grep "Server Version"
可以看到版本为Server Version: 20.10.2
说明docker安装成功

### Step 8: 配置国内加速镜像以及内网仓库地址
vim /etc/docker/daemon.json
写入如下内容
```
{
    "registry-mirrors": ["https://registry.docker-cn.com"],
    "insecure-registries":["http://image.registry.ai.ipanel.cn:5000"]
}
```

registry-mirrors指定镜像加速地址，这里采用docker中国官方地址，常用的还有阿里腾讯华为清华源等
insecure-registries为私有镜像仓库地址,需要在内网dns或者服务器hosts指向部署harbor的服务器

### Step 9：重新加载配置并重启docker
sudo systemctl  daemon-reload
sudo systemctl  restart docker

## 二、部署docker-compose
### step 1: 配置epel源
yum install epel-release

### step 2: 安装docker-compose
yum install docker-compose

## 二、部署Harbor
### step 1: 下载软件
https://github.com/goharbor/harbor
选择合适的版本，本文选择v2.4.1，下载安装包harbor-offline-installer-v2.4.1.tgz

### step 2: 修改配置
#### a.解压缩
cd /r2/soft/
tar xf harbor-offline-installer-v2.4.1.tgz

#### b.修改配置文件
cd harbor
cp harbor.yml.tmpl harbor.yml
vim harbor.yml 修改配置文件

```
   hostname: reg.mydomain.com 
```

将hostname: reg.mydomain.com修改为服务器ip

```
# https related config
https:
  # https port for harbor, default is 443
  port: 443
  # The path of cert and key files for nginx
  certificate: /your/certificate/path
  private_key: /your/private/key/path
```

生产环境推荐用https协议加域名，如无法购买证书就把这段注释掉，直接用ip+http算了

```shell
harbor_admin_password: Harbor12345
```

修改harbor管理员密码

```
# Harbor DB configuration
database:
  # The password for the root user of Harbor DB. Change this before any production use.
  password: root123
  # The maximum number of connections in the idle connection pool. If it <=0, no idle connections are retained.
  max_idle_conns: 100
  # The maximum number of open connections to the database. If it <= 0, then there is no limit on the number of open connections.
  # Note: the default number of connections is 1024 for postgres of harbor.
  max_open_conns: 900
```

修改数据库密码

```
# The default data volume
data_volume: /data
```

修改harbor的存储目录，建议改为一个较大的目录，比如homed服务器常用的/r2/harbordata

### step 3: 安装harbor
./install.sh

提示
***✔ ----Harbor has been installed and started successfully.----***
则安装成功

### step 4: 浏览器登录镜像仓库
image.registry.ai.ipanel.cn