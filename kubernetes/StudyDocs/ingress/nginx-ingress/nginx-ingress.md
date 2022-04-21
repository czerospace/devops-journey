# nginx-ingress部署流程

## 一、创建deployment

```shell
kubectl apply -f nginx-deployment.yaml
```

## 二、创建svc将deployment暴露出来

```shell
kubectl expose deployment nginx-deployment-demo --type=NodePort --port=80
```

## 三、创建ingress-nginx-controller资源

```shell
kubectl apply -f ingress-nginx-controller.yaml
```

## 四、创建Ingress资源

```shell
kubectl apply -f ingress.yaml
```

## 五、查看Ingress资源

```shell
kubectl get ingress

NAME              CLASS    HOSTS              ADDRESS      PORTS   AGE
example-ingress   <none>   hello-world.info   10.0.0.240   80      115s
```

## 六、访问

```shell
#修改/etc/hosts

10.0.0.240 hello-world.info

#访问
curl hello-world.info
```

## 七、登录nginx-ingress-controller pod查看nginx配置已经被nginx controller更新

```shell
#查看pod
kubectl get pod -n ingress-nginx

NAME                                        READY   STATUS    RESTARTS   AGE
nginx-ingress-controller-84865c44d9-44mcz   1/1     Running   0          8m35s
#登陆pod
kubectl exec -it nginx-ingress-controller-84865c44d9-44mcz /bin/bash  -n ingress-nginx
```

```nginx
## start server hello-world.info
	server {
		server_name hello-world.info ;
		
		listen 80  ;
		listen [::]:80  ;
		listen 443  ssl http2 ;
		listen [::]:443  ssl http2 ;
		
		set $proxy_upstream_name "-";
		
		ssl_certificate_by_lua_block {
			certificate.call()
		}
		
		location / {
			
			set $namespace      "default";
			set $ingress_name   "example-ingress";
			set $service_name   "nginx-deployment-demo";
			set $service_port   "80";
			set $location_path  "/";
			
			rewrite_by_lua_block {
				lua_ingress.rewrite({
					force_ssl_redirect = false,
					ssl_redirect = true,
					force_no_ssl_redirect = false,
					use_port_in_redirects = false,
				})
				balancer.rewrite()
				plugins.run()
			}
			
			header_filter_by_lua_block {
				
				plugins.run()
			}
			body_filter_by_lua_block {
				
			}
			
			log_by_lua_block {
				
				balancer.log()
				
				monitor.call()
				
				plugins.run()
			}
			
			port_in_redirect off;
			
			set $balancer_ewma_score -1;
			set $proxy_upstream_name "default-nginx-deployment-demo-80";
			set $proxy_host          $proxy_upstream_name;
			set $pass_access_scheme  $scheme;
			set $pass_server_port    $server_port;
			set $best_http_host      $http_host;
			set $pass_port           $pass_server_port;
			
			set $proxy_alternative_upstream_name "";
			
			client_max_body_size                    1m;
			
			proxy_set_header Host                   $best_http_host;
			
			# Pass the extracted client certificate to the backend
			
			# Allow websocket connections
			proxy_set_header                        Upgrade           $http_upgrade;
			
			proxy_set_header                        Connection        $connection_upgrade;
			
			proxy_set_header X-Request-ID           $req_id;
			proxy_set_header X-Real-IP              $remote_addr;
			
			proxy_set_header X-Forwarded-For        $remote_addr;
			
			proxy_set_header X-Forwarded-Host       $best_http_host;
			proxy_set_header X-Forwarded-Port       $pass_port;
			proxy_set_header X-Forwarded-Proto      $pass_access_scheme;
			
			proxy_set_header X-Scheme               $pass_access_scheme;
			
			# Pass the original X-Forwarded-For
			proxy_set_header X-Original-Forwarded-For $http_x_forwarded_for;
			
			# mitigate HTTPoxy Vulnerability
			# https://www.nginx.com/blog/mitigating-the-httpoxy-vulnerability-with-nginx/
			proxy_set_header Proxy                  "";
			
			# Custom headers to proxied server
			
			proxy_connect_timeout                   5s;
			proxy_send_timeout                      60s;
			proxy_read_timeout                      60s;
			
			proxy_buffering                         off;
			proxy_buffer_size                       4k;
			proxy_buffers                           4 4k;
			
			proxy_max_temp_file_size                1024m;
			
			proxy_request_buffering                 on;
			proxy_http_version                      1.1;
			
			proxy_cookie_domain                     off;
			proxy_cookie_path                       off;
			
			# In case of errors try the next upstream server before returning an error
			proxy_next_upstream                     error timeout;
			proxy_next_upstream_timeout             0;
			proxy_next_upstream_tries               3;
			
			proxy_pass http://upstream_balancer;
			
			proxy_redirect                          off;
			
		}
		
	}
	## end server hello-world.info
```

