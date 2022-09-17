
## mysql
### index
1. 多路平衡二叉树
    a. 稳定 -> IO的次数固定 -> 都在叶子结点上
    b. 叶子结点存储 主键ID + Record
    c. 叶子结点是 顺序链表，并且靠近的ID都在一个磁盘块上
    d. 辅助索引存储的是主键
2. 回表
 select * from tab where id = 9; 这种不需要回表 -> 覆盖索引
 select * from tab where num = 9; 这种就需要回表，想不回表也可以，就要查询字段在索引上
3. 查询优化器 -> 会优化我们的sql
4. 索引太多-> 优化器 会对每个索引进行执行计划评估 -> 也会降低查询性能
5. 短索引 -> 构成索引的字段长度短，可以让单个存储块存储更多的索引值，提升IO效率
6. 左前缀 -> 没必要每个都建立索引
explain几种情况
system > const(唯一索引，非关联查询)> >eq_ref(唯一索引，关联查询)>ref(非唯一索引)> index(遍历索引树)< all(从硬盘读取)
key: 要使用的索引，null 就是没有使用
key_len ： 越短越好
extra: using filesort(order by 无法利用索引排序，外部文件排序，效率低)
       using temporary 临时表，对查询结果进行排序，效率低
       using index ，直接select 索引中就有想要的数据，不需要回表，效率不错
       using index condition 索引下推， 查找使用了索引，但是需要回表，查询列的问题
       using where 5.7 需要在server层过滤，就是 where 条件中有非索引列

7. >< ! 都会中断索引，索引列做计算，字符串不加'',select * 不会使用覆盖索引。or , like,
    有时候全表可能更快，这个看数据。

8. 很多情况下，我们并没有办法把using filesort优化为using index，只能退而求其次，尽量从filesort的角度去优化，通过外部条件。
 - order by的字段不是索引
 - order by 字段是索引字段，但是 select 中没有使用覆盖索引
 - order by 中同时存在 ASC 升序排序和 DESC 降序排序
 - order by中用到的是复合索引，但没有保持复合索引中字段的先后顺序（即违背了最左前缀原则）