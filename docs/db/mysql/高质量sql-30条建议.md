
1. 不要是用 select * ,尽量使用 具体的字段名字
2. 如果知道查询的结果只有 一条记录，建议使用 limit

```sql
select id,name from test where name='hou' limit 1;
```

- 加上limit 1后,只要找到了对应的一条记录,就不会继续向下扫描了,效率将会大大提高。
- 当然，如果name是唯一索引的话，是不必要加上limit 1了，因为limit的存在主要就是为了防止全表扫描，从而提高性能,如果一个语句本身可以预知不用全表扫描，有没有limit ，性能的差别并不大。

3. 应尽量避免在where子句中使用or来连接条件

反例： select * from user where userId = 1 or age =18;

使用 union all
```sql
select * from user where userId=1
union all
select * from user where age =18;

// 或者分开写 SQL

```
使用or可能会使索引失效，从而全表扫描。

> 对于or+没有索引的age这种情况，假设它走了userId的索引，但是走到age查询条件时，它还得全表扫描，也就是需要三步过程：全表扫描+索引扫描+合并 如果它一开始就走全表扫描，直接一遍扫描就完事。mysql是有优化器的，处于效率与成本考虑，遇到or条件，索引可能失效，看起来也合情合理。

4. 优化limit分页

分页的使用limit，当 偏移量特别大的时候，查询效率变得特别低下。

反例
```sql
select id,name,age from employee limit 10000,10;
```

正例：
```sql
// 方案1：
select id,name from user where id>1000 limit 10;
// 方案2 order by+ 索引
select id,name from user order by id limit 1000,10;
// 方案3： 在业务允许的方位 限制页数
```
理由：

- 当偏移量最大的时候，查询效率就会越低，因为Mysql并非是跳过偏移量直接去取后面的数据，而是先把偏移量+要取的条数，然后再把前面偏移量这一段的数据抛弃掉再返回的。
- 如果使用优化方案一，返回上次最大查询记录（偏移量），这样可以跳过偏移量，效率提升不少。
- 方案二使用order by+索引，也是可以提高查询效率的。
- 方案三的话，建议跟业务讨论，有没有必要查这么后的分页啦。因为绝大多数用户都不会往后翻太多页。

5. 优化 like语句

反例：

```sql
select userId,name from user where userId like `%123`;
```

正例：
```sql
select userId,name from user where userId like `%123`;
```

以下是自己使用 explain的命令做了几个测试：

```
mysql> explain select * from account where id like '1%'
    -> ;
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+
| id | select_type | table   | type | possible_keys | key  | key_len | ref  | rows | Extra       |
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+
|  1 | SIMPLE      | account | ALL  | PRIMARY       | NULL | NULL    | NULL |    6 | Using where |
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+
1 row in set (0.00 sec)

mysql> explain select * from account where id like '%1'
    -> ;
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+
| id | select_type | table   | type | possible_keys | key  | key_len | ref  | rows | Extra       |
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+
|  1 | SIMPLE      | account | ALL  | NULL          | NULL | NULL    | NULL |    6 | Using where |
+----+-------------+---------+------+---------------+------+---------+------+------+-------------+
1 row in set (0.00 sec)

mysql> explain select * from account where id=1;
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+
| id | select_type | table   | type  | possible_keys | key     | key_len | ref   | rows | Extra |
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+
|  1 | SIMPLE      | account | const | PRIMARY       | PRIMARY | 4       | const |    1 | NULL  |
+----+-------------+---------+-------+---------------+---------+---------+-------+------+-------+
1 row in set (0.00 sec)


```

7. 尽量避免再 索引列上使用 MySQL的内置函数

反例：

```sql
select userId,loginTime 
from
 loginuser 
where
Date_ADD(
loginTime,Interval 7 DAY)>=now();
```

8. 尽量避免在 where子句中对字段进行表达式操作，会导致放弃索引进行全表扫描。



参考： https://mp.weixin.qq.com/s/Gid96Ivb0I3yGAiieyKRdg