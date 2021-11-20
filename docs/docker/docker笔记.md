# 相关链接
https://www.bilibili.com/video/BV1og4y1q7M4?t=400&p=1  b站地址

## 阿里云镜像加速
https://docs.docker.com/engine/install/centos/  docker for centos
国内换成阿里云的镜像，然后继续install
```shell
 sudo yum install -y yum-utils
 sudo yum-config-manager \
    --add-repo \
    https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

sudo systemctl start docker

docker version

# 查看hello-world镜像
docker images # 查看当前镜像

REPOSITORY    TAG       IMAGE ID       CREATED        SIZE
hello-world   latest    d1165f221234   4 months ago   13.3kB
```
阿里云镜像加速

https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors

阿里云 -> 镜像加速器  -> 每个人的不同

```shell
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://0jxi0lne.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```

docker run的流程图
运行命令 -> 判断本机是否有镜像 -> 使用这个镜像
        -> 去dockerhub（阿里镜像） 下载
                -> 找不到这个镜像 -> 返回错误
                -> 下载镜像到本地 
                        -> 运行镜像

# docker
## 底层原理
1. Docker是一个 cs结构系统，Docker的守护进程运行的主机上，通过 Socket从客户端访问
2. Docker server 从 client 接收到指令，就会执行这个指令。

## Docker 相比 VM
1. docker 有着比 vm更少的抽象层
2. docker 用的宿主机的内核，vm需要 guestos（费劲）
新建一个容器的时候，docker不需要 像是vm家在一个操作系统内核，避免引导，docker利用 宿主机的操作系统，速度比较快。
3. 打包带上环境 -> images -> 下载镜像 -> 直接运行即可
4. docker 集装箱 互相隔离 端口冲突 项目是交叉的，隔离是docker的核心思想。
5. 可以把 Linux服务器压榨到极致
6. 2013 docker 开源，2014年4月 docker1.0发布 轻量级 

## DevOps 开发运维
1. 应用更快速的交付和部署，docker 打包镜像发布测试，一键运行
2. 更便捷的升级和扩所容，部署应用就和搭建积木一样。-> 可以整体升级
3. 更简单的系统运维，开发和测试环境是高度一致的。
4. 更高效的资源计算利用
5. docker 是内核级别的虚拟化，可以在一个物理机运行很多容器实例。

## docker 基本组成
1. ![docker架构图](https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fupload-images.jianshu.io%2Fupload_images%2F14854885-99420ba8b2d81151.png&refer=http%3A%2F%2Fupload-images.jianshu.io&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1639471327&t=bfddfd5c0e34a90f4040ca8e420f62ac)


client + server + registry

client -> images -> 容器

images:
docker 相当于一个模板，通过模板创建一个 容器服务，tomcat 镜像 -> run -> tomcat01容器（题容服务器）
通过这个镜像可以创建多个镜像
container:
通过容器技术，可以独立运行一个或者一组应用，通过镜像来创建
启动，停止，删除基本命令
可以理解问一个简易的linux系统
repository:
存放镜像的地方，
仓库分为公有仓库和私有仓库
配置镜像加速

## 常用命令

官方文档： https://docs.docker.com/reference/
仓库地址： https://hub.docker.com/ 


## docker 是怎样工作的
1. docker 是一个cs结构系统，docker的守护进程运行在主机上，
2. 通过守护进程 运行容器
3. 容器内的8080 (eg)   
![](./images/d1.png)

### 帮助命令
docker version  显示版本信息
docker info     更加详细的信息 系统信息 镜像和容器的数量
docker  --help 
帮助文档地址：https://docs.docker.com/reference/

### 镜像命令
1. docker images 查看本地本机上看
https://docs.docker.com/engine/reference/commandline/images/

2. docker search 搜索镜像 -> 

```shell
[root@aliyun soft]# docker search mysql
NAME                              DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                             MySQL is a widely used, open-source relation…   11168     [OK]
mariadb                           MariaDB Server is a high performing open sou…   4238      [OK]
```
https://docs.docker.com/engine/reference/commandline/search/

### docker pull

docker pull mysql[:tag] 下载镜像
docker pull mysql:5.7 下载指定版本

分层下载，只会下载差异部分

联合文件系统 -> 共用文件

Docker 删除镜像  docker rmi -f 镜像ID或者镜像名称

i = images 通过imageid 删除

删除容器 docker rm 

docker rmi -f d1165f221234 d1165f221235 空格分割

删除所有镜像
docker rmi -f $(docker images -aq)



### 容器命令

1. 下载一个 centos 进行学习
docker pull centos
查看是否下载完毕 docker images centos
2. 新建容器并启动
docker run [可选参数] imageId 参数说明

-- name=“Name”   tomcat01 02 用来区分容器
-d 后台方式运行 -> nohup
-it  使用交互方式进行，进入容器查看内容
-p  指定容器端口 -p 8080:8080
    -p 主机端口:容器端口（常用）
    -p 容器端口
    容器端口
-P 随机指定端口 

启动并进入容器
[root@aliyun ~]# docker run -it centos /bin/bash

exit 退出容器 -> 会停止

control + P + Q 容器不停止 退出 
配合 docker ps 使用 
```bash
# zhenguo.hou @ C026M in ~ [17:53:16] C:127
$ docker run -it centos /bin/bash
[root@8d4f22eb5ad4 /]# %

# zhenguo.hou @ C02D6M in ~ [17:53:49]
$ docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS          PORTS     NAMES
8d4f22eb5ad4   centos    "/bin/bash"   32 seconds ago   Up 32 seconds             fervent_wilbur
```

列出所有的运行的容器
docker ps -a   -a 是历史运行过的参数，
docker ps 当前在运行的
docker ps -n=2 显示最近创建的2个容器
docker ps -a -n2
docker ps -q 只显示编号

启动和停止容器的操作
1. docker start 容器ID    docker start d15822c4c59c
2. docker restart 容器ID
3. docker stop 容器ID     docker stop d15822c4c59c
4. docker kill  容器ID 

常用命令
1.  后台启动 docker run -d centos

```bash
# zhenguo.hou @ C02FRH64MD6M in ~ [18:00:03] C:127
$ docker run -d centos
360f4348bc83197bd5c012c0e5f10aa1a16a72b1b4bd64ebd2d360c955f12f58

# zhenguo.hou @ C02FRH64MD6M in ~ [18:00:12]
$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES 
```

docker ps 发现 centos 停止了，docker 容器使用后台运行，就必须要有
一个前台进程，容器发现没有应用，就会自动停止。 


查看日志
1. 
docker run -it centos ./bin/bash -c "while true;do echo houzhenguo;sleep 1;done"
2. docker logs -t -f --tail 10 817e328df432

top 

docker top 817e328df432

查看镜像元数据
docker inspect --help
docker inspect 817e328df432

## 进入容器
进入当前正在运行的容器
我们容器通常使用后台运行，需要进入容器，修改一些配置 
docker exec -it 容器ID 
 docker exec -t 817e328df432 /bin/bash
进入容器后开启一个新的终端，可以在里面操作

docker attach 容器ID

进入融资正在执行的终端，不会启动新的进程

docker attach 817e328df432

从容器copy 文件到主机
1. docker cp 容器ID:容器内路径 目的主机路径
docker cp e6765f4b6869:/home/test.java ./
容器可以没有在运行，需要进入dockern内部
未来可以使用卷技术，可以实现和主机的打通


docker exec -ti d5d039df430b redis-cli

## 练习
docker 安装ng
docker search nginx
docker run -d --name nginx01 -p:3344:80 nginx
curl localhost:3344
 外部端口：容器内部端口  映射端口暴露

 进入 docker exec -it 3bfe7527b094 /bin/bash

 ![](./images/d3.png)


tomcat 

docker run -d --name tomcat01 -p:8089:8080 tomcat

测试访问没有问题

进入容器 docker exec -it tomcat01 /bin/bash

发现阉割版 tomcat,保证最小可运行环境

cp -r webapps.dist/* ./webapps 就可以访问了

## 部署es + kibana
1. es 暴露端口多
2. es十分耗内存
3. es 的数据一般需要放置到安全目录，挂载

步骤 
1. docker pull elasticsearch:7.6.2
2. docker run -d  --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2
3. es十分耗费内存，

docker stats 查看 docker 状态

docker status 容器ID



curl localhost:9200


## portainer

1. docker 图形化界面管理工具，提供一个后台面板 提供给我们操作

docker run -d -p 8088:9000 --restart=always -v /var/run/docker.sock:/var/run/docker.sock --privileged=true portainer/portainer

http://localhost:8088/#/init/admin


### 镜像是什么
镜像是一种轻量级，可执行的独立软件包，用来打包软件运行环境和基于运行环境开发的软件，包含运行某个软件所有的内容
包括 代码 运行时，库，环境变量和配置文件。

如何得到镜像
1. from origin rep
2. copy
3. 制作镜像 dockerFile

## docker 镜像加载原理
1. unionFS 联合文件系统
下载的时候看到的一层一层的就是 unionFS.分层，轻量级且高性能的文件系统，支持对文件系统的修改作为一次提交来一层层的叠加，
同时可以将不同目录挂载到同一饿虚拟文件系统下。
UFS是docker 镜像的基础。
bootfs 系统启动引导加载，主要包含 bootloader 和kernel.bootloader 主要是引导加载kernel.linux 刚启动时候，会加载 bootfs
文件系统，在docker镜像最底层就是 bootfs。当boot 加载完成之后就在整个内核的内存中了，此时内存多使用权已由bootfs 交给内核，此时系统
也会卸载bootfs
rootfs，在bootfs 之上，包含的就是典型的Linux系统中的 /dev/ /proc/ /bin /etc 等标准目录和文件，rootfs 就是各种不同的操作系统发行版。
比如 ubuntu,centos

对于一个精简OS,rootfs很小，只包含最基本的命令，工具和程序库就可以了，底层还是使用的是主机的内核，自己只需要 提供rootfs就可以了。

## 理解
1. 所有的docker 镜像都起始于一个基础镜像层，当进行修改或增加新的内容时，就会在当前的镜像层之上，创建新的镜像层。
eg. 加入基于 ubuntu linux 创建一个新的镜像，这是新镜像的第一层，如果在该镜像中添加py包，就会添加第二层，如果继续添加一个安全补丁，
就会创建第三个镜像层。

![](./images/d2.png)
![](./images/d4.png)


## commit 镜像
docker commit 提交容器成为一个新的副本。
docker commit -m="提交的描述信息" -a="作者" 容器id 镜像名字 

### 实战
docker run -it -p 8089:8080 tomcat    # 启动一个默认的tomcat
docker exec -it ccc6e0582540 /bin/bash # 进入容器
cp -r webapps.dist/* ./webapps # copy到webapps才能访问
docker commit -a="zhenguo" -m="add webapps" ccc6e0582540 tomcat02:1.0  # 提交

![](./images/d5.png)
之后就可以直接使用修改过的镜像。

tomcat + 我们的一些操作 =》打包成一个新的镜像
![](./images/d6.png)