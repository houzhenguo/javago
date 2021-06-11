package main

import "fmt"

// & 是取出地址
// *是指针 ，可以用来取值
func main() {
	a := 10
	b := &a                              // 取出a变量的地址存储在b中
	fmt.Printf("type of b is : %T\n", b) // *int
	c := &b
	fmt.Printf("c type is %T\n", c)
	fmt.Println("b value is", b)
	fmt.Println("c value is", c)

	d := 40
	modify1(d)
	fmt.Println("d is ", d)
	modify2(&d)
	fmt.Println("modify2", d)

	// make 主要用于 slice, map, chan
	// func make (t Type, size ...IntegerType) Type

	// map
	var m = make(map[string]int)
	m["liz"] = 1
	m["liz2"] = 2
	for k, v := range m {
		fmt.Println("k is", k, "v", v)
	}
}

func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 500
}
