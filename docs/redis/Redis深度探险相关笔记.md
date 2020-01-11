
本文笔记 依赖 老钱的 Redis深度探险

# 基础

## string

最大长度 512M,小于 1M 加倍扩容 现有，大于 1M 扩容 1M

可以批量读写 

SDS 预分配 惰性回收

```
// 单次读写
set name houzhenguo
get name
// 批量读写
mset name1 zhangsan name2 lisi name3 wangwu
mget name1 name2 name3

// 过期时间 
expire name 5 // 5秒过期
// setnx

set age 5
incrby age -5
```

## list

相当于 Java 中的linkedlist 是链表而不是数组

用于做异步队列使用 , 使用 index  相关的操作要慎重，因为 会遍历链表。

快速列表：在数据元素较少的时候使用 ziplist ，分配的是一块连续的内存。

quicklist 就是 ziplist 结合连接起来。

```
rpush books java python
lpop books
```

## hash

字典，数组 + 链表 。 rehash 与 hashMap不同。

渐进式 rehash 。 保留新旧 两个 hash结构，查询时候会同时 查询两个 hash结构，而后在后续的定时
任务中以及 hash的子指令中，在后续的定时任务以及 hash 的子指令中，循序渐进的将 旧的 hash的内容
一点点迁移到新的 hash结构中。当hash 移除最后的元素之后，该数据结构会被自动回收。

hash 可以用来存储用户的信息，如果存储 字符串的话，反序列化会耗费网络流量。

```
hset books java "think in java"
hget books java
```

## Set

去重的功能。

## ZSet

字典，保证唯一性。 score ,内部跳跃表的实现。zset 可以用来存储粉丝列表，key为 id,score为 粉丝关注的事件

跳跃表的优点就是 比链表 查询比较快，定位。

随机分层。

## 容器的规则

1. 没有的时候创建，没有元素的时候回收。
2. 过期 以对象为单位。比如 一个 hash 结构的过期是整个 hash对象的过期，而不是其中的某个 子 key.

# 应用
## 分布式锁

重入锁的问题

锁超时的问题

我并不觉得 redis 实现分布式锁 好，建议还是 zk

获取锁失败 

1. 抛出异常
2. sleep  重试
3. 将请求转移到延迟队列，一会儿再试

## 延时队列

rpush lpop 实现 延迟队列

轮询，sleep 1 s 避免 pop 队列为空的时候一直轮询。 

替代方案： blpop/ brpop -> blocking pop

空先连接断开，注意 捕获异常重试。




## 位图

应用场景： 可以记录 每个用户 365 天签到的记录 ，1 为 签到 0 为 未签到
这样的存储 节省内存资源

## HyperLogLog  

统计 UV,用户访问量，可以使用 set，最后求 set的size。
但是对于 数据量特别大，几千万 上亿的时候，可能对于 UV的数值并不要求特别的精确，可以使用 pfadd

## BloomFilter

用户 读取过的新闻，推送的时候不能再次推送了。这个时候怎么做？

误判率。结果显示存在时候，它有可能不存在。结果显示不存在，它一定不存在。

需要安装插件。

爬虫 URL 的去重

## 限流

## 地图

附近的人 ，GeoHash算法。将 地球 映射到 二维平面，将 二维平面 编码成为 一维线性的。set score

地图的Redis 单独部署，最好是可以按照 城市划分。

## scan

可以有限制的匹配 某些key 

## PipLine

将 write read write read -> write write read  read

## Redis 事务

Redis 的事务 比较简单，不像 关系型 数据库那样。

也不支持原子性，中间发生错误之后还是会继续的执行。