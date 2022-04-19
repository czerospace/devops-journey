# docker笔记

## 一、docker概念

师从慕课网 麦兜搞IT 全面的Docker 系统性入门+进阶实践（2021最新版）

课程笔记地址:https://www.docker.tips

## 二、安装docker

```
安装环境centos7
```

```
必要系统工具安装sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```

### 1、通过get-docker.com脚本自动化安装

#### 1）获取脚本
```
 curl -fsSL get.docker.com -o get-docker.sh
```
#### 2）执行脚本
```
 sh get-docker.sh
```
### 2、通过yum源安装

本人在深圳，所以选择腾讯源比较快一点

#### 1)获取docker.repo文件 

```
wget -O /etc/yum.repos.d/docker-ce.repo https://download.docker.com/linux/centos/docker-ce.repo
```

#### 2）把软件仓库地址替换为腾讯源:

```
sudo sed -i 's+download.docker.com+mirrors.cloud.tencent.com/docker-ce+' /etc/yum.repos.d/docker-ce.repo
```

#### 3）安装:

```
sudo yum makecache fast
sudo yum install -y docker-ce
```

## 三、Docker快速上手

### 1、启动docker

```
systemctl start docker
```

### 2、使用七牛云镜像加速

```
vim  /etc/docker/daemon.json 加入如下内容
```

```
{"registry-mirrors":["https://reg-mirror.qiniu.com/"]}
```

### 3、重启docker

```
$ sudo systemctl daemon-reload
$ sudo systemctl restart docker
```

### 4、docker命令行的基本使用

```
docker + 管理的对象（比如容器，镜像） + 具体操作（比如创建，启动，停止，删除）
```

例如

- `docker image pull nginx` 拉取一个叫nginx的docker image镜像
- `docker container stop nginx` 停止一个叫web的docker container容器

### 5、docker container run 背后发生了什么？

以运行一个nginx为例：

```
$ docker container run -d --publish 80:80 --name webhost nginx
1)在本地查找是否有nginx这个image镜像，如果没有发现

2)去远程的image registry查找nginx镜像（默认的registry是Docker Hub)

3)下载最新版本的nginx镜像 （nginx:latest 默认)

4)基于nginx镜像来创建一个新的容器，并且准备运行

5)docker engine分配给这个容器一个虚拟IP地址

6)在宿主机上打开80端口并把容器的80端口转发到宿主机上

7)启动容器，运行指定的命令（这里是一个shell脚本去启动nginx）
```

## 四、镜像管理

### 1、镜像的获取

#### 1)pull from registy(online)

#### 2)build from Dockerfile(onlie)

#### 3)load from file (offline)

### 2、镜像的registry

#### 1)公有

```
最常用:hub.docker.com
红帽的:quay.io
```

#### 2)私有

```
公司内部搭建一个私有的registry,常见harbor
```

#### 3)离线

image导出

```
docker image save nginx:1.20.0 -o nginx.image
```

image导入

```
docker image load -i nginx.image
```

### 3、镜像的构建

```
编写好Dockerfile之后，在Dockerfile所在的目录下执行
docker image build -t imageName .
```

```
镜像打tag
docker image tag imageName dockerhubID/imageName:版本号
```



## 五、Dockerfile

### 1.基础镜像

- 官方镜像优于非官方的镜像，如果没有官方镜像，则尽量选择Dockerfile开源的
- 固定版本tag而不是每次都使用latest
- 尽量选择体积小的镜像

### 2.RUN

```
主要用于在Image里执行指令，比如安装软件，下载文件等。
每一行的RUN命令都会产生一层image layer, 导致镜像的臃肿。
```

### 3.文件的复制和目录操作

``````
COPY 把local的一个文件复制到镜像里，如果目标目录不存在，则会自动创建
ADD  把local的一个文件复制到镜像里，如果目标目录不存在，则会自动创建,如果复制的是一个gzip等压缩文件时，ADD会帮助我们自动去解压缩文件。
``````

```
所有的文件复制均使用 COPY 指令，仅在需要自动解压缩的场合使用 ADD。
```

### 4.构建参数和环境变量

```
ARG 可以在镜像build的时候动态修改value, 通过 --build-arg
```

```
ENV 设置的变量可以在Image中保持，并在容器中的环境变量里
```

### 5.CMD

```
CMD可以用来设置容器启动时默认会执行的命令。
```

- 容器启动时默认执行的命令
- 如果docker container run启动容器时指定了其它命令，则CMD命令会被忽略
- 如果定义了多个CMD，只有最后一个会被执行。

### 6.ENTRYPOINT

### 7.Shell 格式和 Exec 格式

### 8.其它

## 六、Docker的存储

## 七、Docker的网络

## 八、Docker-compose

## 九、Docker-swarm





