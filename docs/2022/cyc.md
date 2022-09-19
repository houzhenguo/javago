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