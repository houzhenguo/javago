
# Arthas


## 概述
> Arthas 是Alibaba开源的Java诊断工具。在线排查问题，无需重启；动态跟踪Java代码；实时监控JVM状态。Arthas 支持JDK 6+，支持Linux/Mac/Windows，采用命令行交互模式，同时提供丰富的 Tab 自动补全功能，进一步方便进行问题的定位和诊断。

- Github: https://github.com/alibaba/arthas
- 文档: https://alibaba.github.io/arthas/
- 安装方式： https://alibaba.github.io/arthas/install-detail.html


# Demo初级教程 

## 1. 启动 arthas-demo

下载`arthas-demo.jar`，再用`java -jar`命令启动：

```java
wget https://alibaba.github.io/arthas/arthas-demo.jar
java -jar arthas-demo.jar
```
`arthas-demo`是一个很简单的程序，它随机生成整数，再执行因式分解，把结果打印出来。如果生成的随机数是负数，则会打印提示信息。

## 2. 启动arthas-boot

在新的`Terminal 2`里，下载`arthas-boot.jar`，再用`java -jar`命令启动：

```java
wget https://alibaba.github.io/arthas/arthas-boot.jar
java -jar arthas-boot.jar --target-ip 0.0.0.0
```
`arthas-boot`是`Arthas`的启动程序，它启动后，会列出所有的Java进程，用户可以选择需要诊断的目标进程。

选择第一个进程，输入 1 ，再Enter/回车：

Attach成功之后，会打印Arthas LOGO。输入 help 可以获取到更多的帮助信息。
```bash
[INFO] Start download arthas from remote server: https://repo1.maven.org/maven2/com/taobao/arthas/arthas-packaging/3.1.4/arthas-packaging-3.1.4-bin.zip

[INFO] Download arthas success.
[INFO] arthas home: /home/houzhenguo/.arthas/lib/3.1.4/arthas
[INFO] Try to attach process 4596
[INFO] Attach process 4596 success.
[INFO] arthas-client connect 0.0.0.0 3658
  ,---.  ,------. ,--------.,--.  ,--.  ,---.   ,---.                           
 /  O  \ |  .--. ''--.  .--'|  '--'  | /  O  \ '   .-'                          
|  .-.  ||  '--'.'   |  |   |  .--.  ||  .-.  |`.  `-.                          
|  | |  ||  |\  \    |  |   |  |  |  ||  | |  |.-'    |                         
`--' `--'`--' '--'   `--'   `--'  `--'`--' `--'`-----'    

```

# 简单命令

## Dashboard

`dashboard` 命令可以查看当前系统的实时数据面板。输入 Q 或者 Ctrl+C 可以退出dashboard命令。

## Thread

`thread 1` 命令会打印线程ID 1的栈。

Arthas支持管道，可以用 `thread 1 | grep 'main('` 查找到main class。

可以看到main class是demo.MathGame：
```bash
[arthas@4596]$ thread 1 | grep 'main'
"main" Id=1 TIMED_WAITING
    at demo.MathGame.main(MathGame.java:17)
```

## Sc

可以通过 `sc` 命令来查找JVM里已加载的类：

通过`-d`参数，可以打印出类加载的具体信息，很方便查找类加载问题。

sc支持通配

```bash
sc -d *MathGame

# 打印出来的info
class-info        demo.MathGame                                                                                                                                                                                
 code-source       /home/houzhenguo/server/arthas/arthas-demo.jar                   
 name              demo.MathGame                                                    
 isInterface       false                                                            
 isAnnotation      false                                                            
 isEnum            false                                                            
 isAnonymousClass  false                                                            
 isArray           false                                                            
 isLocalClass      false                                                            
 isMemberClass     false                                                            
 isPrimitive       false                                                            
 isSynthetic       false                                                            
 simple-name       MathGame                                                         
 modifier          public                 
 annotation                               
 interfaces                               
 super-class       +-java.lang.Object     
 class-loader      +-sun.misc.Launcher$AppClassLoader@55f96302                      
                     +-sun.misc.Launcher$ExtClassLoader@f7b4d77                     
 classLoaderHash   55f96302                                                         
Affect(row-cnt:1) cost in 81 ms.
```

## sm

`sm`命令则是查找类的具体函数。比如：
```bash
sm java.math.RoundingMode
```

通过`-d`参数可以打印函数的具体属性：
```bash
sm -d java.math.RoundingMode
```

## Jad

可以通过 `jad` 命令来反编译代码：
```bash
jad demo.MathGame
```

## Watch

通过`watch`命令可以查看函数的参数/返回值/异常信息。

```bash
# q 或者 ctrl + c 退出watch命令
watch demo.MathGame primeFactors returnObj
```

## Exit/Shutdown 退出Arthas

用 `exit` 或者 `quit` 命令可以退出Arthas。

退出Arthas之后，还可以再次用 `java -jar arthas-boot.jar` 来连接。

**彻底退出Arthas**

exit/quit命令只是退出当前session，arthas server还在目标进程中运行。

想完全退出Arthas，可以执行 `shutdown` 命令

# Arthas 进阶

## 1. 启动 demo

下载demo-arthas-spring-boot.jar，再用java -jar命令启动：

```bash
wget https://github.com/hengyunabc/katacoda-scenarios/raw/master/demo-arthas-spring-boot.jar

# 注意切换root 否则因为程序中占用的端口为80端口，需要root权限才可以启动
java -jar demo-arthas-spring-boot.jar
```

`demo-arthas-spring-boot`是一个很简单的spring boot应用，源代码：[git地址](https://github.com/hengyunabc/spring-boot-inside/tree/master/demo-arthas-spring-boot)

启动之后，可以访问80端口：localhost

```
输出内容为
Date: 2019-10-28
Time: 23:13:23
Message: Hello World

    ok 200, user 1：user 1
    Error 500, user 0：user 0
    Error 404, a.txt：a.txt
    Error 401, admin：admin
    hello jsp：hello jsp

```

在 `/home/houzhenguo/server/arthas` 这个路径执行 
```
java -jar arthas-boot.jar --target-ip 0.0.0.0
```
否则参考 初级版本的启动，需要下载的过程。不过对于内网机来说，这些jar包需要提前下载完成。

# 查看JVM信息

下面介绍Arthas里查看`JVM`信息的命令。

## 1. sysprop

`sysprop` 可以打印所有的System Properties信息。

也可以指定单个key： sysprop java.version

也可以通过grep来过滤：sysprop | grep user

可以设置新的value： sysprop testKey testValue

## 2. sysenv
sysenv 命令可以获取到环境变量。和sysprop命令类似。

## 3. jvm

jvm 命令会打印出JVM的各种详细信息。

# Tips

为了更好使用Arthas，下面先介绍Arthas里的一些使用技巧。

## 1. Help
Arthas里每一个命令都有详细的帮助信息。可以用`-h`来查看。帮助信息里有EXAMPLES和WIKI链接。

比如： sysprop -h

## 2. 自动补全

Arthas支持丰富的自动补全功能，在使用有疑惑时，可以输入`Tab`来获取更多信息。

比如输入 sysprop java. 之后，再输入Tab，会补全出对应的key：

```bash
$ sysprop java.
java.runtime.name             java.protocol.handler.pkgs    java.vm.version
java.vm.vendor                java.vendor.url               java.vm.name
```

## 3. readline的快捷键支持

Arthas支持常见的命令行快捷键，比如`Ctrl + A`跳转行首，`Ctrl + E`跳转行尾。

更多的快捷键可以用 keymap 命令查看。

## 4. pipeline

Arthas支持在pipeline之后，执行一些简单的命令，比如：

```bash
sysprop | grep java
sysprop | wc -l

```

# Ognl
## 调用static函数

```bash
ognl '@java.lang.System@out.println("hello ognl")'
```

可以检查Terminal 1里的进程输出，可以发现打印出了hello ognl。
## 获取静态类的静态字段


# 其他

辅助命令：
```bash
 netstat -nalp | grep 80
```
参考：

[Springboot无法启动80端口问题](https://community.atlassian.com/t5/Jira-questions/Failed-to-initialize-connector-Connector-HTTP-1-1-80-after/qaq-p/647418)

copyright @houzhenguo 20191029