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

## kubernetes 的几种 project(投射) volume
> 在 Kubernetes 中，有几种特殊的 Volume，它们存在的意义不是为了存放容器里的数据，也不是用来进行容器和宿主机之间的数据交换。
> 
> 这些特殊 Volume 的作用，是为容器提供预先定义好的数据。
> 
> 所以，从容器的角度来看，这些 Volume 里的信息就是仿佛是被 Kubernetes“投射”（Project）进入容器当中的。
> 
> 这正是 Projected Volume 的含义

### 1、Secret
secret 的作用是把 Pod 需要访问到的加密数据，存放到 etcd 中，这样子就可以通过 Pod 容器里挂载 Volume 的方式，访问到保存的信息。

下面给出使用例子

使用 Secret 对象的Pod：
```yaml
apiVersion: v1
kind: Pod
metadata: 
  name: test-secret-pod
spec: 
  containers: 
    - name: test-secret-pod
      image: busybox
      args: 
        - sleep
        - "86400"
      volumeMounts: 
        - name: mysql-cred
          mountPath: "/mysql-cred"
          readOnly: true
  # 声明挂载了类型为 projected 的 volume，来源为 secret
  volumeMounts: 
    - name: mysql-cred
      projected: 
        sources: 
          - secret: 
              name: user
          - secret: 
              name: password
```

创建 secret 对象(命令创建和 yaml 文件)：
``` shell
$ cat ./username.txt
admin
$ cat ./password.txt
c1oudc0w!

$ kubectl create secret generic user --from-file=./username.txt
$ kubectl create secret generic pass --from-file=./password.txt

$ kubectl get secrets
NAME           TYPE                                DATA      AGE
user          Opaque                                1         51s
pass          Opaque                                1         51s
```
```yaml
apiVersion: v1
kind: Secret
metadata: 
  name: mysecret
type: Opaque
data: 
  user: YWRtaW4=
  password: MWYyZDFlMmU2N2Rm
```

通过这种方式 挂载进入容器的 Secret，会同步更新改动，因为 kubelet 组件在定时维护这些 Volume

但这个更新可能有时延，所以在编写链接数据库的逻辑时，要主要加上重试和超时的逻辑

### 2、configMap
ConfigMap 与 Secret 类似，区别是 ConfigMap 用于存储不需要加密的，应用所需的配置信息。

例子：
```yaml
appVersion: v1
kind: ConfigMap
data:
  ui.properties: | 
    color.good=purple 
    color.bad=yellow 
    allow.textmode=true 
    how.nice.to.look=fairlyNice
metadata: 
  name: ui-config
.......
```
```shell

# .properties文件的内容
$ cat example/ui.properties
color.good=purple
color.bad=yellow
allow.textmode=true
how.nice.to.look=fairlyNice

# 从.properties文件创建ConfigMap
$ kubectl create configmap ui-config --from-file=example/ui.properties
```
> 备注：kubectl get -o yaml 这样的参数，会将指定的 Pod API 对象以 YAML 的方式展示出来。

### 3、downwardAPI
作用是让 Pod 中的容器可以访问到 Pod Api 对象本身的信息。
```yaml
apiVersion: v1
kind: Pod
metadata: 
  name: test-downwardapi-volume
  labels: 
    zone: us-est-coast 
    cluster: test-cluster1 
    rack: rack-22
spec: 
  containers: 
    - name: client-container
      image: busybox
      command: ["sh", "-c"]
      args: 
        - while true; do
            if [[ -e /etc/podinfo/labels ]]; then
              echo -en '\n\n'; cat /etc/podinfo/labels; fi;
            sleep 5;
          done;
      volumeMounts: 
        - name: podinfo
          mountPath: /etc/podinfo
          readOnly: false
  volumes: 
    - name: podinfo
      projected: 
        sources: 
          - downwardAPI: 
            items: 
              - path: "labels"
                fieldRef: 
                  fieldPath: metadata.labels
```
通过上面例子的方式，当前 Pod 的 Labels 的值，就会被 kubernetes 自动挂载到成为容器的 /etc/podinfo/labels 文件

Downward API 能够获取到的信息，一定是 Pod 里的容器进程启动之前就能够确定下来的信息

Downward API 支持的部分字段
```yaml
1. 使用fieldRef可以声明使用:
  spec.nodeName - 宿主机名字
  status.hostIP - 宿主机IP
  metadata.name - Pod的名字
  metadata.namespace - Pod的Namespace
  status.podIP - Pod的IP
  spec.serviceAccountName - Pod的Service Account的名字
  metadata.uid - Pod的UID
  metadata.labels['<KEY>'] - 指定<KEY>的Label值
  metadata.annotations['<KEY>'] - 指定<KEY>的Annotation值
  metadata.labels - Pod的所有Label
  metadata.annotations - Pod的所有Annotation

2. 使用resourceFieldRef可以声明使用:
  容器的CPU limit
  容器的CPU request
  容器的memory limit
  容器的memory request
```

### 4、ServiceAccountToken（一种特殊的 Secret）
使用场景： 从pod里面调用k8s API来控制集群，需要使用 serviceAccountToken 保存的授权信息，才可以合法地访问 API Service

##

