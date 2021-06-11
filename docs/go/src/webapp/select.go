package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)
	select {
	case s1 := <-ch1:
		fmt.Println("s1 is", s1)

	case s2 := <-ch2:
		fmt.Println("s2 is", s2)
	}
}

// selct 在没有数据的时候一直阻塞，当又一个ready的时候则返回，
// 以下代码是 直接返回server1
// 测试 select
func server1(ch1 chan string) {
	time.Sleep(3 * time.Second)
	ch1 <- "hello server1"
}
func server2(ch2 chan string) {
	time.Sleep(5 * time.Second)
	ch2 <- "hello server2"
}
