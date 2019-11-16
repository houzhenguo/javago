# Spring学习笔记（五）基于注解的IOC

## 五、基于注解的IOC实例

基于xml实例，代码重复太多，就不贴出来了，主要是对于注解的应用，建议也手动敲一遍代码，对记忆和理解的加深有帮助。

1. Spring 中的新注解：
    - @Configuration
        - 作用：指定当前类是一个配置类
        - 细节：当配置类作为 AnnotationConfigurationApplicationContext 对象创建的参数时，该注解可以不写
    - @ComponentScan
        - 作用：用于通过注解指定 Spring 在创建容器时要扫描的包
        - 属性：

            value : 它和 basepackages 的作用是一样的，都是用于指定创建容器时要扫描的包

            使用此注解就等同于在 xml 中配置了：

            <context:component-scan base-package="com.greyson"></context:component-scan>

    - @Bean
        - 作用：用于把当前方法的返回值作为 bean 对象放入 Spring 的IOC容器中
        - 属性：

            name : 用于指定 bean 的 id，当不写时，默认值为当前方法的名称

        - 细节：

            当我们使用注解配置方法时，如果方法有参数，Spring 框架会去容器中查找有没有可用的 bean 对象，

            查找的方式和 Autowired 注解的作用是一样的

    - @I*mport*
        - 作用：用于导入其他的配置类
        - 属性：

            value : 用于指定其他配置类的字节码

            当我们使用 Import 的注解之后，有 Import 注解的类就是父配置类，而导入的都是子配置类

    - @Properties
        - 作用：用于指定 properties 文件的位置
        - 属性：

            value : 指定文件的名称和路径

            关键字：classpath , 表示类路径下

# [Spring学习笔记（六）Spring整合Junit](./Spring学习笔记（六）Spring整合Junit.md)