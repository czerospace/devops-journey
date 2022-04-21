# What

```html
ReplicaSet 的目的是维护一组在任何时候都处于运行状态的 Pod 副本的稳定集合。 因此，它通常用来保证给定数量的、完全相同的 Pod 的可用性。
```



# How

```html
虽然 ReplicaSets 可以独立使用，但今天它主要被Deployments 用作协调 Pod 创建、删除和更新的机制。 使用 Deployment 时会自动创建 ReplicaSet。Deployment 会拥有并管理它们的 ReplicaSet。
参考下面demo，不在此过多赘述
```

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      name: nginx
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx
          ports:
            - containerPort: 80
```

