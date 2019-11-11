
## 死锁

```java
public class DeadLock {
    public static void main(String[] args) {
        Object object1= new Object();
        Object object2= new Object();
        new Thread(()->{
            synchronized (object1) {
                try {
                    Thread.sleep(1000L);
                }catch (Exception e) {

                }
                synchronized (object2) {

                }

            }
        },"Thread1").start();

        new Thread(()->{
            synchronized (object2) {
                try {
                    Thread.sleep(1000L);
                }catch (Exception e) {

                }
                synchronized (object1) {

                }
            }
        },"Thread2").start();
    }
}

```

## AtomicInteger 
```java
public class AtomicTest {
    public static AtomicInteger val =new AtomicInteger();
    public static int noSafeVal = 0;

    public static void main(String[] args) throws Exception{
        CountDownLatch countDownLatch = new CountDownLatch(10000);
        for (int i=0; i<10000; i++) {
            new Thread(()->{
                try {
                    Thread.sleep(10);
                }catch (Exception e) {

                }

                noSafeVal++;
                val.getAndIncrement();
                countDownLatch.countDown();
            }).start();
        }
        countDownLatch.await();
        System.out.println("noSafe"+noSafeVal);
        System.out.println("safe"+val.get());
    }
}

```