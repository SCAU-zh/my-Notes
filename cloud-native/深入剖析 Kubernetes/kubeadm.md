# 使用 kubeadm 部署 kubernetes 集群

## 一、安装 kubeadm 和 Docker

``` shell
# 添加 kubeadm 的源，然后使用 apt-get 安装即可
$ curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
$ cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
$ apt-get update
$ apt-get install -y docker.io kubeadm
```

在上述安装 kubeadm 的过程中，kubeadm 和 kubelet、kubectl、kubernetes-cni 这几个二进制文件都会被自动安装好。



## 二、部署 kubernetes 的 Master 节点

通过配置文件部署，首先编写名为 kubeadm.yaml 的给 kubeadm 用的配置文件。

``` yaml
apiVersion: kubeadm.k8s.io/v1alpha1
kind: MasterConfiguration
controllerManagerExtraArgs:
# 将来部署的 kube-controller-manager 能够使用自定义资源（Custom Metrics）进行自动水平扩展
  horizontal-pod-autoscaler-use-rest-clients: "true"
  horizontal-pod-autoscaler-sync-period: "10s"
  node-monitor-grace-period: "10s"
apiServerExtraArgs:
  runtime-config: "api/all=true"
# 部署的 Kubernetes 版本号
kubernetesVersion: "stable-1.11"
```

然后，只需要执行一句指令

```shell
$ kubeadm init --config kubeadm.yaml
# 成功后的打印
kubeadm join 10.168.0.2:6443 --token 00bwbx.uvnaa2ewjflwu1ry --discovery-token-ca-cert-hash sha256:00eb62a2a6020f94132e3fe1ab721349bbcd3e9b94da9654cfe15f2985ebd711
```

这个 kubeadm join 命令，就是用来给这个 Master 节点添加更多工作节点（Worker）的命令。

我们在后面部署 Worker 节点的时候马上会用到它，所以找一个地方把这条命令记录下来。

