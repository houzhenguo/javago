
参考： http://bingohuang.nos-eastchina1.126.net/effective-go-zh-en-gitbook.pdf

https://github.com/uber-go/guide/blob/master/style.md#avoid-init

https://github.com/golang/go/wiki/CodeReviewComments

uber go https://studygolang.com/articles/23941
## 注释

1. 包级别注释,每个包都应该有一个包级别的注释，包含多个文件的，出现在其中一个文件即可。

## 命名
1. 包名 
    a. 小写
    b. 举例 bufio.Reader 而不是 buf.BufReader ,前面那个已经有那个含义了，不需要重复定义。
       bufio.Reader  和io.Reader
2. Getter
    a. 不需要getter ，通过大写控制。应该是 Owner 而非 GetOwner 多余
    b. 可以提供 set ,eg. SetOwner

3. 接口
 只包含一个方法的接口应该以改方法的名称上加上 -er 后缀命名，eg. Reader, Writer, Formater, CloseNotifier

4. 驼峰
5. 分号，go 只有 for 这种才会有分号，它是通过某些关键字来确定分号，要求我们在coding的时候一定要注意规范。
    ` break continue fallthrough return ++ -- ) }`
6. for 
for init,condition,post {}
for condition = while
for {} = for(;;)
for k,v :=range map {
    如果只需要k,那么v可以省略，如果只需要v,k可以用_代替
}
7. switch
```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
}
return 0 }
```

可以用来判断断言
```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T", t)
case bool:
    fmt.Printf("boolean %t\n", t)
case int:
    fmt.Printf("integer %d\n", t)
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t)
case *int:
    fmt.Printf("pointer to integer %d\n", *t)
// %T 输出 t 是什么类型 // t 是 bool 类型
// t 是 int 类型
// t 是 *bool 类型
// t 是 *int 类型
}
```

## function 
go 的返回值可以当作 形参 被命名，他们会在function开始执行的时候，初始化为零值，按照自己风格来吧。
```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
}
return
}
```

defer 放在 开启的地方，类似finally


## new
https://studygolang.com/articles/3363

来初始化一个对象，并且返回该对象的首地址．其自身是一个指针．可用于初始化任何类型

```go
// 一般采用以下形式
 return &File{fd: fd, name: name}
```

## make
返回一个初始化的实例，返回的是一个实例，而不是指针，其只能用来初始化：slice,map和channel三种类型
```go
make([]int, 10,100)
```
make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用 new 分配内存。

## slice

切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。 
若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 
这 可以理解为传递了底层数组的指针。因此，Read 函数可接受一个切片实参 而非一个指针和 一个计数;切片的长度决定了可读取数据的上限。


## context 
作为链路追踪和整个api/rpc 方法链，尽量把它放在第一个参数位置


## print 
```go
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
```


## 定义
1. 常量的定义 是数字，字符，字符串，bool,编译期创建。
定义的表达式必须是编译器求值的常量表达式。例如 1<<3 就是一个常量表达式，而 math.Sin(math.Pi/4) 则不是，因为对 math.Sin 的函数调用在运行时才会发生。

2. 变量的定义可以是加载的过程中才进行赋值
```go
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```