# Spring学习笔记（二）Bean的装配与管理

## 二、Spring 中 Bean 的细节

### （一）、三种创建 bean 对象的方式

1.  使用默认构造函数创建

    在spring的配置文件中，使用 id 和 class 属性之后，且没有其他属性和标签时，采用的就是默认构造函数创建 bean 对象，此时如果类中没有默认构造函数，则对象无法创建。

        <bean id = "accountService" class = "com.itheima.service.impl.AccountServiceImpl"></bean>

2. 使用普通工厂中的方法创建对象（使用某个类中的方法创建对象，并存入 Spring容器）,如下

        /**
         *模拟一个工厂类，该类可能存在于jar包中，无法通过修改源码的方式来提供默认构造函数
         * 
         */
        public class InstanceFactory {
            public IAccountService getAccountService() {
                return new AccountServiceImpl();
            }
        }

    配置方式如下：

        <bean id = "instanceFactory" class = "com.itheima.factory.InstanceFactory"></bean>
            <bean id = "accountService" factory-bean="instanceFactory" factory-method="getAccountService"></bean>

3. 使用工厂中的静态方法创建对象（使用某个类中的静态方法创建对象，并存入spring容器），如下：

        public class StaticFactory {
            public  static IAccountService getAccountService() {
        
                return new AccountServiceImpl();
            }
        }

    配置方式如下：

        <bean id = "accountService" class = "com.itheima.factory.StaticFactory" factory-method="getAccountService"></bean>

### （二）、bean 的作用范围调整

1. bean 标签的 scope 属性

    作用：用于指定 bean 的作用范围

    取值：常用的就是单例和多例

    - singletond : 单例的（default） (常用)
    - prototype : 多例的 (常用)
    - request : 作用于 web 应用的请求范围
    - session : 作用于 web  应用的会话范围
    - global-session : 作用于集群的会话范围（全局会话范围），当不是集群范围时，它就是 session
    - gloabl-session 示意图：

        ![](../images/2.png)

2. bean对象的生命周期

    单例对象：(立即创建，启动)

    - 出生：当容器创建时发生
    - 活着：只要容器还在对象就一直活着
    - 死亡：容器销毁，对象消亡

    总结：单例对象的声明周期和容器相同

    多例对象：（使用创建）

    - 出生：当我们使用对象时 Spring 框架为我们创建
    - 活着：对象只要是在使用过程中就活着
    - 死亡：当对象长时间不用，且没有别的对象引用时，由 Java 的GC回收

```xml
        <bean id = "accountService" class = "com.itheima.service.impl.AccountServiceImpl" scope="singleton" init-method=
        "" destroy-method=""></bean>
```

# [Spring学习笔记（三）依赖注入](Spring学习笔记（三）依赖注入.md)