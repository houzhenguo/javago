
参考： https://golang.design/go-questions/slice/vs-array/
https://github.com/lifei6671/interview-go/blob/master/question/q009.md

## 在 golang 协程和channel配合使用
写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
```golang
func main() {
    out := make(chan int)
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            out <- rand.Intn(5)
        }
        close(out)
    }()
    go func() {
        defer wg.Done()
        for i := range out {
            fmt.Println(i)
        }
    }()
    wg.Wait()
}
```

## 1. 数组和切片有什么异同
1. slice 底层数据是数组 -> slice 是对数组的封装 -> 两者都是通过下标来访问单个元素
2. 数组是定长的 而slice可以动态扩容.
3. 数组是连续的内存, slice 是结构体,长度，容量，底层数组.
4. slice底层可以指向同一个数组，所以操作一个slice，可能堆其他的也有影响。

## 2. 切片的容量是怎么增长的？
1. append 之后才可能引起扩容.append 可变参数. 根据golang版本小于 18的，小于1024的每次扩容2倍，超过1024 每次是1.25倍, 有内存对齐的过程，所以扩容之后是大于等于 目标值.
2. len,cap,array -> 底层是指针 -> 所以修改数据会修改到底层. -> 都是 值传递.
3. 其实都是复制的值传递，但是你要是修改数组底层数组，用的同一个数组地址，所以底层数据会修改；但是你要是修改slice的什么的，这个不会改到外面的，因为他已经复制出来了。 https://golang.design/go-questions/slice/as-func-param/

Go 的 slice 底层数据结构是由一个 array 指针指向底层数组，len 表示切片长度，cap 表示切片容量。slice 的主要实现是扩容。对于 append 向 slice 添加元素时，假如 slice 容量够用，则追加新元素进去，slice.len++，返回原来的 slice。当原容量不够，则 slice 先扩容，扩容之后 slice 得到新的 slice，将元素追加进新的 slice，slice.len++，返回新的 slice。对于切片的扩容规则：当切片比较小时（容量小于 1024），则采用较大的扩容倍速进行扩容（新的扩容会是原来的 2 倍），避免频繁扩容，从而减少内存分配的次数和数据拷贝的代价。当切片较大的时（原来的 slice 的容量大于或者等于 1024），采用较小的扩容倍速（新的扩容将扩大大于或者等于原来 1.25 倍），主要避免空间浪费，网上其实很多总结的是 1.25 倍，那是在不考虑内存对齐的情况下，实际上还要考虑内存对齐，扩容是大于或者等于 1.25 倍。

（关于刚才问的 slice 为什么传到函数内可能被修改，如果 slice 在函数内没有出现扩容，函数外和函数内 slice 变量指向是同一个数组，则函数内复制的 slice 变量值出现更改，函数外这个 slice 变量值也会被修改。如果 slice 在函数内出现扩容，则函数内变量的值会新生成一个数组（也就是新的 slice，而函数外的 slice 指向的还是原来的 slice，则函数内的修改不会影响函数外的 slice。）


## hash表
1. 冲突的时候 -> 链表法和开放寻址法
2. 源码不想读了。
3. val,ok := map["key"] -> 不同的返回值 他相当于做了函数的重载。
4.  float可以作为key麽？golang只要是可以比较的，都可以作为key。除了 slice,map,functions几种类型。 包含 bool,number,pointer,string,channel,interface,struct,
任何都可以作为 value.
5. 可以边遍历边删除麽？


## 基础
1. new make区别
相同点: 给变量分配内存
不同点: new给string,int 和数组分配空间,make给切片,map,channel分配内存. 
   new 返回指向变量的指针，make返回变量本身. 
   make 分配的空间会被初始化.
2. for range 的时候地址如果发生变化？
  for a,b := range c,在 遍历的时候a，b在内存中是一份地址，只是把值赋值给这个地址。a,b在内存的地址始终不变，所以我们在for 开携程，不要直接把a或者b的地址传递给协程，应该创建临时变量。
3. defer的一些问题
a. defer释放资源，关闭文件/关闭连接/捕获panic defer延迟函数
b. defer的调用是 单链表，每次都是将 defer实例插在头部，函数结束再一次从头部取出，从而形成后进先出的效果 ; 首先return/ return value/defer

```golang
package main

import "fmt"

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 1
	}()
	return i //或者直接写成return
}
func main() {
	fmt.Println("return:", b()) // 2
}
```

4. uint 类型溢出问题  超过最大存储值如uint8最大是255
5. rune 
a. golang的字符串底层是通过byte数组，中文在unicode占2个字节，在utf-8编码下占3个字节，golang的默认编码是 utf-8
byte等同于 int8常用来处理 ascii字符，rune 相当于int32常用来处理unicode 和utf-8字符,
"hello 你好"用len 就是12byte，用rune就是 8,比较符合常理。
b. 遍历字符串用for range 底层 会将 str转换 rune类型切片.
6. golang 是怎么解析tag的
a. 通过反射实现 -> 反射可以用来修改一个变量的值,前提是这个值可以被修改。
```golang
type User struct {
	name string `json:name-field`
	age  int
}
func main() {
	user := &User{"John Doe The Fourth", 20}

	field, ok := reflect.TypeOf(user).Elem().FieldByName("name")
	if !ok {
		panic("Field not found")
	}
	fmt.Println(getStructTag(field))
}

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}
```

7. 讲讲 Go 的 select 底层数据结构和一些特性？
go 的 select 为 golang 提供了多路 IO 复用机制，和其他 IO 复用一样，用于检测是否有读写事件是否 ready。linux 的系统 IO 模型有 select，poll，epoll，go 的 select 和 linux 系统 select 非常相似。

select 结构组成主要是由 case 语句和执行的函数组成 select 实现的多路复用是：每个线程或者进程都先到注册和接受的 channel（装置）注册，然后阻塞，然后只有一个线程在运输，当注册的线程和进程准备好数据后，装置会根据注册的信息得到相应的数据。

select 的特性

1）select 操作至少要有一个 case 语句，出现读写 nil 的 channel 该分支会忽略，在 nil 的 channel 上操作则会报错。

2）select 仅支持管道，而且是单协程操作。

3）每个 case 语句仅能处理一个管道，要么读要么写。

4）多个 case 语句的执行顺序是随机的。

5）存在 default 语句，select 将不会阻塞，但是存在 default 会影响性能。

8. 单引号，双引号，反引号的区别？

单引号，表示byte类型或rune类型，对应 uint8和int32类型，默认是 rune 类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。

双引号，才是字符串，实际上是字符数组。可以用索引号访问某字节，也可以用len()函数来获取字符串所占的字节长度。

反引号，表示字符串字面量，但不支持任何转义序列。字面量 raw literal string 的意思是，你定义时写的啥样，它就啥样，你有换行，它就换行。你写转义字符，它也就展示转义字符。



## Map
1. map不是线程安全的,并发读写可能造成panic.
2.  map 中删除一个 key，它的内存会释放么？  不会。需要重建 ？设置成nil会被回收。
3. nil map不能赋值， 空map 取出来的东西为空。
4. slices/maps/functions 都不能作为 map的key.


## context
1. context 中包含 deadline，Done,Error,Value 方法返回一个time.Time.主要用来传递上下文，
比如 我们region 和 header信息的传递。 context.Background(); 超时等待

## channel
1. 底层数据结构:buf，发送队列，接收队列，lock。
2. 读写nil 管道会永久阻塞/关闭的管道可以继续读/往关闭的管道写会panic/关闭为nil的管道会panic
3. 无换冲和有换冲的区别？
    a. 没有缓冲区，冲管道读取数据会阻塞 -> 直到有协程向管道中写入数据
    b. 向管道中写入数据会阻塞 -> 直到有协程从管道中读取数据
    c. 有缓冲区: 缓冲区空了，读数据阻塞。缓冲区满了，写数据阻塞。
4. 使用场景 -> 消息传递/消息过滤/限流

## 进程/线程/协程 的区别 
1. 进程 -> 引用程序的启动实例
2. 线程 -> 从属于进程,线程是CPU调度的基本单位 -> 

## 抢占式调度是如何抢占的。
像操作系统负责线程的调度一样，Go的runtime 要负责goroutine的调度.时间片的抢夺，避免CPU被少数线程占用。

## 锁相关
1. mutex,  将 共享变量 读写放到一个 goroutine中，其他goroutine通过channel进行读写操作./可以使用个数为1的semaphore实现互斥
2. atomic 提供原子的读取
3. 原子锁 对某值的操作/互斥锁 -> 对某些关键位置的 操作.


