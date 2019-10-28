
本文使用的是netty4.1版本 ， gradle 4.4版本，ideal 201802版本，需要全部提前配置

# 代码部分
## 代码1 启动类
```java
package com.houzhenguo.netty.firstexample;

import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioServerSocketChannel;
/**
 *  20191027
 */
public class TestServer {
    public static void main(String[] args) throws Exception{
        // 就是两个死循环
        EventLoopGroup bossGroup = new NioEventLoopGroup();  // 事件循环组 两个线程组 boss接收链接，交给worker
        EventLoopGroup workerGroup = new NioEventLoopGroup(); // 真正完成

        try {
            ServerBootstrap serverBootstrap = new ServerBootstrap();// 用于启动服务端的一个类
            serverBootstrap.group(bossGroup, workerGroup)// acceptor boss / worker
                    .channel(NioServerSocketChannel.class) // 反射实现
                    .childHandler(new TestServerInitlializer()); // 请求处理器（自定义的服务器初始化器）

            ChannelFuture channelFuture = serverBootstrap.bind(8899).sync();
            channelFuture.channel().closeFuture().sync();
        }finally {
            bossGroup.shutdownGracefully(); // shutdown
            workerGroup.shutdownGracefully();
        }
    }
}

```

## 代码2 initlializer
```java
package com.houzhenguo.netty.firstexample;

import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.codec.http.HttpServerCodec;
// 20191027 2
/**
 *  链接建立之后就会被创建，就会执行 initChannel方法
 */
public class TestServerInitlializer extends ChannelInitializer<SocketChannel> {
    @Override
    protected void initChannel(SocketChannel ch) throws Exception {
        ChannelPipeline pipeline = ch.pipeline(); // 管道，可以有很多channel handler 相当于拦截器
        pipeline.addLast("httpServerCodec",new HttpServerCodec()); // netty提供 可以自己提供处理器 编解码用的
        pipeline.addLast("testHttpServerHandler", new TestHttpServerHandler()); // 自定义的handler
    }
}

```

## 3 代码 自定义的handler

```java
package com.houzhenguo.netty.firstexample;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.handler.codec.http.*;
import io.netty.util.CharsetUtil;
/**
 *  自己定义的处理器
 */
public class TestHttpServerHandler extends SimpleChannelInboundHandler<HttpObject> {
    @Override // 读取客户端发过来的请求，响应客户端的方法 相当于messageReceived mina
    protected void channelRead0(ChannelHandlerContext ctx, HttpObject msg) throws Exception {
        if (msg instanceof HttpRequest) {
            ByteBuf content = Unpooled.copiedBuffer("hello world", CharsetUtil.UTF_8); // 返回的内容
            // http 版本 ，ok，内容
            FullHttpResponse response = new DefaultFullHttpResponse(HttpVersion.HTTP_1_0, HttpResponseStatus.OK, content);
            response.headers().set(HttpHeaderNames.CONTENT_TYPE, "text/plain");
            response.headers().set(HttpHeaderNames.CONTENT_LENGTH, content.readableBytes());

            ctx.writeAndFlush(response); // 返回
        }
    }
}

```

最后测试部分：

```bash
curl "http://localhost:8899" 
curl -X POST "http://localhost:8899"

```
chrome浏览器 可能会发送两次请求