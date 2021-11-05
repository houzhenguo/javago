1. goroutine 成本低
2. channel 进行数据通信，避免内存共享问题。
3. 多组合，无继承，接口简单，不需要声明，直接实现。

极客时间地址 https://time.geekbang.org/column/article/378076?utm_term=pc_interstitial_1267

## 环境变量
- goroot 编译工具，标准库的安装路径
- gopath go的工作目录，编译后二进制文件的存放地方
- go111module 。on 开启，会让编译器忽略$gopath 和vendor文件夹，只根据go.mod下载 。
    off 是关闭，在gopath 和vendor目录查找以来关系。
- goproxy 设置国内代理
- goprivate 制定不走代理的go包域名，go get 可以不经过代理直接下载
- 