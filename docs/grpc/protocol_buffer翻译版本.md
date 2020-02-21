
https://developers.google.com/protocol-buffers/docs/overview

https://developers.google.com/protocol-buffers/docs/javatutorial

官网的重点翻译：

## 简介
protocol-buffer 更小，更快，更灵活，多语言，跨平台，自动生成 源码的一款序列化工具。

需要定义 信息结构体 在 `.proto` files 。包含一些 基础信息，key-value 。下面是一个简单的例子


## 优点
1. effective Java 中告诉我们，Java自带的序列化 在 跨平台 或者 跨本的 时候表现不好，有可能数据该表导致不兼容。
2. 可以将数据序列化称为 xml .

```proto
syntax = "proto2";

package tutorial; // 这个是避免在不同的项目中命名冲突的 ，如果没有明确的定义 java_package，这个就是包名

option java_package = "com.example.tutorial"; // Java 文件生成的地方,不定义就是上面的package
option java_outer_classname = "AddressBookProtos"; 定义一个类包含这个文件中所有的类,不给出的就是用文件名驼峰处理，比如 `my_proto.proto` 会生成 MyProto 作为外部类名字

message Person {
    required string name = 1; // 后面的数字 标明 每个 属性在二进制编码中。这个一旦使用了旧不能乱改
    required int32  id   = 2;
    optional string email = 3;

    enum PhoneType {
        MOBILE = 0;
        HOME = 1;
        WORK = 2;
    }
    message PhoneNumber {
        required string number = 1;
        optional PhoneType type = 2 [default = HOME];
    }
    repeated PhoneNumber phone = 4;
}
```

> `bool`,`int32`,`float`,`double`,`string` 也可以自定义，比如 PhoneNumber，特殊的是：1-15 占用更少的字节。 `required`,`optional`,`optioinal` 是必须存在其中一个的

- required 是属性必须提供的描述之一，表示该值是必须要设置的,除此之外，这个属性可能要考虑未被初始化。尝试创建一个 违背初始化的消息会抛出 `RuntimeException` ，尝试解析会抛出 `IOException` .慎用。
- optional 字段可以设置，也可以不设置. 如果不设置，那么就会使用默认值。
- repeated 可以重复的， 元素的顺序是保留的
`.proto` 文件的格式化比较简单，可以自定义结构体。定义完结构体之后，可以使用 protocol buffer 的编译器生成对应的代码的类。你可以在结构体中`新增` 属性而不需要考虑兼容性，在解析的时候会自动的忽略新增的属性。

## Start ProtocolBuffer

1. 下载相关的 release 包， https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.4

    在这里我下载的这个版本 ： protoc-3.11.4-win64.zip
    配置好 path 环境变量 ，xxxprotoc/bin

2. proto3

最新的发行版本是 3.  新版本与旧的2版本不是完全的兼容，只是增加了一些新的特性。定义好自己的 proto文件，接下来就是生成源代码。

可以在 xx.proto 文件下 执行这个命令 ` protoc  --java_out=. addressbook.proto`

```java
// required string name = 1;
public boolean hasName();
public String getName();

// required int32 id = 2;
public boolean hasId();
public int getId();

// optional string email = 3;
public boolean hasEmail();
public String getEmail();

// repeated .tutorial.Person.PhoneNumber phones = 4;
public List<PhoneNumber> getPhonesList();
public int getPhonesCount();
public PhoneNumber getPhones(int index);
```