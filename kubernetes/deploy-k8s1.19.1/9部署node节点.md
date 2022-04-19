# 部署Node节点组件

# 9.1部署 kubelet 组件

## 9.1.1 从Master节点复制Kubernetes 文件到Node

```shell
cd /soft
for i in node-01 node-02;do rsync -av kubernetes/server/bin/kubelet kubernetes/server/bin/kube-proxy $i:/usr/local/bin/;done
```

## 9.1.2 创建kubelet bootstrap.kubeconfig 文件

```shell
#Maste-1节点
mkdir /root/config ; cd /root/config
cat >environment.sh<<EOFL
# 创建kubelet bootstrapping kubeconfig
BOOTSTRAP_TOKEN=ae11d6f3bc311001e8b54a139554bcf0
KUBE_APISERVER="https://106.55.94.54:6443"
# 设置集群参数
kubectl config set-cluster kubernetes \
  --certificate-authority=/etc/kubernetes/ssl/ca.pem \
  --embed-certs=true \
  --server=\${KUBE_APISERVER} \
  --kubeconfig=bootstrap.kubeconfig
# 设置客户端认证参数
kubectl config set-credentials kubelet-bootstrap \
  --token=\${BOOTSTRAP_TOKEN} \
  --kubeconfig=bootstrap.kubeconfig
# 设置上下文参数
kubectl config set-context default \
  --cluster=kubernetes \
  --user=kubelet-bootstrap \
  --kubeconfig=bootstrap.kubeconfig
# 设置默认上下文
kubectl config use-context default --kubeconfig=bootstrap.kubeconfig
#通过 bash environment.sh获取 bootstrap.kubeconfig 配置文件。
EOFL

#执行脚本
sh environment.sh
```

## 9.1.3创建kube-proxy kubeconfig文件 （master-1）

```shell
cat  >env_proxy.sh<<EOF
# 创建kube-proxy kubeconfig文件
BOOTSTRAP_TOKEN=f89a76f197526a0d4bc2bf9c86e871c3
KUBE_APISERVER="https://106.55.94.54:6443"

kubectl config set-cluster kubernetes \
  --certificate-authority=/etc/kubernetes/ssl/ca.pem \
  --embed-certs=true \
  --server=\${KUBE_APISERVER} \
  --kubeconfig=kube-proxy.kubeconfig

kubectl config set-credentials kube-proxy \
  --client-certificate=/etc/kubernetes/ssl/kube-proxy.pem \
  --client-key=/etc/kubernetes/ssl/kube-proxy-key.pem \
  --embed-certs=true \
  --kubeconfig=kube-proxy.kubeconfig

kubectl config set-context default \
  --cluster=kubernetes \
  --user=kube-proxy \
  --kubeconfig=kube-proxy.kubeconfig

kubectl config use-context default --kubeconfig=kube-proxy.kubeconfig
EOF

#执行脚本
[root@master-1 bin]# sh env_proxy.sh
```

## 9.1.4 复制kubeconfig文件与证书到所有Node节点

```shell
#将bootstrap kubeconfig kube-proxy.kubeconfig 文件复制到所有Node节点
#远程创建目录 (master-1)
for i in node-01 node-02;do ssh $i "mkdir -p /etc/kubernetes/{cfg,ssl}";done

#复制证书文件ssl  (master-1)
for i in node-01 node-02;do rsync -av /etc/kubernetes/ssl/* $i:/etc/kubernetes/ssl/;done

#复制kubeconfig文件  (master-1)
cd /root/config
for i in node-01 node-02;do rsync -av bootstrap.kubeconfig kube-proxy.kubeconfig $i:/etc/kubernetes/cfg/;done
```

## 9.1.5 创建kubelet参数配置文件

```shell
#不同的Node节点, 需要修改IP地址 （node节点操作）
cat >/etc/kubernetes/cfg/kubelet.config<<EOF
kind: KubeletConfiguration
apiVersion: kubelet.config.k8s.io/v1beta1
address: 10.0.4.10
port: 10250
readOnlyPort: 10255
cgroupDriver: cgroupfs
clusterDNS: ["10.0.0.2"]
clusterDomain: cluster.local.
failSwapOn: false
authentication:
  anonymous:
    enabled: true
EOF
```

## 9.1.6 创建kubelet配置文件

```shell
#不同的Node节点, 需要修改IP地址
#/etc/kubernetes/cfg/kubelet.kubeconfig 文件自动生成
[root@node-1 bin]# cat >/etc/kubernetes/cfg/kubelet<<EOF
KUBELET_OPTS="--logtostderr=true \
--v=4 \
--hostname-override=10.0.4.10 \
--kubeconfig=/etc/kubernetes/cfg/kubelet.kubeconfig \
--bootstrap-kubeconfig=/etc/kubernetes/cfg/bootstrap.kubeconfig \
--config=/etc/kubernetes/cfg/kubelet.config \
--cert-dir=/etc/kubernetes/ssl \
--pod-infra-container-image=docker.io/kubernetes/pause:latest"
EOF
```

## 9.1.7 创建kubelet系统启动文件(node节点)

```shell
cat >/usr/lib/systemd/system/kubelet.service<<EOF
[Unit]
Description=Kubernetes Kubelet
After=docker.service
Requires=docker.service

[Service]
EnvironmentFile=/etc/kubernetes/cfg/kubelet
ExecStart=/usr/local/bin/kubelet \$KUBELET_OPTS
Restart=on-failure
KillMode=process

[Install]
WantedBy=multi-user.target
EOF
```

## 9.1.8 将kubelet-bootstrap用户绑定到系统集群角色(master节点)

```shell
kubectl create clusterrolebinding kubelet-bootstrap \
  --clusterrole=system:node-bootstrapper \
  --user=kubelet-bootstrap
```

## 9.1.9 启动kubelet服务（node节点）

```shell
systemctl start kubelet
systemctl status kubelet
systemctl enable kubelet
```

# 9.2 服务端批准与查看CSR请求

```shell
#查看CSR请求
#Maste-1节点操作
kubectl get csr
结果如下:
NAME                                                   AGE     SIGNERNAME                                    REQUESTOR           CONDITION
node-csr-J0rASsXgYvRRZs5QhmDfv56aIbgi7vs-wUgUd9NqXHY   6m45s   kubernetes.io/kube-apiserver-client-kubelet   kubelet-bootstrap   Pending
node-csr-MV_bTXY_UBDrM9btoUXvTuOrhqWmW1wmHnPIlRXMNTU   6m45s   kubernetes.io/kube-apiserver-client-kubelet   kubelet-bootstrap   Pending
```

## 9.2.1 批准请求

```shell
kubectl certificate approve node-csr-J0rASsXgYvRRZs5QhmDfv56aIbgi7vs-wUgUd9NqXHY
kubectl certificate approve node-csr-MV_bTXY_UBDrM9btoUXvTuOrhqWmW1wmHnPIlRXMNTU
```

# 9.3 节点重名处理

```shell
#如果出现节点重名, 可以先删除证书, 然后重新申请
#Master节点删除csr
[root@master-1 bin]# kubectl delete csr node-csr-U4v31mc3j_xPq5n1rU2KdpyugqfFH_0g1wOC66oiu04

#Node节点删除kubelet.kubeconfig
#客户端重启kubelet服务, 再重新申请证书
[root@node-1 bin]# rm -rf /etc/kubernetes/cfg/kubelet.kubeconfig
```

# 9.4 部署kube-proxy 组件

## 查看节点状态 

```shell
#所有的Node节点状态必须为Ready （master）
kubectl get nodes
```

## 9.4.1 创建kube-proxy配置文件

```shell
#注意修改hostname-override地址, 不同的节点则不同。
cat >/etc/kubernetes/cfg/kube-proxy<<EOF
KUBE_PROXY_OPTS="--logtostderr=true \
--v=4 \
--metrics-bind-address=0.0.0.0 \
--hostname-override=10.0.4.10 \
--cluster-cidr=10.0.0.0/24 \
--kubeconfig=/etc/kubernetes/cfg/kube-proxy.kubeconfig"
EOF
```

## 9.4.2 创建kube-proxy systemd unit 文件

```shell
cat >/usr/lib/systemd/system/kube-proxy.service<<EOF
[Unit]
Description=Kubernetes Proxy
After=network.target

[Service]
EnvironmentFile=/etc/kubernetes/cfg/kube-proxy
ExecStart=/usr/local/bin/kube-proxy \$KUBE_PROXY_OPTS
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF
```

## 9.4.3 启动kube-proxy 服务

```shell
systemctl start kube-proxy
systemctl status kube-proxy
systemctl enable kube-proxy
```

