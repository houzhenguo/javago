
[原文连接](https://www.jianshu.com/p/4001ab555c19)

# ReentrantLock lockInterruptibly()
ReentrantLock的加锁方法Lock()提供了无条件地轮询获取锁的方式，lockInterruptibly()提供了可中断的锁获取方式。这两个方法的区别在哪里呢？通过分析源码可以知道lock方法默认处理了中断请求，一旦监测到中断状态，则中断当前线程；而lockInterruptibly()则直接抛出中断异常，由上层调用者区去处理中断。

## lock操作

lock获取锁过程中，忽略了中断，在成功获取锁之后，再根据中断标识处理中断，即selfInterrupt中断自己。 acquire操作源码如下：

源码入口 : `ReentrantLock.lock()`

```java
    public final void acquire(int arg) {
        if (!tryAcquire(arg) &&
            acquireQueued(addWaiter(Node.EXCLUSIVE), arg))
            selfInterrupt();
    }

    final boolean acquireQueued(final Node node, int arg) {
        try {
            boolean interrupted = false;
            for (;;) {
                final Node p = node.predecessor();
                if (p == head && tryAcquire(arg)) {
                    setHead(node);
                    p.next = null; // help GC
                    return interrupted;
                }
                if (shouldParkAfterFailedAcquire(p, node) &&
                    parkAndCheckInterrupt())
                    interrupted = true; // 好像是这里产生的不同。这里只是设置了标志位，继续获取锁。
            }
        } catch (Throwable t) {
            cancelAcquire(node);
            throw t;
        }
    }    
```

备注：

1. `selfInterrupt`在线程 t 调用自己的 t.interrupt() 方法后，此线程中断标志就变成true。但是，中断标志为true实际上不会对正常运行的线程产生影响，因为正常运行的线程不会自己去检查自己的中断标志。
2. `acquireQueued`，在for循环中无条件重试获取锁，直到成功获取锁，同时返回线程中断状态。该方法通过for循正常返回时，必定是成功获取到了锁。

## lockInterruptibly操作

可中断加锁，即在锁获取过程中不处理中断状态，而是直接抛出中断异常，由上层调用者处理中断。源码细微差别在于锁获取这部分代码，这个方法与acquireQueue差别在于方法的返回途径有两种，一种是for循环结束，正常获取到锁；另一种是线程被唤醒后检测到中断请求，则立即抛出中断异常，该操作导致方法结束。

```java
    public final void acquireInterruptibly(int arg)
            throws InterruptedException {
        if (Thread.interrupted())
            throw new InterruptedException();
        if (!tryAcquire(arg))
            doAcquireInterruptibly(arg); 
    }

    private void doAcquireInterruptibly(int arg)
        throws InterruptedException {
        final Node node = addWaiter(Node.EXCLUSIVE);
        try {
            for (;;) {
                final Node p = node.predecessor();
                if (p == head && tryAcquire(arg)) {
                    setHead(node);
                    p.next = null; // help GC
                    return;
                }
                if (shouldParkAfterFailedAcquire(p, node) &&
                    parkAndCheckInterrupt())
                    throw new InterruptedException(); // 好像是这里产生的不同。这里抛出了异常，中断了
            }
        } catch (Throwable t) {
            cancelAcquire(node);
            throw t;
        }
    }
```
## 结论
 ：ReentrantLock的中断和非中断加锁模式的区别在于：线程尝试获取锁操作失败后，在等待过程中，如果该线程被其他线程中断了，它是如何响应中断请求的。lock方法会忽略中断请求，继续获取锁直到成功；而lockInterruptibly则直接抛出中断异常来立即响应中断，由上层调用者处理中断。

     那么，为什么要分为这两种模式呢？这两种加锁方式分别适用于什么场合呢？根据它们的实现语义来理解，我认为lock()适用于锁获取操作不受中断影响的情况，此时可以忽略中断请求正常执行加锁操作，因为该操作仅仅记录了中断状态（通过Thread.currentThread().interrupt()操作，只是恢复了中断状态为true，并没有对中断进行响应)。如果要求被中断线程不能参与锁的竞争操作，则此时应该使用lockInterruptibly方法，一旦检测到中断请求，立即返回不再参与锁的竞争并且取消锁获取操作（即finally中的cancelAcquire操作）
