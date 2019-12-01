
以下为 sql 语句练习：

+----+---------+------+--------+
| id | name    | age  | gender |
+----+---------+------+--------+
|  1 | wangsan |   20 |      1 |
|  2 | lisi    |   18 |      1 |
|  3 | xiaohua |   22 |      0 |
+----+---------+------+--------+


```sql
show databases;
create database test;
use test;
show tables;
-- create table;
create table student(id int not null auto_increment,name varchar(10),age int default 18,primary key(id));
-- insert one record
insert into  student(name) values("wangsan");
-- limit
select * from student limit 1,1;
-- desc 降序 / asc 升序
select * from student order by id desc;
-- avg
select avg(age) from student;
-- count
select count(id) from student;
-- add col gender
alter table student add gender int(1);
-- group by
select gender as '性别', count(id) as '人数' from student group by gender;
select gender as '性别', count(id) as '人数' from student group by gender order by '人数' desc;
-- 
```