# 查看CPU个数

```shell
grep 'model name' /proc/cpuinfo | wc -l
```

# stress命令模拟一个 CPU 使用率 100% 场景

```shell
stress --cpu 1 --timeout 600
```

# mpstat 查看  CPU  使用率的变化情况

```shell
# -P ALL 表示监控所有CPU，后面数字5表示间隔5秒后输出一组数据
mpstat -P ALL 5
```

# watch 查看 系统负载变化

```shell
# -d 参数表示高亮显示变化的区域
watch -d uptime
```

# pidstat 查询CPU变化

```shell
# 间隔5秒后输出一组数据
pidstat -u 5 1
```

