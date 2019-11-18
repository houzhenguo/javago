
# Robot
## Resource 



1. ark_resource/resource/develop/common/data 都是以json的格式配置的。
2. 配置在 main -> reloadAll -> AllCfgs的构造方法中加载的。数据校验在 onLoad方法中校验，以error日志的形式。

数据格式有 intMap 形式，有 class 对象的形式。

## RobotManager

1. 所有的bean 在 robot.bean
2. RobotManager中的构造方法中 反射得到所有 robot bean 的实例
3. start的时候，同时启动 ThreadCount个线程。同时传入robot的userId -> bind Server -> 启动Client(异步操作在NetManager中) -> getConnector.addListener->complete->exception
4. 在robot的prop中的 robotLinks中配置连接服务器的地址列表。可以目前是写死的。==-》
5. 重点： RobotBase robot.ai.Robot这两个类
          自己的bean需要在RobotBase的CmdTypes中注册。

6. ActorThread 是发送协议的底层线程。
        a. 定时任务线程 守护线程
        b. postDelayMsg 延迟投放消息, 将消息放入到队列中。
        c. 死循环 blockingQueue中take
        d. 根据msg Class 取出相应的dispatcher策略。RobotThread的构造方法中 RobotThread 传入了策略
        e. RobotThread 中 onRobotEvent .getMsgHandler, on()方法支持机器人。


## RobotManager 处理消息的过程

1. RobotManager.messageReceived()
        a. session中塞这一个Robot实例
        b. 根据 robot的id找到相应的Thread进行处理 (自定义的 RobotThread)
2.  RobotThread.postDelayMsg 发送消息。 new 一个RobotEvent 对象 。RobotEvent 消息实例会进入到 ActorThread 的 msgQueue 队列中
3. ActorThread 中 run 方法 死循环 去 take。获取 到RobotEvent 。在dispather中获取到 RobotEvent.class 进行处理。 处理方法在 RobotThread 
   的构造方法中传入。 on(RobotEvent.class, this:: onRobotEvent) 意味着所有的消息在 RobotEvent进行处理。最终都是调用 RobotEvent 的 onRobotEvent方法。
4. RobotThread.onRobotEvent(RobotEvent event) 
        a. pre(msg) 预处理方法 -> 生成相应的 Bean<?> 这个是在robot.bean下面定义的。
        b. 查找相应bean的on方法。这个东西是在 RobotManager中通过反射生成的。然后进行 反射调用。
        c. 接下来就是 RobotBase中 on login 一系列的流程处理。每个on有自己发送消息的流程。
        
5. Robot 当 最后一个受到 服务器发送的 CmdEnterworld_Re 时候，进行move() 操作 。满血 满速度 无饥饿


## 在 启动时候配置 自己需要使用的 Robot的类是哪个
1. 这个是在 RobotManager 的 onAddSession 中实现。将 相关 的Robot类绑定在Session上面。


## 如何在 robot.bean 下面生成相应的 Bean<?>

CmdEnterworld_Re 在哪里定义的没有搜到 ====================================》Q


## 如何查看机器人 发送协议的效果

  





## Move的过程

1. 在 ../res/robot/move.log 中解析出各个参数。



## GM命令



Q ：

1. 这个log是如何生成的？
2. RobotBase 中的 CmdDebugGsCmd 中的Id 以及参数的含义
3. 看代码，发现 每次都有一个 new P().submit();的过程，是多线程处理么。 FDebugCmd 比如 onLeaveTeam submit两次?
4. Procedure 中 lock之前都会 unlock ,这时候不是别人主动释放的锁？ 工会的锁，role的锁？
-14 182

侯振国：
1. 熟悉服务器 机器人部分的代码
2. 体验游戏
3. 看服务器其他模块的代码 协议处理流程部分


## 学到的东西 

Prop = property 参数 


## 其他

协议开始都是在 ds.handler.link下面。

NetManager -> messageReceived -> onProcess

所有的协议都在 AlllBeans下面，自动生成。

lib.util.TimeUtil.curMs(); 使用的时候尽量使用这个，因为底层使用的是 System.currentMis() + 时间; 支持调时间。
