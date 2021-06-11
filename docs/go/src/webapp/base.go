package main

import "fmt"

// 基础练习
func main() {
	// testArr()
	// testSlice()
	// testMap()
	// testFuc()
	var num int = testFuncParam(2, 4, add)
	fmt.Println("num is", num)
	var num1 int = testFuncParam(2, 4, sub) // 函数式编程
	fmt.Println("num1 is", num1)
	// 测试函数当参数传入
	var num3 int = testFuncParam1(1, 2, add)
	fmt.Println("num3 is", num3)

	// 匿名函数
	add := func(x, y int) {
		fmt.Println("x + y ", x+y)
	}
	add(3, 5)

	// 闭包
	ss := add2(2)
	fmt.Println(ss(3))

	fmt.Println("f1", f1())
	fmt.Println("f2", f2())
	fmt.Println("f3", f3())
	fmt.Println("f4", f4())

}

func testArr() {
	// 数组练习
	// 1. 定义定长数组
	var a = [...]int{1, 2, 3}
	fmt.Println("a is ", a)
	var b = [...]string{"hel", "bei", "jing"}
	fmt.Println(b)
	c := [...]int{4, 5, 6}
	fmt.Println("c is ", c)
	// 数组遍历
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	for i, v := range c {
		fmt.Println("i is", i, "v is", v)
	}
	var d = [...]int{1, 3, 5, 7, 8}
	var total int = 0
	for _, v := range d {
		total += v
	}
	fmt.Println("total is", total)
}

// test 切片
func testSlice() {
	// 因为数组有长度，所以有很多局限性
	var b = [...]int{1, 2, 3, 4, 5, 6}
	var a = []bool{false, true} // 没有长度的就是切片
	c := b[1:4]
	fmt.Println(a == nil)
	fmt.Println("c cap is", cap(c)) // 容量 从index 开始，c 是从1 开始,到最后结束
	fmt.Println("c len is", len(c)) // len 就是当前元素到数量 ，index = 1,2,3 共三个值

	// 其他创建切片到方式
	d := make([]int, 2, 5) // make([]T, size, cap)
	//fmt.Printf("d len is %d, cap is %d", len(d), cap(d))
	d = append(d, 3)
	d = append(d, 3)
	d = append(d, 3)
	d = append(d, 3)
	d = append(d, 4, 5, 6, 6, 6)
	fmt.Println(d)
	fmt.Print("d len", len(d), "cap is", cap(d)) // 变成2 倍 ,源码 ： 当 len < 1024 扩容2倍，当 >1024 扩 1.25倍(for 循环)
	// 判断切片为空 len(d) == 0 不能使用nil

	// 从切片中删除元素
	f := []int{1, 2, 3, 4, 5, 6}
	f = append(f[:2], f[3:]...) // 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	fmt.Println("f is ", f)
}

func testMap() {
	// map[KeyType]ValueType
	scoreMap := make(map[string]int32)
	scoreMap["zs"] = 60
	scoreMap["ls"] = 80
	scoreMap["ww"] = 90
	fmt.Println(scoreMap)
	// 判断是否存在和取之
	v, ok := scoreMap["zs"]
	fmt.Println("ok", ok, "v", v) // ok判断是否存在
	// map的遍历
	for k, v := range scoreMap {
		fmt.Println("k is", k, "v is", v)
	}
	// 删除
	delete(scoreMap, "ww") // map的删除
	for k, v := range scoreMap {
		fmt.Println("k1 is", k, "v1 is", v)
	}
}

// 函数 把函数当作变量
type calculation func(int, int) int

func add(x int, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func testFuc() {
	var a calculation = add
	var s calculation = sub

	var sb1 int = a(3, 5)
	var sb2 int = s(5, 3)
	fmt.Println("sum is", sb1, "sub is", sb2)
}

// 把函数当作参数
func testFuncParam(x, y int, cal calculation) int {
	var v int = cal(x, y)
	return v
}
func testFuncParam1(x, y int, f func(x, y int) int) int {
	return f(x, y)
}

// 闭包
func add2(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
