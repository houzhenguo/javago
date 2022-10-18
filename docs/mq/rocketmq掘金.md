https://juejin.cn/post/6844904008629354504

## 简介
1. Java 分布式 队列模型的mq 阿里系的；高性能，高可靠，高实时，分布式
2. 金融领域 -> 高可靠，高可用，低延迟 -> 订单/充值/交易/消息推送
## 核心模块
1. broker -> 接受生产者发来的消息并存储
2. client -> 提供发送/消费消息的客户端
3. nameserver -> 保留topicName，meta信息
4. 底层采用netty fastjson序列化 和自定义的二进制
NameServer/Broker/Consumer/Producer
NameServer 是一个topic和路由注册中心，Producer发送消息之前会去 nameserver获取broker的路由信息（定期的，其实心跳中就可以），consumer也是。
Broker是主要提供服务的-> 上亿消息的堆积能力？
5. 支持push/pull
6. 除了topic以外，还有tag
7. 顺序消费 -> 队列-> 如果想要全局唯一的顺序则确保使用的主题只有一个消息队列

## 优点
1. 稳定 高可用高性能 10w+的吞吐，消息可靠 0丢失，Java开发。

## 消息可靠性
1. 同步刷盘/异步刷盘
2. 每个broker用一个 commitlog存储？

1. producer提交halfmsg -> half msg success -> 执行本地事务 -> 给mq commit/rollback -> 
定期扫描那些长期处于 halfmsg状态的消息 -> 消息回查 -> 给 producer检查，然后 commit/rollback

2. 回溯消费
3. 定时消息，支持固定 level