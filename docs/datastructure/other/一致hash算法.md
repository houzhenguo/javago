
[原文链接](https://crossoverjie.top/2018/01/08/Consistent-Hash/)

# 一致性 Hash 算法分析

## 背景

当我们在做数据库分库分表或者是分布式缓存时，不可避免的都会遇到一个问题:如何将数据均匀的分散到各个节点中，并且尽量的在加减节点时能使受影响的数据最少。

## Hash 取模

随机放置就不说了，会带来很多问题。通常最容易想到的方案就是 hash 取模了。

可以将传入的 Key 按照 index = hash(key) % N 这样来计算出需要存放的节点。其中 hash 函数是一个将字符串转换为正整数的哈希映射方法，N 就是节点的数量。

这样可以满足数据的均匀分配，但是这个算法的<font color='red'>容错性和扩展性都较差</font>。
比如增加或删除了一个节点时，所有的 Key 都需要重新计算，显然这样成本较高，为此需要一个算法满足分布均匀同时也要有良好的容错性和拓展性。

## 一致Hash算法

一致 Hash 算法是将所有的哈希值构成了一个环，其范围在 0 ~ 2^32-1。如下图：

![Hash一致性](../images/consistent-hash-1.jpg)

之后将各个节点散列到这个环上，可以<font color='red'>用节点的 IP、hostname 这样的唯一性字段作为 Key </font>进行 hash(key)，散列之后如下：

![Hash一致性](../images/consistent-hash-2.jpg)

之后需要将数据定位到对应的节点上，使用同样的 hash 函数 将 Key 也映射到这个环上。

![Hash一致性](../images/consistent-hash-3.jpg)

这样按照顺时针方向就可以把 k1 定位到 N1节点，k2 定位到 N3节点，k3 定位到 N2节点。

## 容错性

这时假设 N1 宕机了：

![Hash一致性](../images/consistent-hash-4.jpg)

依然根据顺时针方向，k2 和 k3 保持不变，只有 k1 被重新映射到了 N3。这样就很好的保证了容错性，当一个节点宕机时只会影响到少少部分的数据。


## 拓展性

当新增一个节点时:

![Hash一致性](../images/consistent-hash-5.jpg)

在 N2 和 N3 之间新增了一个节点 N4 ，这时会发现受印象的数据只有 k3，其余数据也是保持不变，所以这样也很好的保证了拓展性。

## 虚拟节点

到目前为止该算法依然也有点问题:

当节点较少时会出现数据分布不均匀的情况：

![Hash一致性](../images/consistent-hash-6.jpg)

这样会导致大部分数据都在 N1 节点，只有少量的数据在 N2 节点。

为了解决这个问题，一致哈希算法引入了虚拟节点。将每一个节点都进行多次 hash，生成多个节点放置在环上称为虚拟节点:

![Hash一致性](../images/consistent-hash-7.jpg)

计算时可以在 IP 后加上编号来生成哈希值。

这样只需要在原有的基础上多一步由虚拟节点映射到实际节点的步骤即可让少量节点也能满足均匀性。

## 代码实现，可以给每台服务器加权，映射到不同 性能级别的服务器上
[参考链接](https://www.cnblogs.com/parryyang/p/8431100.html)

```java
/**
 *  一致Hash协议的的一个demo
 */
public class ConsistencyHashing {

    // 虚拟节点的个数
    private static final int VIRTUAL_NUM = 5;
    // 虚拟节点的分配，key是hash值，value是虚拟节点服务器的名称
    private static SortedMap<Integer, String> shards = new TreeMap<>();
    // 真实节点列表
    private static List<String> realNodes = new LinkedList<>();
    // 模拟初始服务器
    private static String[] servers = { "192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.5", "192.168.1.6" };

    static {
        for (String server: servers) {
            realNodes.add(server);
            System.out.println("真实节点 ["+server+"] 被添加");
            for (int i = 0; i< VIRTUAL_NUM; i++) {
                String virtualNode = server + "&&VN" +i;
                int hash = getHash(virtualNode);
                shards.put(hash, virtualNode);
                System.out.println("虚拟节点[" + virtualNode + "] hash:" + hash + "，被添加");
            }
        }
    }

    public static String getServer(String node) {
        int hash = getHash(node);
        Integer key = null;
        SortedMap<Integer, String> subMap = shards.tailMap(hash);
        if (subMap.isEmpty()) { // 没有比hash值大的
            key = shards.lastKey(); // 取出 全局最大的那个
        } else {
            key = subMap.firstKey(); // 取出当前map的第一个
        }
        String virtualNode = shards.get(key);
        return virtualNode.substring(0, virtualNode.indexOf("&&"));
    }


    /**
     * 添加节点
     *
     * @param node
     */
    public static void addNode(String node) {
        if (!realNodes.contains(node)) {
            realNodes.add(node);
            System.out.println("真实节点[" + node + "] 上线添加");
            for (int i = 0; i < VIRTUAL_NUM; i++) {
                String virtualNode = node + "&&VN" + i;
                int hash = getHash(virtualNode);
                shards.put(hash, virtualNode);
                System.out.println("虚拟节点[" + virtualNode + "] hash:" + hash + "，被添加");
            }
        }
    }

    /**
     * 删除节点
     *
     * @param node
     */
    public static void delNode(String node) {
        if (realNodes.contains(node)) {
            realNodes.remove(node);
            System.out.println("真实节点[" + node + "] 下线移除");
            for (int i = 0; i < VIRTUAL_NUM; i++) {
                String virtualNode = node + "&&VN" + i;
                int hash = getHash(virtualNode);
                shards.remove(hash);
                System.out.println("虚拟节点[" + virtualNode + "] hash:" + hash + "，被移除");
            }
        }
    }

    /**
     * FNV1_32_HASH算法
     */
    private static int getHash(String str) {
        final int p = 16777619;
        int hash = (int) 2166136261L;
        for (int i = 0; i < str.length(); i++)
            hash = (hash ^ str.charAt(i)) * p;
        hash += hash << 13;
        hash ^= hash >> 7;
        hash += hash << 3;
        hash ^= hash >> 17;
        hash += hash << 5;
        // 如果算出来的值为负数则取其绝对值
        if (hash < 0)
            hash = Math.abs(hash);
        return hash;
    }

    public static void main(String[] args) {

        //模拟客户端的请求
        String[] nodes = { "127.0.0.1", "10.9.3.253", "192.168.10.1" };

        for (String node : nodes) {
            System.out.println("[" + node + "]的hash值为" + getHash(node) + ", 被路由到结点[" + getServer(node) + "]");
        }

        // 添加一个节点(模拟服务器上线)
        addNode("192.168.1.7");
        // 删除一个节点（模拟服务器下线）
        delNode("192.168.1.2");

        for (String node : nodes) {
            System.out.println("[" + node + "]的hash值为" + getHash(node) + ", 被路由到结点[" + getServer(node) + "]");
        }
    }
}


```



