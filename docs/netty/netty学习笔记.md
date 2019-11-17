
[B站netty链接](https://www.bilibili.com/video/av33707223/?p=2)

[网友的netty笔记Github](https://github.com/ganxinming/nettyTest.git)

[netty官方网站](https://netty.io/)

[netty in action](https://manning.com)

## 1. netty概述

参考：https://netty.io/

[Google Protocol Buffer](https://developers.google.com/protocol-buffers/) 是一款序列化的工具，可以实现跨语言 rpc的数据序列化以及反序列化工作。需要程序手动定义数据结构，通过工具生成对应的语言代码。

[Thrift](http://thrift.apache.org/) 它所完成的功能与google 的 protocol buffer的功能类似。

WebSocket 相比之前的HTTP1.0 ，是一种升级，可以实现长链接，通信的时候不需要重复建立链接。

[Gradle](https://gradle.org/install/) 下载，配置相应的环境变量，下载completely版本。


## Gradle 

gradle 与maven一样，都是一款包的管理工具，但是比maven更加轻量级。需要下载与自己的idea 匹配版本的Gradle，否则可能无法正常加载相关依赖。需要配置相关的环境变量，以及使用auto_import and 使用本地自己部署的gradle。注意，不同版本的gradle 对应的build.gradle的配置文件引入包的方式不同，可以参考给的junit的样子写。至于 包的地址，可以参考 https://search.maven.org

## Git

## Netty

netty 的三种应用场景：

netty可以作为RPC的框架 socket开发，

netty可以作为长连接 websocket的服务器

netty可以作为http服务器 例如 tomcat 不遵守servlet规范 请求路由需要自己处理

## curl 安装 
可以使用bash win 版本

## httpServer的搭建
- [01httpserver的创建.md](./code/01httpserver的创建.md)

curl -X POST "http://localhost:8899"


## netty学习方法

不要陷入到细节中，应该先学习应用。否则陷入蒙蔽中。

## WebSocket
 html5的规范，解决http的无状态问题。
 http: request -> response 建立连接 。 client -> request -> server handle -> response -> client -> 连接断掉 (http1.0)
 http 1.1 keepAlive ，建立连接。request -> response 在指定时间内，client还要请求server，不需要重新建立连接。 无法实现服务器的 push技术。 轮询技术早期

 poll 技术：1.无意义返回 2. header 的大小。

 html5 socket的长连接。client and server 平等。 不需要发送 header信息，节省带宽。双向数据通信。

## ProtocolBuff

1. RMI Remote method invocation, 只针对Java

client: 序列化字节码 

server： 反序列化

rpc 几乎都有 代码生成的过程。在client 称为 stub(庄) server:skeleton(骨架)

序列化与反序列化：也叫做 编码与解码。encode,decode.对象 <---> 字节

RPC: Remote Procedure Call 远程过程调用。与 RMI相比 跨语言。

（序列化的过程？）
1. 定义一个接口说明文件 ： 描述了对象（结构体），对象成员，接口方法的一系列信息。
```
message Person {
    required int32 id  =1;
    optional string name =2;
}
```
2. 通过 RPC框架所提供的编译器将接口说明文件编译成具体语言文件
3. 在客户端与服务器段分别引入 RPC编译器所生成的文件，即可像调用本地方法一样调用远程方法。

编码效率，压缩比例，解码效率，影响传输。这是影响 RPC的选择。

服务与服务之间的调用使用 RPC。

required 必须提供

optional 不是必须的。 后面的数字代表顺序。


## 使用 Git作为 版本控制系统 
Q: Client 和 Server 需要同一份生成出来的 Proto Buf 的Java 文件，为了避免手动拷贝，可以使用 git进行控制。

### 解决方案1（不是特别完美）

1. git submodule ： Git 仓库里面的仓库.

ServerProject : 已经在 Git版本之中了。这个工程用到了 Protocol Buf 。

Protobuf-Java: 独立的git的工程。通过 编译器生成的Java代码推送到这个仓库中。protoc

通过 submodule 将 Protobuf-Java 引入到 ServerProject 中。通过git命令.

ClientProject: 与 ServerProject 相似。

这种情况对于 分支情况不友好。
branch:
    develop
    test
    master

两个仓库，容易出现 分支错乱。

## Git subtree (推荐)

ServerProject

Protobuf-Java:(公用)

ClientProject

将公用的代码拉取到 ServerProject 中，仅仅是一个仓库，产生了合并，产生了一次提交。

# Thrift

