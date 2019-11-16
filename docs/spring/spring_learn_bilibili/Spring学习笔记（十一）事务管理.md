# Spring学习笔记（十一）事务管理

## 十一、Spring的事务管理

### （一）、声明式事务管理

1. Spring中基于 xml 的声明式事务控制配置步骤

    ( 1 ) 配置事务管理器

    ( 2 ) 配置事务的通知

    - 1 ) 导入事务的约束
    - 2 ) 使用tx:advice标签配置事务通知
        - 属性：
            - id：给事务通知起一个唯一标识
            - transaction-manager：给事务通知提供一个事务管理器引用
    - 3 ) 配置AOP中的通用切入点表达式
    - 4 ) 建立十五通知和切入点表达式的对应关系
    - 5 ) 配置事务的属性（在事务的通知 tx : advice标签的内部）
        - isolation：用于指定书屋的隔离级别，默认值为 DEFAULT，表示数据库的默认隔离级别
        - propagation：用于指定事务的传播行为。默认值是REQUIRED，表示一定会有事务，增删改的选择。查询方法可以选择SUPPORTS。
        - read-only：用于指定事务是否只读。只有查询方法才能设置为true。默认值是false，表示读写。
        - timeout：用于指定事务的超时时间，默认值是-1，表示永不超时。如果指定了数值，以秒为单位。
        - rollback-for：用于指定一个异常，当产生该异常时，事务回滚，产生其他异常时，事务不回滚。没有默认值。表示任何异常都回滚。
        - no-rollback-for：用于指定一个异常，当产生该异常时，事务不回滚，产生其他异常时事务回滚。没有默认值。表示任何异常都回滚。
2. 基于xml配置实例
    - bean.xml

            <?xml version="1.0" encoding="UTF-8"?>
            <beans xmlns="http://www.springframework.org/schema/beans"
                   xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                   xmlns:aop="http://www.springframework.org/schema/aop"
                   xmlns:tx="http://www.springframework.org/schema/tx"
                   xsi:schemaLocation="
                    http://www.springframework.org/schema/beans
                    http://www.springframework.org/schema/beans/spring-beans.xsd
                    http://www.springframework.org/schema/tx
                    http://www.springframework.org/schema/tx/spring-tx.xsd
                    http://www.springframework.org/schema/aop
                    http://www.springframework.org/schema/aop/spring-aop.xsd">
            
                <!--configure service layer-->
                <bean id="accountService" class="com.itheima.service.AccountServiceImpl">
                    <property name="accountDao" ref="accountDao"></property>
                </bean>
            
                <!-- 配置账户的持久层-->
                <bean id="accountDao" class="com.itheima.dao.impl.AccountDaoImpl">
                    <property name="dataSource" ref="dataSource"></property>
                </bean>
            
            
                <!-- 配置数据源-->
                <bean id="dataSource" class="org.springframework.jdbc.datasource.DriverManagerDataSource">
                    <property name="driverClassName" value="com.mysql.jdbc.Driver"></property>
                    <property name="url" value="jdbc:mysql://localhost:3306/eesy"></property>
                    <property name="username" value="root"></property>
                    <property name="password" value="HotteMYSQL"></property>
                </bean>
            
                <!--Declarative transaction management-->
                <bean id="transactionManager" class="org.springframework.jdbc.datasource.DataSourceTransactionManager">
                    <property name="dataSource" ref="dataSource"></property>
                </bean>
            
                <!--configure transaction advice-->
                <tx:advice id="txAdvice" transaction-manager="transactionManager">
                    <tx:attributes>
                        <tx:method name="transfer" propagation="REQUIRED" read-only="false"/>
                        <tx:method name="find*" propagation="SUPPORTS" read-only="true"></tx:method>
                    </tx:attributes>
                </tx:advice>
            
                <!--configure pointcut expression-->
                <aop:config>
                    <aop:pointcut id="pt1" expression="execution(* com.itheima.service.*.*(..))"></aop:pointcut>
                    <aop:advisor advice-ref="txAdvice" pointcut-ref="pt1"></aop:advisor>
                </aop:config>
            </beans>

    - IAccountDao

            public interface IAccountDao {
            
                /**
                 * 根据Id查询账户
                 * @param accountId
                 * @return
                 */
                Account findAccountById(Integer accountId);
            
                /**
                 * 根据名称查询账户
                 * @param accountName
                 * @return
                 */
                Account findAccountByName(String accountName);
            
                /**
                 * 更新账户
                 * @param account
                 */
                void updateAccount(Account account);
            }

    - AccountDaoImpl

            public class AccountDaoImpl extends JdbcDaoSupport implements IAccountDao {
            
                @Override
                public Account findAccountById(Integer accountId) {
                    List<Account> accounts = super.getJdbcTemplate().query("select * from account where id = ?",new BeanPropertyRowMapper<Account>(Account.class),accountId);
                    return accounts.isEmpty()?null:accounts.get(0);
                }
            
                @Override
                public Account findAccountByName(String accountName) {
                    List<Account> accounts = super.getJdbcTemplate().query("select * from account where name = ?",new BeanPropertyRowMapper<Account>(Account.class),accountName);
                    if(accounts.isEmpty()){
                        return null;
                    }
                    if(accounts.size()>1){
                        throw new RuntimeException("结果集不唯一");
                    }
                    return accounts.get(0);
                }
            
                @Override
                public void updateAccount(Account account) {
                    super.getJdbcTemplate().update("update account set name=?,money=? where id=?",account.getName(),account.getMoney(),account.getId());
                }
            }

    - Accont

            public class Account implements Serializable {
            
                private Integer id;
                private String name;
                private Float money;
            
                public Integer getId() {
                    return id;
                }
            
                public void setId(Integer id) {
                    this.id = id;
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
                            "id=" + id +
                            ", name='" + name + '\'' +
                            ", money=" + money +
                            '}';
                }
            }

    - IAccountService

            /**
             * Service layer interface of the account
             */
            public interface IAccountService {
            
                /**
                 * query account information by id
                 * @param accountId
                 * @return
                 */
                Account findAccountById(Integer accountId);
            
                /**
                 * transfer
                 * @param sourceName    transfer out account
                 * @param targetName    transfer to account
                 * @param Money     transfer amount
                 */
                void transfer(String sourceName, String targetName, Float Money);
            }

    - AccountServiceImpl

            public class AccountServiceImpl implements IAccountService{
            
                private IAccountDao accountDao;
            
                public void setAccountDao(IAccountDao accountDao) {
                    this.accountDao = accountDao;
                }
            
                @Override
                public Account findAccountById(Integer accountId) {
                    return accountDao.findAccountById(accountId);
                }
            
                @Override
                public void transfer(String sourceName, String targetName, Float money) {
                    System.out.println("transfer....");
                    //2.1根据名称查询转出账户
                    Account source = accountDao.findAccountByName(sourceName);
                    //2.2根据名称查询转入账户
                    Account target = accountDao.findAccountByName(targetName);
                    //2.3转出账户减钱
                    source.setMoney(source.getMoney()-money);
                    //2.4转入账户加钱
                    target.setMoney(target.getMoney()+money);
                    //2.5更新转出账户
                    accountDao.updateAccount(source);
            
                            int i=1/0;
            
                    //2.6更新转入账户
                    accountDao.updateAccount(target);
                }
            }

    - AccountServiceTest

            @RunWith(SpringJUnit4ClassRunner.class)
            @ContextConfiguration(locations = "classpath:bean.xml")
            public class AccountServiceTest {
            
                @Autowired
                private  IAccountService as;
            
                @Test
                public  void testTransfer(){
                    as.transfer("aaa","bbb",100f);
            
                }
            
            }

3. Spring 中基于注解的声明式事务控制配置步骤
    - 1 ) 配置事务管理器
    - 2 ) 开启spring对注解事务的支持
    - 3 ) 在需要事务支持的地方使用@Transactional注解

### （二）、编程式事务管理