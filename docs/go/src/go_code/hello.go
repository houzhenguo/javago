package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("hello world")
	// var n = 10
	// fmt.Println(n) // 定义的变量要用到，用不到会报错
	// 函数
	sum, sub := getVal(8, 5)
	sum2, _ := getVal(8, 5)
	fmt.Println("sum = ", sum, "sub=", sub)
	fmt.Println("sum = ", sum2)

	// 变量定义
	// 变量的定义和盛名
	var i int // 指定数据类型 使用默认值
	// i = 10
	fmt.Println("test 变量 默认值", i)
	var j = 10.22 // 自行推导
	fmt.Println("自行推导", j)
	x, y, z := "zhangsan", 111, 33.3 // 一次性声明三个变量
	fmt.Println("x = ", x, "y=", y, "z=", z)

	var str = "hello" + "world"
	fmt.Println(str)

	var ui uint32
	ui = 2222
	fmt.Println("ui", ui, "占用字节数：", unsafe.Sizeof(ui))

	// 显式转换
	var i1 int32 = 100
	fmt.Println(" i1 占用字节数：", unsafe.Sizeof(i1)) // 4
	var n1 = float64(i1)
	fmt.Println(" n1 占用字节数：", unsafe.Sizeof(n1)) // 8

	// function
	var result = calBill(4, 5)
	fmt.Println("total price is", result)

	var reF = calBillFloat(2.5, 4.0)
	fmt.Println("float is", reF)
	var reIn = calBillIn(3, 4)
	fmt.Println("in result is", reIn)

	// 函数返回多值情况
	var vSum, vSub = returnMultiple(4, 5)
	var vSum1, _ = returnMultiple(4, 5)
	var _, vSub2 = returnMultiple(4, 5) // 占位符
	fmt.Println("vsum = ", vSum, "vSub", vSub)
	fmt.Println("vsum1 = ", vSum1, "vSub1", vSub2)
	// 测试条件
	var flag = testCondition(5)
	fmt.Println("flag is ", flag)
	// 测试for
	var total = testFor(10)
	fmt.Println("total is", total)
	// 测试switch
	var tw = testSwitch1(5)
	fmt.Println("switch is ", tw)
	// 测试数组
	nums := [...]int{2, 3, 4, 5, 6}
	testArr(nums)
	// 测试分片
	testSlince()
	// 测试pointer
	testPointer()
	// 测试struct
	testPersonStruct()
	// 测试方法体
	testMethod()
	// 测试接口
	testInterface()
	// 测试协程
	testThread()
	// 测试channel
	testChannel()
	// 测试waitGroup
	testWaitGroup()
}

/**
变量 + 返回值
*/
func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

// 普通
func calBill(price int, num int) int {
	return price * num
}

// float 转换
func calBillFloat(price float64, num int) float64 {
	return price * float64(num)
}

// 入参
func calBillIn(price, num int) int {
	return price * num
}

// 返回多个值
func returnMultiple(a, b int) (int, int) {
	return a + b, a - b
}

// test if
func testCondition(num int) int {
	if num == 5 {
		return 25
	} else if num == 6 {
		return 36
	} else { // 注意else 的位置 在 } 的后面一行
		return 1
	}
}

func testFor(num int32) int32 {
	var total int32
	var i int32
	for i = 1; i <= num; i++ {
		total = total + num
	}
	return total
}
func testSwitch(num int) string {
	switch num {
	case 1, 5: // 支持多值
		return "hello 1"
	case 2:
		return "hello 2"
	case 3:
		return "hello 3"
	default:
		return "default"
	}
}
func testSwitch1(num int) string {
	var result = ""
	switch {
	case num > 0: // 支持条件表达式 但是switch 后面不需要写
		result = result + "hello >0"
		fallthrough // 继续往下执行 fallthrough 应该放在 case 的最后一个
	case num >= 5 && num < 10:
		result = result + "hello 5-10"
	default:
		return "default"
	}
	return result
}

// 注意 [5]int 和[25]int 是不同的类型，可以使用slices调整这个问题
func testArr(nums [5]int) { // 数组这里要指定长度，这里不太好
	for i := 0; i < len(nums); i++ {
		fmt.Println("i is", i, ";val is ", nums[i])
	}
	// 第二种range的遍历方式 类似Java中的 foreach
	for i, v := range nums { // 可以使用 for _,v
		fmt.Println("i is :", i, ";v is", v)
	}
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

// 切片 类似python中的
func testSlince() {
	a := [5]int{1, 2, 3, 4, 5} // 操作的是原来的数组
	var b []int = a[1:4]       // 左闭右开
	a[1] = 66                  // 对切片的任何修改都会反应的底层数组之上
	dslice := a[2:5]
	for i := range dslice {
		dslice[i]++ // 作用于之前的数组上 每个元素 +1
	}
	fmt.Println("test slice ", b)
	fmt.Println("test dslice ", a)
}

// test 指针
func testPointer() {
	// *T 表示该变量存储的是指针 但是有些情况下可以省略
	b := 233
	var a *int = &b
	fmt.Println("b value is", b, "address is ", a)
	// test pointer nil
	c := 23
	var d *int
	if d == nil { // pointer 的null 为 nil
		fmt.Println("d is", d)
		d = &c
		fmt.Println("d is", d)
	}
	// 指针的解引用
	fmt.Println("指针的解引用 值是多少呢 ", *a) // 使用 * 进行解引用
	// 使用指针引用修改之前的值
	fmt.Println("a 之前 所引用地址 b 的值是", b)
	*a++ // 可以进行修改
	fmt.Println("a 之后 所引用地址 b 的值是", b)
	// change Value
	f := 33
	changeVal(&f)
	fmt.Println("f after change is", f)
}

// 向函数传递指针参数 对于数组来说，还是尽量传递切片
func changeVal(val *int) {
	*val++
}

// struct 结构体
type Person struct {
	name string
	age  int
	man  bool
}
type Address struct {
	province string
	city     string
}
type Student struct {
	name    string
	address Address
}

func testPersonStruct() {
	p1 := Person{ // 初始化的第一种方式
		name: "zhangsan",
		age:  22,
		man:  true, // 注意逗号

	}
	fmt.Println("person name is , age is, ", p1.name, p1.age)
	p2 := Person{"lisi", 15, false} // 注意这里是大括号
	fmt.Println("p2 is", p2)
	// 创建匿名结构体 ？
	p3 := struct {
		name string
		age  int
	}{
		name: "lidan", // 如果不进行赋值，则是 ""空字符串
		age:  34,
	}
	fmt.Println("p3 匿名结构体", p3)
	// 访问person中的字段
	fmt.Println("p3 的 name是", p3.name)
	// 结构体中的指针
	p4 := &p3
	p3.name = "wangdan"
	fmt.Println("p4 指针结构体", p4)
	// test 嵌套结构
	s1 := Student{
		name: "wang",
		address: Address{
			province: "shand",
			city:     "qd",
		},
	}
	// 如果Address 在student中是匿名字段，则可以提升为 s1.city
	fmt.Printf("student is %s city is %s", s1.name, s1.address.city)
}

// 方法
// 1.方法其实是一个函数
// 2.方法和函数不同的点: 入参前后
func (p Person) sayhello() {
	fmt.Println("one person name is", p.name, "say hello")
}

// 对于值类型的 内部改动对外不可见
func (p Person) changename() {
	p.name = "wb"
}

// 对于指针类型的是改变的同一个地址
func (p *Person) changeAge() {
	p.age = 27
}
func testMethod() {
	p1 := Person{
		name: "hzg",
		age:  12,
		man:  true,
	}
	p1.sayhello()
	fmt.Println("p before", p1)
	p1.changename()
	fmt.Println("p after change name", p1)
	p1.changeAge()
	fmt.Println("p after change age", p1)
}

// 接口 interface

type Phone interface {
	call()
}
type HuaweiPhone struct {
}
type ApplePhone struct {
}

func (huawei HuaweiPhone) call() {
	fmt.Println("I am huawei")
}
func (apple ApplePhone) call() {
	fmt.Println("I am Apple")
}
func testInterface() {
	var phone Phone
	phone = HuaweiPhone{}
	phone.call()
	// test interface
	phone = ApplePhone{}
	phone.call()
}

// test go 协程
func testThread() {
	fmt.Println("main begin")
	go testHello()

	go testThreadNum()
	go testThreadNum2()
	time.Sleep(20 * time.Second) // sleep 1s
	fmt.Println("main end")
}
func testHello() {
	fmt.Println("test hello 我是单独的协程")
}

// 并发启动多个协程
func testThreadNum() {
	for i := 0; i < 20; i++ {
		fmt.Println("num2 is", i)
		time.Sleep(50 * time.Millisecond)
	}
}
func testThreadNum2() {
	for i := 0; i < 20; i++ {
		fmt.Println("num1 is", i)
		time.Sleep(50 * time.Millisecond)
	}
}

// test channel
// 管道 T
// data :=<- a // 读取信道a
// a <-data 写入信道a
// 发送+ 接收阻塞 ？ 等待有数据被消费 或者等待数据被生产
func testChannel() {
	var c chan bool
	if c == nil {
		fmt.Println("channel is nil")
		c = make(chan bool)
		//fmt.Printf("Type of c is %T", c) // chan int
	} else {
		fmt.Println("channel is not nil")
	}

	go testChannelBlock(c)
	<-c
	fmt.Println("start test channle main testChannel done")
	s := make(chan int)
	go testChannelCal(s)
	go testChannelRead(s)

	// 测试有缓冲的信道
	// 如果在主线程中进行 channel2中的内容则会报 deadlock
	chb := make(chan int, 5)
	go testBufferChannel2(chb)
	go testBufferChannel(chb)

	time.Sleep(10 * time.Second)
	fmt.Println("main testChannel done")
}

// 使用channel 进行通信 阻塞主线程
func testChannelBlock(flag chan bool) {
	fmt.Println("我是go协程")
	time.Sleep(3 * time.Second)
	flag <- true
}

// 信道测试2  你发我收
func testChannelCal(s chan int) {
	for i := 0; i < 10; i++ {
		sum := i * i
		s <- sum
		if i == 6 {
			close(s)
			break
		}
		fmt.Println("cal i ", i)
		time.Sleep(10 * time.Millisecond)
	}
}
func testChannelRead(s chan int) {
	for {
		v, ok := <-s // ok表示有没有关闭通道
		if !ok {
			fmt.Println("通道关闭了.当前v", v) // 当ok = false ，读取信道为 0
			break
		}
		fmt.Println("v 输出 is", v)
	}
}
func testBufferChannel(chb chan int) {
	for {
		v, ok := <-chb
		if ok {
			fmt.Println("有buffer channel", v)
		}
	}
}
func testBufferChannel2(chb chan int) {
	for i := 0; i < 10; i++ {
		chb <- i
		fmt.Println("buffer in ", i)
	}
}

// WaitGroup 类似Java中的 CountDownLatch
func testWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go testWaitProcess(i, &wg)
	}
	wg.Wait()
	fmt.Println("testWaitGroup main done")
}
func testWaitProcess(i int, wg *sync.WaitGroup) {
	fmt.Println("wait group i is", i)
	time.Sleep(10 * time.Millisecond)
	wg.Done()
}
