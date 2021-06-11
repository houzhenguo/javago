package main

import "fmt"

// go 中 database/sql 包定义了对数据库对一系列操作，但是没有提供任何的官方数据库驱动
// 需要使用第三方的驱动包
func main() {
	fmt.Println("hello db")
}
