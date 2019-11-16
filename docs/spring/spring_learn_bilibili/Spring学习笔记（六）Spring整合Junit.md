# Spring学习笔记（六）Spring整合Junit

## 六、Spring 整合 Junit（后续补充）

### 1、Spring 整合 Junit 的配置过程：

1. 导入 Spring 整合 Junit 的 jar ( 坐标 )
2. 使用 Junit 提供的一个注解把原有的 main 方法替换了，替换成 Spring 提供的

    @Runwith

3. 告知 Spring 的运行器， Spring 和 ioc 创建是基于 xml 还是注解的，并且说明位置，用到的注解如下

    @ContextConfiguration

    Locations : 指定 xml 文件的位置，加上 classpath 关键字，表示在类路径下

    classes : 指定注解类所在地位置

4. 使用@Autowired 给测试类中的变量注入数据

# [Spring学习笔记（七）动态代理分析](Spring学习笔记（七）动态代理分析.md)