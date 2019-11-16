# Spring学习笔记（七）动态代理分析

## 七、动态代理分析

### 1、什么是代理？

简单理解，本来厂商可以自产自销，但是由于各种开销，最后厂商选择只生产产品，销售则交由各级经销商完成。

![](-d7c45db4-1f69-4b5c-b5f5-a9624176e6f8.png)

- 特点：字节码随用随创建，随用随加载
- 作用：不修改源码的基础上对方法增强
- 分类：

    基于接口的动态代理

    基于子类的动态代理

### 2、基于接口的动态代理

1. 基于接口的动态代理：

    涉及的类：Proxy

    提供者：JDK官方

2. 如何创建代理对象：

    使用Proxy类中的newProxyInstance方法

3. 创建代理对象的要求：

    被代理类最少实现的一个接口，如果没有则不能使用

4. newProxyInstance方法的参数：

    ClassLoader : 用于加载代理对象字节码，和被代理对象使用相同的类加载器，固定写法

    Class [ ] : 用于让代理对象和被代理对象有相同的方法，固定写法

    InvocationHandler : 用于提供增强的代码

    它是让我们写如何代理。我们一般是写一个该接口的实现类，通常是匿名内部类，但不是必须的。此接口的实现类都是谁用谁写。

- 生产厂家接口IProducer

        /**
         * 对生产厂家要求的接口
         */
        public interface IProducer {
            /**
             * 销售
             * @param money
             */
            public void saleProduct(float money);
        
            /**
             * 售后
             * @param money
             */
            public void afterService(float money);
        }

- 生产者

        /**
         * 一个生产者
         */
        public class Producer implements IProducer {
        
            /**
             * 销售
             * @param money
             */
            public void saleProduct(float money) {
                System.out.println("销售产品，并拿到钱：" + money);
            }
        
            /**
             * 售后
             * @param money
             */
            public void afterService(float money) {
                System.out.println("提供售后服务，并拿到钱：" + money);
            }
        }

- 消费者

        /**
         * 模拟一个消费者
         */
        public class Client {
            public static void main(String[] args) {
                final Producer producer = new Producer();
        
                IProducer proxyProducer = (IProducer) Proxy.newProxyInstance(producer.getClass().getClassLoader(),
                        producer.getClass().getInterfaces(), new InvocationHandler() {
                            /**
                             * 作用：执行被代理对象的任何接口方法都会经过该方法
                             * 方法参数的含义：
                             * @param proxy         代理对象的含义
                             * @param method        当前执行的方法
                             * @param args          当前执行方法的参数
                             * @return              和被代理对象方法有相同的返回值
                             * @throws Throwable
                             */
                            public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                                // 提供增强的代码
                                Object returnValue = null;
                                // 1.获取方法执行的参数
                                Float money = (Float)args[0];
                                // 2.判断当前方法是不是销售
                                if ("saleProduct".equals(method.getName())) {
                                    returnValue = method.invoke(producer,money * 0.8f);
                                }
                                return returnValue;
                            }
                        });
            }
        }

### 3、基于子类的动态代理

1. 基于子类的动态代理：

    涉及的类：Enhancer

    提供者：第三方 cglib 库

2.  如何创建代理对象：

    使用 Enhancer 类中的 create 方法

3. 创建代理对象的要求：

    被代理类不能是最终类

4. create 方法的参数：

    Class : 它是用于被指定代理对象的字节码

    callback : 用于提供增强的代码

    它是让我们写如何代理。我们一般是写一个该接口的实现类，通常是匿名内部类，但不是必须的。此接口的实现类都是谁用谁写。我们一般写的都是该接口的子实现类：MethodInterCeptor

- 生产者

        public class Producer {
        
            /**
             * 销售
             * @param money
             */
            public void saleProduct(float money) {
                System.out.println("销售产品，并拿到钱：" + money);
            }
        
            /**
             * 售后
             * @param money
             */
            public void afterService(float money) {
                System.out.println("提供售后服务，并拿到钱：" + money);
            }
        }

- 消费者

        /**
         * 模拟一个消费者
         */
        public class Client {
            public static void main(String[] args) {
                final Producer producer = new Producer();
                Producer cglibProducer = (Producer) Enhancer.create(producer.getClass(), new MethodInterceptor() {
                    public Object intercept(Object proxy, Method method, Object[] args, MethodProxy methodProxy) throws Throwable {
                        // 提供增强的代码
                        Object returnValue = null;
                        // 1.获取方法执行的参数
                        Float money = (Float)args[0];
                        // 2.判断当前方法是不是销售
                        if ("saleProduct".equals(method.getName())) {
                            returnValue = method.invoke(producer,money * 0.8f);
                        }
                        return returnValue;
                    }
                });
                cglibProducer.saleProduct(12000f);
            }
        }

# [Spring学习笔记（八）AOP概念](Spring学习笔记（八）AOP概念.md)