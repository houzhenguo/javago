
## 20200215

Redis 设计与实现的笔记重点

因为在 新东方的面试中，有被问到底层的重点，所以抽出一天的时间，
迅速读一遍。

Redis3.0

Redis 数据库中的键总是一个字符串对象。

list,set,zset,string,hash

```
set msg "helloworld"
```

## String

SDS, redis 的key 是键，

value 是一个列表对象。

SDS,保留着 length,预分配，懒回收。

```c
struct sds {
    int len;
    int free;
    char[] buff;
}
```

## 链表

listNode 保存了 前置节点，后置节点。 以及节点的值。

链表虽然是 节点组成的，但是它保留了 len 的长度这个字段。

## 字典

map 是使用 数组 + 链表的方式解决冲突。

skiplist

## 跳跃表 
O(logn)
最坏 On

