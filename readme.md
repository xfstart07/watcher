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

配置 Redis 服务地址和密码，`watch_path` 是监控的文件夹路径。

### 运行

```text
./watcher
```

## 项目

### 包版本管理

使用 glide 管理包版本。

### 日志

使用 zap 日志库