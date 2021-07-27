https://github.com/OpenAtomFoundation/pika/wiki

https://zhuanlan.zhihu.com/p/90891875

Pika相对于Redis，最大的不同就是Pika是持久化存储，数据存在磁盘上，而Redis是内存存储，由此不同也给Pika带来了相对于Redis的优势和劣势

优势：
容量大：Pika没有Redis的内存限制, 最大使用空间等于磁盘空间的大小
加载db速度快：Pika在写入的时候, 数据是落盘的, 所以即使节点挂了, 不需要rdb或者oplog，Pika重启不用加载所有数据到内存就能恢复之前的数据, 不需要进行回放数据操作。
备份速度快：Pika备份的速度大致等同于cp的速度（拷贝数据文件后还有一个快照的恢复过程，会花费一些时间），这样在对于百G大库的备份是快捷的，更快的备份速度更好的解决了主从的全同步问题
劣势：
由于Pika是基于内存和文件来存放数据, 所以性能肯定比Redis低一些, 但是我们一般使用SSD盘来存放数据, 尽可能跟上Redis的性能