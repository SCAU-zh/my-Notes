## 如何在本地进行远程调试 k8s 集群中的某个 JAVA 服务
> 当线上服务发生异常时，往往首先我们会先通过看线上的日志来定位问题，但是也可能存在日志无法定位的问题，我们可以通过 remote debug 来尝试定位，本文记录如何在本地通过 idea 工具 debug 线上 k8s 集群中的某个 JAVA 服务。

### 服务 Deployment yaml 文件配置
```yaml
spec:
  replicas: 1
  selector:
    matchLabels:
      app: 
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: 
    spec:
      containers:
        - name: 
          image:
          ports:
            - containerPort: 8080
              protocol: TCP
            # 对外暴露 5005 端口 用于调试 
            - containerPort: 5005
              protocol: TCP
          env:
            # 开启 debug 模式
            - name: ENABLE_DEBUG
              value: 'true'
```

### k8s 转发端口
```shell
# 转发服务 service 端口 5005 -> 本地 5005:远程 5005 
kubectl -n default port-forward service/xxx-xxx 5005:5005

输出:
Forwarding from 127.0.0.1:5005 -> 5005
Forwarding from [::1]:5005 -> 5005
代表成功转发
```

### idea 进行远程调试
idea Run/Debug Configurations 中添加 Remote JVM Debug,设置目标地址和端口 127.0.0.1:5005（上一步已经完成端口转发）
接着进行允许该 debug，即可实现 remote debug
