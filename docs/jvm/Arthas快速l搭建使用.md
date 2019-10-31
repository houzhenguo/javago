
# Arthas 内网环境安装

[官方安装教程](https://alibaba.github.io/arthas/install-detail.html) : 全量安装

## 全量安装

1. 下载 arthas 包 [arthas-packaging-3.1.4-bin.zip](http://repository.sonatype.org/service/local/artifact/maven/redirect?r=central-proxy&g=com.taobao.arthas&a=arthas-packaging&e=zip&c=bin&v=LATEST)

2. 上传下载的jar包到Linux服务器，使用 unzip xxx.zip -d arthas-boot/ 将文件解压到这个文件夹，可以使用 `ls -alh | wc -l` 统计数量，一共13个（可能与版本有关）  这个时候启动 `java -jar arthas-boot.jar` 就可以正常使用了。

3. 编写arthas-start.sh

```bash
[root@localhost arthas-boot]# cat arthas-start.sh 
#!/bin/bash
java -jar $ARTHASPATH/arthas-boot.jar

chomod +777 xxx.sh
```

4. 

别名
```bash
vim /etc/profile #新加变量 文件位置到环境变量中

## 内容如下

source profile 

vim /etc/bashrc

# 内容如下
alias as='arthas-start.sh'
```

5. 通过别名 as即可使用 archas



Arthas目前支持Web Console，用户在attach成功之后，可以直接访问：http://127.0.0.1:8563/

# 远程连接部分

使用arthas tunnel server连接远程arthas

## 下载部署arthas tunnel server

[下载地址](https://github.com/alibaba/arthas/releases)

Q:

user 看不到root的进程？
环境变量的配置方案


firewall-cmd --add-port=8563/tcp


yum install telnet


https://github.com/alibaba/arthas/issues/868


https://alibaba.github.io/arthas/web-console.html#arthas-tunnel-server

http://192.168.85.131:8080/