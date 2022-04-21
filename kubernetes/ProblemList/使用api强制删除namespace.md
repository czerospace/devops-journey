# 一、查看已存在的namespace

![](images\ns.png)

```shell
如图，monitoring一直处于Terminating状态，kubectl delete ns monitoring无法删除，这里使用api强制删除的方式
```

# 二、获取需要强制删除的ns信息

```shell
kubectl get namespace monitoring -o json > monitoring.json
cat monitoring.json
```

![](images\json.png)

# 三、修改json文件信息

```shell
删除上图中的三行信息
```

# 四、运行kube proxy

```shell
新开一个终端，运行kube proxy
[root@master-1 monitor]# kubectl proxy
Starting to serve on 127.0.0.1:8001
```

# 五、通过API强制删除ns

```shell
curl -k -H "Content-Type: application/json" -X PUT --data-binary @monitoring.json http://127.0.0.1:8001/api/v1/namespaces/monitoring/finalize
```

# 六、验证是否已删除

```shell
kubectl get ns
```

