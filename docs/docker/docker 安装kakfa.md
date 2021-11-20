https://www.jianshu.com/p/e8c29cba9fae

## 安装 zk
1. docker pull wurstmeister/zookeeper
2. 启动镜像生成容器 docker run -d --name zookeeper -p 2181:2181 -v /etc/localtime:/etc/localtime wurstmeister/zookeeper

## 安装kafka
1. docker pull wurstmeister/kafka

2. 启动
docker run -d --name kafka1 -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=10.12.179.175:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://10.12.179.175:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -v /etc/localtime:/etc/localtime wurstmeister/kafka

-e KAFKA_BROKER_ID=0  在kafka集群中，每个kafka都有一个BROKER_ID来区分自己

-e KAFKA_ZOOKEEPER_CONNECT=192.168.155.56:2181/kafka 配置zookeeper管理kafka的路径192.168.155.56:2181/kafka

-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://10.12.179.175:9092  把kafka的地址端口注册给zookeeper

-e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 配置kafka的监听端口

-v /etc/localtime:/etc/localtime 容器时间同步虚拟机的时间

注意ip 通过 ifconfig改为本机地址

## 验证Kafka是否可用

docker exec -it kafka /bin/sh

进入路径：/opt/kafka_2.11-2.0.0/bin下

运行kafka生产者发送消息

./kafka-console-producer.sh --broker-list localhost:9092 --topic sun

发送消息

{"datas":[{"channel":"","metric":"temperature","producer":"ijinus","sn":"IJA0101-00002245","time":"1543207156000","value":"80"}],"ver":"1.0"}


运行kafka消费者接收消息

kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sun --from-beginning

## zk查看
1. docker exec -it zookeeper /bin/sh
2. cd bin
3. ./zkCli.sh
4.  ls /
```
[kafka, zookeeper]
```
5. ls /kafka/brokers/topics/sun/partitions
    返回 [0]
6. get /kafka/brokers/topics/sun  查看该节点的节点数据内容和属性信息
```
{"version":2,"partitions":{"0":[0]},"adding_replicas":{},"removing_replicas":{}}
cZxid = 0x23
ctime = Fri Nov 19 07:51:32 UTC 2021
mZxid = 0x23
mtime = Fri Nov 19 07:51:32 UTC 2021
pZxid = 0x24
cversion = 1
dataVersion = 0
aclVersion = 0
ephemeralOwner = 0x0
dataLength = 80
numChildren = 1
```

## idea
1. 

