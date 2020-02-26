
# gRPC 学习

参考学习链接 ：  http://doc.oschina.net/grpc?t=58008

Github 地址： https://github.com/grpc/grpc

中文官网： http://doc.oschina.net/grpc?t=60134 

## 简介
 客户端像是调用本地方法一样直接调用另外一台不同机器上服务端应用的方法。可以使我们更容易的创建分布式的应用和服务。 全称： Google Remote Procedure Call 远程过程调用。


![grpc](./images/grpc_1.png)

grpc 可以很容易的使用 Java 创建一个 gRPC的服务器，使用其他语言创建 客户端进行交互。

## 通信协议

使用 protocol buffer ，protocol buffer 就是一种序列化的机制。具体可以参考 protocol 官网 或者 相关的github。 proto3 最好是使用这个 版本，避免版本 兼容问题。

protocol buffer 官网 https://developers.google.com/protocol-buffers/docs/overview

## 官网翻译

服务的定义：

gRPC 使用 protocol buffers 作为接口定义的方式。就像下面这样:

```java
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string greeting = 1;
}

message HelloResponse {
  string reply = 1;
}
```
gRPC 可以定义4中服务的类型：
1. client给 服务器发送一个参数，server 给client 回复一个参数
```java
rpc SayHello(HelloRequest) returns (HelloResponse) {
}
```
2. client 给 server 发送 流数据，server读取流数据。client 返回流式数据。
```java
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse) {
}
```
3. 剩下的两种都是 流式。翻译的不太好。 https://grpc.io/docs/guides/concepts/
以上不同就是 client 与 server 的交互是否是流式交换。


- 在 server端，server实现定义的服务，处理客户端的请求。
- 在 client 端，实现与 服务器 相同的协议。client 可以在本地调用方法，

同步调用 以及 异步调用。

同步调用就是等待 server 的response.异步调用： 缓存 通信的id，在返回的时候 调用。

