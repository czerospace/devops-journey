# What

```shell
Ingress 是对集群中服务的外部访问进行管理的 API 对象，典型的访问方式是 HTTP。
Ingress 可以提供负载均衡、SSL 终结和基于名称的虚拟托管。

```

# How

```shell
Ingress 公开了从集群外部到集群内服务的 HTTP 和 HTTPS 路由。 流量路由由 Ingress 资源上定义的规则控制。
```

![image-20220406181458485](image\ingress.png)

```shell
#1.用户请求到ingress负载均衡器
#2.ingress根据请求的hosts/path匹配相对应的service
#3.ingress通过路由规则将数据发送到service
#4.service获取到pod的endpoint ip和port
#5.service通过iptbales/ipvs将数据转发给pod
```

