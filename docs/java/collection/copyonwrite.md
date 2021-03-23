1. 读写分离 写操作在⼀一个复制的数组上进⾏行行，读操作还是在原始数组中进⾏行行，读写分离，互不不影响。
写操作需要加锁，防⽌止并发写⼊入时导致写⼊入数据丢失。 写操作结束之后需要把原始数组指向新的复制数组。

```java
    public boolean add(E e) {
        final ReentrantLock lock = this.lock;
        lock.lock();
        try {
            Object[] elements = getArray();
            int len = elements.length;
            Object[] newElements = Arrays.copyOf(elements, len + 1);
            newElements[len] = e;
            setArray(newElements);
            return true;
        } finally {
            lock.unlock();
        }
    }

    final Object[] getArray() {
        return array;
    }

    /**
     * Sets the array.
     */
    final void setArray(Object[] a) {
        array = a;
    }
    public int indexOf(E e, int index) {
        Object[] elements = getArray();
        return indexOf(e, elements, index, elements.length);
    }
```
## 应用场景

读比较多，写比较少 

但是 CopyOnWriteArrayList 有其缺陷: 内存占⽤用:在写操作时需要复制⼀一个新的数组，使得内存占⽤用为原来的两倍左右;
数据不不⼀一致:读操作不不能读取实时性的数据，因为部分写操作的数据还未同步到读数组中。 所以 CopyOnWriteArrayList 不不适合内存敏敏感以及对实时性要求很⾼高的场景。