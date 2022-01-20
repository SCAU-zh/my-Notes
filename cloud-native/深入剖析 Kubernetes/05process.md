## 例子引入
```shell
$ docker run -it busybox /bin/sh
/ #
```
在上面的命令中，我们通过 docker run 启动了一个容器，在容器中执行 /bin/sh ，并且分配了一个命令行终端跟容器交互。
-it：告诉了 Docker 项目在启动容器后，需要给我们分配一个文本输入 / 输出环境，也就是 TTY，

```shell
/ # ps
PID  USER   TIME COMMAND
  1 root   0:00 /bin/sh
  10 root   0:00 ps
```

通过 ps 命令发现1号进程是 容器启动后执行的 /bin/sh ，说明这个容器已经被 docker 隔离在一个跟宿主机完全不同的世界中了(宿主机默认为 Linux 操作系统，并不是，通过 namespace 限制)。

## 如何实现与宿主机的隔离？
实际上，是通过 linux 的 namespace 机制使得进程只能看到重新计算过的进程编号。

namespace 的使用方式实际上是 Linux 创建新进程的一个可选参数：

```shell
# 在linux中创建进程的系统调用为 clone()，这个系统调用就会为我们创建一个新的进程，并且返回它的进程号 pid
int pid = clone(main_function, stack_size, SIGCHLD, NULL); 
# 在参数中指定 CLONE_NEWPID，新创建的这个进程将会“看到”一个全新的进程空间，在这个进程空间里，它的 PID 是 1。
int pid = clone(main_function, stack_size, CLONE_NEWPID | SIGCHLD, NULL);
# 新进程会看到一个新的进程空间，PID 是 1，实际上在宿主机真实的进程空间里，进程的 PID 还是真实的数值。
```
除了我们刚刚用到的 PID Namespace，Linux 操作系统还提供了 Mount、UTS、IPC、Network 和 User 这些 Namespace，用来对各种不同的进程上下文进行“障眼法”操作。

比如，Mount Namespace，用于让被隔离进程只看到当前 Namespace 里的挂载点信息；Network Namespace，用于让被隔离进程看到当前 Namespace 里的网络设备和配置。

所以说，**容器实际上是特殊的进程**。

### 容器
![avatar](https://notes-1303113205.cos.ap-guangzhou.myqcloud.com/my-notes-image/d1bb34cda8744514ba4c233435bf4e96.webp)
在 MacOS 和 windows 上跑 Linux 容器实际上是用到了虚拟化的技术，docker engine在宿主机上跑了个 Linux 虚拟机，再在虚拟机中跑 docker engine。
