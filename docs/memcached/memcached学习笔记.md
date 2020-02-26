
官网： http://memcached.org/

Linux版本 memchched安装： https://www.runoob.com/memcached/memcached-install.html

个人比较喜欢 安装包安装的方式，可控。

## 安装
yum install memcached
注意：如果使用自动安装 memcached 命令位于 /usr/local/bin/memcached。

安装比较复杂，需要编译，我在 Linux上没有安装成功。

## 命令相关

## set命令
Memcached set 命令用于将 value(数据值) 存储在指定的 key(键) 中。`如果set的key已经存在，该命令可以更新该key所对应的原来的数据`，也就是实现更新的作用。

基本语法：

```
set key flags exptime bytes [noreply] 
value // 注意 value 在第二行
// 以下是例子
set runoob 0 900 9
memcached
STORED

get runoob
VALUE runoob 0 9
memcached

END
```

- key：键值 key-value 结构中的 key，用于查找缓存值。
- flags：可以包括键值对的整型参数，客户机使用它存储关于键值对的额外信息 。
- exptime：在缓存中保存键值对的时间长度（以秒为单位，0 表示永远）
- bytes：在缓存中存储的字节数
- noreply（可选）： 该参数告知服务器不需要返回数据
- value：存储的值（<font color="red">始终位于第二行</font>）（可直接理解为key-value结构中的value）

STORED：保存成功后输出。

ERROR：在保存失败后输出。

## add命令

Memcached add 命令用于将 value(数据值) 存储在指定的 key(键) 中。

如果 add 的 key 已经存在，则`不会更新数据`(过期的 key 会更新)，之前的值将仍然保持相同，并且您将获得响应 NOT_STORED。(注意与set命令区别。这个不会更新 ，set是更i性能数据) 其他的并没有发现与 set命令的区别。

## replace 命令

Memcached replace 命令用于替换已存在的 key(键) 的 value(数据值)。

如果 key 不存在，则替换失败，并且您将获得响应 NOT_STORED。

```
replace key flags exptime bytes [noreply]
value
```

## add set replace区别 

用过 memcache 的人都有一个疑惑，那就是 memcache 中为什么会有一个 add 方法、一个 set 方法、一个 replace 呢，这几个方法又有着什么样的区别呢，下边我们来分析下这几个方法的不同之处：

- memcache::add 方法：add 方法用于向 memcache 服务器添加一个要缓存的数据。
- memcache::set 方法：set 方法用于设置一个指定 key 的缓存内容，set 方法是 add 方法和 replace 方法的集合体。
- mmecache::replace 方法： replace 方法用于替换一个指定 key 的缓存内容，如果 key 不存在则返回 false。

方法 | 当key存在 |  当key不存在  
-|-|-
add | false | true |
replace | 替换(true) | false |
set | 替换(true) | true |

## append命令

Memcached append 命令用于向已存在 key(键) 的 value(数据值) 后面追加数据 。

```
append key flags exptime bytes [noreply]
value
```

实例：
- 首先我们在 Memcached 中存储一个键 runoob，其值为 memcached。
- 然后，我们使用 get 命令检索该值。
- 然后，我们使用 append 命令在键为 runoob 的值后面追加 "redis"。
- 最后，我们再使用 get 命令检索该值。

```
set runoob 0 900 9
memcached
STORED
get runoob
VALUE runoob 0 9
memcached
END
append runoob 0 900 5
redis
STORED
get runoob
VALUE runoob 0 14
memcachedredis
END
```

## prepend 

Memcached prepend 命令用于向已存在 key(键) 的 value(数据值) 前面追加数据 。

如果命令不存在，返回 NOT_STORED .用法与 append类似

https://www.runoob.com/memcached/memcached-prepend-data.html


## CAS

Memcached CAS（Check-And-Set 或 Compare-And-Swap） 命令用于执行一个"检查并设置"的操作

它仅在当前客户端最后一次取值后，该key 对应的值没有被其他客户端修改的情况下， 才能够将值写入。

检查是通过cas_token参数进行的， 这个参数是Memcach指定给已经存在的元素的一个唯一的64位值。

```
cas key flags exptime bytes unique_cas_token [noreply]
value
```

## 取出命令 

## get 指令
Memcached get 命令获取存储在 key(键) 中的 value(数据值) ，如果 key 不存在，则返回空。

```
get key
get key1 key2 key3
```

## gets 指令

Memcached gets 命令获取带有 CAS 令牌存 的 value(数据值) ，如果 key 不存在，则返回空。

```
set runoob 0 900 9
memcached
STORED
gets runoob
VALUE runoob 0 9 1
memcached
END
```

在 使用 gets 命令的输出结果中，在最后一列的数字 1 代表了 key 为 runoob 的 `CAS 令牌`。

## delete 命令

Memcached delete 命令用于删除已存在的 key(键)。

```
delete key [noreply]
```
返回值：

- DELETED：删除成功。
- ERROR：语法错误或删除失败。
- NOT_FOUND：key 不存在。

## incr/decr 命令
```java
incr/decr key increment_value
```

实例：
在以下实例中，我们使用 visitors 作为 key，初始值为 10，之后进行加 5 操作。
```
set visitors 0 900 2
10
STORED
get visitors
VALUE visitors 0 2
10
END
incr visitors 5
15
get visitors
VALUE visitors 0 2
15
END
```

- NOT_FOUND：key 不存在。
- CLIENT_ERROR：自增值不是对象。
- ERROR其他错误，如语法错误等。

## 统计相关的命令

## stats

stats 命令用于 返回统计信息例如 PID(进程号)、版本号、连接数等。

相关参数 参考： https://www.runoob.com/memcached/memcached-stats.html


## Java 相关的实战应用

https://www.runoob.com/memcached/java-memcached.html

