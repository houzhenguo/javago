快速总结

## 数据库
1. 事务
ACID: atomicity, consistency,isolation,durability
原子性（最小单元，要么成功，要么失败），一致性（所有事务读取数据一致），隔离性（对其他事务不可见），持久性（永久保存到DB）
Mysql 自动提交事务，如果不显式指定。
2. 并发情况下DB是怎么处理的？
a. 表锁 行锁
b. 间隙锁 mvcc + 版本号
脏读（读取到别人没有提交的数据）
不可重复读， T1 之后修改的数据又被撤销了，T2就会读取到 T1撤销之前到数据。
幻读 -> where 条件一样，读取到数据不一样

产生并发情况的原因是 破坏了隔离性，通过控制并发来实现隔离性。
尽量少的加锁。
读写锁： 写锁，读锁 x,s
加了x锁之后不能再添加任何锁
加了s锁，其他事务可以加s锁，但是不能x
意向锁，获取 x/s锁之前都先获取一下对应的IX，IS锁，如果获取不到就获取锁失败。

隔离级别： read uncommited(脏), read commited（不可重复） ,repeateable read（幻）,serializable read

MVCC -> read commited/repeateable read ， 系统版本号，事务版本号
创建版本 （创建时候的系统版本），删除版本 < 系统版本（真的删除了）。
+ undo log,mvcc 把日志存在undo log中。
record 的快照 通过指针的形式存在 undo log中，我们主要通过版本号来读取最近的一个快照。举例。当前 T 读取，只能读取 《=T 事务版本号 + 删除版本号 > 自身的，这批数据才是有效的数据。

update的后，将当前版本号作为删除的版本号，创建新的快照，创建版本号，先delete -> insert

mvcc 在快照读的情况下可以解决幻读的问题，但是在 在repeateabel read 当前读的情况下，需要配合 next keys 一起解决 幻读的问题。

https://blog.csdn.net/Edwin_Hu/article/details/124392174
举例：
1. 快照读 where id > 3 ，当时 其中中途被T2事务插入了一条数据，T1 update where id = 6 导致 当前读，再次使用mvcc
机制的时候，出现了幻读。

select class from c group by class having count(distinct student) > 5
select email from c group by emails having count(*) >2
delete p1 from p1,p2 where p1.email = p2.email and p1.id > p2.id

## Btree
Balance Tree. 复杂度稳定，因为它所有的叶子结点都是主键ID，并且有顺序性，对于临近
查询可以预加载，旋转。
红黑树 ，调整频繁，度为2 ，就会比较深，log2n,
减少磁盘IO,顺序读不需要磁盘寻到。

索引可以将随机IO变成顺序IO.

explain: system>const>eq_ref>ref>range>index>all

## index 优点
顺序io，减少数据量扫描，覆盖索引减少回表，但是对于数据量非常小的表，全表扫描可能要比建立索引高效，所以分情况。索引也是占用空间的。
explain -> sample,union,

1. 查询的时候只返回需要的列，最好不要使用select * 
2. limit的使用，减少复杂查询已经 联表查询
3. 

sharding , 按照 itemID/ labelID 进行sharding.
目前来看，存储到tidb也不是特别好的方案，查询速度慢，所以可以采用 MySQL，
我们采用分表的方式进行查询。

双写，重新rehash,加版本号，我们之前是按照item/label维度，提前评估好数据量，做好定期的归档，
避免单表数据量过大，至于数据量有多少，看我们集器的性能。

事务问题，采用的是 单个操作的事务，使用同一个orm，
但是分布式的事务其实用的比较少，这个的保证比较困难，
可以采用重试机制和业务自己保证他的最终一致性。
sharding的问题：数据排查问题，

## binlog
master
1. binlog write main server to binary log
slave
2. io read binlog and write to relay log
3. sql thread 读取 relay log 
有延迟
master 写完binlog到本地， slave来读取binlog，
然后写入到自己的relaylog中，之后自己的sql thread
读取relaylog，在服务器中重放，
bus.

读写分离，增加冗余

使用代理。

## Redis

1. string,list,set,zset,hash
proxy, sharding,持久化，数据同步，扩展写

热点key,这种可以复制多份，
大key，可以把key 拆分成多份，维护 meta信息

redis集群主要考虑数据容灾稳定，以及数据的切换。
string: 字符串，自增自减
list: lpush,lpop
set
zset： score 对于排行可以这么操作
hash:用来存储 基本信息

set hello world
get hello
del hello
rpush list-key item
rpush list-key item2
rpush list-key item
lrange list-key 0 -1

hash: hash key name -> k,v
zset: zset key name -> k,score, 跳表

hash, 两个hash表，方便rehash操作，两个dict的。
rehash不是一次性操作完成，渐进式，避免 一次性阻塞

skip table 快速查找，有顺序，二分查找，数据式随机的，redis 的zset就是这样的。

有序链表，可以看成多个有序链表的组合。

插入速度快，不需要维护平衡性，更容易实现。

4. Redlock
5. set可以实现交集并集操作
现在用 redis比较多，不怎么使用memcache,mem只支持 string类型

redis 持久化 RDB，AOF,混合使用。
redis可以持久化，mem只是放在内存中。redis 4.0

数据淘汰策略： LRU，expire key,也不是到点就删除

hash这种只能给整个设置过期时间，无法给某个元素设置。

redis 数据淘汰策略：
a. 已经设置过期时间最少使用
b. xx 将要过期的
c. 任意数据
d. 所有数据最少使用
为了性能和内存的消耗考虑，allkeys -lfu
LFU,最近访问频次少，并且设置过期时间的。

AOF，append of file -> always, per second, no
RDB 快照 -> 定期备份 -> 复制到其他服务器 创建相同数据丢服务器副本
如果宕机，将丢失最后一次 创建快照之后的数据。
数据量大， 保存备份的时间太长。 slave 传递的就是整个快照，然后AOF
every second 兼顾效率和性能。
AOF重写

Reactor socket IO多路复用模型， 文件事件处理器， 命令请求处理器 命令回复处理器
连接应答处理器

主从同步， master RDB快照， 同步给 slave，如果这个时候快照太大，并且 命令写缓冲区已经满了，意味着数据没有同步完成，这个时候可能就要重新同步，这个是一个问题。
slave 接受到 快照数据之后，就会丢弃旧的数据，使用快照数据，
先发送快照数据，再发送 缓存的命令数据。

主从，不论式MySQL还是redis，其实都有可能面临主从同步的问题，占用master的资源，这时随着slave的数量增多，就需要创建中间层来分担主的压力。

哨兵

redis sharding 分片 proxy分片

zset, set. list,hash
ar:9999  title, content, link
set 存储点赞的人
zset 存储score,id

## java
1. bool 0/1 jvm 支持 bool数组，通过  byte数组实现
2. 包装类型，拆箱，装箱-》 -128 -127
Integer.valueOf, 缓存池，可以调节，但是一般不要动。除了new Integer以外，其他的都是
同一个对象。
AutoBoxCacheMax=size
3. string 比较特殊，final，不可变，不可被继承。
4. Java8 使用char数组存储数据，Java9用byte数组，并且标记哪中字符编码。
不变可以作为hash的可以，字符串常量池，heap,作为参数传递，没啥问题。
string thread safe
stringbuilder  stringbuffer -> 内部使用synchronized 进行同步

string pool, heap 字符串常量池

switch 
string ,底层hash equals
private protectd public

子类的访问级别不能低于 父类，防止 用父类对象不能访问子类。尽量private，减少对外界的暴露。

抽象类，接口。
接口定义行为，方法。
抽象类不能被实例化，需要子类继承抽象类才可以。

## object
1. hashCode/ toString /wait /notify /notify all /wait(time)

equals 

```java
public boolean equals(Object o) {
    if (this == o)
     return true
     if o == null || getClass != o.getClass
       return false;
     Exaple ex = (Example)o
     // 判断值是不是相等
}
```

重写equals方法的时候要重写 hashCOde

shallow copy 
deep copy

final 基本类型不能修改，引用类型不能修改，被引用的对象本身可以修改。

声明的方法不能被子类重写，
final的类不能被继承
static 静态变量，类变量，可以通过类名进行访问他
实例变量， 和实例同生存


父类 静态变量 静态语句块
子类 静态变量 静态语句块
父类 实例变量 普通语句块
父类 构造函数
子类 实例变量 普通语句块
子类 构造函数

## 反射
1. Class 对象， 包含类的有关信息，编译新类的时候，产生一个同名的.class文件，该
文件内存储的class对象
2. 类加载相当于class对象的加载，类再第一次使用的时候才会动态加载到jvm中，
也可以使用 class.forName 这种方式进行控制类的加载，该方法返回一个 class 对象
3. 反射是运行时的类信息，在运行的时候加载进来，编译期，.class 不存在也可以加载进来。

反射可以用来调试 和测试，动态扩展。全限定名。

性能开销，安全限制

io/net/out of bound/file/npe/socket/cast

##  泛型

<? extends T> 必须是T的子类
<? super T> 必须是T的父类 设定界限

类型擦除

## 集合
Set,List,Queue
SortSet -> TreeSet
ArrayList, LinkedList,Vector,
PriorityQueue

TreeSet 红黑树，旋转,支持有序的操作，但是 查找范围效率不如hashSet

hashset 其实是一个 hashmap,所以查找起来快。O1的时间复杂度，
用来去重比较优秀，go中没有这种集合帮忙去重

LinkedHashSet 有序 去重
LinkedList 只能顺序访问，所以遍历的时候必坑

LinkedList/LinkedHashSet/Vector

Map -> TreeMap, HashMap, LinkedHashMap,HashTable

ConcurrentHashMap
ConcurrentHashSet
TreeMap, HashMap

## 集合中的设计模式
1. interator 迭代器模式
2. interable <- collection list/set/queue

产生interator对象，stream 流。

迭代器模式，适配器模式。

## arraylist
1. RandomAccess
2. default cap = 10 / 1.5倍 指定 cap,避免频繁 迁移。
3. 删除元素的复杂度 On ,因为要把后面的移过去
4. fail-fast 在迭代的时候，不做删除操作， concurrentModifiedExe
5. transient 不会被序列化
vector -> synchronized -> 每次扩容 2倍
## copyonwriteArraylist
无锁的方式，读写分离，适合一些配置项，对于数据一致性要求不敏感的场景

## 并发
1. thread -> runable running wating terminate/stop/block
start -> runnable
running
sleep time waiting
wait waiting
lock/synchronized block
exception  terminated
runnable / running/ blocked/ wait/ time wait / terminated

## in
interruptedException,对于处于阻塞的thread可以抛出这个错误。
thread的错误不会 传递回 main

使用wait会释放锁，属于object

wait 等待， 其他线程调用notify， notify all
wait 挂起期间，会释放锁，不然其他线程就没办法进入，进而就没办法 notify释放锁

## CoutDownLatch
CyclicBarier 用来控制多个线程的互相等待，reset之后可以重新利用。
Semaphore 信号量，acquire

## BlockQueue
1. FIFO LinkedBlockQueue
2. ArrayBlockingQueue

## JVM
主要屏蔽不同操作系统之间的差异，在不同的平台下达到统一的效果。

主内存与工作内存。

处理器读取速度比内存快几个数量级，
缓存一致性高速缓存。

jvm保证 内存模型的原子性， read, load,use, asign,store,write,lock,unlock

volatile 通过内存屏障 禁止指令重排序。

## thread safe

1. 不可变，final,string,枚举
2. synchronized/Reetrantlock
3. cas /cow 非阻塞的同步

1. 不可变，2. 无锁机制 3. relock, sychronized cow,枚举

## sychronized
互斥同步，总是认为只要不去做措施，肯定会有问题。共享数据一定要加锁.

## 非阻塞同步
1. CAS 乐观并发策略,先进行操作,如果没有其他线程竞争共享数据,那么操作就成功了.重试知道成功.
compare and sway . 内存地址, 旧的预期值,新的值B,只有当旧的值等于A,才将V更新为B.
2. AutoInteger JUC

ABA的问题,通过 带有标记的原子引用类来解决这个问题,他是通过控制变量的版本号来实现的,不过大多数情况下,ABA不会影响程序的正确性,如果解决ABA的问题,改用传统的互斥可能比原子类更加高效.

栈封闭 -> 局部变量不会出现线程安全问题,因为他是属于线程私有的.

ThreadLocal,比如我们的日期.Web交互模型中,一个请求对应一个服务器线程的处理方式,这种处理方式的广泛应用是的Web服务端使用 Threadlocal 解决线程安全问题.

Thread对象中有个ThreadLocalMap的结构体.
可能存在内存泄露,尽可能在使用之后手动 remove,

## 锁优化
1. 自适应自旋锁 -> 根据上一次的旋转的次数和拥有者的状态
2. 锁消除 通过逃逸分析,发现堆上的数据并不能被其他线程访问到,所以把他们当作私有数据对待.
轻量级锁
1.  无锁 偏向锁 轻量级锁 重量级锁

锁对象头中 markword,标记 是否已经加锁,如果没有加锁,使用CAS修改 markword.
在 虚拟机栈中创建lock record ,将 对象 stack pointer指向自己的,如果指向自己的,则重入,没有问题,如果不是自己,那么直接膨胀为重量级锁.

## JVM
1. 程序计数器: 字节码指令的地址
2. 虚拟机栈: 局部变量, 方法的调用 从入栈到出战的过程栈帧,操作数栈,常量引用.
StackOverFlowError, OutOfMemoryError
3. 堆, GC对象分配的主要区域,也是垃圾回收的主要区域.OutOfMemoryError 异常
4. 本地方法栈
5. 方法区
6. 运行时常量池

## 系统设计
1. 响应时间，吞吐量，并发用户数

## 分布式事务
1. 2pc 单点/不一致/阻塞
2. 本地消息表 ， rocketmq
3. cap 
cp保证一致性
ap保证可用性
4. base 基本可用，最终一致性。

Raft协议， Follower，Candidate， Leader

## mq
1. 点对点模式
2. 发布订阅模式
异步 削峰 填谷， 解耦

## 设计
```java
private static final Singleton instance = new Singleton();
public static getInstance(){
  return instance;
}
// lazy
public class Singleton {
  // 防止指令重排序
  private volatile static Singleton instance;
  public static Singleton getInstance(){
    if (instance ==null) {
      sychronized(Singleton.class) {
        if (instance == null) {
          // 1. 分配内存空间
          // 2. 初始化
          // 3. 将对象指向内存分配的地址， 防止他jvm指令重排序
          instance = new Singleton();
        }
      }
    }
    return instance;
  }
}
// 枚举的方式，更加安全，也可以避免通过反射修改。
// 其他的方式可以通过 setAccessible 将私有的构造函数访问级别设置为public
// 初始化实例对象，
```
简单工厂类
一个接口，大家都实现一下，最后用接口接，每个类自己具体实现，在 方法中
可以通过 if判断输入的参数，来决定返回哪个子类。


```java
public String replaceStr(SringBuffer str) {
  int p1 = str.length() -1;
  for (int i=0;i<=p1;i++) {
    if (str.charAt(i) == ' ') {
      str.append("  ");
    }
  }
  int p2 = str.length() -1;
  while (p1>=0; p1<p2) {
    char c = str.charAt(p1--);
    if (c == ' ') {
      str.setCharAt(p2--, '0');
      str.setCharAt(p2--, '2');
      str.setCharAt(p2--, '%');
    }else {
      str.setCharAt(p2--, c);
    }
  }
  return str.toString();
}
```

## 逆序打印链表
```java
// 递归的方式,这种方式简单，比较好理解。
public ArrayList<Integer> printListFromTailToHead(ListNode node) {
  ArrayList<Integer> ret = new ArrayList<>();
  if (node != null) {
    ret.addAll(printListFromTailToHead(node.next));
    ret.add(node.val);
  }
  return ret;
}
// 使用栈的方式
public ArrayList<Integer> printListFromTailToHead(ListNode node) {
  Stack<Integer> stack = new Stack<>();
  while(node != null) {
    stack.add(node.val);
    node = node.next;
  }
  ArrayList<Integer> ret = new ArrayList<>();
  while(!stack.isEmpty()) {
    ret.add(stack.pop());
  }
  return ret;
}
```

## 重建二叉树
```java
// 缓存中序遍历数组每个值对应的索引
private Map<Integer,Integer> indexForInOrders = new HashMap<>();
public TreeNode reConstructBinaryTree(int[] pre, int[] in) {
    for (int i=0; i< in.length;i++) {
      indexForInOrders.put(in[i], i);
    }
    return reConstructBinTree();
}
private TreeNode reConstructBinTree(int[] pre, int preL, int preR, int inL) {
  if (preL > preR) {
    return null;
  }
  TreeNode root = new TreeNode(pre[preL]);
  int inindex = indexForInOrders(root.val);
  int leftSize = inIndex-preL;
  root.left = reConstructBinTree(pre, preL+1, preL+leftSize, inL);
  root.right = reConstructBinTree(pre,preL+1+leftsize,preR, inINdex+1);
  return root;
}
```
## 二叉树的下一个结点
中序遍历的下一个结点
```java
public TreeLinkNode GetNext(TreeLinkNode pNode) {
  if (pNode.right != null) {
    TreeLinkNode node = pNode.right;
    while (node.left != null) {
      node = node.left;
      return node;
    }
  } else {
    while (pnode.next != null) {
      TreeLinkNode parent = pNode.next;
      if (parent.left == pNode) {
        return parent;
      }
      pnode = pnode.next;
    }
  }
}
```

## 用两个栈实现队列
```java
Stack<Integer> in = new Stack<Integer>();
Stack<Integer> out = new Stack<Integer>();
public void push(int node) {
  in.push(node);
}
public int pop(){
  if (out.isEmpty()) {
    while(!in.isEmpty()){
      out.push(in.pop());
    }
  }
}
return out.pop();
```

## 删除链表
```java
// 1. 删除的是中间结点，直接让 next.value 赋值给toDelete.value
// 2. 删除的是 
public ListNode deleteNode(ListNode head, ListNode toBeDelete) {
  if (head == null || toBeDelete == null) {
    return null;
  }
  if toBeDelete.next != null {
    next = tobedelete.next;
    tobeDelete.val = next.val
    tobeDelete.next = next.next;
  }else {
    if tobedelete == head {
      head = null
    }
    else
     ListNode cur = head;
     while cur.next != tobedelete {
       cur = cur.next
     } 
     cur.next = null
  }
  return head;
}
```

## 偶数在前奇数在后

```java
public void reOrderArray(int nums) {
  int N = nums.length;
  for (int i = N-1; i>=0;i--) {
    for int j = 0; j<i;j++ {
      if isEven(nums[i]) && !isEven(nums[j+1]) {
        swap(nums, j, j+1)
      }
    }
  }
}
private boolean isEven(int x) {
  return x %2 == 0;
}
private void swap(int[] nums, int i, int j) {

}
```

## 链表中倒数第K个结点
双指针
```java
public ListNode FindKthToTail(ListNode head, int k) {
  if (head == null) {
    return null;
  }
  ListNode P1 = head;
  while (P1 != null && k--> 0) {
    P1 = P1.next;
  }
  if (k >0 ){
    return null;
  }
  ListNode P2 = head;
  while(P1 != null) {
    P1 = P1.next;
    P2 = P2.next;
  }
  return P2;
}
```

## 链表中环的入口结点
```java
// 双指针， 快慢指针
public ListNode find(ListNode pHead) {
  if (pHead == null || pHead.next == null) {
    return null;
  }
  ListNode fast = pHead;
  ListNode slow = pHead;
  do {
    fast = fast.next.next;
    slow = slow.next;
  }while(slow != fast);

}

// 
x + 2y +z = x_y
```

## 反转链表
```java
public ListNode ReverseList(ListNode head) {
  if (head == null || head.next == null) {
    return head;
  }
  ListNode next = head.next;
  head.next = null;
  ListNode newHead = ReverseList(next);
  next.next = head;
  return newHead;
}
```

## 合并两个排序链表
```java
public ListNode Merge(ListNode list1, ListNode list2) {
    if (list1 == null) {
      return list2;
    }
    if (list2 == null) {
      return list1;
    }
    if (list1.val < list2.val) {
        list1.next = Merge(list1.next, list2);
        return list1;
    } else {
        list2.next = Merge(list1, list2.next);
        return list2;
    }
}
```

## 判断是否是子树
```java
public boolean HasSubTree(TreeNode root1, TreeNode root2) {
  if (root1 == null || root2 == null) {
    return false;
  }
}
private boolean isSubtreeWithRoot(TreeNode r1, TreeNode r2){
  if r2 == null {
    return true;
  }
  if r1 == null {
    return false;
  }
  if (r1.val != r2.val) {
    return false;
  }
  return isSubtreeWithRoot(r1.left, r2.left) &&isSubtreeWithRoot(r1.right, r2.right);
}
``` 

## 二叉树镜像
```java
public void Mirror(TreeNode root) {
  if root  == null {
    return
  }
  swap(root);
  Mirror(root.left);
  Mirror(root.right);
}
private void swap(TreeNode root){
  TreeNode t = root.left;
  root.left = root.right;
  root.right = t;
}
```
## 对称二叉树
```java
public boolean isDuichen(TreeNode proot){
  if proot == null {
    return true;
  }
  return isDuichen(proot.left, proot.right);
}
private boolean isDuichen(TreeNode t1, TreeNode t2){
  if t1 ==null && t2 == null {
    return true;
  }
  if t1  || t2  ==null {
    return false;
  }
  if t1.val != t2.val {
    return false;
  }
  return duichen (t1.left, t2.right) && (t1.right, t2.left)
}
```
## 包含min函数的栈
```java
private Stack<Integer> dataStack = new Stack<>();
private Stack<Integer> minStack = new Stack<>();
public void push(int node) {
  dataStack.push(node);
  minStack.push(minStack.isEmpty()? node: math.min(minStack.peek, node));
}
public void pop() {
  dataStack.pop;
  minStack.pop;
}
public int top(){
  return dataStack.peek;
}
public int min(){
  return minStack.peek;
}
```

## 树的遍历
```java
public ArrayList<Integer> print(TreeNode root){
  Queue<TreeNode> queue = new LinkedList<>(;)
  ArrayList<Integer> ret = new ArrayList<>();
  queue.add(root);
  while (!queue.isempty) {
    int cnt = queue.size();
    while (cnt -- >0 ) {
      TreeNode t = queue.poll();
      if t == null {
        continue;
      }
    }
    ret.add(root.val);
    queue.add(root.left);
    queue.add(root.right);
  }
  return ret;
}
```
## 蛇形打印二叉树
```java
public ArrayList<ArrayList<Integer>> Print(TreeNode root) {
  ArrayList<ArrayList<Integer>> ret = new ArrayList<>();
  Queue<TreeNode> queue = new LinkedList<>();
  queue.add(root);
  boolean reverse = false;
  while (!queue.isempty) {
    ArrayList<Integer> list = new ArrayList<>();
    int cnt = queue.size;
    while (cnt-- >0) {
      TreeNode node = queue.poll;
      if node == null
        continue;
      list.add (node.val);
      queue.add left
      queue.add right
    }
    if reverse {
      collections.reverse list
    }
    reverse = !reverse;
    if list.size != 0 
       ret.add list
    return ret
  }
}
```