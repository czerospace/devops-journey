# 下载etcd二进制安装文件（所有master）

```shell
mkdir -p /soft/etcd && cd /soft/etcd
wget https://github.com/etcd-io/etcd/releases/download/v3.3.10/etcd-v3.3.10-linux-amd64.tar.gz
tar -xvf etcd-v3.3.10-linux-amd64.tar.gz
cd etcd-v3.3.10-linux-amd64/
cp etcd etcdctl /usr/local/bin/
#复制到所有master节点
for i in master-2 master-3;do rsync -av /usr/local/bin/etcd* $i:/usr/local/bin/;done
```

# 5.1 编辑etcd配置文件（所有master）

```shell
#注意修改每个节点的ETCD_NAME
#注意修改每个节点的监听地址
mkdir -p /etc/etcd/{cfg,ssl}

cat  >/etc/etcd/cfg/etcd.conf<<EOFL
#[Member]
ETCD_NAME="master-1"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://10.0.4.16:2380"
ETCD_LISTEN_CLIENT_URLS="https://10.0.4.16:2379,http://10.0.4.16:2390"

#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://10.0.4.16:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://10.0.4.16:2379"
ETCD_INITIAL_CLUSTER="master-1=https://10.0.4.16:2380,master-2=https://10.0.4.2:2380,master-3=https://10.0.4.9:2380"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_INITIAL_CLUSTER_STATE="new"
EOFL
```

# 5.2创建ETCD的系统启动服务（所有master）

```shell
cat > /usr/lib/systemd/system/etcd.service<<EOFL
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile=/etc/etcd/cfg/etcd.conf
ExecStart=/usr/local/bin/etcd \
--name=\${ETCD_NAME} \
--data-dir=\${ETCD_DATA_DIR} \
--listen-peer-urls=\${ETCD_LISTEN_PEER_URLS} \
--listen-client-urls=\${ETCD_LISTEN_CLIENT_URLS},http://127.0.0.1:2379 \
--advertise-client-urls=\${ETCD_ADVERTISE_CLIENT_URLS} \
--initial-advertise-peer-urls=\${ETCD_INITIAL_ADVERTISE_PEER_URLS} \
--initial-cluster=\${ETCD_INITIAL_CLUSTER} \
--initial-cluster-token=\${ETCD_INITIAL_CLUSTER_TOKEN} \
--initial-cluster-state=\${ETCD_INITIAL_CLUSTER_STATE} \
--cert-file=/etc/etcd/ssl/server.pem \
--key-file=/etc/etcd/ssl/server-key.pem \
--peer-cert-file=/etc/etcd/ssl/server.pem \
--peer-key-file=/etc/etcd/ssl/server-key.pem \
--trusted-ca-file=/etc/etcd/ssl/ca.pem \
--peer-trusted-ca-file=/etc/etcd/ssl/ca.pem
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOFL
```

# 5.3 复制etcd证书到指定目录

```shell
rsync -av  /root/etcd/*pem /etc/etcd/ssl/
#复制etcd证书到每个节点
for i in master-2 master-3 node-01 node-02;do ssh $i mkdir -p /etc/etcd/{cfg,ssl};done
for i in master-2 master-3 node-01 node-02;do rsync -av /etc/etcd/ssl/* $i:/etc/etcd/ssl/;done
for i in master-2 master-3 node-01 node-02;do echo $i "------>"; ssh $i ls /etc/etcd/ssl;done
```

# 5.4 启动etcd (所有节点)

```shell
systemctl start etcd
systemctl status etcd
systemctl enable etcd
```

# 5.5 检查etcd 集群是否运行正常

```shell
etcdctl --ca-file=/etc/etcd/ssl/ca.pem --cert-file=/etc/etcd/ssl/server.pem \
--key-file=/etc/etcd/ssl/server-key.pem --endpoints="https://10.0.4.16:2379"  cluster-health
```

# 5.6 创建Docker所需分配POD 网段 (任意master节点)

```shell
etcdctl --ca-file=/etc/etcd/ssl/ca.pem \
--cert-file=/etc/etcd/ssl/server.pem --key-file=/etc/etcd/ssl/server-key.pem \
--endpoints="https://10.0.4.16:2379,https://10.0.4.2:2379,https://10.0.4.9:2379" \
 set /coreos.com/network/config  \
 '{ "Network": "172.17.0.0/16", "Backend": {"Type": "vxlan"}}'
 
 #检查是否建立网段
 etcdctl \
--endpoints=https://10.0.4.16:2379,https://10.0.4.2:2379,https://10.0.4.9:2379 \
--ca-file=/etc/etcd/ssl/ca.pem \
--cert-file=/etc/etcd/ssl/server.pem \
--key-file=/etc/etcd/ssl/server-key.pem \
get /coreos.com/network/config
```