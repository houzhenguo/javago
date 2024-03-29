



## 什么是索引？

索引是一种能帮助 MySQL 提高查询效率的数据结构。

## 索引分别有哪些优点和缺点？

索引的优点如下：

- 快速放问数据表中的特定信息，提高检索速度。
- 创建唯一性索引，保证数据表中每一行数据的唯一性。
- 加速表与表之间的连接。
- 使用分组和排序进行数据检索时，可以显著减少查询中分组和排序的时间。

索引的缺点：

- 虽然提高了的查询速度，但却降低了更新表的速度，比如 update、insert，因为更新数据时，MySQL 不仅要更新数据，还要更新索引文件；
- 建立索引会占用磁盘文件的索引文件。

使用索引注意事项：

- 索引不会包含 NULL 值的列；
- 使用短索引，短索引不仅可以提高查询速度，更能节省磁盘空间和 I/O 操作；
- 索引列排序，MySQL 查询只使用一个索引，因此如果 where 子句中已经使用了索引的话，那么 order by 中的列是不会使用索引的，因此数据库默认排序可以符合要求的情况下，不要进行排序操作；尽量不要包含多个列的排序，如果需要最好给这些列创建复合索引；
- like 语句操作，一般情况下不鼓励使用 like 操作，如果非使用不可， 注意 like "%aaa%" 不会使用索引，而like "aaa%"可以使用索引；
- 不要在列上进行运算；
- 不适用 NOT IN 和 <> 操作

## 以下 SQL 有什么问题？该如何优化？
```sql
select * from t where f/2=100;
```
该 SQL 会导致引擎放弃索引而全表扫描，尽量避免在索引列上计算。可改为：
```sql
select * from t where f=100*2;
```

## 为什么 MySQL 官方建议使用自增主键作为表的主键？

因为自增主键是连续的，在插入过程中尽量减少页分裂，即使要进行页分裂，也只会分裂很少一部分；并且自增主键也能减少数据的移动，每次插入都是插入到最后，所以自增主键作为表的主键，对于表的操作来说性能是最高的。

## 自增主键有哪些优缺点?

优点：

- 数据存储空间很小；
- 性能最好；
- 减少页分裂。
缺点：

- 数据量过大，可能会超出自增长取值范围；
- 无法满足分布式存储，分库分表的情况下无法合并表；
- 主键有自增规律，容易被破解；

综上所述：是否需要使用自增主键，需要根据自己的业务场景来设计。如果是单表单库，则优先考虑自增主键，如果是分布式存储，分库分表，则需要考虑数据合并的业务场景来做数据库设计方案。

## 索引有几种类型？分别如何创建？

MySQL 的索引有两种分类方式：逻辑分类和物理分类。

 按照逻辑分类，索引可分为：

- 主键索引：一张表只能有一个主键索引，不允许重复、不允许为 NULL；
- 唯一索引：数据列不允许重复，允许为 NULL 值，一张表可有多个唯一索引，但是一个唯一索引只能包含一列，比如身份证号码、卡号等都可以作为唯一索引；
- 普通索引：一张表可以创建多个唯一索引，可以包含多列；
- 全文索引：让搜索关键词更高效的一种索引。

按照物理分类，索引可分为：

- 聚集索引：一般是表中的主键索引，如果表中没有显示指定主键，则会选择表中的第一个不允许为 NULL 的唯一索引，如果还是没有的话，就采用 Innodb 存储引擎为每行数据内置的 6 字节 ROWID 作为聚集索引。每张表只有一个聚集索引，因为聚集索引的键值的逻辑顺序决定了表中相应行的物理顺序。聚集索引在精确查找和范围查找方面有良好的性能表现（相比于普通索引和全表扫描），聚集索引就显得弥足珍贵，聚集索引选择还是要慎重的（一般不会让没有语义的自增 id 充当聚集索引）；
- 非聚集索引：该索引中索引的逻辑顺序与磁盘上行的物理存储顺序不同（非主键的那一列），一个表中可以拥有多个非聚集索引。
各种索引的创建脚本如下：

```sql
-- 创建主键索引
alter table t add primary key add (`id`);
-- 创建唯一索引
alter table t add unique (`username`);
-- 创建普通索引
alter table t add index index_name (`username`);
-- 创建全文索引
alter table t add fulltext (`username`);
```

## 主索引和唯一索引有什么区别？

- 主索引不能重复且不能为空，唯一索引不能重复，但可以为空；
- 一张表只能有一个主索引，但可以有多个唯一索引；
- 主索引的查询性能要高于唯一索引。

## 在 InnDB 中主键索引为什么比普通索引的查询性能高？

因为普通索引的查询会多执行一次检索操作。比如主键查询 select * from t where id=10 只需要搜索 id 的这棵 B+ 树，而普通索引查询 select * from t where f=3 会先查询 f 索引树，得到 id 的值之后再去搜索 id 的 B+ 树，因为多执行了一次检索，所以执行效率就比主键索引要低。

## 什么叫回表查询？

普通索引查询到主键索引后，回到主键索引树搜索的过程，我们称为回表查询。

参考SQL：

```sql
mysql> create table T(
id int primary key, 
k int not null, 
name varchar(16),
index (k))engine=InnoDB;
```

如果语句是 select * from T where ID=500，即主键查询方式，则只需要检索主键 ID 字段。

```sql
mysql>  select * from T where ID=500;
+-----+---+-------+
| id  | k | name  |
+-----+---+-------+
| 500 | 5 | name5 |
+-----+---+-------+
```

如果语句是 select * from T where k=5，即普通索引查询方式，则需要先搜索 k 索引树，得到 ID 的值为 500，再到 ID 索引树搜索一次，这个过程称为回表查询。
```sql
mysql> select * from T where k=5;
+-----+---+-------+
| id  | k | name  |
+-----+---+-------+
| 500 | 5 | name5 |
+-----+---+-------+
```
也就是说，基于非主键索引的查询需要多扫描一棵索引树。因此，我们在应用中应该尽量使用主键查询。

## 如何查询一张表的所有索引？

SHOW INDEX FROM T 查询表 T 所有索引。
show index from table;

## MySQL 最多可以创建多少个索引列？

MySQL 中最多可以创建 16 个索引列。

## 以下 like 查询会使用索引的是哪一个选项？为什么？

A.like '%A%' 

B.like '%A' 

C.like 'A%' 

D.以上都不是 

答：C 题目解析：like 查询要走索引，查询字符不能以通配符（%）开始。

## 如何让 like %abc 走索引查询？

我们知道如果要让 like 查询要走索引，查询字符不能以通配符（%）开始，如果要让 like %abc 也走索引，可以使用 REVERSE() 函数来创建一个函数索引，查询脚本如下：

```sql
select * from t where reverse(f) like reverse('%abc');
```

## MySQL 联合索引应该注意什么？

联合索引又叫复合索引，MySQL 中的联合索引，遵循最左匹配原则，比如，联合索引为 key(a,b,c)，则能触发索引的搜索组合是 a|ab|abc 这三种查询。

## 联合索引的作用是什么？

联合索引的作用如下：

- 用于多字段查询，比如，建了一个 key(a,b,c) 的联合索引，那么实际等于建了 key(a)、key(a,b)、key(a,b,c) 等三个索引，我们知道，每多一个索引，就会多一些写操作和占用磁盘空间的开销，尤其是对大数据量的表来说，这可以减少一部分不必要的开销；
- 覆盖索引，比如，对于联合索引 key(a,b,c) 来说，如果使用 SQL：select a,b,c from table where a=1 and b = 1 ，就可以直接通过遍历索引取得数据，而无需回表查询，这就减少了随机的 IO 操作，减少随机的 IO 操作，可以有效的提升数据库查询的性能，是非常重要的数据库优化手段之一；
- 索引列越多，通过索引筛选出的数据越少。

## 什么是最左匹配原则？它的生效原则有哪些？

最左匹配原则也叫最左前缀原则，是 MySQL 中的一个重要原则，说的是索引以最左边的为起点任何连续的索引都能匹配上，当遇到范围查询（>、<、between、like）就会停止匹配。 生效原则来看以下示例，比如表中有一个联合索引字段 index(a,b,c)：

- where a=1 只使用了索引 a；
- where a=1 and b=2 只使用了索引 a,b；
- where a=1 and b=2 and c=3 使用a,b,c；
- where b=1 or where c=1 不使用索引；
- where a=1 and c=3 只使用了索引 a；
- where a=3 and b like 'xx%' and c=3 只使用了索引 a,b。

## 列值为 NULL 时，查询会使用到索引吗？

在 MySQL 5.6 以上的 InnoDB 存储引擎会正常触发索引。但为了兼容低版本的 MySQL 和兼容其他数据库存储引擎，不建议使用 NULL 值来存储和查询数据，建议设置列为 NOT NULL，并设置一个默认值，比如 0 和空字符串等，如果是 datetime 类型，可以设置成 1970-01-01 00:00:00 这样的特殊值。

## 以下语句会走索引么？
```sql
select * from t where year(date)>2018;
```
不会，因为在索引列上涉及到了运算。

## 能否给手机号的前 6 位创建索引？如何创建？

可以，创建方式有两种：

- alter table t add index index_phone(phone(6));
- create index index_phone on t(phone(6));

## 什么是前缀索引？

前缀索引也叫局部索引，比如给身份证的前 10 位添加索引，类似这种给某列部分信息添加索引的方式叫做前缀索引。

## 为什么要用前缀索引？

前缀索引能有效减小索引文件的大小，让每个索引页可以保存更多的索引值，从而提高了索引查询的速度。但前缀索引也有它的缺点，不能在 order by 或者 group by 中触发前缀索引，也不能把它们用于覆盖索引。

## 什么情况下适合使用前缀索引？

当字符串本身可能比较长，而且前几个字符就开始不相同，适合使用前缀索引；相反情况下不适合使用前缀索引，比如，整个字段的长度为 20，索引选择性为 0.9，而我们对前 10 个字符建立前缀索引其选择性也只有 0.5，那么我们需要继续加大前缀字符的长度，但是这个时候前缀索引的优势已经不明显，就没有创建前缀索引的必要了。

## 什么是页？

页是计算机管理存储器的逻辑块，硬件及操作系统往往将主存和磁盘存储区分割为连续的大小相等的块，每个存储块称为一页。主存和磁盘以页为单位交换数据。数据库系统的设计者巧妙利用了磁盘预读原理，将一个节点的大小设为等于一个页，这样每个节点只需要一次磁盘 IO 就可以完全载入。

## 索引的常见存储算法有哪些？

- 哈希存储法：以 key、value 方式存储，把值存入数组中使用哈希值确认数据的位置，如果发生哈希冲突，使用链表存储数据；
- 有序数组存储法：按顺序存储，优点是可以使用二分法快速找到数据，缺点是更新效率，适合静态数据存储；
- 搜索树：以树的方式进行存储，查询性能好，更新速度快。

## InnoDB 为什么要使用 B+ 树，而不是 B 树、Hash、红黑树或二叉树？

因为 B 树、Hash、红黑树或二叉树存在以下问题：

- B 树：不管叶子节点还是非叶子节点，都会保存数据，这样导致在非叶子节点中能保存的指针数量变少（有些资料也称为扇出），指针少的情况下要保存大量数据，只能增加树的高度，导致IO操作变多，查询性能变低；
- Hash：虽然可以快速定位，但是没有顺序，IO 复杂度高；
- 二叉树：树的高度不均匀，不能自平衡，查找效率跟数据有关（树的高度），并且 IO 代价高；
- 红黑树：树的高度随着数据量增加而增加，IO 代价高。

## 为什么 InnoDB 要使用 B+ 树来存储索引？

B+Tree 中的 B 是 Balance，是平衡的意思，它在经典 B Tree 的基础上进行了优化，增加了顺序访问指针，在B+Tree 的每个叶子节点增加一个指向相邻叶子节点的指针，就形成了带有顺序访问指针的 B+Tree，这样就提高了区间访问性能：如果要查询 key 为从 18 到 49 的所有数据记录，当找到 18 后，只需顺着节点和指针顺序遍历就可以一次性访问到所有数据节点，极大提到了区间查询效率（无需返回上层父节点重复遍历查找减少 IO 操作）。

索引本身也很大，不可能全部存储在内存中，因此索引往往以索引文件的形式存储的磁盘上，这样的话，索引查找过程中就要产生磁盘 IO 消耗，相对于内存存取，IO 存取的消耗要高几个数量级，所以索引的结构组织要尽量减少查找过程中磁盘 IO 的存取次数，从而提升索引效率。 综合所述，InnDB 只有采取 B+ 树的数据结构存储索引，才能提供数据库整体的操作性能。

## 唯一索引和普通索引哪个性能更好？


- 对于查询来说两者都是从索引数进行查询，性能几乎没有任何区别；
- 对于更新操作来说，因为主键索引需要先将数据读取到内存，然后需要判断是否有冲突，因此比唯一所以要多了判断操作，所以性能就比普通索引性能要低。

## 优化器选择查询索引的影响因素有哪些？

优化器的目的是使用最小的代价选择最优的执行方案，影响优化器选择索引的因素如下：

- 扫描行数，扫描的行数越少，执行代价就越少，执行效率就会越高；
- 是否使用了临时表；
- 是否排序。

## MySQL 是如何判断索引扫描行数的多少？

MySQL 的扫描行数是通过索引统计列（cardinality）大致得到并且判断的，而索引统计列（cardinality）可以通过查询命令 show index 得到，索引扫描行数的多少就是通过这个值进行判断的。

## MySQL 是如何得到索引基数的？它准确吗？

MySQL 的索引基数并不准确，因为 MySQL 的索引基数是通过采样统计得到的，比如 InnoDb 默认会有 N 个数据页，采样统计会统计这些页面上的不同值得到一个平均值，然后除以这个索引的页面数就得到了这个索引基数。

## MySQL 如何指定查询的索引？

在 MySQL 中可以使用 force index 强行选择一个索引，具体查询语句如下：
```sql
select * from t force index(index_t)
```

## 在 MySQL 中指定了查询索引，为什么没有生效？

我们知道在 MySQL 中使用 force index 可以指定查询的索引，但并不是一定会生效，原因是 MySQL 会根据优化器自己选择索引，如果 force index 指定的索引出现在候选索引上，这个时候 MySQL 不会在判断扫描的行数的多少直接使用指定的索引，如果没在候选索引中，即使 force index 指定了索引也是不会生效的。

## 以下 or 查询有什么问题吗？该如何优化？

```sql
select * from t where num=10 or num=20;
```

答：如果使用 or 查询会使 MySQL 放弃索引而全表扫描，可以改为：

```sql
select * from t where num=10 union select * from t where num=20;
```

## 以下查询要如何优化？

表中包含索引：

- KEY mid (mid)
- KEY begintime (begintime)
- KEY dg (day,group)

使用以下 SQL 进行查询：

```sql
select f from t where day='2010-12-31' and group=18 and begintime<'2019-12-31 12:14:28' order by begintime limit 1;
```

答：此查询理论上是使用 dg 索引效率更高，通过 explain 可以对比查询扫描次数。由于使用了 order by begintime 则使查询放弃了 dg 索引，而使用 begintime 索引，从侧面印证 order by 关键字会影响查询使用索引，这时可以使查询强制使用索引，改为以下SQL：

```sql
select f from t use index(dg) where day='2010-12-31' and group=18 and begintime< '2019-12-31 12:14:28' order by begintime limit 1;
```

## 如何解决 MySQL 错选索引的问题？

- 删除错选的索引，只留下对的索引；
- 使用 force index 指定索引；
- 修改 SQL 查询语句引导 MySQL 使用我们期望的索引，比如把 order by b limit 1 改为 order by b,a limit 1 语义是相同的，但 MySQL 查询的时候会考虑使用 a 键上的索引。

## 如何优化身份证的索引？

在中国因为前 6 位代表的是地区，所以很多人的前六位都是相同的，如果我们使用前缀索引为 6 位的话，性能提升也并不是很明显，但如果设置的位数过长，那么占用的磁盘空间也越大，数据页能放下的索引值就越少，搜索效率也越低。针对这种情况优化方案有以下两种：

- 使用身份证倒序存储，这样设置前六位的意义就很大了；
- 使用 hash 值，新创建一个字段用于存储身份证的 hash 值。

## desc explain

```sql
desc select * from account where id =1;
explain select * from account where id =1;
mysql> explain select * from account where id =1;
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+
| id | select_type | table   | type  | possible_keys | key     | key_len | ref   | rows | Extra |
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+
|  1 | SIMPLE      | account | const | PRIMARY       | PRIMARY | 4       | const |    1 | NULL  |
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+

// 可以看到 以下没有使用索引的扫描行数 为 4行 
mysql> explain select * from account where id =1;
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+
| id | select_type | table   | type  | possible_keys | key     | key_len | ref   | rows | Extra |
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+
|  1 | SIMPLE      | account | const | PRIMARY       | PRIMARY | 4       | const |    1 | NULL  |
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+

```

![](../images/mysql-3.png)

如上图，咱们把 user_id + org_id 联合索引的数据结构理解成一棵树，MySQL查找时，从左往右去查找。且纵向的每列数据都是从小到大排序的，比如 
user_id 从上到下是排好序的。
接下来我们以SQL语句为例，看看它的性能是否高效。
1. select * from user where user_id = 1 and org_id = 10 【高效，能精确找到指定的行数】
2. select * from user where org_id = 11 【没法查找了，性能低，可能全表扫描。但针对类似Case，MySQL可能会做个优化，就
是把所有userId取出来，然后遍历下，因此扫描的数据行数可能等于userId的数量】
3. select * from user where user_id = 3 【上图user_id只有1和2两个值，没有3，所以这个SQL性能高效，一次就查出结果为空】
4. select * from user where user_id = 1 【会走索引，但返回数据的行数可能较多，需酌情处理】 
5. select * from user where org_id > 1  【性能低，全表扫描】
6. select * from user where user_id > 1  【会走索引，性能可能低】
7. select * from user where user_id != 1  【只会查等于或者大于小于，不等于就没法查了，所以全表扫描，性能低】




### 覆盖索引使用实例 

现在我创建了索引(username,age)，我们执行下面的 sql 语句

```sql
select username , age from user where username = 'Java' and age = 22
```

在查询数据的时候：要查询出的列在叶子节点都存在！所以，就不用回表。


## 查询优化器偷偷干了哪些事儿

1、如果建的索引顺序是 (a, b)。而查询的语句是 where b = 1 AND a = ‘陈哈哈’; 为什么还能利用到索引？

  理论上索引对顺序是敏感的，但是由于 MySQL 的查询优化器会自动调整 where 子句的条件顺序以使用适合的索引，所以 MySQL 不存在 where 子句的顺序问题而造成索引失效。当然了，SQL书写的好习惯要保持，这也能让其他同事更好地理解你的SQL。
