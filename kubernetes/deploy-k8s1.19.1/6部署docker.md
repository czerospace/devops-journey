# 在所有的Node节点安装

```shell
#安装CE版本
yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

#把软件仓库地址替换为腾讯云
sudo sed -i 's+download.docker.com+mirrors.cloud.tencent.com/docker-ce+' /etc/yum.repos.d/docker-ce.repo

#安装19.03.6版本
yum install -y docker-ce-19.03.6 docker-ce-cli-19.03.6 containerd.io

#启动Docker服务
systemctl start docker
systemctl enable docker

#配置镜像加速器(所有node节点)

tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://mirror.ccs.tencentyun.com"]
}
EOF

#重启docker
systemctl daemon-reload
systemctl restart docker
```

