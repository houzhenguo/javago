
## Collection

1. 说说 ArrayList,Vector, LinkedList 的存储性能和特性。
2. 快速失败 (fail-fast) 和安全失败 (fail-safe) 的区别是什么？
3. hashmap , ConcurrentHashMap 的数据结构 工作原理 扩容
4. List、Map、Set 三个接口，存取元素时，各有什么特点？
5. Java 集合类框架的基本接口有哪些？ Collection -> List, Set , Map
6. HashSet 和 TreeSet 有什么区别？ Hash <-> 红黑树
7. LinkedHashMap 的实现原理?
8. 什么是迭代器 (Iterator)？Iterator 和 ListIterator 的区别是什么？
9. Comparable 和 Comparator 接口是干什么的？列出它们的区别
    
     > comparator 是比较器，可以直接执行。采用了策略模式，一个实体对象根据需要设计可以设计多个比较器。隔离性好，方便;comparable 通过比较实体对象来调用的。但是一个实体只能实现一个接口，所以扩展性不是很好，跟类绑定了。
10. Collection 和 Collections 的区别。

## 序列化问题
1. 为什么集合类没有实现 Cloneable 和 Serializable 接口？
2. transient变量有什么特点

## 克隆问题

## Java8问题

## Java 基础问题
1. Interface 与 abstract 类的区别
2. a = a + b 与 a += b 的区别？
3. float和double的默认值是多少 0.0  3.14f
4. char 型变量中能不能存贮一个中文汉字
5. 基础类型(Primitives)与封装类型(Wrappers)的区别在哪里
6. 基本数据类型的大小，以及他们的封装类
7. int 和 Integer 哪个会占用更多的内存 有争议
8. char 型变量中能不能存贮一个中文汉字
9. int 强制转换为 byte 类型的变量吗？如果该值大于 byte 类型的范围，将会出现什么现象？ 高 24位去掉
10. 

## JVM问题
1. Java 类加载过程？
2. 描述一下 JVM 加载 Class 文件的原理机制?
3. Java 内存分配。Java内存模型
4. GC 是什么? 为什么要有 GC？
5. 简述 Java 垃圾回收机制
6. 如何判断一个对象是否存活？（或者 GC 对象的判定方法）
7. 垃圾回收的优点和原理。并考虑 2 种回收机制
8. 垃圾回收器的基本原理是什么？垃圾回收器可以马上回收内存吗？有什么办法主动通知虚拟机进行垃圾回收？
9. Java 中会存在内存泄漏吗，请简单描述
10. 深拷贝和浅拷贝。
11. System.gc() 和 Runtime.gc() 会做什么事情？
12. finalize() 方法什么时候被调用？析构函数 (finalization) 的目的是什么？
13. 如果对象的引用被置为 null，垃圾收集器是否会立即释放对象占用的内存？
14. 什么是分布式垃圾回收（DGC）？它是如何工作的？
15. 串行（serial）收集器和吞吐量（throughput）收集器的区别是什么？
16. 在 Java 中，对象什么时候可以被垃圾回收？
17. 简述 Java 内存分配与回收策率以及 Minor GC 和 Major GC。
18. JVM 的永久代中会发生垃圾回收么？
19. Java 中垃圾收集的方法有哪些？
20. 什么是类加载器，类加载器有哪些？
21. 类加载器双亲委派模型机制？
22. 静态变量在什么时候加载？静态代码块加载的时机

## 高级部分

1. Java 的反射机制


## 并发相关

1. Synchronized 用过吗，其原理是什么？
2. 什么是可重入性，为什么说 Synchronized 是可重入锁？
3. JVM 对 Java 的原生锁做了哪些优化？
4. 为什么说 Synchronized 是非公平锁？
5. 什么是锁消除和锁粗化？
6.  Synchronized 是一个悲观锁？乐观锁的实现原理又是什么？什么是 CAS，它有什么特性？
7. 乐观锁一定就是好的吗？
8. Synchronized 相比，可重入锁 ReentrantLock 其实现原理有什么不同？
9. AQS 框架是怎么回事儿？
10. Synchronized 和 ReentrantLock 的异同。
11. ReentrantLock 是如何实现可重入性的？
12. 你还接触过 JUC 中的哪些并发工具？
13. ReadWriteLock 和 StampedLock。
14. 如何让 Java 的线程彼此同步？你了解过哪些同步器？请分别介绍下。
15. CyclicBarrier 和 CountDownLatch 看起来很相似，请对比下呢？
16. Java 线程池相关问题
17. Java 中的线程池是如何实现的？
18. 创建线程池的几个核心构造参数？
19. 线程池中的线程是怎么创建的？是一开始就随着线程池的启动创建好的吗？
20. 既然提到可以通过配置不同参数创建出不同的线程池，那么 Java 中默认实现好的线程池又有哪些呢？请比较它们的异同
21. 如何在 Java 线程池中提交线程？
22. 什么是 Java 的内存模型，Java 中各个线程是怎么彼此看到对方的变量的？
23. 请谈谈 volatile 有什么特点，为什么它能保证变量对所有线程的可见性？
24. 请对比下 volatile 对比 Synchronized 的异同。
25. 请谈谈 ThreadLocal 是怎么解决并发安全的？
26. 很多人都说要慎用 ThreadLocal，谈谈你的理解，使用 ThreadLocal 需要注意些什么？
27. BlockingQueue 简述 ConcurrentLinkedQueue LinkedBlockingQueue 的用处和不同之处。



## Spring
1. 什么是 Spring 框架？Spring 框架有哪些主要模块？
2. 使用 Spring 框架能带来哪些好处？
3. 什么是控制反转(IOC)？什么是依赖注入？
4. 请解释下 Spring 框架中的 IoC？
5. BeanFactory 和 ApplicationContext 有什么区别？
6. Spring 有几种配置方式？
7. 如何用基于 XML 配置的方式配置 Spring？
8. 如何用基于 Java 配置的方式配置 Spring？
9. 怎样用注解的方式配置 Spring？
10. 请解释 Spring Bean 的生命周期？
11. Spring Bean 的作用域之间有什么区别？
12. 什么是 Spring inner beans？
13. Spring 框架中的单例 Beans 是线程安全的么？
14. 请举例说明如何在 Spring 中注入一个 Java Collection？
15. 如何向 Spring Bean 中注入一个 Java.util.Properties？
16. 请解释 Spring Bean 的自动装配？
17. 请解释自动装配模式的区别？
18. 如何开启基于注解的自动装配？
19. 请举例解释@Required 注解？
20. 请举例解释@Autowired 注解？
21. 请举例说明@Qualifier 注解？
22. 构造方法注入和设值注入有什么区别？
23. Spring 框架中有哪些不同类型的事件？
24. FileSystemResource 和 ClassPathResource 有何区别？
25. Spring 框架中都用到了哪些设计模式？

## SpringBoot

1. 什么是 Spring Boot？
2. Spring Boot 有哪些优点？
3. 什么是 JavaConfig？
4. 如何重新加载 Spring Boot 上的更改，而无需重新启动服务器？
5. Spring Boot 中的监视器是什么？
6. 如何在 Spring Boot 中禁用 Actuator 端点安全性？
7. 如何在自定义端口上运行 Spring Boot 应用程序？
8. 什么是 YAML？
9. 如何实现 Spring Boot 应用程序的安全性？
10. 如何集成 Spring Boot 和 ActiveMQ？
11. 如何使用 Spring Boot 实现分页和排序？
12. 什么是 Swagger？你用 Spring Boot 实现了它吗？
13. 什么是 Spring Profiles？
14. 什么是 Spring Batch？
15. 什么是 FreeMarker 模板？
16. 如何使用 Spring Boot 实现异常处理？
17. 您使用了哪些 starter maven 依赖项？
18. 什么是 CSRF 攻击？
19. 什么是 WebSockets？
20. 什么是 AOP？
21. 什么是 Apache Kafka？
22. 我们如何监视所有 Spring Boot 微服务？


## 设计模式
1. 请列举出在 JDK 中几个常用的设计模式？
2. 什么是设计模式？你是否在你的代码里面使用过任何设计模式？
3. Java 中什么叫单例设计模式？请用 Java 写出线程安全的单例模式
4. 在 Java 中，什么叫观察者设计模式（observer design pattern）？
5. 使用工厂模式最主要的好处是什么？在哪里使用？
6. 举一个用 Java 实现的装饰模式(decorator design pattern)？它是作用于对象层次还是类
层次？
7. 在 Java 中，为什么不允许从静态方法中访问非静态变量？
8. 设计一个 ATM 机，请说出你的设计思路？
9. 在 Java 中，什么时候用重载，什么时候用重写？
10. 举例说明什么情况下会更倾向于使用抽象类而不是接口
11. 策略模式 观察者模式 单例模式 代理模式
12. 生产者 消费者模式。


## Netty
1. BIO、NIO和AIO的区别？
2. NIO的组成？
3. Netty的特点？
4. Netty的线程模型？
5. TCP 粘包/拆包的原因及解决方法？
6. 了解哪几种序列化协议？
7. 如何选择序列化协议？
8. Netty的零拷贝实现？
9. Netty的高性能表现在哪些方面？
10. NIOEventLoopGroup源码？

## Redis

1. Reids的特点
2. Redis支持的数据类型
3. Redis是单进程单线程的
4. Redis锁
5. 读写分离模型
6. 数据分片模型
7. Redis的回收策略
8. 使用Redis有哪些好处？
9. redis相比memcached有哪些优势？
10. redis常见性能问题和解决方案
11. MySQL里有2000w数据，redis中只存20w的数据，如何保证redis中的数据都是热点数据
12. Memcache与Redis的区别都有哪些？
13. Redis 常见的性能问题都有哪些？如何解决？
14. Redis 最适合的场景

## 算法

1. 一致hash算法
2. 什么是尾递归，为什么需要尾递归

## 数据库

1. ACID
2. 





