# nginx-ingress部署流程

## 一、deploy the ingress controller

```shell
#优雅安装
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.0/deploy/static/provider/cloud/deploy.yaml
#src中deploy.yaml已下载好并改为阿里云镜像，可直接使用
```

## 二、使用hostport暴露nginx-ingress端口

```shell
1.给节点打标签
kubectl label node 192.168.51.223 app=ingress_nginx
2.修改yaml里面的网络模式和标签
hostNetwork: true
nodeSelector:
  app: ingress-nginx
```

## 三、创建ingress-nginx-controller资源

```shell
kubectl apply -f deploy.yaml
```

## 四、创建后端nginx资源

```shell
kubectl apply -f nginx-demo.yaml
```

## 五、创建Ingress规则

```shell
#如果不需要动态准入控制，直接删除
kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission
#创建ingress规则
kubectl apply -f ingress.yaml

kubectl get ingress

NAME              CLASS    HOSTS              ADDRESS      PORTS   AGE
example-ingress   <none>   www.czerospace.com             80      2m46s
```

## 六、访问

```shell
#修改/etc/hosts

10.0.0.240 www.czerospace.com

#访问
curl www.czerospace.com
```

