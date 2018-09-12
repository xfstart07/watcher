# 监控文件计算文件MD5值

将文件夹中的文件的MD5计算出来，并存储Redis。

### 部署运行

 #### 构建
 
```
 go build
```

### 配置

```ini
redis_uri=127.0.0.1:6379
redis_pass=
watch_path=/Users/x/golang/src/adserver-cloud, /Users/x/golang/src/adserver-cloud/app
``` 

配置 Redis 路径和密码，监控的文件夹数组

### 运行

```text
./watcher
```

