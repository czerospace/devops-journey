# What

```shell
将运行在一组 Pods 上的应用程序公开为网络服务的抽象方法。
使用 Kubernetes 服务无需修改应用程序即可使用通用的服务发现机制。 Kubernetes 为 Pods 提供自己的 IP 地址，并为一组 Pod 提供相同的 DNS 名， 并且可以在它们之间进行负载均衡。
```



# Why

```shell
每个 Pod 都有自己的 IP 地址，但是在 Deployment 中，在同一时刻运行的 Pod 集合可能与稍后运行该应用程序的 Pod 集合不同。比如某个 Pod 此时运行在 node1 上，下个时刻被调度到了 node2 上，Pod ip 地址发生了改变。
这导致了一个问题： 如果一组 Pod（称为“后端”）为群集内的其他 Pod（称为“前端”）提供功能， 那么前端如何找出并跟踪要连接的 IP 地址，以便前端可以使用后端部分？
```



# How

```shell
Service 在 Kubernetes 中是一个 REST 对象，和 Pod 类似。 像所有的 REST 对象一样，Service 定义可以基于 POST 方式，请求 API server 创建新的实例。
例如，假定有一组 Pod，它们对外暴露了 9376 端口，同时还被打上 app=MyApp 标签。
具体见src
```

