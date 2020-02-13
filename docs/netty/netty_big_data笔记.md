
# NIO
## Buffer

Buffer 与 Channel 交互。数据之间互相读取。

ByteBuffer
CharBuffer
ShortBuffer

flip 从 写模式 调整为 读模式

clear 清空buffer

capcity > limit > position

ByteBuffer buf = ByteBuffer.allocate(23);

int bytesRead = inChannel.read(buf);

flip()

## Channel

所有的IO 都是从 Channel 开始的。

从通道进行数据的读取。

FileChannel 

SocketChannel

ServerSocketChannel

Scatter / Gathter

在多个缓冲区 实现 简单的 IO 操作。

Scatter : 从一个 Channel 读取信息分散到 N 个缓冲区

Gathter: 将 N 个 Buffer 里面的内容 按顺序发送到一个 Channel

## Selector

选择器，多路复用器。 用于检查一个 或 多个 NIO Channel 状态是否可读 可写。

如此可以实现 单线程管理多个 Channel.

channel 配置为 非阻塞的。 FileChannel 是无法配置为 非阻塞的


事件有： OP_connect ,op_accept, op_read,op_write


# Netty


Netty 的 buffer 专为 网络通讯而生。

## TCP 与 Buffer

TCP 会把应用层的数据拆开成字节。TCP 包最大长度限制。1400多个字节，总体长度 1500多个字节。

## ChannelBuffer 

ByteBuf

netty 中的 channelbuffer 是 把 nio 中的buffer 重新实现了一遍。

readIndex and writeIndex 这种方式比flip更友好一点。


可以自动扩容


CompositeChannelBuffer 可以组合多个 buffer，保存了 所有 ChannelBuffer的引用。从而实现 ZeroCopy

## Reactor

dispather ? 

CPU 的处理速度 远远快于 IO 速度的。如果 CPU 为了    IO 阻塞


事件驱动。回调方式。。应用业务 向中间人注册一个回调（event handler） 当IO 就绪之后，就
中间人产生要给事件，并且通知 handler 处理。

Don't call us, we'll call you


BossGroup

workerGroup 开启线程池去处理

## pipline

Channel 是通讯的载体， ChannelHandler 是负责 Channel中的逻辑处理


ChannelPipLine 是 ChannelHandler的容器。分为 上行 下行 
