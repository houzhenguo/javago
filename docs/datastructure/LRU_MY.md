
## 自己实现版本的LRU 要熟练掌握

```java

public LRUCache {
    private static class DLinkedNode {
        int key;
        int value;
        DLinkedNode pre;
        DLinkedNode next; 
    }
    private HashMap<Integer, DLinkedNode> 
            cache = new HashMap();
    private int count;
    private int capacity;
    private DLinkedNode head,tail;

    public LRUCache(int capacity){
        this.count = 0;
        this.capcity = capacity;

        this.head = new DLinkedNode();
        this.head.pre = null;

        this.tail = new DLinkedNode();
        this.tail.next = null;

        this.head.next = tail;
        this.tail.head = head;
    }

    public int get(int key){
        DLinkedNode node = cache.get(key);
        if (null == node){
            return -1;
        }
        // cache中有，将该节点挪到开头
        moveToHead(node);
        return node.value;
    }

    private void moveToHead(DLindeNode node){

        //1. 摘除节点
        node.pre.next = node.next;
        node.next.pre = node.pre;

        addNode(node);
    }
    private void addNode(DLindedNode node){
        //2.插入头节点
        node.pre = head;
        node.next = head.next;
        head.next.pre = node;
        head.next = node;
    }

    public int put(int key, int value){
         DLinkedNode node = cache.get(key);
         if(null== node){
            node = new DLinkedNode();
            node.key = key;
            node.value = value;

            cache.put(key, node);

            addNode(node);

            ++count;
            if (count>capacity){
                DLinedNode delNode = tail.pre;
                delNode.pre.next = tail;
                tail.pre = delNode.pre;
                cache.remove(delNode.key);
                delNode = null;
                --count;
            }
         }else{
             node.value = value;
             moveToHead(node);
         }
    }


}
```