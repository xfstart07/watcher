# 监控多个文件夹，计算文件的MD5值

将文件夹中的文件MD5值计算并存储入Redis。

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
