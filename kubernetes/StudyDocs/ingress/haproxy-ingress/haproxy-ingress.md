# haproxy-ingress版本v0.10

# 一、给node设置label

```shell
kubectl label node 10.0.4.7 role=ingress-controller
#查看
kubectl get node -l role=ingress-controller
```

# 二、创建haproxy-ingress-controller资源

```shell
kubectl apply -f haproxy-ingress.yaml

#查看pod
kubectl get pods -n ingress-controller 
NAME                    READY   STATUS    RESTARTS   AGE
haproxy-ingress-4h89z   1/1     Running   0          16s
```

# 三、创建一个后端 deployment 并暴露为 svc

```shell
kubectl create deployment nginx-http --image nginx:alpine
kubectl expose deployment nginx-http --port=80
```

# 四、创建Ingress资源

```shell
kubectl apply -f nginxdemo-ingress.yaml
#查看ingress资源
kubectl get ingress
NAME    CLASS    HOSTS            ADDRESS   PORTS   AGE
nginx   <none>   nginx.test.com             80      92s
```

# 五、查看

```shell
#将nginx.test.com域名解析到haproxy所在的node
10.0.4.7  nginx.test.com
#curl访问
curl nginx.test.com
```

# 六、haproxy.cfg

```nginx
backend default_nginx-http_80
    mode http
    balance roundrobin
    acl https-request ssl_fc
    http-request set-header X-Original-Forwarded-For %[hdr(x-forwarded-for)] if { hdr(x-forwarded-for) -m found }
    http-request del-header x-forwarded-for
    option forwardfor
    http-response set-header Strict-Transport-Security "max-age=15768000"
    server srv001 172.17.93.8:80 weight 1 check inter 2s
    server srv002 127.0.0.1:1023 disabled weight 1 check inter 2s
    server srv003 127.0.0.1:1023 disabled weight 1 check inter 2s
    server srv004 127.0.0.1:1023 disabled weight 1 check inter 2s
    server srv005 127.0.0.1:1023 disabled weight 1 check inter 2s
    server srv006 127.0.0.1:1023 disabled weight 1 check inter 2s
    server srv007 127.0.0.1:1023 disabled weight 1 check inter 2s
```

从配置上可以看出，haproxy-ingress 直接转发数据到pod 的ip+port了，没有经过svc