

### 1 AQS 介绍

AQS (AbstractQueuedSynchronizer) ，位置 java.util.concurrent.locks包。AQS是一个用来构建 锁和同步器的框架，使用AQS可以简单 高效的构造出应用广泛的大量同步器，比如我们提到的 ReentrantLock，Semaphore 其他的比如 ReetrantReadWriteLock ， SynchronousQueue,FutureTask 都是基于 AQS的。

### 2 AQS原理

#### 2.1 原理概览

**AQS核心思想是：若果被请求的共享资源空闲，则将当前请求资源的线程设置为有效的工作线程，并且将共享资源设置为锁定的状态。如果被请求的共享资源被占用，那么就需要一套线程阻塞等待以及被唤醒时锁分配的机制，这个机制AQS是用CLH队列锁实现的，级将暂时获取不到锁的线程加入到队列中**

> CLH 队列是一个虚拟的双向队列（虚拟的双向队列即不存在队列实例，仅存在节点之间的关联关系）AQS是将每条请求共享资源的线程封装成一个CLH锁队列的一个节点 Node 来实现锁的分配。


AQS 原理图：

![AQS_CLH](../images/AQS_CLH.png)

AQS使用一个int成员变量来表示同步状态，通过内置的FIFO队列来完成获取资源线程的排队工作。AQS使用CAS对该同步状态进行原子操作实现对其值的修改。

```java
// 返回同步状态的当前值
protected final int getState()
{
    return state;
}
// 设置同步状态的值
protected final void setState(int newState)
{
    state = newState;
}
// 原子的（CAS操作）将同步状态值设置外给定值update,如果当前同步状态的值等于expect（期望值）
protected final boolean compareAndSetState(int expect, int update)
{
    return unsafe.compareAndSwapInt(this, stateOffset, expect, update);
}
```

#### 2.2 AQS对资源的共享方式

- **Exclusive**（独占）:只有一个线程能执行，如 ReetrantLock. 又可分为 公平锁和非公平锁：
    - 公平锁：按照线程在队列中排队的顺序，先到者先拿到锁
    - 非公平锁：当线程去获取锁时， 无视队列的顺序直接去抢锁，谁抢到就是谁的。
- **Share**(共享)：多个线程可以同时执行，如Semaphore/CountDownLatch. Semaphore,CountDownLatch,CyclicBarrier, ReadWriteLock 我们后面会讲到。

ReentrantReadWriteLock 可以看成是组合式，因为 ReentrantReadWriteLock也就是读写锁 允许多个线程同时对某一资源进行读。

不同的自定义同步器争用共享资源的方式也不同。自定义同步器在实现时只需要实现资源 state 的获取与释放方式 即可。至于具体线程等待队列的维护（如获取资源失败入队/唤醒出队等）AQS已经在上层帮我们实现好了。

#### 2.3 AQS底层使用了模板方法模式

同步器的实际是基于模板方法模式的，如果需要自定义同步器一般的方式是这样（模板方法模式是很经典的一个应用）:

1. 使用者 继承 AbstractQueuedSynchronizer 并重写指定的方法。（这些重写方法很简单，无非是对于共享资源state的获取和释放）

2. 将AQS组合在自定义同步组建的实现中，并调用其模板方法，而这些模板方法会调用使用者重写的方法。

这和我们意外实现接口的方式有很大的区别，这是模板方法一个很经典的应用。下面简单的介绍一个模板方法，这个设计模式也是很好理解的之一。

> 模板方法模式是基于”继承“的，主要是为了在不改变模板结构的前提下在子类中重新定义模板中的内容以实现复用代码。举个很简单的例子假如我们要去一个地方的步骤是：购票`buyTicket()`->安检`securityCheck()`->乘坐某某工具回家`ride()`->到达目的地`arrive()`。我们可能乘坐不同的交通工具回家比如飞机或者火车，所以除了`ride()`方法，其他方法的实现几乎相同。我们可以定义一个包含了这些方法的抽象类，然后用户根据自己的需要继承该抽象类然后修改 `ride()`方法。

**AQS使用了模板方法模式，自定义同步器需要重写下面几个AQS提供的模板方法：**

```java
inHeldExclusively(); // 该线程是否正在独占资源。只有用到 condition才需要去实现它。
tryAcquire(int); // 独占方式。尝试获取资源，成功则返回true，失败则返回false。
tryRelease(int); // 独占方式。尝试释放资源，成功返回true，失败返回false。
tryAcquireShared(int) // 共享方式。尝试获取资源。负数表示失败，0表示成功，但没有可以用资源；正数表示成功，且有剩余资源。
tryReleaseShared(int) // 共享方式。尝试释放资源，成功返回true，失败返回false。
```

默认情况下，每个方法都抛出 `UnsupportedOperationException`。这些方法的实现必须是内部线程安全的，并且通常应该简短而不阻塞。AQS中其他方法都是final，所以无法被其他类使用，只有这个几个方法可以被其他类使用。

以 ReentrantLock 为例，state初始化为0，表示未锁定的状态。A线程 lock()时，会调用 tryAcquire()独占该锁，并将state+1.此后，其他线程在tryAcquire() 时就会失败，直到A线程unlock()到state=0（即释放锁）为止，其他线程才有机会获得该锁。当然，释放锁之前，A线程自己是可以重复获得此锁的（state会累加）,这就是可重入的概念，但是要注意，获取多少次就要释放多少次，这样才能保证 state是可能回到零态的。


再以CountDownLatch 为例，任务分为 N个子线程去执行，state初始化为N（注意N要与线程个数一致）.这N个子线程是并行执行的，每个子线程执行完之后 countDown()一次，state会CAS 减1，等到所有的子线程都执行完成之后（state=0），会unpark()主调用线程，然后主调用线程就会从await()函数返回，继续后续动作。

 一般来说，自定义同步器要么是独占方法，要么是共享方式，他们只需要实现`tryAcquire-tryRelease`、`tryAcquireShared-tryReleaseShared`中的一种即可。但 AQS也支持自定义同步器同时实现独占和共享两种方式，如 `ReentrantReadWriteLock`.

 推荐两篇 AQS 原理和相关源码分析的文章：

- http://www.cnblogs.com/waterystone/p/4920797.html
- https://www.cnblogs.com/chengxiao/archive/2017/07/24/7141160.html

### 3 Semaphore（信号量）- 允许多个线程同时访问

**synchronized 和 ReentrantLock 都是一次只允许一个线程访问某个资源，Semaphore(信号量)可以指定多个线程同时访问某个资源** 示例代码如下:

```java
public class SemaphoreExample1 {
	// 请求的数量
	private static final int threadCount = 550;

	public static void main(String[] args) throws InterruptedException{
		// 创建一个具有固定线程数量的线程池对象
        //（如果这里线程池的线程数量给太少，会执行的很慢）
		ExecutorService threadPool = Executors.newFixedThreadPool(300);
		// 一次只允许执行的线程数量
		final Semaphore semaphore = new Semaphore(20);
		for (int i=0;i<threadCount;i++)
		{
			final int threadNum = i;
			threadPool.execute(()->{
				try
				{
					semaphore.acquire(); // 获取一个许可，所以可运行线程的数量为20/1 = 20
					test(threadNum);
					semaphore.release(); // 释放一个许可
				}catch (InterruptedException ex)
				{
					ex.printStackTrace();
				}
			});
		}
		threadPool.shutdown();
		System.out.println("finish");
	}
	public static void test(int threadnum) throws InterruptedException {
		Thread.sleep(1000);// 模拟请求的耗时操作
		System.out.println("threadnum:" + threadnum);
		Thread.sleep(1000);// 模拟请求的耗时操作
	}
}

```

执行 `acquire` 方法阻塞，直到有一个许可证可以获得然后拿走一个许可证；每个 `release` 方法增加一个许可证，这可能会释放一个阻塞的acquire方法。然而，其实并没有实际的许可证这个对象，Semaphore只是维持了一个可获得许可证的数量。 Semaphore经常用于限制获取某种资源的线程数量。

当然一次也可以一次拿取和释放多个许可，不过一般没有必要这样做：

```java
          semaphore.acquire(5);// 获取5个许可，所以可运行线程数量为20/5=4
          test(threadnum);
          semaphore.release(5);// 释放5个许可
```

除了 `acquire`方法之外，另一个比较常用的与之对应的方法是`tryAcquire`方法，该方法如果获取不到许可就立即返回false。


Semaphore 有两种模式，公平模式和非公平模式。

- **公平模式：** 调用acquire的顺序就是获取许可证的顺序，遵循FIFO；
- **非公平模式：** 抢占式的。

**Semaphore 对应的两个构造方法如下：**

```java
   public Semaphore(int permits) {
        sync = new NonfairSync(permits);
    }

    public Semaphore(int permits, boolean fair) {
        sync = fair ? new FairSync(permits) : new NonfairSync(permits);
    }
```
**这两个构造方法，都必须提供许可的数量，第二个构造方法可以指定是公平模式还是非公平模式，默认非公平模式。** 

由于篇幅问题，如果对 Semaphore 源码感兴趣的朋友可以看下面这篇文章：

- https://blog.csdn.net/qq_19431333/article/details/70212663

### 4 CountDownLatch(倒计时器)

CountDownLatch 是一个同步工具类，它允许一个或多个线程一直等待，知道其他线程的操作执行完成之后再执行。

#### 4.1 CountDownLatch的三种典型用法

1. 每个线程在开始运行前等待n个线程执行完毕。将CountDownLatch 的计数器初始化为 n,`new CountDownLatch(n)` ，每当一个任务线程执行完毕，将计数器 减1 `countdownLatch.countDown()`,当计数器的值变为 0 时，在`CountDownLatch 上 await()` 的线程就会被唤醒。一个典型的应用场景就是启动一个服务时，主线程需要等待多个组件加载完毕，之后再继续执行。

2. 实现多个线程开始执行任务的最大并行性。注意是并行性，不是并发，强调的是多个线程在某一个时刻同时开始执行。类似于赛跑，将多个线程放到起点，等待发令枪响，然后同时开跑。做法是初始化一个共享的`CountDownLatch` 对象，将其计数器初始化为 1，`new CountDownLatch(1)` ，多个线程开始执行任务前首先 `countdownlatch.await()` ,当主线程调用 countDown()时，计数器变为 0，过个线程同时被唤醒。

3. 死锁检测：一个非常方便的使用场景是，你可以使用n个线程访问共享资源，在每次测试阶段的线程数目是不同的，并尝试产生死锁。


#### 4.2 CountDownLatch 的使用示例

```java
/**
 * @Description: CountDownLatch 使用方法示例
 */
public class CountDownLatchExample1 {
  // 请求的数量
  private static final int threadCount = 550;

  public static void main(String[] args) throws InterruptedException {
    // 创建一个具有固定线程数量的线程池对象（如果这里线程池的线程数量给太少的话你会发现执行的很慢）
    ExecutorService threadPool = Executors.newFixedThreadPool(300);
    final CountDownLatch countDownLatch = new CountDownLatch(threadCount);
    for (int i = 0; i < threadCount; i++) {
      final int threadnum = i;
      threadPool.execute(() -> {// Lambda 表达式的运用
        try {
          test(threadnum);
        } catch (InterruptedException e) {
          // TODO Auto-generated catch block
          e.printStackTrace();
        } finally {
          countDownLatch.countDown();// 表示一个请求已经被完成
        }

      });
    }
    countDownLatch.await();
    threadPool.shutdown();
    System.out.println("finish");
  }

  public static void test(int threadnum) throws InterruptedException {
    Thread.sleep(1000);// 模拟请求的耗时操作
    System.out.println("threadnum:" + threadnum);
    Thread.sleep(1000);// 模拟请求的耗时操作
  }
}

```
上面的代码中，我们定义了请求的数量为550，当这550个请求被处理完成之后，才会执行`System.out.println("finish");`。

与CountDownLatch的第一次交互是主线程等待其他线程。主线程必须在启动其他线程后立即调用CountDownLatch.await()方法。这样主线程的操作就会在这个方法上阻塞，直到其他线程完成各自的任务。

其他N个线程必须引用闭锁对象，因为他们需要通知CountDownLatch对象，他们已经完成了各自的任务。这种通知机制是通过 CountDownLatch.countDown()方法来完成的；每调用一次这个方法，在构造函数中初始化的count值就减1。所以当N个线程都调 用了这个方法，count的值等于0，然后主线程就能通过await()方法，恢复执行自己的任务。

#### 4.3 CountDownLatch 的不足

CountDownLatch是一次性的，计数器的值只能在构造方法中初始化一次，之后没有任何机制再次对其设置值，当CountDownLatch使用完毕后，它不能再次被使用。

#### 4.4 CountDownLatch：

解释一下CountDownLatch概念？

CountDownLatch 和CyclicBarrier的不同之处？

给出一些CountDownLatch使用的例子？

CountDownLatch 类中主要的方法？

## CountDownLatch源码阅读
[参考](https://juejin.im/post/5ae754dd6fb9a07abc29b2ce)

shared -> 可中断的
CountDownLatch 底层使用的是 Sync-> AQS

CountDownLatch() 构造方法指定的是 AQS 中 state 的数量

tryAcquireShared(){(getState() == 0) ? 1 : -1} 当获取的时候，如果state 为 0 才可以获取。

tryAcquireSharedNanos() // 超时


await() -> doAcquireSharedInterruptibly  // for(;;)循环等待，阻塞
// new Node ,tail 到队列 ,

countDown() //
## await步骤
1. 将当前线程包装成一个 Node 对象，加入到 AQS 的队列尾部。
2. 如果他前面的 node 是 head ，便可以尝试获取锁了。
3. 如果不是，则阻塞等待，调用的是 LockSupport.park(this);

CountDown 的 await 方法就是通过 AQS 的锁机制让主线程阻塞等待。而锁的实现就是通过构造器中设置的 state 变量来控制的。当 state 是 0 的时候，就可以获取锁。然后执行后面的逻辑。

## ReentrantLock源码阅读

Sync

state : 对于重入锁来说，是 获取的次数。
ownerExclusiveThread : 所属独占锁的线程是否是当前线程 <->若不是，release会抛出异常。

isLocked() : state > 0 ? true : false;

fair : 底层有两个类的实现

lock(){sync.accquire(1)} -> tryAcquire -> nonfairTryAcquire -> if state ==0 || owner == myThread -> return true

unlock -> release wake up

## Semaphore

permits -> state 同上 