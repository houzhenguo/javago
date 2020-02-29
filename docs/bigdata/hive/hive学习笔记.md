
视频地址： https://www.bilibili.com/video/av65556024

## Hive

Facebook 开源，解决海量 结构化日志的数据统计。

Hive 基于 Hadoop的一个数据仓库工具。将结构化的数据文件映射为一张表，提供类 SQL的查询。

本质： 将 HQL 转成 MapReduce程序。

![](../images/hive-1.png)

1. Hive 处理的数据存储在 HDFS
2. Hive 分析数据底层的实现是 MapReduce
3. 执行程序运行在 Yarn上

相当于 hadoop的客户端。

## Hive 优缺点
1. 操作接口采用 SQL 语法，提供快速开发能力，简单，容易上手
2. 避免去写 MapReduce ，减少开发人员学习成本
3. Hive 延迟比较高，常用于 数据分析，对实时性要求不高的场合
4. Hive的优势在于处理大数据，
5. Hive 支持用户自定义函数，可以根据自己的需求 来实现自己的函数

缺点：

1. 迭代式运算 无法表达
2. 数据挖掘方面不擅长
3. hive自动生成的MapReduce 作业，通常不够智能化
4. 调优比较困难

## 架构

![](../images/hive-2.png)

## 与传统sql数据库

1. 由于 SQL 被广泛的应用在数据仓库中，因此，专门针对 Hive 的特性设计了类 SQL 的
查询语言 HQL。熟悉 SQL 开发的开发者可以很方便的使用 Hive 进行开发。

2. Hive 是建立在 Hadoop 之上的，所有 Hive 的数据都是存储在 HDFS 中的。而数据库则
可以将数据保存在块设备或者本地文件系统中。

3. 由于 Hive 是针对数据仓库应用设计的，而数据仓库的内容是读多写少的。因此，Hive
中不建议对数据的改写，所有的数据都是在加载的时候确定好的。而数据库中的数据通常是
需 要 经 常 进 行 修 改 的

4.  Hive 中大多数查询的执行是通过 Hadoop 提供的 MapReduce 来实现的。而数据库通常
有自己的执行引擎。

5. Hive 在查询数据的时候，由于没有索引，需要扫描整个表，因此延迟较高。另外一个导
致 Hive 执行延迟高的因素是 MapReduce 框架。由于 MapReduce 本身具有较高的延迟，因此
在利用 MapReduce 执行 Hive 查询时，也会有较高的延迟

## Hive 安装地址

1. 官网地址   http://hive.apache.org/
2. 文档地址 https://cwiki.apache.org/confluence/display/Hive/GettingStarted
3. 下载地址 http://archive.apache.org/dist/hive/
4. github地址 https://github.com/apache/hive
