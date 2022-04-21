# 从集群外部访问 Service的方法

## NodePort

```shell
在集群内部IP的基础上，在集群的每一个节点的端口上开放这个服务。你可以在任意:NodePort地址上访问到这个服务。
```

## ExternalIP

```shell
给service绑定一个可以被客户端访问的ip
```

## LoadBalancer

```shell
在使用一个集群内部IP地址和在NodePort上开放一个Service的基础上，还可以向云提供者申请一个负载均衡器，将流量转发到已经以NodePort形式开发的Service上。
```

