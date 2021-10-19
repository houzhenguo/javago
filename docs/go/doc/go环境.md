1. export ENV=test
2. go env -w GOPROXY=https://goproxy.cn,direct  设置国内镜像
3. go mod 命令
```
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

go get -u git.xxxlib@test

go mod支持语义化版本号，比如go get foo@v1.2.3，也可以跟git的分支或tag，比如go get foo@master，当然也可以跟git提交哈希，比如go get foo@e3702bed2。

在项目中执行go get命令可以下载依赖包，并且还可以指定下载的版本。

运行go get -u将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
运行go get -u=patch将会升级到最新的修订版本
运行go get package@version将会升级到指定的版本号version