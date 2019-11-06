
[优秀的博客](https://blog.csdn.net/u010647035/article/details/86375981)


```java
 public V put(K key, V value) {
        return putVal(key, value, false);
    }

    /** Implementation for put and putIfAbsent */
    final V putVal(K key, V value, boolean onlyIfAbsent) {
        if (key == null || value == null) throw new NullPointerException();
        //计算key的hash值
        int hash = spread(key.hashCode());
        //记录链表长度
        int binCount = 0;
        for (Node<K,V>[] tab = table;;) {
            Node<K,V> f; int n, i, fh;
            //数据 tab 为空，初始化数组
            if (tab == null || (n = tab.length) == 0)
                tab = initTable();
            //根据 hash 值，找到对应数组下标，并获取对应位置 的第一个节点 f
            else if ((f = tabAt(tab, i = (n - 1) & hash)) == null) {
            // 如果数组该位置为空，用 CAS 操作将这个新值放入这个位置
            // 如果 CAS 失败，那就是有并发操作，继续下一个循环
                if (casTabAt(tab, i, null,
                             new Node<K,V>(hash, key, value, null)))
                    break;                   // no lock when adding to empty bin
            }
            //f 的 hash 值等于 MOVED，说明在扩容
            else if ((fh = f.hash) == MOVED)
                //进行数据迁移
                tab = helpTransfer(tab, f);
            else {
                //f 是该位置的首节点，而且不为空
                V oldVal = null;
                // 获取数组该位置的首节点的监视器锁
                synchronized (f) {
                    if (tabAt(tab, i) == f) {
                        // 首节点 hash 值大于 0，说明是链表
                        if (fh >= 0) {
                            //记录链表长度
                            binCount = 1;
                            //链表遍历
                            for (Node<K,V> e = f;; ++binCount) {
                                K ek;
                                //// 如果找到了相等的 key，判断是否要进行值覆盖
                                if (e.hash == hash &&
                                    ((ek = e.key) == key ||
                                     (ek != null && key.equals(ek)))) {
                                    oldVal = e.val;
                                    if (!onlyIfAbsent)
                                        e.val = value;
                                    break;
                                }
                                Node<K,V> pred = e;
                                //到了链表的尾部，将新值放到链表的尾部
                                if ((e = e.next) == null) {
                                    pred.next = new Node<K,V>(hash, key,
                                                              value, null);
                                    break;
                                }
                            }
                        }
                        //如果当前位置是红黑树
                        else if (f instanceof TreeBin) {
                            Node<K,V> p;
                            binCount = 2;
                            //红黑树插入新节点
                            if ((p = ((TreeBin<K,V>)f).putTreeVal(hash, key,
                                                           value)) != null) {
                                oldVal = p.val;
                                if (!onlyIfAbsent)
                                    p.val = value;
                            }
                        }
                    }
                }
                //如果链表长度不为0
                if (binCount != 0) {
                    //如果链表长度大于等于临界值【8】，将链表转换为红黑树
                    if (binCount >= TREEIFY_THRESHOLD)
                        //将链表转换成红黑树前判断，如果当前数组的长度小于 64，那么会选择进行数组扩容，而不是转换为红黑树
                        treeifyBin(tab, i);
                    if (oldVal != null)
                        return oldVal;
                    break;
                }
            }
        }
        addCount(1L, binCount);
        return null;
    }
————————————————
版权声明：本文为CSDN博主「IT码客」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/u010647035/article/details/86375981
```