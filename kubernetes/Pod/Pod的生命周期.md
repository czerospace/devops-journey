# Pod 的生命周期

## 容器阶段 Phase

### Pending（挂起） 

Pod 已被 Kubernetes 系统接受，但有一个或者多个容器尚未创建亦未运行。此阶段包括等待 Pod 被调度的时间和通过网络下载镜像的时间。

### Running（运行中） 

Pod 已经绑定到了某个节点，Pod 中所有的容器都已被创建。至少有一个容器仍在运行，或者正处于启动或重启状态。

### Succeeded（成功） 

Pod 中的所有容器都已成功终止，并且不会再重启。

### Failed（失败） 

Pod 中的所有容器都已终止，并且至少有一个容器是因为失败终止。也就是说，容器以非 0 状态退出或者被系统终止。

### Unknown（未知） 

因为某些原因无法取得 Pod 的状态。这种情况通常是因为与 Pod 所在主机通信失败。

## 容器状态 Status

一旦调度器将 Pod 分派给某个节点，kubelet 就通过 容器运行时 开始为 Pod 创建容器。 容器的状态有三种：Waiting（等待）、Running（运行中）和 Terminated（已终止）。
kubectl describe pod <pod 名称>

### Waiting （等待）

如果容器并不处在 Running 或 Terminated 状态之一，它就处在 Waiting 状态。 处于 Waiting 状态的容器仍在运行它完成启动所需要的操作：例如，从某个容器镜像 仓库拉取容器镜像，或者向容器应用 Secret 数据等等。 当你使用 kubectl 来查询包含 Waiting 状态的容器的 Pod 时，你也会看到一个 Reason 字段，其中给出了容器处于等待状态的原因。

### Running（运行中）

Running 状态表明容器正在执行状态并且没有问题发生。 如果配置了 postStart 回调，那么该回调已经执行完成。 如果你使用 kubectl 来查询包含 Running 状态的容器的 Pod 时，你也会看到 关于容器进入 Running 状态的信息。

### Terminated（已终止）

处于 Terminated 状态的容器已经开始执行并且或者正常结束或者因为某些原因失败。 如果你使用 kubectl 来查询包含 Terminated 状态的容器的 Pod 时，你会看到 容器进入此状态的原因、退出代码以及容器执行期间的起止时间。