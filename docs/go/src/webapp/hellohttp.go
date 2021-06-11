package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数 参数是固定的
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("requst = {}", r)
	fmt.Fprintln(w, "hello world", r.URL.Path)
}
func main() {
	//myHandler := MyHandler{}
	fmt.Println("hello")
	http.HandleFunc("/hello", handler) // 这个可以自动给你转
	//http.Handle("/myhandle", &myHandler)
	// 创建路由
	http.ListenAndServe(":8080", nil)
}

// 自定义 handler

type MyHandler struct {
}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的请求处理") // 实现接口
}
