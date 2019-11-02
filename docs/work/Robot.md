
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