# What

探针 是由 kubelet 对容器执行的定期诊断。 要执行诊断，kubelet 调用由容器实现的 Handler （处理程序）。有三种类型的处理程序：

ExecAction： 在容器内执行指定命令。如果命令退出时返回码为 0 则认为诊断成功。

TCPSocketAction： 对容器的 IP 地址上的指定端口执行 TCP 检查。如果端口打开，则诊断被认为是成功的。

HTTPGetAction： 对容器的 IP 地址上指定端口和路径执行 HTTP Get 请求。如果响应的状态码大于等于 200 且小于 400，则诊断被认为是成功的。（最常用）

每次探测都将获得以下三种结果之一：

Success（成功）：容器通过了诊断。
Failure（失败）：容器未通过诊断。
Unknown（未知）：诊断失败，因此不会采取任何行动。

# How

对于所包含的容器需要较长时间才能启动就绪的 Pod 而言，启动探针是有用的

如果你的容器启动时间通常超出 initialDelaySeconds + failureThreshold × periodSeconds 总值，你应该设置一个启动探测，对存活态探针所使用的同一端点执行检查。 periodSeconds 的默认值是 30 秒。