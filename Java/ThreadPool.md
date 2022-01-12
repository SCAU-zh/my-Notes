## Java 的线程池
Java 创建线程池有 2 种方法，一是使用 Executors 工厂类创建提供给我们的几种默认线程池，另外一种是使用 ThreadPoolExecutor 创建自定义线程池。

### 一、Executors 创建
- Executors.newCachedThreadPool();
- Executors.newFixedThreadPool();
- Executors.newSingleThreadExecutor();
- Executors.newFixedThreadPool();
- Executors.newWorkStealingPool();

### 二、ThreadPoolExecutor 创建
- corePoolSize：核心线程数，线程池中始终存活的线程数。


- maximumPoolSize: 最大线程数，线程池中允许的最大线程数。


- keepAliveTime: 存活时间，线程没有任务执行时最多保持多久时间会终止。


- unit: 单位，参数keepAliveTime的时间单位，7种可选。

| 参数      | 描述 |
| ----------- | ----------- |
| TimeUnit.DAYS      | 天       |
| TimeUnit.HOURS   | 小时        |
| ..........|........|

- workQueue: 一个阻塞队列，用来存储等待执行的任务，均为线程安全，7种可选。


| 参数                  | 描述                                                        |
| --------------------- | ------------------------------------------------------------ |
| ArrayBlockingQueue    | 一个由数组结构组成的有界阻塞队列                             |
| LinkedBlockingQueue   | 一个由链表结构组成的有界阻塞队列                             |
| SynchronousQueue      | 一个不存储元素的阻塞队列，即直接提交给线程不保持它们。       |
| PriorityBlockingQueue | 一个支持优先级排序的无界阻塞队列                             |
| DelayQueue            | 一个使用优先级队列实现的无界阻塞队列，只有在延迟期满时才能从中提取元素。 |
| LinkedTransferQueue   | 一个由链表结构组成的无界阻塞队列。与SynchronousQueue类似，还含有非阻塞方法。 |
| LinkedBlockingDeque   | 一个由链表结构组成的双向阻塞队列。                           |


较常用的是LinkedBlockingQueue和Synchronous。线程池的排队策略与BlockingQueue有关。

- threadFactory: 线程工厂，主要用来创建线程，默及正常优先级、非守护线程。


- handler：拒绝策略，拒绝处理任务时的策略，4种可选，默认为AbortPolicy。

| 参数                | 描述                                                     |
| ------------------- | --------------------------------------------------------- |
| AbortPolicy         | 拒绝并抛出异常                                            |
| CallerRunsPolicy    | 重试提交当前的任务，即再次调用运行该任务的execute()方法。 |
| DiscardOldestPolicy | 抛弃队列头部（最旧）的一个任务，并执行当前任务。          |
| DiscardPolicy       | 抛弃当前任务。                                            |


### 建议使用 ThreadPoolExecutor 来创建线程池
> 阿里 Java 开发规范强制不允许 Executors 创建线程池 

【强制】线程池不允许使用Executors去创建，而是通过ThreadPoolExecutor的方式，这样的处理方式让写的同学更加明确线程池的运行规则，规避资源耗尽的风险。
Executors返回的线程池对象的弊端如下:
FixedThreadPool和SingleThreadPool：允许的请求队列长度为Integer.MAX_VALUE，可能会堆积大量的请求，从而导致OOM。
CachedThreadPool和ScheduledThreadPool：允许的创建线程数量为Integer.MAX_VALUE，可能会创建大量的线程，从而导致OOM。
