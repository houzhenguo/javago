
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

gradle 与maven一样，都是一款包的管理工具，但是比maven更加轻量级。需要下载与自己的idea 匹配版本的Gradle，否则可能无法正常加载相关依赖。需要配置相关的环境变量，以及使用auto_import and 使用本地自己部署的gradel。注意，不同版本的gradle 对应的build.gradle的配置文件引入包的方式不同，可以参考给的junit的样子写。至于 包的地址，可以参考 https://search.maven.org

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


