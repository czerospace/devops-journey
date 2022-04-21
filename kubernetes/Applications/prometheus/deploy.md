# k8s中部署Prometheus

```shell
1.部署nfs
#先在服务器部署nfs
#修改nfs 配置文件(修改ip)
nfs-deployment.yaml
kubectl apply -f  .
#验证
kubectl get pvc

2.修改配置文件,将ip改为实际的ip
cd /root/monitor/serviceMonitor
[root@master-1 serviceMonitor]#  ls | xargs grep 91
prometheus-EtcdService.yaml:  - ip: 192.168.91.18
prometheus-EtcdService.yaml:  - ip: 192.168.91.19
prometheus-EtcdService.yaml:  - ip: 192.168.91.20
prometheus-kubeControllerManagerService.yaml:  - ip: 192.168.91.18
prometheus-kubeControllerManagerService.yaml:  - ip: 192.168.91.19
prometheus-kubeControllerManagerService.yaml:  - ip: 192.168.91.20
prometheus-KubeProxyService.yaml:  - ip: 192.168.91.21
prometheus-KubeProxyService.yaml:  - ip: 192.168.91.22
prometheus-kubeSchedulerService.yaml:  - ip: 192.168.91.18
prometheus-kubeSchedulerService.yaml:  - ip: 192.168.91.19
prometheus-kubeSchedulerService.yaml:  - ip: 192.168.91.20


3. install setup (crd)
cd /root/monitor/
kubectl apply -f setup/

4. alertmanager
cd /root/monitor/
kubectl apply -f alertmanager/

#check 
kubectl get pods,svc -n monitoring 

5. node-exporter
kubectl apply -f node-exporter/

#node 检查9100
netstat -anltup | grep 9100


6.kube-state-metrics
kubectl apply -f kube-state-metrics/

7.granfa
kubectl apply -f grafana/

8.prometheus
kubectl apply -f prometheus/

9. servicemonitor
kubectl apply -f serviceMonitor/


10.结果
kubectl get pod,svc -n monitoring
#granfa dashaboard
service/grafana-IP:port/login(默认账号密码admin/admin)
#prometheus 
service/prometheus-k8s-IP:port
#查看prometheus 日志
kubectl logs -f prometheus-k8s-0  -n monitoring -c prometheus
```

