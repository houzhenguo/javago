## 开场
1. 今天给大家分享一下pika 相关的知识。

## 背景知识
1. pika 是 360 公司开源的一款类似 redis 的存储系统，它主要用来解决对于大数据量redis内存不足的瓶颈，底层的存储系统是
rocketdb，一款高性能的存储db。和tidb 底层使用的相同。
2. 既然是 类似 redis,那么它要兼容redis 大多数的协议，目前对于 redis 中的string, hash,list,zset,set 等大多数
的接口pika都会有自己的实现，后面我会大概讲一下 pika是怎么在一个纯kvdb上实现这几种数据结构的。
3. 我们公司的cp 底层也是使用的pika,TODO (怎么使用的)

## pika vs redis
pika 解决的是redis 内存不足的问题，它是使用磁盘存储。相比reids 的内存存储，它有以下几个优点：
1. 容量大,Pika没有Redis的内存限制, 最大使用空间等于磁盘空间的大小
    如果我们的业务场景 超过了50G,可以使用pika
2. 加载db速度快，pika 的数据会进行持久化，不是redis rdb和aof，pika 重启不需要加载所有的数据到内存，
, 不需要进行回放数据操作。rocksdb 启动不需要加载全部数据, 只需要加载几M的log 文件就可以启动, 因此恢复时间非常快
> 大容量redis 遇到的问题
线上的redis 一般同时开启rdb 和 aof. 我们知道aof的作用是实时的记录用户的写入操作, rdb 是redis 某一时刻数据的完整快照. 那么恢复的时候一般是通过 rdb + aof 的方式进行恢复, 根据我们线上的情况 50G redis 恢复时间需要差不多70分钟

3. 备份速度快 拷贝数据文件后还有一个快照的恢复过程，会花费一些时间），这样在对于百G大库的备份是快捷的，更快的备份速度更好的解决了主从的全同步问题。

为了`防止同步缓冲区被复写`，dba给redis设置了2G的巨大同步缓冲区，这对于内存资源来讲代价很大. 当由于机房之间网络有故障, 主从同步出现延迟了大于2G以后, 就会触发全同步的过程. 如果多个从库同时触发全同步的过程, 那么很容易就将主库给拖死

劣势当然也比较明显，性能比较低，官方统计，pika 的性能是redis 50%,但是可以使用SSD存放数据，尽可能 跟上Redis性能。


## 整体结构
1. 接下来介绍一下整体的结构，pika 主要由 4个模块组成。
    a. 网络模块pink
    b. 线程模块
    c. 存储引擎nemo
    d. 日志模块binlog



2. pink 网络模块也是 360基础架构团队开元大网络编程库，支持pb,redis 协议，对网络进行封装，单线程，多线程网络模型
3. 线程模块 之后再进行讲解

## 架构


## 线程模型
接下来讲解一下 pika的线程模型，他使用的线程类型比较多。
1. DispatchThread 它的主要职能就是 accept 客户端的socket 连接，将 connect 封装之后 push 到 
workThread 的一个notfity_queue 队列之中，而后通知 workThread进行处理。
[Pika传火计划之线程模型](https://whoiami.github.io/PIKA_THREAD_MODEL)
2. Thread线程池 执行客户端发过来的指令完成之后，由worker 进行回复消息，work 就是收发消息的，而中间还有个
thredpool 只是处理消息，不涉及接收和回复的IO操作。
线程池中的线程数量由用户配置，执行WorkerThread调度过来的Task, Task的内容主要是写DB和写Binlog

- PikaAuxiliaryThread：辅助线程，处理同步过程中状态机状态的切换，主从之间心跳的发送以及超时检查
- PikaReplClient：本质上是一个Epoll线程(与其他Pika实例的PikaReplServer进行通信)加上一个由若干线程组成的线程数组(异步的处理写Binlog以及写DB的任务)
- PikaReplServer：本质上是一个Epoll线程(与其他Pika实例的PikaReplClient进行通信)加上一个由若干线程组成的线程池(处理同步的请求以及根据从库返回的Ack更新Binlog滑窗)
- MonitorThread：执行了Monitor命令的客户端会被分配在这个线程上，这个线程将目前Pika正在处理的命令返回给挂在这个线程上的客户端
- KeyScanThread：在这个线程中执行info keyspace 1触发的统计Key数量的任务
- BgSaveThread：对指定的DB进行Dump操作，以及全同步的时候发送Dump数据到从库（对一个DB执行全同步是先后向Thread中扔了BgSave以及DBSync两个任务从而保证顺序)
- PurgeThread：用于清理过期的Binlog文件
- PubSubThread：用于支持PubSub相关功能
  
FAQ:
pika有一些比较耗时的任务，PurgeThread 如删binlog，扫描key，备份，同步数据文件等等，为了不影响正常的用户请求，这些任务都是放到后台执行的，并且将能并行的都放到不同线程里来最大程度上提升后台任务的执行速度；

## 存储引擎
1. 接下来介绍一下pika 底层的存储引擎。
pika 底层使用的是 一个叫blackwidow的存储引擎，它本质是对 rocksdb的改造和封装，使其支持多数据结构的存储，rocksdb只支持kv存储,之前的版本是
2. blackwidow支持五种数据结构类型的存储：KV键值对、Hash结构、List结构、Set结构和ZSet结构。因为rocksdb的存储方式只有kv一种结构，所以以上所说的5种数据结构的存储最终都要落盘到rocksdb的kv存储方式上。

介绍一下 blackwidow 是怎么进行数据结构转换的。
### kv存储
String本质上就是Key, Value, 我们知道Rocksdb本身就是支持kv存储的， 我们为了实现Redis中的expire功能，所以在value后面添加了4 Bytes用于存储timestamp, 作为最后Rocksdb落盘的kv格式，下面是具体的实现方式:

![](https://i.imgur.com/KnA707a.png)

`如果我们没有对该String对象设置超时时间，则timestamp存储的值就是默认值0， 否则就是该对象过期时间的时间戳， 每次我们获取一个String对象的时候， 首先会解析Value部分的后四字节， 获取到timestamp做出判断之后再返回结果。`


### hash 存储

blackwidow中的hash表由两部分构成，元数据(meta_key, meta_value), 和普通数据(data_key, data_value), 元数据中存储的主要是hash表的一些信息， 比如说当前hash表的域的数量以及当前hash表的版本号和过期时间(用做秒删功能), 而普通数据主要就是指的同一个hash表中一一对应的field和value，作为具体最后Rocksdb落盘的kv格式，下面是具体的实现方式:
1. 每个hash表的meta_key和meta_value的落盘方式:
![](https://i.imgur.com/YLP48rg.png)

meta_key实际上就是hash表的key, 而meta_value由三个部分构成: 4Bytes的Hash size(用于存储当前hash表的大小) + 4Bytes的Version(用于秒删功能) + 4Bytes的Timestamp(用于记录我们给这个Hash表设置的超时时间的时间戳， 默认为0)

2. hash表中data_key和data_value的落盘方式:
![](https://i.imgur.com/phiBsqd.png)

data_key由四个部分构成: 4Bytes的Key size(用于记录后面追加的key的长度，便与解析) + key的内容 + 4Bytes的Version + Field的内容， 而data_value就是hash表某个field对应的value。

3. 如果我们需要查找一个hash表中的某一个field对应的value, 我们首先会获取到meta_value解析出其中的timestamp判断这个hash表是否过期， 如果没有过期， 我们可以拿到其中的version, 然后我们使用key, version，和field拼出data_key, 进而找到对应的data_value（如果存在的话)


普通kv 是怎么根据version 进行判断的，是什么时候删除的。
=================================================
- version
存在version 的原因？
在我们大量的使用场景中. 对于Hash, zset, set, list这几种多数据机构，当member或者field很多的时候，用户有批量删除某一个key的需求, 那么这个时候实际删除的就是rocksdb 底下大量的kv结构, 如果只是单纯暴力的进行删key操作, 那时间肯定非常的慢, 难以接受. 那我们如何快速删除key？

刚才的nemo 的实现里面我们可以看到, 我们在value 里面增加了version, ttl 字段, 这两个字段就是做这个事情。

pika 没有采用暴力删除，而是Key的元信息增加版本，表示当前key的有效版本；，时间复杂度就是O(1)

=================================================


version字段用于对该键值对进行标记，以便后续的处理，如删除一个键值对时，可以在该version进行标记，后续再进行真正的删除，这样可以减少删除操作所导致的服务阻塞时间，version 字段配合 meta 信息中的version 进行使用。
pika多数据结构的实现主要是“meta key + 普通key”来实现的，所以对于多数据结构的读写，肯定都是对rocksdb进行2次及以上的读写次数




### list 的存储结构
blackwidow中的list由两部分构成，元数据(meta_key, meta_value), 和普通数据(data_key, data_value), 元数据中存储的主要是list链表的一些信息， 比如说当前list链表结点的的数量以及当前list链表的版本号和过期时间(用做秒删功能), 还有当前list链表的左右边界(由于nemo实现的链表结构被吐槽lrange效率低下，所以这次blackwidow我们底层用数组来模拟链表，这样lrange速度会大大提升，因为结点存储都是有序的), 普通数据实际上就是指的list中每一个结点中的数据，作为具体最后Rocksdb落盘的kv格式，下面是具体的实现方式
1. 每个list链表的meta_key和meta_value的落盘方式:
![](https://i.imgur.com/083SjIc.png)

meta_key实际上就是list链表的key, 而meta_value由五个部分构成: 8Bytes的List size(用于存储当前链表中总共有多少个结点) + 4Bytes的Version(用于秒删功能) + 4Bytes的Timestamp(用于记录我们给这个List链表设置的超时时间的时间戳， 默认为0) + 8Bytes的Left Index（数组的左边界) + 8Bytes的Right Index(数组的右边界)

2. list链表中data_key和data_value的落盘方式:
![](https://i.imgur.com/FBBn6kd.png)

data_key由四个部分构成: 4Bytes的Key size(用于记录后面追加的key的长度，便与解析) + key的内容 + 4Bytes的Version + 8Bytes的Index(这个记录的就是当前结点的在这个list链表中的索引)， 而data_value就是list链表该node中存储的值

### set 存储
blackwidow中的set由两部分构成，元数据(meta_key, meta_value), 和普通数据(data_key, data_value), 元数据中存储的主要是set集合的一些信息， 比如说当前set集合member的数量以及当前set集合的版本号和过期时间(用做秒删功能), 普通数据实际上就是指的set集合中的member，作为具体最后Rocksdb落盘的kv格式，下面是具体的实现方式：
1. 每个set集合的meta_key和meta_value的落盘方式:
![](https://i.imgur.com/bQeVvSj.png)

meta_key实际上就是set集合的key, 而meta_value由三个部分构成: 4Bytes的Set size(用于存储当前Set集合的大小) + 4Bytes的Version(用于秒删功能) + 4Bytes的Timestamp(用于记录我们给这个set集合设置的超时时间的时间戳， 默认为0)

2. set集合中data_key和data_value的落盘方式:
![](https://i.imgur.com/d2ctqPo.png)

data_key由四个部分构成: 4Bytes的Key size(用于记录后面追加的key的长度，便与解析) + key的内容 + 4Bytes的Version + member的内容， 由于set集合只需要存储member, 所以data_value实际上就是空串

### zset
blackwidow中的zset由两部部分构成，元数据(meta_key, meta_value), 和普通数据(data_key, data_value), 元数据中存储的主要是zset集合的一些信息， 比如说当前zset集合member的数量以及当前zset集合的版本号和过期时间(用做秒删功能), 而普通数据就是指的zset中每个member以及对应的score, 由于zset这种数据结构比较特殊，需要按照memer进行排序，也需要按照score进行排序， 所以我们对于每一个zset我们会按照不同的格式存储两份普通数据, 在这里我们称为member to score和score to member，作为具体最后Rocksdb落盘的kv格式，下面是具体的实现方式：
1. 每个zset集合的meta_key和meta_value的落盘方式:
![](https://i.imgur.com/RhZ8KMw.png)

meta_key实际上就是zset集合的key, 而meta_value由三个部分构成: 4Bytes的ZSet size(用于存储当前zSet集合的大小) + 4Bytes的Version(用于秒删功能) + 4Bytes的Timestamp(用于记录我们给这个Zset集合设置的超时时间的时间戳， 默认为0)

2. 每个zset集合的data_key和data_value的落盘方式(member to score):
![](https://i.imgur.com/C85Ba5Z.png)

member to socre的data_key由四个部分构成：4Bytes的Key size(用于记录后面追加的key的长度，便与解析) + key的内容 + 4Bytes的Version + member的内容， data_value中存储的其member对应的score的值，大小为8个字节，由于rocksdb默认是按照字典序进行排列的，所以同一个zset中不同的member就是按照member的字典序来排列的(同一个zset的key size, key, 以及version，也就是前缀都是一致的，不同的只有末端的member).

3. 每个zset集合的data_key和data_value的落盘方式(score to member):
![](https://i.imgur.com/QV9XHEk.png)

score to member的data_key由五个部分构成：4Bytes的Key size(用于记录后面追加的key的长度，便与解析) + key的内容 + 4Bytes的Version + 8Bytes的Score + member的内容， 由于score和member都已经放在data_key中进行存储了所以data_value就是一个空串，无需存储其他内容了，对于score to member中的data_key我们自己实现了rocksdb的comparator，同一个zset中score to member的data_key会首先按照score来排序， 在score相同的情况下再按照member来排序

## 数据同步
数量同步采用和mysql 类似的binlog 数据同步，接下来 讲一下 数据同步的一个基本过程。
### pika 主从全量同步
在进行全量数据同步之前，
- slave的向master发送MetaSync请求，在同步之前确保自身db的拓扑结构和master一致

FAQ:
rocksdb提够对当前db快照备份的功能，我们基于此，在dump时先对pika阻住用户的写，然后记录当前的binlog偏移量并且调用rocksdb的接口来拿到当前db的元信息，这个时候就可以放开用户写，然后基于这个元信息来进行快照数据的后台拷贝，阻写的时间很短 


- slave下的每个partition单独的向master端对应的partition发起trysync请求，建立同步关系

1. 在pika实例启动的同时会启动Rsync服务
2. master发现某一个partition需要全同步时，判断是否有备份文件可用，如果没有先dump一份
3. master通过rsync向slave发送对应partition的dump的文件
4. slave的对应partition用收到的文件替换自己的db
5. slave的对应partition用最新的偏移量再次发起trysnc
6. 完成同步

### binlog 增量同步

从库发送BinlogSyncRequest 报文，报文中需说明自己已经收到的BinlogOffset。

2，主库收到BinlogSyncRequest之后会从同步点开始发出一批BinlogSyncResponse。

3，从库在收到BinlogSyncResponse之后，会在写入本地binlog之后再进行1流程。

1. pika 的数据同步是使用binlog完成的。
并且binlog 中存储的是redis的命令
master是先写db再写binlog，让slave通过多个worker来写提高写入速度，不过这时候有一个问题，为了保证主从binlog顺序一致，写binlog的操作还是只能又一个线程来做，也就是receiver，所以slave这边是先写binlog在写db，所以slave存在写完binlog挂掉导致丢失数据的问题，不过redis在master写完db后挂掉同样会丢失数据，所以redis采用全同步的办法来解决这一问题，pika同样，默认使用部分同步来继续，如果业务对数据十分敏感，此处可以强制slave重启后进行全同步即。


## todo
1. binlog
2. 哨兵



## 缺点
 1. 目前没有对数据进行分片，你可以理解成和单机 Redis 类似，支持 master-slave 的架构，因此单个 pika 实例存储的大小的限制是磁盘大小的限制。类似于单机 Redis，那么单机性能是个瓶颈
 2. Pika 目前并没有使用类似 Redis 的 sentinel，pika 前面是挂 LVS 来负责主从切换。目前也没有使用 Codis 这样的 proxy 方案。



 # copi 

 1. 





