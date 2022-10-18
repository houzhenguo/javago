

<!-- TOC -->

- [select加锁分析](#select加锁分析)
    - [总结](#总结)
- [引言](#引言)
    - [问题](#问题)
- [正文](#正文)
    - [锁类型](#锁类型)
    - [加锁算法](#加锁算法)
    - [快照读和当前读](#快照读和当前读)
    - [分析](#分析)
        - [RC/RU+条件列非索引](#rcru条件列非索引)
        - [RC/RU+条件列是聚簇索引](#rcru条件列是聚簇索引)
        - [RC/RU+条件列是非聚簇索引](#rcru条件列是非聚簇索引)
        - [RR/Serializable+条件列非索引](#rrserializable条件列非索引)
        - [<font color='red'>RR/Serializable+条件列是聚簇索引（重要）</font>](#font-colorredrrserializable条件列是聚簇索引重要font)
        - [RR/Serializable+条件列是非聚簇索引](#rrserializable条件列是非聚簇索引)

<!-- /TOC -->

# select加锁分析

## 总结

1. 讨论加锁情况 需要 根据不同的事务隔离级别判断
2. innnoDB 一定存在 聚簇索引，默认主键作为 聚簇索引
3. 有几个索引，就有几棵B+树
4. 聚簇索引的叶子节点为磁盘上的真实数据。非聚簇索引的叶子节点除了保留值还指向聚簇索引。

5. S 共享锁 T2可读; X 排他锁 ；意向S锁，意向 X锁 ->对于表级锁请求检查
6. 加锁算法： 
    - RecordLocks : 行锁。 -> 索引记录 -> 最终落在 聚簇索引
    - Gap Locks: 间隙锁 -> RR Serializable 才有
    - Next-Key Locks: 行锁 + 间隙锁 
7. 快照读 当前读

    执行快照读，读的是 数据库的快照版本，是不加锁的。 -> 在 serial 中是一直加锁的。

8. RU/RC

    - 普通的都是 快照读不加锁
    - 最终落在 聚簇索引加 S，X锁
    - 对于 非聚簇索引 -》 会先在这个索引 + 聚簇索引加锁
9. RR/Serial

    - 需要多考虑间隙锁 -> 间隙锁就是锁 全表了
    - S,X 聚簇索引加 相应锁 ，还有间隙锁
    - = 等号的不加 间隙锁
    
    



# 引言

```sql
select * from table where id = ?
select * from table where id < ?
select * from table where id = ? lock in share mode
select * from table where id < ? lock in share mode
select * from table where id = ? for update
select * from table where id < ? for update
```

这六句sql在不同的事务隔离级别下，是否加锁，加的是共享锁还是排他锁，是否存在间隙锁。

## 问题

- 当前事务隔离级别是什么
- id列是否存在索引
- 如果存在索引是聚簇索引还是非聚簇索引呢？

# 正文

- innodb一定存在聚簇索引，默认以主键作为聚簇索引
- 有几个索引，就有几棵B+树(不考虑hash索引的情形)
- 聚簇索引的叶子节点为磁盘上的真实数据。非聚簇索引的叶子节点还是索引，指向聚簇索引B+树。

基础知识

## 锁类型

`共享锁`(S锁):假设事务T1对数据A加上共享锁，那么事务T2可以`读数据A`，`不能修改数据A`。

`排他锁`(X锁):假设事务T1对数据A加上排他锁，那么事务T2不能读数据A，不能修改数据A。 我们通过update、delete等语句加上的锁都是行级别的锁。只有LOCK TABLE … READ和LOCK TABLE … WRITE才能申请表级别的锁。

`意向共享锁`(IS锁):一个事务在获取（任何一行/或者全表）S锁之前，一定会先在所在的表上加IS锁。

`意向排他锁`(IX锁):一个事务在获取（任何一行/或者全表）X锁之前，一定会先在所在的表上加IX锁。

> 意向锁存在的目的: 假设事务T1，用X锁来锁住了表上的几条记录，那么此时表上存在IX锁，即意向排他锁。那么此时事务T2要进行`LOCK TABLE … WRITE`的表级别锁的请求，可以直接根据意向锁是否存在而判断是否有锁冲突。

## 加锁算法

1. `Record Locks`：简单翻译为`行锁`吧。注意了，该锁是<font color='red'>对索引记录进行加锁</font>！锁是在加索引上而不是行上的。注意了，innodb一定存在聚簇索引，因此行锁最终都会落到聚簇索引上！

2.`Gap Locks`：简单翻译为`间隙锁`，是对索引的间隙加锁，其目的只有一个，`防止其他事物插入数据`。在Read Committed隔离级别下，不会使用间隙锁。这里我对官网补充一下，隔离级别比Read Committed低的情况下，也不会使用间隙锁，如隔离级别为Read Uncommited时，也不存在间隙锁。当隔离级别为Repeatable Read和Serializable时，就会存在间隙锁。

3. `Next-Key Locks`：这个理解为Record Lock+索引前面的Gap Lock。记住了，锁住的是索引前面的间隙！比如一个索引包含值，10，11，13和20。那么，间隙锁的范围如下:<font color='red'>(不包括索引自己)</font>
```sql
(negative infinity, 10]
(10, 11]
(11, 13]
(13, 20]
(20, positive infinity)
```
## 快照读和当前读

在mysql中select分为快照读和当前读，执行下面的语句:
```sql
select * from table where id = ?;
```
执行的是`快照读`，读的是数据库记录的快照版本，是不加锁的。（这种说法在隔离级别为Serializable中不成立，后面我会补充。） 那么，执行
```sql
select * from table where id = ? lock in share mode; -- 会对读取记录加S锁 (共享锁)，执行
```

> 加的是表锁还是行锁呢？

针对这点，我们先回忆一下事务的四个隔离级别，他们由弱到强如下所示:

- `Read Uncommited(RU)`：读未提交，一个事务可以读到另一个事务未提交的数据！
- `Read Committed (RC)`：读已提交，一个事务可以读到另一个事务已提交的数据!
- `Repeatable Read (RR)`:可重复读，加入间隙锁，一定程度上避免了幻读的产生！注意了，只是一定程度上，并没有完全避免!我会在下一篇文章说明!另外就是记住从该级别才开始加入间隙锁(这句话记下来，后面有用到)!
- `Serializable`：串行化，该级别下读写串行化，且所有的select语句后都自动加上`lock in share mode`，即使用了共享锁。因此在该隔离级别下，使用的是当前读，而不是快照读。

为了便于说明，我来个例子，假设有表数据如下，pId为主键索引:

pId(int) |	name(varchar) |	num(int)
-|-|-
1 |	aaa |	100
2 |	bbb |	200
7 |	ccc |	200

执行语句(name列无索引)
```sql
select * from table where name = `aaa` for update
```

那么此时在pId=1,2,7这三条记录上存在行锁(把行锁住了)。另外，在(-∞,1)(1,2)(2,7)(7,+∞)上存在间隙锁(把间隙锁住了)。因此，给人一种整个表锁住的错觉！

注意我上面说的，之所以能够锁表，是 间隙锁来实现的。那么，`RU`和`RC`都不存在间隙锁，这种说法在`RU`和`RC`中还能成立么？
因此，该说法只在`RR`和`Serializable`中是成立的。如果隔离级别为`RU`和`RC`，无论条件列上是否有索引，都不会锁表，只锁行！

## 分析

下面来对开始的问题作出解答，假设有表如下，pId为主键索引:


pId(int)|	name(varchar)|	num(int)
-|-|-
1|	aaa	|100
2|	bbb|	200
3|	bbb	|300
7|	ccc	|200

---
### RC/RU+条件列非索引

(1)`select * from table where num = 200`

不加任何锁，是快照读。

(2)`select * from table where num > 200`

不加任何锁，是快照读。

(3)`select * from table where num = 200 lock in share mode`

当num = 200，有两条记录。这两条记录对应的pId=2，7，因此在pId=2，7的聚簇索引上加行级S锁，采用当前读。

(4)`select * from table where num > 200 lock in share mode`

当num > 200，有一条记录。这条记录对应的pId=3，因此在pId=3的聚簇索引上加上行级S锁，采用当前读。

(5)`select * from table where num = 200 for update`

当num = 200，有两条记录。这两条记录对应的pId=2，7，因此在pId=2，7的聚簇索引上加行级X锁，采用当前读。

(6)`select * from table where num > 200 for update`

当num > 200，有一条记录。这条记录对应的pId=3，因此在pId=3的聚簇索引上加上行级X锁，采用当前读

总结：
> `RU` 和 `RC` 模式下，普通的查询不加任何锁，是`快照读`。对于加锁的情况，分配 S,X ,对应相应的聚簇索引上加锁，采用的都是`当前读`。

### RC/RU+条件列是聚簇索引
大家应该知道pId是主键列，因此pId用的就是聚簇索引。此情况其实和RC/RU+条件列非索引情况是类似的。

(1)`select * from table where pId = 2`

不加任何锁，是快照读。

(2)`select * from table where pId > 2`

不加任何锁，是快照读。

(3)`select * from table where pId = 2 lock in share mode`

在pId=2的聚簇索引上，加S锁，为当前读。

(4)`select * from table where pId > 2 lock in share mode`

在pId=3，7的聚簇索引上，加S锁，为当前读。

(5)`select * from table where pId = 2 for update`

在pId=2的聚簇索引上，加X锁，为当前读。

(6)`select * from table where pId > 2 for update`

在pId=3，7的聚簇索引上，加X锁，为当前读。

> 为什么条件列加不加索引，加锁情况是一样的？

ok,其实是不一样的。在RC/RU隔离级别中，MySQL Server做了优化。在条件列没有索引的情况下，尽管通过聚簇索引来扫描全表，进行全表加锁。但是，MySQL Server层会进行过滤并把不符合条件的锁当即释放掉，因此你看起来最终结果是一样的。但是<font color='red'>RC/RU+条件列非索引比本例多了一个释放不符合条件的锁的过程</font>！

### RC/RU+条件列是非聚簇索引

我们在num列上建上非唯一索引。此时有一棵聚簇索引(主键索引，pId)形成的B+索引树，其叶子节点为硬盘上的真实数据。以及另一棵非聚簇索引(非唯一索引，num)形成的B+索引树，其叶子节点依然为索引节点，保存了num列的字段值，和对应的聚簇索引。
接下来分析开始


(1)`select * from table where num = 200`

不加任何锁，是快照读。

(2)`select * from table where num > 200`

不加任何锁，是快照读。

(3)`select * from table where num = 200 lock in share mode`

当num = 200，由于num列上有索引，因此先在 num =

200的两条索引记录上加行级S锁。接着，去聚簇索引树上查询，这两条记录对应的pId=2，7，因此在pId=2，7的聚簇索引上加行级S锁，采用当前读。

(4)`select * from table where num > 200 lock in share mode`

当num > 200，由于num列上有索引，因此先在符合条件的 num = 300的一条索引记录上加行级S锁。接着，去聚簇索引树上查询，这条记录对应的pId=3，因此在pId=3的聚簇索引上加行级S锁，采用当前读。

(5)`select * from table where num = 200 for update`

当num = 200，由于num列上有索引，因此先在 num = 200的两条索引记录上加行级X锁。接着，去聚簇索引树上查询，这两条记录对应的pId=2，7，因此在pId=2，7的聚簇索引上加行级X锁，采用当前读。

(6)`select * from table where num > 200 for update`
当num > 200，由于num列上有索引，因此先在符合条件的 num = 300的一条索引记录上加行级X锁。接着，去聚簇索引树上查询，这条记录对应的pId=3，因此在pId=3的聚簇索引上加行级X锁，采用当前读。

总结：
> 整体来看 `RU` 和 `RC` 都是对 聚簇索引进行的加锁，最多就是对 非聚簇索引也加了一次锁，但是对于不加锁的情况，都是 快照读。所有加锁的情况都是 当前读。

---

### RR/Serializable+条件列非索引

RR级别需要多考虑的就是<font color ='red'>gap lock，他的加锁特征在于，无论你怎么查都是锁全表</font>。如下所示 接下来分析开始:

(1)`select * from table where num = 200`

在RR级别下，不加任何锁，是快照读。
在Serializable级别下，在pId = 1,2,3,7（全表所有记录）的聚簇索引上加S锁。并且在
聚簇索引的所有间隙(-∞,1)(1,2)(2,3)(3,7)(7,+∞)加gap lock

(2)`select * from table where num > 200`

在RR级别下，不加任何锁，是快照读。
在Serializable级别下，在pId = 1,2,3,7（全表所有记录）的聚簇索引上加S锁。并且在
聚簇索引的所有间隙(-∞,1)(1,2)(2,3)(3,7)(7,+∞)加gap lock

(3)`select * from table where num = 200 lock in share mode`

在pId = 1,2,3,7（全表所有记录）的聚簇索引上加S锁。并且在
聚簇索引的所有间隙(-∞,1)(1,2)(2,3)(3,7)(7,+∞)加gap lock

(4)`select * from table where num > 200 lock in share mode`

在pId = 1,2,3,7（全表所有记录）的聚簇索引上加S锁。并且在
聚簇索引的所有间隙(-∞,1)(1,2)(2,3)(3,7)(7,+∞)加gap lock

(5)`select * from table where num = 200 for update`
在pId = 1,2,3,7（全表所有记录）的聚簇索引上加X锁。并且在
聚簇索引的所有间隙(-∞,1)(1,2)(2,3)(3,7)(7,+∞)加gap lock

(6)`select * from table where num > 200 for update`

在pId = 1,2,3,7（全表所有记录）的聚簇索引上加X锁。并且在
聚簇索引的所有间隙(-∞,1)(1,2)(2,3)(3,7)(7,+∞)加gap lock

总结
> 虽然条件列非索引，但是最红加锁还是  落在了 聚簇索引上。区别就是 显式加锁的时候 会加上间隙锁，其他的普通
查询都是不加锁的。<font color='red'> （不加锁这个东西应该是与 聚簇索引以下最大的区别）</font>

### <font color='red'>RR/Serializable+条件列是聚簇索引（重要）</font>
<font size = 25 color='red'> 全局最重要的部分</font>

大家应该知道pId是主键列，因此pId用的就是聚簇索引。该情况的加锁特征在于，`如果where后的条件为精确查询(=的情况)，那么`<font color='red'>只存在</font>`record lock。如果where后的条件为范围查询(>或<的情况)，那么存在的是record lock+gap lock。`

(1)`select * from table where pId = 2`

在RR级别下，不加任何锁，是快照读。
在Serializable级别下，是当前读，在pId=2的聚簇索引上加S锁，不存在gap lock。

(2)`select * from table where pId > 2`

在RR级别下，不加任何锁，是快照读。
在Serializable级别下，是当前读，在pId=3,7的聚簇索引上加S锁。在(2,3)(3,7)(7,+∞)加上gap lock

(3)`select * from table where pId = 2 lock in share mode`

是当前读，在pId=2的聚簇索引上加S锁，不存在gap lock。

(4)`select * from table where pId > 2 lock in share mode`

是当前读，在pId=3,7的聚簇索引上加S锁。在(2,3)(3,7)(7,+∞)加上gap lock

(5)`select * from table where pId = 2 for update`

是当前读，在pId=2的聚簇索引上加X锁。

(6)`select * from table where pId > 2 for update`

在pId=3,7的聚簇索引上加X锁。在(2,3)(3,7)(7,+∞)加上gap lock

(7)`select * from table where pId = 6 [lock in share mode|for update]`

注意了，pId=6是不存在的列，这种情况会在(3,7)上加gap lock。

(8)`select * from table where pId > 18 [lock in share mode|for update]`

注意了，pId>18，查询结果是空的。在这种情况下，是在(7,+∞)上加gap lock。

### RR/Serializable+条件列是非聚簇索引

这里非聚簇索引，需要区分是否为唯一索引。因为如果是非唯一索引，间隙锁的加锁方式是有区别的。
先说一下，唯一索引的情况。如果是唯一索引，情况和RR/Serializable+条件列是聚簇索引类似，唯一有区别的是:这个时候有两棵索引树，加锁是加在对应的非聚簇索引树和聚簇索引树上！大家可以自行推敲!
下面说一下，非聚簇索引是非唯一索引的情况，他和唯一索引的区别就是通过索引进行精确查询以后，不仅存在record lock，还存在gap lock。而通过唯一索引进行精确查询后，只存在record lock，不存在gap lock。老规矩在num列建立非唯一索引

(1)`select * from table where num = 200`

在RR级别下，不加任何锁，是快照读。
在Serializable级别下，是当前读，在pId=2，7的聚簇索引上加S锁，在num=200的非聚集索引上加S锁，在(100,200)
(200,300)加上gap lock。

(2)`select * from table where num > 200`

在RR级别下，不加任何锁，是快照读。
在Serializable级别下，是当前读，在pId=3的聚簇索引上加S锁，在num=300的非聚集索引上加S锁。在(200,300)(300,+∞)加上gap lock

(3)`select * from table where num = 200 lock in share mode`

是当前读，在pId=2，7的聚簇索引上加S锁，在num=200的非聚集索引上加S锁，在(100,200)(200,300)加上gap lock。

(4)`select * from table where num > 200 lock in share mode`

是当前读，在pId=3的聚簇索引上加S锁，在num=300的非聚集索引上加S锁。在(200,300)(300,+∞)加上gap lock。

(5)`select * from table where num = 200 for update`

是当前读，在pId=2，7的聚簇索引上加S锁，在num=200的非聚集索引上加X锁，在(100,200)(200,300)加上gap lock。

(6)`select * from table where num > 200 for update`

是当前读，在pId=3的聚簇索引上加S锁，在num=300的非聚集索引上加X锁。在(200,300)(300,+∞)加上gap lock

(7)`select * from table where num = 250 [lock in share mode|for update]`

注意了，num=250是不存在的列，这种情况会在(200,300)上加gap lock。

(8)`select * from table where num > 400 [lock in share mode|for update]`

注意了，pId>400，查询结果是空的。在这种情况下，是在(400,+∞)上加gap lock。


[参考链接](https://juejin.im/post/5d5671a2e51d45620821cea7) 原作者： 孤独烟
[MySql死锁](https://www.cnblogs.com/zejin2008/p/5262751.html)





