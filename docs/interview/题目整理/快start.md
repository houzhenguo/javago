# 快速阅读

## mysql
1. ACID -> 修改丢失 脏读 不可重复读 幻读
2. 行表锁 意向锁 s/x锁
3. 隔离级别 readuncommit(脏) readcommit（不可重复） rr （幻）serializable
4. mvcc 事务版本号 undolog 两列 / 快照读 当前读
5. mvcc不能解决幻读，间隙锁 + mvcc解决
    record locks
    gap lock
6. b+tree
   叶子节点在同一层 logn,范围查找
   红黑树度2，高度更高，查找更费劲/ 磁盘预读取，node大小为page = 16k大小
   聚簇索引 data -> id
7. index 检索 排序 做前缀 server优化 explain 区分性强度方前面
8. 主从复制 binlog I/O sql线程，replay log replay 并行复制，读写分离
9. double write

## redis
1. string list set hash zset ziplist
2. 渐进式rehash, skiplist -> 插入速度快，不用维护 旋转 
3. memcached 不支持持久化
4. lru 淘汰 最少使用 已经过期，随机 
5. rdb/aof 快照 /append 1second
6. IO多路复用 命令请求处理器/命令回复处理器/链接应答处理器
7. 哨兵sentinel 
8. 512M 1M >1M->1M SDS 预分配 惰性挥手
9. list linkedlist + ziplist（数据量少）
10. scan
