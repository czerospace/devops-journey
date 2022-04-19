# 10.1部署coredns

```shell
#在master节点
mkdir /root/dns && cd /root/dns
kubectl create clusterrolebinding system:anonymous --clusterrole=cluster-admin --user=system:anonymous
kubectl apply -f coredns-rbac.yaml 
kubectl apply -f coredns.yaml 
```

```shell
#测试（master）
kubectl run -it --rm --restart=Never --image=infoblox/dnstools:latest dnstools
dnstools# nslookup kubernetes
Server:		10.0.0.2
Address:	10.0.0.2#53

Name:	kubernetes.default.svc.cluster.local
Address: 10.0.0.1


dnstools# cat /etc/resolv.conf 
nameserver 10.0.0.2
search default.svc.cluster.local. svc.cluster.local. cluster.local.
options ndots:5
```

# 10.2部署dashboard

```shell
#master
kubectl apply -f recommended-2.0.yaml
#创建用户授权
kubectl create serviceaccount  dashboard-admin -n kube-system

kubectl create clusterrolebinding  \
dashboard-admin --clusterrole=cluster-admin --serviceaccount=kube-system:dashboard-admin

#获取Token
kubectl describe secrets -n kube-system $(kubectl -n kube-system get secret | awk '/dashboard-admin/{print $1}')

#使用token登陆dashboard页面查看
```

