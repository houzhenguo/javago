
# HashMap


## hashmap数组的长度为 2的幂次方
最近在看HashMap的源码时，发现了里面好多很不错的算法，相比Java7从性能上提高了许多。其中tableSizeFor就是一个例子。tableSizeFor的功能（不考虑大于最大容量的情况）是返回大于输入参数且最近的2的整数次幂的数。比如10，则返回16。该算法源码如下：

```java
static final int tableSizeFor(int cap) {
    int n = cap - 1;
    n |= n >>> 1;
    n |= n >>> 2;
    n |= n >>> 4;
    n |= n >>> 8;
    n |= n >>> 16;
    return (n < 0) ? 1 : (n >= MAXIMUM_CAPACITY) ? MAXIMUM_CAPACITY : n + 1;
}
```
总结： 想让他低位全是1 +1 

详解如下：

先来分析有关n位操作部分：先来假设n的二进制为01xxx...xxx。接着

对n右移1位：001xx...xxx，再位或：011xx...xxx

对n右移2为：00011...xxx，再位或：01111...xxx

此时前面已经有四个1了，再右移4位且位或可得8个1

同理，有8个1，右移8位肯定会让后八位也为1。

综上可得，该算法让最高位的1后面的位全变为1。

最后再让结果n+1，即得到了2的整数次幂的值了。

现在回来看看第一条语句：

```java
int n = cap - 1;
```

让cap-1再赋值给n的目的是另找到的目标值大于或等于原值。例如二进制1000，十进制数值为8。如果不对它减1而直接操作，将得到答案10000，即16。显然不是结果。减1后二进制为111，再进行操作则会得到原来的数值1000，即8。

### 扩展
如何判断2 的幂次方？ 因为2的幂次方只有高位 为 1. n>0 && (n&(n-1) == 0) 

xxxx1000
&
xxxx0111 = xxxx0000

[n-1的妙用](https://www.cnblogs.com/skillking/p/9930095.html)

hash & (n-1) n为 2的次幂 

[为什么 %与&一样](https://stackoverflow.com/questions/46111975/binary-arithmetic-why-hashn-is-equivalent-to-hashn-1?r=SearchResults)

```java
 1 = 00000000000000000000000000000001  // Shift by 4 to get...
16 = 00000000000000000000000000010000  // Subtract 1 to get...
15 = 00000000000000000000000000001111
```
So just the lowest 4 bits are set in 15. If you `&` this with another int, it will only allow bits in the last 4 bits of that number to be set in the result,` so the value will only be in the range 0-15`, so it's like doing % 16.


## 其他
1. 链表法 和开放地址法

defalut 16 = 1<< 4
factor  = 0.75
tree_threshold = 8 转红黑树 +  min_tree_capacity = 64

```java
  static final int hash(Object key) {
        int h;
        return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>> 16);
    }
```
由于和（length-1）运算，length 绝大多数情况小于2的16次方。所以始终是hashcode 的低16位（甚至更低）参与运算。要是高16位也参与运算，会让得到的下标更加散列。

### getNode
```java
    final Node<K,V> getNode(int hash, Object key) {
        Node<K,V>[] tab; Node<K,V> first, e; int n; K k;
        if ((tab = table) != null && (n = tab.length) > 0 &&
            (first = tab[(n - 1) & hash]) != null) {
            if (first.hash == hash && // always check first node
                ((k = first.key) == key || (key != null && key.equals(k))))
                return first;
            if ((e = first.next) != null) {
                if (first instanceof TreeNode)
                    return ((TreeNode<K,V>)first).getTreeNode(hash, key);
                do {
                    if (e.hash == hash &&
                        ((k = e.key) == key || (key != null && key.equals(k))))
                        return e;
                } while ((e = e.next) != null);
            }
        }
        return null;
    }
```

