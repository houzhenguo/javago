package main

import (
	"fmt"
	"unsafe"
)

type Myint int
type Alisint = int

func main() {
	var a Myint = 2
	fmt.Printf("a type is %T", a) // main.Myint% 这个是新类型
	var b Alisint = 3
	fmt.Printf("b type is %T", b) // int 这个是别名

	var p1 person = person{
		"zhangsan",
		12,
	}
	fmt.Println(p1)
	fmt.Println("p1.name is", p1.name)
	// 匿名结构体
	var c struct {
		Name string
		Age  int
	}
	c.Name = "wangwu"
	c.Age = 32
	fmt.Println("c is ", c)
	fmt.Println("p name add", &p1.name)
	fmt.Println("p age add", &p1.age) // 地址是连续的

	// 空结构体是不占用空间的
	var v struct{}
	fmt.Println("v size is", unsafe.Sizeof(v))
	// 测试方法
	result := p1.sayHello(p1.name)
	fmt.Print("p1 sayHello result ", result)
	p1.changeAge(222) // 指针
	fmt.Println("p1 after change age is", p1)

	dog := Dog{
		"ss",
		12,
	}
	fmt.Println("dog 匿名字段", dog)

	// 匿名结构体 + 嵌套结构体
	h := House{
		"王武",
		Dog{"ss", 4},
		Cat{ // 非匿名结构体
			name: "ca1",
			age:  2,
		},
	}
	fmt.Print("house is", h)
	fmt.Print("house cat name is ", h.cat.name)
}

type person struct {
	name string
	age  int
}

// 方法 与 函数区别，方法作用于特殊的接收者
// func (接收者) 方法名(参数列表)(返回参数)
func (p person) sayHello(name string) string {
	fmt.Println("p 地址", &p)
	return "hello " + name
}

// 当需要修改接收者中的值
// 接收者cp 代价比较高
//
// 指针类型
func (p *person) changeAge(age int) {
	fmt.Println("p指针入参 地址", &p)
	p.age = age
}
func (p person) changeAge1(age int) {
	p.age = age
}

// 结构体匿名字段
type Dog struct {
	string
	int
}
type Cat struct {
	name string
	age  int
}
type House struct {
	name string
	Dog
	cat Cat
}
