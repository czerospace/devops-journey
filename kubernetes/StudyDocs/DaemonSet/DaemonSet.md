# WHAT

```html
DaemonSet 确保全部（或者某些）节点上运行一个 Pod 的副本。 当有节点加入集群时， 也会为他们新增一个 Pod 。 当有节点从集群移除时，这些 Pod 也会被回收。删除 DaemonSet 将会删除它创建的所有 Pod。
```

# WHERE

DaemonSet 的一些典型用法：

- 在集群的每个节点上运行存储 Daemon，比如 glusterd 或 ceph。
- 在每个节点上运行日志收集 Daemon，比如 flunentd 或 logstash。
- 在每个节点上运行监控 Daemon，比如 Prometheus Node Exporter 或 collectd。

# HOW

