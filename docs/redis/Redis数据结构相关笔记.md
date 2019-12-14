
# String 

sds :simple dynamic string

redis底层使用心得 sds 数据结构，可以动态的扩展C的字符串。
- c的字符串空格问题 -> 二进制
- c的字符串长度问题 -> len
- 空间预分配 -> 1M以下-> +len  以上
- 空间懒回收

# list

底层就是一个双向链表，pre,next,head,tail,len;

# 字典

底层就是个 Entry -> 里面放着键值对 数组，记录着 used, size . 
Entry -> next -> 指向下一个的指针（链地址法解决hash冲突 -> 头插法）

# 跳跃表

查找的平均复杂度 ：O(logn) -> 最差 O(n)

zset and 集群节点中用到。




