
```bash
# 创建 houzhenguo 并且在指定home目录
useradd -d /home/houzhenguo houzhenguo
[root@localhost /]# passwd haha   #为该用户设置密码

# 查看当前用户的分组
groups

## bash 终端
ssh-keygen 在 c: user: admin.. : .ssh id_rsa.pub 上传到服务器

# 在home/user下面创建 .ssh
mkdir /home/houzhenguo/.ssh
touch /home/cola/.ssh/authorized_keys
# 修改hostname
hostname
hostnamectl set-hostname aliyun

# 安装rz sz
yum -y install lrzsz

# 上传不乱码
re -be 

# 解压
unzip arthas-packaging-3.1.4-bin.zip -d arthas-packag
# 安装jdk https://cloud.tencent.com/developer/article/1447083   获取安装目录，你发现在/usr/lib/jvm目录下可以找到他们。
yum install -y java-1.8.0-openjdk-devel.x86_64

# 安装 open jdk11 https://blog.csdn.net/ringliwei/article/details/85260801

# 搜索历史命令
ctrl + r 输入关键字 然后回车


```

tomcat 
https://www.cnblogs.com/yw-ah/p/9770971.html

/usr/tomcat/apache-tomcat-9.0.27/webapps/docs/appdev

## linux记录登录ip方法
PS：Linux用户操作记录一般通过命令history来查看历史记录，但是如果因为某人误操作了删除了重要的数据，这种情况下history命令就不会有什么作用了。以下方法可以实现通过记录登陆IP地址和所有用户登录所操作的日志记录！

在/etc/profile配置文件的末尾加入以下脚本代码就可以实现，通过上面的代码可以看出来，在系统的/tmp新建个history目录（这个目录可以自定义），在目录中记录了所有的登陆过系统的用户和IP地址，这也是监测系统安全的方法之一。

```bash
PS1="`whoami`@`hostname`:"'[$PWD]'
history
USER_IP=`who -u am i 2>/dev/null| awk '{print $NF}'|sed -e 's/[()]//g'`
if [ "$USER_IP" = "" ]
then
USER_IP=`hostname`
fi
if [ ! -d /tmp/history ]
then
mkdir /tmp/history
chmod 777 /tmp/history
fi
if [ ! -d /tmp/history/${LOGNAME} ]
then
mkdir /tmp/history/${LOGNAME}
chmod 300 /tmp/history/${LOGNAME}
fi
export HISTSIZE=4096
DT=`date +"%Y%m%d_%H%M%S"`
export HISTFILE="/tmp/history/${LOGNAME}/${USER_IP} history.$DT"
chmod 600 /tmp/history/${LOGNAME}/*history* 2>/dev/null
```