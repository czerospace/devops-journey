# Helm简介

- 类似Centos的yum

- Kubernetes的包管理器

- 一包一Chart目录

# Helm部署

**首先需要保证部署helm的节点必须可以正常执行kubectl**

### 1. Helm客户端安装

##### 下载

Helm是一个二进制文件，直接到github的release去下载就可以，地址如下： https://github.com/helm/helm/releases

##### 安装

```bash
# 解压
$ tar -zxvf helm-v2.16.12-linux-amd64.tar.gz
$ mv linux-amd64/helm /usr/local/bin/

# 没配置环境变量的需要先配置好
$ export PATH=$PATH:/usr/local/bin/

# 验证
$ helm version
```

### 2. Tiller安装

Tiller 是以 Deployment 方式部署在 Kubernetes 集群中的，由于 Helm 默认会去 storage.googleapis.com 拉取镜像，我们这里就默认无法科学上网的情况：

```bash
# 指向阿里云的仓库
$ helm init --client-only --stable-repo-url https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts/
$ helm repo add incubator https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts-incubator/
$ helm repo update

# 因为官方的镜像无法拉取，使用-i指定自己的镜像
$ helm init --service-account tiller --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.16.12  --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
 
# 创建TLS认证服务端
$ helm init --service-account tiller --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.16.12 --tiller-tls-cert /etc/kubernetes/ssl/server.pem --tiller-tls-key /etc/kubernetes/ssl/server-key.pem --tls-ca-cert /etc/kubernetes/ssl/ca.pem --tiller-namespace kube-system --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
```

### 3. 给Tiller授权

因为 Helm 的服务端 Tiller 是一个部署在 Kubernetes 中的 Deployment，它会去访问ApiServer去对集群进行操作。目前的 Tiller 部署时默认没有定义授权的 ServiceAccount，这会导致访问 API Server 时被拒绝。所以我们需要明确为 Tiller 部署添加授权。

```bash
# 创建serviceaccount
$ kubectl create serviceaccount --namespace kube-system tiller
# 创建角色绑定
$ kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
```

### 4. 验证

```bash
# 查看Tiller的serviceaccount，需要跟我们创建的名字一致：tiller
$ kubectl get deploy --namespace kube-system tiller-deploy -o yaml|grep serviceAccount

# 验证pods
$ kubectl -n kube-system get pods|grep tiller

# 验证版本
$ helm version
```