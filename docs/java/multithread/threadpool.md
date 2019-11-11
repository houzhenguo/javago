
# 线程池的相关知识

## 线程池的组成

一般的线程池主要分为以下4 个组成部分：

1. 线程池管理器：用于创建并管理线程池
2. 工作线程：线程池中的线程
3. 任务接口：每个任务必须实现的接口，用于工作线程调度其运行
4. 任务队列：用于存放待处理的任务，提供一种缓冲机制

Java中的线程池是通过 Executor 框架实现的，该框架中用到了 Executor， Executors, ExecutorService, ThreadPoolExecutor, Callable, Future, FutureTask等

![threadpool](../images/thread-pool-1.png)

ThreadPoolExecutor 的构造方法如下：

```java
public ThreadPoolExecutor(int corePoolSize,int maximumPoolSize, long keepAliveTime,
    TimeUnit unit, BlockingQueue<Runnable> workQueue) {

    this(corePoolSize, maximumPoolSize, keepAliveTime, unit, workQueue
    ,Executors.defaultThreadFactory(), defaultHandler);
}
```

1. corePoolSize：指定了线程池中的线程数量。
2. maximumPoolSize：指定了线程池中的最大线程数量。
3. keepAliveTime：当前线程池数量超过corePoolSize 时，多余的空闲线程的存活时间，即多
次时间内会被销毁。
4. unit：keepAliveTime 的单位。
5. workQueue：任务队列，被提交但尚未被执行的任务。
6. threadFactory：线程工厂，用于创建线程，一般用默认的即可。
7. handler：拒绝策略，当任务太多来不及处理，如何拒绝任务。

备注：  对于计算密集型的任务，在拥有N个CPU处理器的系统上，当线程池的大小为 N+1个CPU个数时候，通常能实现最优的利用率。

可以通过以下方式获取CPU的数目:
```java
int N_CPUS = Runtime.getRuntime().availableProcessors();
```

## 拒绝策略

线程池中的线程已经用完了，无法继续为新任务服务，同时，等待队列也已经排满了，再也
塞不下新任务了。这时候我们就需要拒绝策略机制合理的处理这个问题。
JDK 内置的拒绝策略如下：
1. AbortPolicy ： 直接抛出异常，阻止系统正常运行。newFixedThreadPool默认使用。
2. CallerRunsPolicy ： 只要线程池未关闭，该策略直接在调用者线程中，运行当前被丢弃的
任务。显然这样做不会真的丢弃任务，但是，任务提交线程的性能极有可能会急剧下降。
3. DiscardOldestPolicy ： 丢弃最老的一个请求，也就是即将被执行的一个任务，并尝试再
次提交当前任务。
4. DiscardPolicy ： 该策略默默地丢弃无法处理的任务，不予任何处理。如果允许任务丢
失，这是最好的一种方案。
以上内置拒绝策略均实现了RejectedExecutionHandler 接口，若以上策略仍无法满足实际
需要，完全可以自己扩展RejectedExecutionHandler 接口。

## Java线程池的工作过程 （重要）

1. 线程池刚创建时，里面没有一个线程。任务队列是作为参数传进来的。不过，就算队列里面
有任务，线程池也不会马上执行它们。

2. 当调用 execute() 方法添加一个任务时，线程池会做如下判断：
- a) 如果正在运行的线程数量小于 corePoolSize，那么马上创建线程运行这个任务；
- b) 如果正在运行的线程数量大于或等于 corePoolSize，那么将这个任务放入队列；
- c) 如果这时候队列满了，而且正在运行的线程数量小于 maximumPoolSize，那么还是要
创建非核心线程立刻运行这个任务；
- d) 如果队列满了，而且正在运行的线程数量大于或等于 maximumPoolSize，那么线程池
会抛出异常RejectExecutionException。

3. 当一个线程完成任务时，它会从队列中取下一个任务来执行。

4. 当一个线程无事可做，超过一定的时间（keepAliveTime）时，线程池会判断，如果当前运
行的线程数大于 corePoolSize，那么这个线程就被停掉。所以线程池的所有任务完成后，它
最终会收缩到 corePoolSize 的大小。

## Executors创建返回ThreadPoolExecutor对象

`Executors`创建返回ThreadPoolExecutor对象的方法共有三种：

- Executors#newCachedThreadPool => 创建可缓存的线程池
- Executors#newSingleThreadExecutor => 创建单线程的线程池
- Executors#newFixedThreadPool => 创建固定长度的线程池

## Executors#newCachedThreadPool方法

```java
public static ExecutorService newCachedThreadPool() {
    return new ThreadPoolExecutor(0, Integer.MAX_VALUE,
                                  60L, TimeUnit.SECONDS,
                                  new SynchronousQueue<Runnable>());
}

```
`CachedThreadPool`是一个根据需要创建新线程的线程池

- corePoolSize => 0，核心线程池的数量为0
- maximumPoolSize => Integer.MAX_VALUE，可以认为最大线程数是无限的
- keepAliveTime => 60L
- unit => 秒
- workQueue => SynchronousQueue

当一个任务提交时，`corePoolSize`为0不创建核心线程，`SynchronousQueue`是一个不存储元素的队列，可以理解为队里永远是满的，因此最终会创建非核心线程来执行任务。对于非核心线程空闲60s时将被回收。因为Integer.MAX_VALUE非常大，可以认为是可以无限创建线程的，在资源有限的情况下容易引起OOM异常

## Executors#newSingleThreadExecutor方法

```java
public static ExecutorService newSingleThreadExecutor() {
    return new FinalizableDelegatedExecutorService
        (new ThreadPoolExecutor(1, 1,
                                0L, TimeUnit.MILLISECONDS,
                                new LinkedBlockingQueue<Runnable>()));
}

```
`SingleThreadExecutor`是单线程线程池，只有一个核心线程

- corePoolSize => 1，核心线程池的数量为1 (重要)
- maximumPoolSize => 1，只可以创建一个非核心线程(重要)
- keepAliveTime => 0L
- unit => 毫秒
- workQueue => LinkedBlockingQueue

当一个任务提交时，首先会创建一个核心线程来执行任务，如果超过核心线程的数量，将会放入队列中，因为`LinkedBlockingQueue`是长度为Integer.MAX_VALUE的队列，可以认为是无界队列，因此往队列中可以插入无限多的任务，在资源有限的时候容易引起OOM异常，同时因为无界队列，`maximumPoolSize`和`keepAliveTime`参数将无效，压根就不会创建非核心线程

## Executors#newFixedThreadPool方法

```java
public static ExecutorService newFixedThreadPool(int nThreads) {
    return new ThreadPoolExecutor(nThreads, nThreads,
                                  0L, TimeUnit.MILLISECONDS,
                                  new LinkedBlockingQueue<Runnable>());
}

```
`FixedThreadPool`是固定核心线程的线程池，固定核心线程数由用户传入

## 如何定义线程池参数

- CPU密集型 => 线程池的大小推荐为`CPU`数量 + 1，`CPU`数量可以根据`Runtime.availableProcessors`方法获取

- IO密集型 => `CPU数量` * `CPU利用率` * (1 + 线程等待时间/线程CPU时间)

- 混合型 => 将任务分为`CPU密集型`和`IO密集型`，然后分别使用不同的线程池去处理，从而使每个线程池可以根据各自的工作负载来调整

- 阻塞队列 => 推荐使用有界队列，有界队列有助于避免资源耗尽的情况发生

- 拒绝策略 => 默认采用的是`AbortPolicy`拒绝策略，直接在程序中抛出

`RejectedExecutionException`异常【因为是运行时异常，不强制catch】，这种处理方式不够优雅。处理拒绝策略有以下几种比较推荐：


## 参考
[阿里巴巴禁止使用Excutors创建线程池](https://juejin.im/post/5dc41c165188257bad4d9e69)

