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

## var
```golang
s, n := "abc", 123
// 占位
i := 0
_ = i
// 批量定义const
const (
	a =1
	b ="ss"
)
const (
	a1   byte = 127       // int to byte ,128会越界
	//b1   int  = 1e20      // float64 to int, overflows
)
// iota的使用，自增+1
const (
	// 0-6
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```

## 类型转换
1. 不支持隐式转换,需要做显式转换
```golang
var b byte = 100
// var n int = b // Error: cannot use b (type byte) as type int in assignment \
var n int = int(b) // 显式转换
```
## string
1. 默认空字符串
2. 可以使用索引访问某个字节
3. 不可变
```golang
// 使用index
s := "abc"
println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
true,true,true
// 转义，原封不动输出
	s1 := `a
b\r\n\x00
c`
	println(s1)
```
+ 拼接必须在上一行末尾，否则编译不通过

```golang
s := "Hello, " +
     "World!"
s2 := "Hello, "
    + "World!"    // Error: invalid operation: + untyped string
```

要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。⽆无论哪种转 换，都会重新分配内存，并复制字节数组。

## 指针
1. ⽀支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。
- 默认值nil
- 操作符 & 取变量地址， * 通过指针访问目标对象，注意空指针

```golang
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d // 取地址

	fmt.Printf("%p, %v\n", p, p.a) // 直接⽤用指针访问目标对象成员，无须转换。
	fmt.Println(*p) // 访问对象
```

## 自定义类型
具有相同声明的未命名类型被视为同⼀一类型。
• 具有相同基类型的指针。
• 具有相同元素类型和⻓长度的 array。
• 具有相同元素类型的 slice。
• 具有相同键值类型的 map。
• 具有相同元素类型和传送⽅方向的 channel。
• 具有相同字段序列 (字段名、类型、标签、顺序) 的匿名 struct。 • 签名相同 (参数和返回值，不包括参数名称) 的 function。
• ⽅方法集相同 (⽅方法名、⽅方法签名相同，和次序⽆无关) 的 interface。

```golang
	var a3 struct { x int `a` }
	var b3 struct { x int `a` }
	fmt.Println(a3 == b3) // 这样是相等的，如果把其中一个a 改成ab 就编译不通过
```