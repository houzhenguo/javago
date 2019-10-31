
```bash
# 创建 houzhenguo 并且在指定home目录
useradd -d /home/houzhenguo houzhenguo

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
```

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