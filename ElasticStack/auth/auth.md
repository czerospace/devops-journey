# Elasticsearch开启安全认证

## 一、生成ca

```shell
bin/elasticsearch-certutil ca
```

## 二、生成证书

```shell
bin/elasticsearch-certutil cert --ca elastic-stack-ca.p12
```

## 三、分发证书

```shell
#将elastic-certificates.p12分发到节点的config目录下
cp elastic-certificates.p12 config/
```

## 四、修改配置

```shell
http.cors.enabled: true
http.cors.allow-origin: "*"
http.cors.allow-headers: Authorization
xpack.security.enabled: true
xpack.security.transport.ssl.enabled: true
xpack.security.transport.ssl.verification_mode: certificate
xpack.security.transport.ssl.keystore.path: elastic-certificates.p12
xpack.security.transport.ssl.truststore.path: elastic-certificates.p12
```

## 五、启动es

## 六、设置密码

```shell
bin/elasticsearch-setup-passwords interactive
```

