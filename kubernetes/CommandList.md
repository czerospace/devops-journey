# 获取 dashboard token
```shell
kubectl describe secrets -n kube-system $(kubectl -n kube-system get secret | awk '/dashboard-admin/{print $1}')
```



# 污点
```shell
给节点增加一个污点
kubectl taint nodes 192.168.51.221  key=value:NoSchedule
查看污点
kubectl describe node 192.168.51.221|grep Taints
删除污点
kubectl taint nodes 192.168.51.221  key=value:NoSchedule-
```

# 标签
```shell
给节点增加一个标签
kubectl   label nodes 192.168.51.222 app=ingress-nginx
查看标签
kubectl get node --show-labels
删除标签
kubectl   label nodes 192.168.51.222 app-
```


# 按 label 查找 Pod
```shell
kubectl get node --show-labels
按标签过滤
kubectl get pod -l run=my-nginx
显示为yaml格式
kubectl get pod -l run=my-nginx -o yaml
```

# 强制删除处于 Terminating 状态的 pod
```shell
kubectl delete pods <pod> --grace-period=0 --force
```