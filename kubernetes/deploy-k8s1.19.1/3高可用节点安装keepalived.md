# 这里简单使用 keepalived，还可用haproxy，lvs，nginx等lb

# 3.1高可用节点安装Keepalived
```shell
yum install -y keepalived
```



# 3.2master节点配置文件

```shell
cat >/etc/keepalived/keepalived.conf <<EOL
global_defs {
   router_id KUB_LVS
}
vrrp_script CheckMaster {
    script "curl -k https://ip:6443"
    interval 3
    timeout 9
    fall 2
    rise 2
}
vrrp_instance VI_1 {
    state MASTER
    interface ens32
    virtual_router_id 18
    priority 100
    advert_int 1
    nopreempt
    authentication {
        auth_type PASS
        auth_pass 111111
    }
    virtual_ipaddress {
        ip/24 dev ens32
    }
    track_script {
        CheckMaster
    }
}
EOL
```



# 3.3slave节点配置文件

```shell
cat >/etc/keepalived/keepalived.conf <<EOL
global_defs {
   router_id KUB_LVS
}
vrrp_script CheckMaster {
    script "curl -k https://ip:6443"
    interval 3
    timeout 9
    fall 2
    rise 2
}
vrrp_instance VI_1 {
    state SLAVE
    interface ens32
    virtual_router_id 18
    priority 90
    advert_int 1
    nopreempt
    authentication {
        auth_type PASS
        auth_pass 111111
    }
    virtual_ipaddress {
        ip/24 dev ens32
    }
    track_script {
        CheckMaster
    }
}
EOL
```



# 3.4启动keepalived
```shell
systemctl start keepalived && systemctl enable keepalived
```



# 3.5检查keepalived状态，停掉master验证vip是否切换到slave上

# 3.6其它

## nginx stream 代理配置

需要安装stream模块

```nginx
stream {
    log_format  main  '$remote_addr $upstream_addr - [$time_local] $status $upstream_bytes_sent';
    access_log  /var/log/nginx/access.log  main;
    upstream apiserver {
    	server apiserver-ip:6443;
    	server apiserver-ip:6443;
    	server apiserver-ip:6443;
	}
	server {
    	listen 0.0.0.0:6443;
	allow kubelet-ip;
	allow kubelet-ip;
	deny all;
    	proxy_connect_timeout 3s;
    	proxy_timeout 3s;
    	proxy_pass apiserver;
	}
}
```

