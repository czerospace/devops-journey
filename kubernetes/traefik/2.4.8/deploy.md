# 部署步骤

```shell
kubectl create -f traefik-crd.2.4.8.yaml 
kubectl get CustomResourceDefinition
kubectl create -f traefik-rbac.2.4.8.yaml 
kubectl apply -f traefik-config.2.4.8.yaml -n kube-system

#获取节点名称
kubectl get nodes

#注意修改节点名称
kubectl label nodes node-01 IngressProxy=true

#部署
kubectl apply -f traefik-deploy.2.4.8.yaml -n kube-system

#获取pod
kubectl get pod -n kube-system 


#创建管理界面访问
kubectl apply -f traefik-dashboard-route.2.4.8.yaml 

#在没有dns解析的情况下通过修改PC hosts 访问管理界面
C:\Windows\System32\drivers\etc 

node-01-ip ingress.czerospace.com

# 浏览器输入 ingress.czerospace.com
```

