https://docs.pingcap.com/zh/tidb/v4.0/overview  TIDB 中文官网


https://pingcap.com/blog-cn/tidb-internal-1 三篇文章了解 TiDB 技术内幕 - 说存储

https://pingcap.com/blog-cn/tidb-internal-2 三篇文章了解 TiDB 技术内幕 - 说计算

https://pingcap.com/blog-cn/tidb-internal-3 三篇文章了解 TiDB 技术内幕 - 谈调度

https://docs.pingcap.com/zh/tidb/v4.0/tidb-architecture  TIDB整体架构

https://tech.meituan.com/2018/11/22/mysql-pingcap-practice.html 美团技术团队

https://juejin.cn/post/6844903622598197256 

## 特点
1. 存储数据量大
2. 写的快 LSM-Tree
3. 写多读少
4. 计算和存储分离，可以扩容缩容
5. raft + replica -> 存储靠谱
6. 

## 数据存储

1. 如何持久化落盘 -> 备份
2. 如何保持一致性 -> 
3. 如何跨数据中心
4. 如何写入速度快
5. 如何方便读取
6. 如何修改数据？如何支持并发修改数据
7. 如何原子修改多条记录

## k-v
1. Key_Value 存储模型，提供有序便利方法。k,v都是原始 Byte数组
2. 排序就是考 key的二进制进行的，
3. 底层使用 RocksDB（FB） -> 简单理解成为一个kv 
## Raft
保证数据复制的一致性，不丢失。 -> 一致性算法  TiKV 源码解析系列 - Raft 的优化 https://zhuanlan.zhihu.com/p/25735592

1. Leader选举
2. 成员变更
3. 日志复制

## region
1. 将 key 不是采用hash 的形式，而是将 kv空间划分成为很多段，每个段是一系列连续的key-> 我们把每段叫做一个region -> 有大小限制 -> 有[startKey,endKey) 
    来描述。
2. 将 region 作为基本单位，散列在所有的节点上 -> 保证每个节点上region数量差不多。
3.  以region 为单位做 Raft 的复制和成员管理。

4. 一个region 会有多个 replica 它们会保存在不同的节点上，之间通过 raft保证一致性。
5. region + replica = raft group 。leader 负责读写，再由 Leader同步给Follower

## MVCC
1. 对于同一个 Key 的多个版本，我们把版本号较大的放在前面，版本号小的放在后面
2. 当用户通过一个 Key + Version 来获取 Value 的时候，可以将 Key 和 Version 构造出 MVCC 的 Key，也就是 Key-Version。然后可以直接 Seek(Key-Version)，定位到第一个大于等于这个 Key-Version 的位置。


# 计算
1. TIDB 主要面向 OLTP 业务，可以快速的读取，保存，修改，删除 一行数据。

## 查询
2. 一个全局有序的分布式 Key-Value 引擎 

TiDB 对每个表分配一个 TableID，每一个索引都会分配一个 IndexID，每一行分配一个 RowID（如果表有整数型的 Primary Key，那么会用 Primary Key 的值当做 RowID），其中 TableID 在整个集群内唯一，IndexID/RowID 在表内唯一，这些 ID 都是 int64 类型。

编码格式： table_[tableId]_index[indexId]_rowId[rowId] 组成一个唯一的key ，并且通过编码的形式尽量保证 顺序。

举个例子：
三行记录
```
1, "TiDB", "SQL Layer", 10
2, "TiKV", "KV Engine", 20
3, "PD", "Manager", 30
```
那么首先每行数据都会映射为一个 Key-Value pair，注意这个表有一个 Int 类型的 Primary Key，所以 RowID 的值即为这个 Primary Key 的值。假设这个表的 Table ID 为 10，其 Row 的数据为：

```
t10_r1 --> ["TiDB", "SQL Layer", 10]
t10_r2 --> ["TiKV", "KV Engine", 20]
t10_r3 --> ["PD", "Manager", 30]
```

除了 Primary Key 之外，这个表还有一个 Index，假设这个 Index 的 ID 为 1，则其数据为：

```
t10_i1_10_1 --> null
t10_i1_20_2 --> null
t10_i1_30_3 --> null
```

`问题：如果有多个索引怎么办？`

## 数据库，表 元信息怎么存储

key 前面 + m (meta) ， value 就是序列化之后的信息


# 调度
- 作为一个分布式高可用存储系统，必须满足的需求，包括四种：

副本数量不能多也不能少
副本需要分布在不同的机器上
新加节点后，可以将其他节点上的副本迁移过来
节点下线后，需要将该节点的数据迁移走
作为一个良好的分布式系统，需要优化的地方，包括：

- 维持整个集群的 Leader 分布均匀
维持每个节点的储存容量均匀
维持访问热点分布均匀
控制 Balance 的速度，避免影响在线服务
管理节点状态，包括手动上线/下线节点，以及自动下线失效节点

每个 TiKV 节点会定期向 PD 汇报节点的整体信息 -> 心跳包中携带

- 总磁盘容量
- 可用磁盘容量
- 承载的 Region 数量
- 数据写入速度
- 发送/接受的 Snapshot 数量（Replica 之间可能会通过 Snapshot 同步数据）
- 是否过载
- 标签信息（标签是具备层级关系的一系列 Tag）


leader 也会想 PD汇报

- Leader 的位置
- Followers 的位置
- 掉线 Replica 的个数
- 数据写入/读取的速度
