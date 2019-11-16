# Spring学习笔记（十）JdbcTemplate

## 十、Jdbc Template

1. 编码方式
    - pom.xml

            <?xml version="1.0" encoding="UTF-8"?>
            <project xmlns="http://maven.apache.org/POM/4.0.0"
                     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                     xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
                <modelVersion>4.0.0</modelVersion>
            
                <groupId>com.greyson</groupId>
                <artifactId>jdbctTemplates</artifactId>
                <version>1.0-SNAPSHOT</version>
                <packaging>jar</packaging>
            
                <dependencies>
                    <dependency>
                        <groupId>org.springframework</groupId>
                        <artifactId>spring-context</artifactId>
                        <version>5.0.7.RELEASE</version>
                    </dependency>
            
                    <dependency>
                        <groupId>org.springframework</groupId>
                        <artifactId>spring-jdbc</artifactId>
                        <version>5.0.2.RELEASE</version>
                    </dependency>
            
                    <dependency>
                        <groupId>org.springframework</groupId>
                        <artifactId>spring-tx</artifactId>
                        <version>5.0.2.RELEASE </version>
                    </dependency>
            
                    <dependency>
                        <groupId>mysql</groupId>
                        <artifactId>mysql-connector-java</artifactId>
                        <version>5.1.6</version>
                    </dependency>
            
                </dependencies>
            
            
            </project>

    - Account

            
            /**
             * Account domain class
             *
             */
            public class Account implements Serializable {
            
                private Integer Id;
                private String name;
                private Float money;
            
                public Integer getId() {
                    return Id;
                }
            
                public void setId(Integer id) {
                    Id = id;
                }
            
                public String getName() {
                    return name;
                }
            
                public void setName(String name) {
                    this.name = name;
                }
            
                public Float getMoney() {
                    return money;
                }
            
                public void setMoney(Float money) {
                    this.money = money;
                }
            
                @Override
                public String toString() {
                    return "Account{" +
                            "Id=" + Id +
                            ", name='" + name + '\'' +
                            ", money=" + money +
                            '}';
                }
            }

    - JdbcTemplateDemo

            /**
             * The Basic Usage of JdbcTemplate
             */
            public class JdbcTemplateDemo {
            
                public static void main(String[] args) {
                    // prepare dataSource
                    DriverManagerDataSource dataSource = new DriverManagerDataSource();
                    dataSource.setDriverClassName("com.mysql.jdbc.Driver");
                    dataSource.setUrl("jdbc:mysql://localhost:3306/eesy");
                    dataSource.setUsername("root");
                    dataSource.setPassword("HotteMYSQL");
                    // 1. create the object of JdbcTemplate
                    JdbcTemplate jdbcTemplate = new JdbcTemplate();
                    jdbcTemplate.setDataSource(dataSource);
                    // 2. execute operation
                    jdbcTemplate.execute("insert into  account(name, money)values('aaa', 1000)");
                }
            }

2. 配置方式

    添加配置文件`ApplicationContext.xml`以及修改`JdbcTemplateDemo`

    - ApplicationContext.xml

            <?xml version="1.0" encoding="UTF-8"?>
            <beans xmlns="http://www.springframework.org/schema/beans"
                   xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                   xsi:schemaLocation="http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans.xsd">
            
                <!-- congigure JdbcTemplate-->
                <bean id="jdbcTemplate" class="org.springframework.jdbc.core.JdbcTemplate">
                    <property name="dataSource" ref="dataSource"></property>
                </bean>
            
                <!--configure dataSource-->
                <bean id="dataSource" class="org.springframework.jdbc.datasource.DriverManagerDataSource">
                    <property name="driverClassName" value="com.mysql.jdbc.Driver"></property>
                    <property name="url" value="jdbc:mysql://localhost:3306/eesy"></property>
                    <property name="username" value="root"></property>
                    <property name="password" value="HotteMYSQL"></property>
                </bean>
            </beans>

    - JdbcTemplateDemo2

            /**
             * The Basic Usage of JdbcTemplate
             */
            public class JdbcTemplateDemo {
            
                public static void main(String[] args) {
                    // prepare dataSource
                    DriverManagerDataSource dataSource = new DriverManagerDataSource();
                    dataSource.setDriverClassName("com.mysql.jdbc.Driver");
                    dataSource.setUrl("jdbc:mysql://localhost:3306/eesy");
                    dataSource.setUsername("root");
                    dataSource.setPassword("HotteMYSQL");
                    // 1. create the object of JdbcTemplate
                    JdbcTemplate jdbcTemplate = new JdbcTemplate();
                    jdbcTemplate.setDataSource(dataSource);
                    // 2. execute operation
                    jdbcTemplate.execute("insert into  account(name, money)values('aaa', 1000)");
                }
            }

# [Spring学习笔记（十一）事务管理](Spring学习笔记（十一）事务管理.md)