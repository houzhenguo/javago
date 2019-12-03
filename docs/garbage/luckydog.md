
推送系统：
    系统内部发生了一个事件，各个消息系统去消费。
    world -> mq -> 开启持久化 -> pushserver -> redis (msgId) 重复消费的问题
    push -> apple, huawei,mi 等。push相关的统计。

    弱校验 -> messageId -> 缓存 redis -> 设置过期事件 -> 同时校验系统的发送时间。这样就可以避免消息在极短的时间内重复消费。

    订单系统 -> mq -> 活动相关的计数

充值系统 -> 订单系统 -> 分布式事务？ 这个可以不说，场景不明显，但是还是要整理一部分的分布式事务的内容。



## 分布式相关

分布式锁使用reids实现还是比较少，一般使用zookeeper

# 备忘录：

1. MySQL是如何实现事务机制的？
    在mysql/事务中进行完善
https://juejin.im/post/5b5a0bf9f265da0f6523913b#heading-4