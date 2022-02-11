# pod 对象解析
> pod 是 kubernetes 中的最小编排单位。
> 凡是网络、存储、安全、调度，基本是 pod 级别的。

## pod 的几个重要字段

### NodeSelector：供用户将 Pod 和 Node 进行绑定的字段。
```yaml
apiVersion: v1
kind: Pod
...
spec:
 nodeSelector:
   disktype: ssd
```
该 yaml 配置的 Pod 永远只能运行在携带了 「disktype: ssd」标签(label)的节点上，否则会调度失败。

### NodeName 
一般由调度器负责设置，表示该 Pod 已经经过调度，结果为该标签的值。

（也可以用于用户设置，骗过调度器，一般在测试或者调试时使用）

### HostAliases：定义了 Pod 的 hosts 文件(比如/etc/hosts) 里的内容

```yaml
apiVersion: v1
kind: Pod
...
spec:
  hostAliases:
  - ip: "10.1.2.3"
    hostnames:
    - "foo.remote"
    - "bar.remote"
...
```
Pod 启动后 ，/etc/hosts 文件的内容将如下所示

```shell
$ cat /etc/hosts
# Kubernetes-managed hosts file.
127.0.0.1 localhost
...
10.244.135.10 hostaliases-pod
10.1.2.3 foo.remote
10.1.2.3 bar.remote
```
注意：需要通过这种方法来设置 Pod 的 hosts， 直接进入容器修改，在 Pod 被删除重建后会丢失。

### imagePullPolicy 字段，定义镜像拉去的策略
- Always 总是拉取镜像
- IfNotPresent 本地有则使用本地镜像,不拉取
- Never 只使用本地镜像，从不拉取，即使本地没有
- 如果省略imagePullPolicy， 策略为always

### LifeCycle， 是在容器状态发生变化时触发一系列“钩子”
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: lifecycle-demo
spec:
  containers:
    - name: lifecycle-demo-container
      image: nginx
      lifecycle:
        # 在容器启动后，立刻执行一个指定操作
        postStart:
          exec:
            command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
        
        # 容器在被杀死之前，会阻塞当前的容器杀死流程，直到这个 Hook 定义操作完成之后，才允许容器被杀死
        preStop:
          exec:
            command: ["/usr/sbin/nginx","-s","quit"]
```

## Pod 对象在 Kubernetes 中的生命周期。
Pod 对象的生命周期主要体现在 Pod API 对象的 Status 部分，其中 pod.status.phase 就是Pod 的当前状态。
- Pending。这个状态意味着，Pod 的 YAML 文件已经提交给了 Kubernetes，API 对象已经被创建并保存在 Etcd 当中。但是，这个 Pod 里有些容器因为某种原因而不能被顺利创建。比如，调度不成功。


- Running。这个状态下，Pod 已经调度成功，跟一个具体的节点绑定。它包含的容器都已经创建成功，并且至少有一个正在运行中。


- Succeeded。这个状态意味着，Pod 里的所有容器都正常运行完毕，并且已经退出了。这种情况在运行一次性任务时最为常见。


- Failed。这个状态下，Pod 里至少有一个容器以不正常的状态（非 0 的返回码）退出。这个状态的出现，意味着你得想办法 Debug 这个容器的应用，比如查看 Pod 的 Events 和日志。


- Unknown。这是一个异常状态，意味着 Pod 的状态不能持续地被 kubelet 汇报给 kube-apiserver，这很有可能是主从节点（Master 和 Kubelet）间的通信出现了问题。


更进一步，Pod 的 Status 字段，还可细分出一组 Conditions，包括 PodScheduled、Ready、Initialized，以及 Unschedulable。

- PodScheduled：Pod 已经被调度到某节点；
- ContainersReady：Pod 中所有容器都已就绪；
- Initialized：所有的 Init 容器 都已成功启动；
- Ready：Pod 可以为请求提供服务，并且应该被添加到对应服务的负载均衡池中。
