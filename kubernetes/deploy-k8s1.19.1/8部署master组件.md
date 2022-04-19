# Master端需要安装的组件如下:

kube-apiserver
kube-scheduler
kube-controller-manager

# 8.1 安装Api Server服务

## 8.1.1下载Kubernetes二进制包

```shell
cd /soft
wget https://dl.k8s.io/v1.19.1/kubernetes-server-linux-amd64.tar.gz
tar xvf kubernetes-server-linux-amd64.tar.gz 
cd kubernetes/server/bin/
cp kube-scheduler kube-apiserver kube-controller-manager kubectl /usr/local/bin/

#复制执行文件到其他的master节点
for i in master-2 master-3;do rsync -av  /usr/local/bin/kube* $i:/usr/local/bin/;done
```

## 8.1.2 配置Kubernetes证书

```shell
#Kubernetes各个组件之间通信需要证书,需要复制个每个master节点（master-1）
[root@master-1  soft]# mkdir -p /etc/kubernetes/{cfg,ssl}
[root@master-1  soft]# cp /root/kubernetes/*.pem /etc/kubernetes/ssl/

#复制到其他的节点
[root@master-1  soft]# for i in master-2 master-3 node-01 node-02;do ssh $i mkdir -p /etc/kubernetes/{cfg,ssl};done
[root@master-1  soft]# for i in master-2 master-3 node-01 node-02;do rsync -av /etc/kubernetes/ssl/* $i:/etc/kubernetes/ssl/;done
[root@master-1 bin]# for i in master-2 master-3 node-01 node-02;do echo $i "---------->"; ssh $i ls /etc/kubernetes/ssl;done
```

## 8.1.3 创建 TLS Bootstrapping Token

```shell
# TLS bootstrapping 功能就是让 kubelet 先使用一个预定的低权限用户连接到 apiserver，然后向 apiserver 申请证书，kubelet 的证书由 apiserver 动态签署
#Token可以是任意的包涵128 bit的字符串，可以使用安全的随机数发生器生成
head -c 16 /dev/urandom | od -An -t x | tr -d ' '
ae11d6f3bc311001e8b54a139554bcf0
```

## 8.1.4 编辑Token 文件(master-1)

```shell
#ae11d6f3bc311001e8b54a139554bcf0:随机字符串,自定义生成; kubelet-bootstrap:用户名; 10001:UID; system:kubelet-bootstrap：用户组
vim /etc/kubernetes/cfg/token.csv
ae11d6f3bc311001e8b54a139554bcf0,kubelet-bootstrap,10001,"system:kubelet-bootstrap"

#复制到其他的master节点
for i in master-2 master-3;do rsync -av  /etc/kubernetes/cfg/token.csv $i:/etc/kubernetes/cfg/token.csv;done
```

## 8.1.5创建Apiserver配置文件(所有的master节点)

```shell
#配置文件内容基本相同, 如果有多个节点, 那么需要修改IP地址即可
cat >/etc/kubernetes/cfg/kube-apiserver.cfg <<EOFL
KUBE_APISERVER_OPTS="--logtostderr=true \
--v=4 \
--insecure-bind-address=0.0.0.0 \
--insecure-port=8080 \
--etcd-servers=https://10.0.4.16:2379,https://10.0.4.2:2379,https://10.0.4.9:2379 \
--bind-address=0.0.0.0 \
--secure-port=6443 \
--advertise-address=0.0.0.0 \
--allow-privileged=true \
--service-cluster-ip-range=10.0.0.0/24 \
--enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,ResourceQuota,NodeRestriction \
--authorization-mode=RBAC,Node \
--enable-bootstrap-token-auth \
--token-auth-file=/etc/kubernetes/cfg/token.csv \
--service-node-port-range=30000-50000 \
--tls-cert-file=/etc/kubernetes/ssl/server.pem  \
--tls-private-key-file=/etc/kubernetes/ssl/server-key.pem \
--client-ca-file=/etc/kubernetes/ssl/ca.pem \
--service-account-key-file=/etc/kubernetes/ssl/ca-key.pem \
--etcd-cafile=/etc/etcd/ssl/ca.pem \
--etcd-certfile=/etc/etcd/ssl/server.pem \
--etcd-keyfile=/etc/etcd/ssl/server-key.pem"
EOFL
```

## 8.1.6 配置kube-apiserver 启动文件(所有的master节点)

```shell
cat >/usr/lib/systemd/system/kube-apiserver.service<<EOFL
[Unit]
Description=Kubernetes API Server
Documentation=https://github.com/kubernetes/kubernetes

[Service]
EnvironmentFile=/etc/kubernetes/cfg/kube-apiserver.cfg
ExecStart=/usr/local/bin/kube-apiserver \$KUBE_APISERVER_OPTS
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOFL
```

## 8.1.7 启动kube-apiserver服务

```shell
systemctl start kube-apiserver
systemctl status kube-apiserver
systemctl enable kube-apiserver

#查看加密的端口是否已经启动
netstat -anltup | grep 6443
```

# 8.2 部署kube-scheduler 服务

```shell
#创建kube-scheduler配置文件（所有的master节点）
cat >/etc/kubernetes/cfg/kube-scheduler.cfg<<EOFL
KUBE_SCHEDULER_OPTS="--logtostderr=true --v=4 --bind-address=0.0.0.0 --master=127.0.0.1:8080 --leader-elect"
EOFL
```

## 8.2.1 创建kube-scheduler 启动文件

```shell
#创建kube-scheduler systemd unit 文件（所有的master节点）
cat >/usr/lib/systemd/system/kube-scheduler.service<<EOFL
[Unit]
Description=Kubernetes Scheduler
Documentation=https://github.com/kubernetes/kubernetes

[Service]
EnvironmentFile=/etc/kubernetes/cfg/kube-scheduler.cfg
ExecStart=/usr/local/bin/kube-scheduler \$KUBE_SCHEDULER_OPTS
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOFL
```

## 8.2.2 启动kube-scheduler服务（所有的master节点）

```shell
systemc	start kube-scheduler
systemc	enable kube-scheduler
```

## 8.2.3查看Master节点组件状态（任意一台master）

```shell
kubectl get cs
```

# 8.3 部署kube-controller-manager

## 8.3.1创建kube-controller-manager配置文件(所有节点)

```shell
cat >/etc/kubernetes/cfg/kube-controller-manager.cfg<<EOFL
KUBE_CONTROLLER_MANAGER_OPTS="--logtostderr=true \
--v=4 \
--master=127.0.0.1:8080 \
--leader-elect=true \
--address=0.0.0.0 \
--service-cluster-ip-range=10.0.0.0/24 \
--cluster-name=kubernetes \
--cluster-signing-cert-file=/etc/kubernetes/ssl/ca.pem \
--cluster-signing-key-file=/etc/kubernetes/ssl/ca-key.pem  \
--root-ca-file=/etc/kubernetes/ssl/ca.pem \
--service-account-private-key-file=/etc/kubernetes/ssl/ca-key.pem"
EOFL
```

## 8.3.2 创建kube-controller-manager 启动文件

```shell
cat  >/usr/lib/systemd/system/kube-controller-manager.service<<EOFL
[Unit]
Description=Kubernetes Controller Manager
Documentation=https://github.com/kubernetes/kubernetes

[Service]
EnvironmentFile=/etc/kubernetes/cfg/kube-controller-manager.cfg
ExecStart=/usr/local/bin/kube-controller-manager \$KUBE_CONTROLLER_MANAGER_OPTS
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOFL
```

## 8.3.3启动kube-controller-manager服务

```shell
systemctl start kube-controller-manager
systemctl status kube-controller-manager
systemctl enable kube-controller-manager
```

# 8.4 查看Master 节点组件状态

```shell
#必须要在各个节点组件正常的情况下, 才去部署Node节点组件.（master节点）
kubectl get cs
```

