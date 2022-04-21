# WAHT

StatefulSet 是用来管理 *有状态* 应用的工作负载 API 对象。

StatefulSet 用来管理 Deployment 和扩展一组 Pod，并且能为这些 Pod 提供序号和唯一性保证。

和 Deployment 相同的是，StatefulSet 管理了基于相同容器定义的一组 Pod。但和 Deployment 不同的是，StatefulSet 为它们的每个 Pod 维护了一个固定的 ID。这些 Pod 是基于相同的声明来创建的，但是不能相互替换：无论怎么调度，每个 Pod 都有一个永久不变的 ID。

# WHY

StatefulSets 可以满足以下一个或多个需求的应用程序：

- 稳定的、唯一的网络标识符。
- 稳定的、持久的存储。
- 有序的、优雅的部署和缩放。
- 有序的、自动的滚动更新。
  稳定意味着 Pod 调度或重调度的整个过程是有持久性的。如果应用程序不需要任何稳定的标识符或有序的部署、删除或伸缩，则应该使用由一组无状态的副本控制器提供的工作负载来部署应用程序，比如 Deployment 或者 ReplicaSet 可能更适用于您的无状态应用部署需要。

# WHERE

headless使用场景：有时候我们创建的服务不想走负载均衡，想直接通过pod-ip链接后端，怎么办呢，使用headless service接可以解决。headless service 是将service的发布文件中的clusterip=none ，不让其获取clusterip ， DNS解析的时候直接走pod。

# HOW

见src