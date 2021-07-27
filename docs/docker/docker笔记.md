https://www.bilibili.com/video/BV1og4y1q7M4?t=400&p=1  b站地址

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


## 底层原理
1. Docker是一个 cs结构系统，Docker的守护进程运行的主机上，通过 Socket从客户端访问
2. Docker server 从 client 接收到指令，就会执行这个指令。

## Docker 相比 VM
1. docker 有着比 vm更少的抽象层
2. docker 用的宿主机的内核，vm需要 guestos（费劲）
新建一个容器的时候，docker不需要 像是vm家在一个操作系统内核，避免引导，docker利用 宿主机的操作系统，速度比较快。

## 常用命令

官方文档： https://docs.docker.com/reference/

### 帮助命令
docker version  显示版本信息
docker info     更加详细的信息 系统信息 镜像和容器的数量
docker  --help 

### 镜像命令
docker images 查看本地本机上看
docker search 搜索镜像 -> 

```shell
[root@aliyun soft]# docker search mysql
NAME                              DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                             MySQL is a widely used, open-source relation…   11168     [OK]
mariadb                           MariaDB Server is a high performing open sou…   4238      [OK]
```

docker pull mysql[:tag] 下载镜像
docker pull mysql:5.7 下载指定版本

联合文件系统 -> 共用文件

Docker 删除镜像  docker rmi -f 镜像ID或者镜像名称

删除容器 docker rm 

docker rmi -f d1165f221234
### 容器命令

1. 下载一个 centos 进行学习
docker pull centos
2. 新建容器并启动

docker run [] imageId 参数说明

-- name=“Name” 容器名字 tomcat01 02 用来区分容器
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

列出所有的运行的容器
docker ps -a;;; -a 是历史运行过的参数，
docker ps 当前在运行的
docker ps -n=2 显示最近创建的2个容器
docker ps -q 只显示编号

启动和停止容器的操作
1. docker start 容器ID    docker start d15822c4c59c
2. docker restart 容器ID
3. docker stop 容器ID     docker stop d15822c4c59c
4. docker kill  容器ID 

常用命令
1.  后台启动 docker run -d centos