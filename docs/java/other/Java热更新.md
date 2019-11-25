[文章链接](https://ask.csdn.net/questions/772031)

可以参考 jane 框架中的

java.lang.instrument这个类很早就出了，redefineClasses这个方法可以更新方法级别的代码，但是不会触发一个类的初始化方法。游戏服务器的bug基本上只需要方法级别的更新就可以了，因为很多重大的bug基本在测试阶段被修复了，少量偶线的bug出现之后有些时候仅仅只需要改动一行代码却有时不得不需要重启所有应用程序，代价都似乎有点大。
现在开始从instrument入手

public static void premain(String agentArgs, Instrumentation inst);

public static void agentmain(String agentArgs, Instrumentation inst);
这两个方法都可以获取到Instrumentation对象，通过redefineClasses方法就可以更新对应的方法了
如果是调用premain这个方法 则需要在程序启动的时候指定对应的jar 同时项目里必须引用这个jar 因为获取到这个引用
java -javaagent:agent.jar -jar xx.jar 例如这样 执行这条命令后程序会查找 agent.jar 里的MANIFEST.MF文件里的Premain-Class参数 设置对应的代理类路径就行。例如：Premain-Class: com.test.JavaAgent 还需要加上 Can-Redefine-Classes: true这个参数才能调用redefineClasses方法。同时 可以拦截对应的类添加标记 做性能分析
agentmain 是通过指定对应的进程pid来加载对应的agent.jar 很典型的jconsule jvisualvm都是通过选择java进程来做一个简单的内存 和cpu分析 ，线程dump .Agent-Class 和上面一样

package com.test;

import java.lang.instrument.Instrumentation;

public class JavaAgent {
    public static Instrumentation INST = null;

    public static void premain(String agentArgs, Instrumentation inst){
        INST = inst;
    }
}

这里保存下引用就可以了 ，单独打成agent.jar

package com.test;

import java.io.FileInputStream;
import java.lang.instrument.ClassDefinition;

public class Test {

    public static void main(String[] args) {
        getInfo();
        testhot();
    }

    public final static void testhot(){
        new Thread(new Runnable() {

            @Override
            public void run() {
                // TODO Auto-generated method stub
                while(true){
                    try {
                    if(JavaAgent.INST != null){
                        FileInputStream is = new FileInputStream("/Users/xxxx/Downloads/Student.class");
                        byte[] array = new byte[is.available()];
                        is.read(array);
                        is.close();
                        Class cls = Class.forName("com.test.Student");
                        JavaAgent.INST.redefineClasses(new ClassDefinition(cls,array));
                    }
                        Thread.sleep(1000);
                    } catch (Exception e) {
                        // TODO Auto-generated catch block
                        e.printStackTrace();
                    }
                }
            }
        }).start();
    }




    public final static void getInfo(){
        new Thread(new Runnable() {

            @Override
            public void run() {
                // TODO Auto-generated method stub
                while(true){
                    //System.out.println("=============="+JavaAgent.INST);
                    new Student().getName();
                    try {
                        Thread.sleep(1000);
                    } catch (InterruptedException e) {
                        // TODO Auto-generated catch block
                        e.printStackTrace();
                    }
                }
            }
        }).start();

    }
}

上面就是一个很简单的例子，一个线程在不停的循环检测更新这个类，另外的一个线程在不停的输出这个对象对应的方法输出信息。
测试之后可以发现 ，方法的输出信息已经改变了。

如果你是单点部署的话，就用你找的这种方法。如果你是集群部署，就按照楼上说的，部署一台服务器启动一台，没有部署的可以一直处于运行状态。

部署时停掉其中一个更新另一个。这个启动了再更新另一个，中间始终有一个提供服务