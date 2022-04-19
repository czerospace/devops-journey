# 7.1 下载Flannel二进制包

```shell
mkdir /soft/flannel ; cd /soft/flannel
wget https://github.com/coreos/flannel/releases/download/v0.11.0/flannel-v0.11.0-linux-amd64.tar.gz
tar xvf flannel-v0.11.0-linux-amd64.tar.gz
mv flanneld mk-docker-opts.sh /usr/local/bin/

#复制flanneld到其他的所有节点
for i in master-2 master-3 node-01 node-02;do rsync -av  /usr/local/bin/flanneld $i:/usr/local/bin/;done
for i in master-2 master-3 node-01 node-02;do rsync -av  /usr/local/bin/mk-docker-opts.sh $i:/usr/local/bin/;done
```

# 7.2 配置Flannel (所有节点)

```shell
mkdir -p /etc/flannel
cat > /etc/flannel/flannel.cfg<<EOF
FLANNEL_OPTIONS="-etcd-endpoints=https://10.0.4.16:2379,https://10.0.4.2:2379,https://10.0.4.9:2379 -etcd-cafile=/etc/etcd/ssl/ca.pem -etcd-certfile=/etc/etcd/ssl/server.pem  -etcd-keyfile=/etc/etcd/ssl/server-key.pem  --healthz-ip=0.0.0.0 --healthz-port=7100"
EOF
```

# 7.3 配置Flannel配置文件

```shell
cat > /usr/lib/systemd/system/flanneld.service <<EOF
[Unit]
Description=Flanneld overlay address etcd agent
After=network-online.target network.target
Before=docker.service

[Service]
Type=notify
EnvironmentFile=/etc/flannel/flannel.cfg
ExecStart=/usr/local/bin/flanneld --ip-masq \$FLANNEL_OPTIONS
ExecStartPost=/usr/local/bin/mk-docker-opts.sh -k DOCKER_NETWORK_OPTIONS -d /run/flannel/subnet.env
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF
```

# 7.4 启动Flannel

```shell
systemctl start flanneld
systemctl status flanneld
systemctl enable flanneld

#所有的节点都需要有172.17.0.0/16 网段IP
```

# 7.5 修改Docker启动文件（node节点）

```shell
cat >/usr/lib/systemd/system/docker.service<<EOFL
[Unit]
Description=Docker Application Container Engine
Documentation=https://docs.docker.com
After=network-online.target firewalld.service
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile=/run/flannel/subnet.env
ExecStart=/usr/bin/dockerd  \$DOCKER_NETWORK_OPTIONS
ExecReload=/bin/kill -s HUP \$MAINPID
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
TimeoutStartSec=0
Delegate=yes
KillMode=process
Restart=on-failure
StartLimitBurst=3
StartLimitInterval=60s

[Install]
WantedBy=multi-user.target
EOFL
```

# 7.6 重启Docker服务

```shell
systemctl daemon-reload
systemctl  restart flanneld
systemctl  restart docker
#检查IP地址, docker 与flanneld 是同一个网段
```

