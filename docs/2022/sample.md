
## jvm
1. 内存区域： 程序计数器，虚拟机栈，本地方法栈，堆，方法区（meta元空间）-> 运行时常量池(string,引用，final,普通变量
a. oom, sof,nio, 不受jvm限制，zero copy 减少 复制，maxMeta
2. Heap 对象 和数组分配 ， 新生代，老年代。s0,s1,eden, 15age header, 大对象。
3. new ->
    a. 类加载检查 -> 运行时常量池 是否能定位到引用 -> 类是否加载解析初始化
    b. 分配内存空间 heap 指针碰撞 空闲列表 线程安全问题，本地线程 TLAB thread local allocate buffer, cas + 失败重试， 分配空间大小在类加载的时候确定
    c. 初始化 0值 
    d. 设置对象头 ,class meta信息，hash,分代，偏向锁
    e. init