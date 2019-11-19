# Clone

Java 中 clone为 shadow copy(浅拷贝),对于一些属性是对象的情况下只能copy引用，如果想做到deep copy ，需要递归实现。慎用.

> Java可以把对象序列化写进一个流里面，反之也可以把对象从序列化流里面读取出来，但这一进一出，这个对象就不再是原来的对象了，就达到了克隆的要求。

下面提供了 `org.apache.commons.lang3.SerializationUtils` 方法

```java
// 必须实现序列化
public class Foo implements Serializable {
    private int id;
    private String name;
    private Person person;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Person getPerson() {
        return person;
    }

    public void setPerson(Person person) {
        this.person = person;
    }
}
// Person
public class Person implements Serializable {
    private int age;

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        this.age = age;
    }
}

public class FooTest {
    public static void main(String[] args) {
        Person p = new Person();
        p.setAge(12);
        Foo foo = new Foo();
        foo.setId(1);
        foo.setName("tom");
        foo.setPerson(p);
        Foo foo1 = SerializationUtils.clone(foo); // 调用方法实现 deep copy
        System.out.println(foo1.getPerson().getAge());

        // 序列化测试
        byte[] bytes = SerializationUtils.serialize(foo);
        Foo foo2 = SerializationUtils.deserialize(bytes);
        System.out.println(foo2.getPerson().getAge() +"p2");
    }
}

```

about maven pom.xml

```xml
<dependency>
    <groupId>org.apache.commons</groupId>
    <artifactId>commons-lang3</artifactId>
    <version>3.8.1</version>
</dependency>
```

## 参考

[About Java cloneable](https://stackoverflow.com/questions/4081858/about-java-cloneable)

[Class SerializationUtils](https://commons.apache.org/proper/commons-lang/javadocs/api-3.4/org/apache/commons/lang3/SerializationUtils.html)

[EffectiveJava]()