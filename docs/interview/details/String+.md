
# String + 底层原理

[博客链接](https://blog.csdn.net/qq_36771269/article/details/80818200)

## 说起String拼接，大家会想到几个对比：
String：常量，不可变，不适合用来字符串拼接，每次都是新创建的对象，消耗较大。
StringBuffer：适合用来作字符串拼接
StringBuilder：JDK1.5引入，适合用来作字符串拼接，与StringBuffer区别是他不是线程安全的

## 接下来进入正题String”+”拼接底层实现原理
曾见过这样一道题：

```java
String s=null;
s=s+"abc";
System.out.println(s);
```

这道题答对结果的很少，我第一次也没有答对，后来是在编译器上运行之后才知道自己错了。

> String拼接，有字符串变量参与时，中间会产生StringBuilder对象（JDK1.5之前产生StringBuffer）
字符串拼接原理：运行时， 两个字符串str1, str2的拼接首先会调用 String.valueOf(obj)，这个Obj为str1，而String.valueOf(Obj)中的实现是return obj == null ? “null” : obj.toString(), 然后产生StringBuilder， 调用的StringBuilder(str1)构造方法， 把StringBuilder初始化，长度为str1.length()+16，并且调用append(str1)！ 接下来调用StringBuilder.append(str2), 把第二个字符串拼接进去， 然后调用StringBuilder.toString返回结果！

## StringBuilder(str) 底层调用

```java
public StringBuilder(String str) {
    super(str.length() + 16);
    append(str);
}
```

## StringBuilder.toString 底层调用

```java
@Override
public String toString() {
 // Create a copy, don't share the array
    return new String(value, 0, count);
}
```

所以答案就是：StringBuilder(“null”).append(“abc”).toString();

---

str += "c";等效于：str = new StringBuffer(str).append("c").toString();
虽然编译器对字符串加号做了优化，它会用StringBuffer的append方法进行追加。再是通过toString方法转换成String字符串的。
它与纯粹的append方法是不同的：
一是每次都要创建一个StringBuilder对象；
二是每次执行完毕都要调用toString方法将其转换为字符串。