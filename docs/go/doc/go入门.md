## 方向
1. golang = golanguage
2.  https://golang.org/ 官方网站
3.  https://studygolang.com/pkgdoc  中文在线标准库文档 可以找一个离线版本
4. 优秀网站 https://www.liwenzhou.com/posts/Go/Gin_framework/

## download
1. https://golang.org/doc/install?download=go1.16.4.darwin-amd64.tar.gz
2. https://blog.csdn.net/weixin_43931792/article/details/98070995
3. sudo tar -C /usr/local -xvzf ~/Downloads/go1.16.4.darwin-amd64.tar.gz 用这个命令不会报错
4. 查看golang 当前版本 go version

5. 终端命令出问题的时候 export PATH=/usr/local/opt/coreutils/libexec/gnubin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/Applications/Wireshark.app/Contents/MacOS

### 环境变量配置
export GO_HOME=/usr/local/go
export PATH=/usr/local/go/bin

## 编译启动
```bash
go run hello.go  # 直接运行  编译 -> 运行
go build hello.go  -> ./hello # 编译之后运行  可执行文件 -> 可以cp到 没有go环境运行
```
编译之后文件会变大很多 -> 所依赖的库文件在其中。 可以指定编译的二进制文件的名字

shift + option + 上下方向 是 复制
```go
package main
import "fmt"
func main()  {
	fmt.Println("hello world")
}
```

## 开发注意事项
1. 源文件 以 .go 结尾 类似Java
2. 执行入口为main 类似Java
3. 严格区分大小写 类似Java
4. 语句后面不需要加分号。   go自动加上
5. go是以行进行编译。主要还是分号,加上分号就可以写多个
6. `go 定义的变量或者引用的包没有用到就会报错` 需要注意

## 转义字符
1. \t
2. \n 换行
3. \\ \
4. \"
5. \r 回车

## 格式化
1. shift + table 向右
2. gofmt -w hello.go 

## 变量
变量 = 变量名 + 值 + 数据类型； := 只有在 声明的时候用到
```go
// 变量定义
// 变量的定义和盛名
var i int // 指定数据类型 使用默认值
// i = 10
fmt.Println("test 变量 默认值", i)
var j = 10.22 // 自行推导
fmt.Println("自行推导", j)
x,y,z := "zhangsan", 111, 33.3 // 一次性声明三个变量
fmt.Println("x = ",x ,"y=",y,"z=",z)
// 字符串拼接
var str = "hello" + "world"
fmt.Println(str)

```

## 数据类型
1. int int8 int16 int32 int64 unit unit8 unit16 unit32 unit64 byte 无符号 + 有符号
2. float32 float64
3. byte 保存字符串 当字符串很长的时候 。+ 保留在上一行 
4. bool 

```go
	var ui uint32
	ui = -2222 // 校验不通过
	fmt.Println("ui", ui)
```
占用字节数
```golang
import (
"fmt" 
"unsafe"
)
var ui uint32
ui = 2222
fmt.Println("ui", ui, "占用字节数：", unsafe.Sizeof(ui))
```
golang 浮点型默认 float64

```golang
	// 显式转换
	var i1 int32 = 100
	fmt.Println(" i1 占用字节数：", unsafe.Sizeof(i1)) // 4
	var n1 float64 = float64(i1)
	fmt.Println(" n1 占用字节数：", unsafe.Sizeof(n1)) // 8
```
