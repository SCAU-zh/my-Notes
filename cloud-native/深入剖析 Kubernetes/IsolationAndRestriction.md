## 容器的隔离与限制
一个正在运行的 Docker 容器，

其实就是一个启用了多个 Linux Namespace 的应用进程，

而这个进程能够使用的资源量，则受 Cgroups 配置的限制。
### 隔离
> Namespace 技术实际上修改了应用进程看待整个计算机“视图”，即它的“视线”被操作系统做了限制，只能“看到”某些指定的内容

**真正对隔离环境负责的是宿主机操作系统本身**，用户运行在容器中的进程，实际上和宿主机中的进程一样，都由宿主机进行管理，只不过这些进程设置了额外的 namespace 参数。

而虚拟化技术需要先通过 Hypervisor 技术负责创建虚拟机，再在虚拟机上启动一个 guest OS ，对于宿主机的系统调用需要经过虚拟化软件的兰姐和处理，

因此 **敏捷与高性能**是容器的优势。

而 **隔离得不够彻底**，是容器相比较虚拟机的劣势。

在 Linux 内核中，有很多资源和对象是不能被 Namespace 化的，最典型的例子就是：时间。

挂载 localtime 时要设置read-only。

### 限制
> Linux Cgroups 就是 Linux 内核中用来为进程设置资源限制的一个重要功能
> 
> Linux Cgroups 的全称是 Linux Control Group。
> 
> 它最主要的作用，就是限制一个进程组能够使用的资源上限，包括 CPU、内存、磁盘、网络带宽等等

Linux Cgroups 的设计还是比较易用的，简单粗暴地理解呢，它就是一个子系统目录加上一组资源限制文件的组合。

而对于 Docker 等 Linux 容器项目来说，它们只需要在每个子系统下面，为每个容器创建一个控制组（即创建一个新目录），然后在启动容器进程之后，把这个进程的 PID 填写到对应控制组的 tasks 文件中就可以了。

docker run 实现 ：
```shell
$ docker run -it --cpu-period=100000 --cpu-quota=20000 ubuntu /bin/bash
```

通过查看 Cgroups 文件系统下，CPU 子系统中，“docker”这个控制组里的资源限制文件的内容来确认：

```shell
$ cat /sys/fs/cgroup/cpu/docker/5d5c9f67d/cpu.cfs_period_us 
100000
$ cat /sys/fs/cgroup/cpu/docker/5d5c9f67d/cpu.cfs_quota_us 
20000
```

### Cgroups 的问题
> Linux 下的 /proc 目录存储的是记录当前内核运行状态的一系列特殊文件，用户可以通过访问这些文件，查看系统以及当前正在运行的进程的信息，
> 
> 比如 CPU 使用情况、内存占用率等，这些文件也是 top 指令查看系统信息的主要数据来源。
> 但是，你如果在容器里执行 top 指令，就会发现，它显示的信息居然是宿主机的 CPU 和内存数据，而不是当前容器的数据。
> 
> 造成这个问题的原因就是，/proc 文件系统并不知道用户通过 Cgroups 给这个容器做了什么样的资源限制，即：/proc 文件系统不了解 Cgroups 限制的存在。

修正：
top 是从 /prof/stats 目录下获取数据，所以道理上来讲，容器不挂载宿主机的该目录就可以了。

lxcfs就是来实现这个功能的，做法是把宿主机的 /var/lib/lxcfs/proc/memoinfo 文件挂载到Docker容器的/proc/meminfo位置。

kubernetes 环境下，也能用，以ds 方式运行 lxcfs ，自动给容器注入正确的 proc 信息。
