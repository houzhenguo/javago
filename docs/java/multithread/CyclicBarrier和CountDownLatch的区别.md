

## 两者区别

两个看上去有点像的类，都在java.util.concurrent下，都可以用来表示代码运行到某个点上，二者的区别在于：

1）`CyclicBarrier`的某个线程运行到某个点上之后，该线程即停止运行，直到所有的线程都到达了这个点，所有线程才重新运行；
`CountDownLatch`则不是，某线程运行到某个点上之后，只是给某个数值-1而已，`该线程继续运行`。

2）CyclicBarrier只能唤起一个任务，CountDownLatch可以唤起多个任务。

3) CyclicBarrier可重用，CountDownLatch不可重用，计数值为0该CountDownLatch就不可再用了。



其实都可以 拦截多个线程。只要调用await方法即可。不一定是拦截一个线程。


---

以下描述的合理：

CountDownLatch : 一个线程(或者多个)， 等待另外N个线程完成某个事情之后才能执行。  CyclicBarrier        : N个线程相互等待，任何一个线程完成之前，所有的线程都必须等待。
这样应该就清楚一点了，对于CountDownLatch来说，重点是那个“一个线程”, 是它在等待， 而另外那N的线程在把“某个事情”做完之后可以继续等待，可以终止。而对于CyclicBarrier来说，重点是那N个线程，他们之间任何一个没有完成，所有的线程都必须等待。
