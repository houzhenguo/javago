
# netcat

## 介绍

基于 tcp/ip 协议，小而精。

下载地址：
     
     windows: https://eternallybored.org/misc/netcat
     Linux： yum install -y nc

## 使用
1. 基础使用

    // 相当于 客户端与服务器互换角色
    
    创建一个服务器端方法
    - nc -l -p [localport]
    [houzhenguo@aliyun book]$ nc -l -p 5555 // 开启监听

    创建一个 客户端连接（连接服务器端）
    - nc [remote_addr][remoteport]
    [houzhenguo@aliyun soft]$ nc 127.0.0.1 5555
    在这边输入 hello 服务器端的nc就能收到，可以相互交互
    可以双方进行 信息交互
2. 返回shell 引用

    创建一个服务器端的方法
    - nc -l -p [localport] -e cmd.exe
    创建一个客户端（连接服务器端）
    - nc [remote_addr] [remoteport]

    可以操作 服务器上的某些命令

    [houzhenguo@aliyun test]$ nc -l -p 5555 -e ./test.sh// 服务器
    [houzhenguo@aliyun soft]$ nc 127.0.0.1 5555
    我是 ./test.sh的返回值
3. 文件传输

    nc中的数据传输，使用的是标准的输入输出流，所以可以直接使用命令进行操作。
    服务器端：
        - nc -l -p [localport] > outfile
        [houzhenguo@aliyun test]$ nc -l -p 5555 >out.txt  // 客户端与服务器端可以互相传输数据。


    客户端：
        - nc [remote_addr][remoteport] < infile
        [houzhenguo@aliyun soft]$ nc 127.0.0.1 5555 <test1.txt 

4. 信息探测

    内网扫描，当获得目标权限之后，如果目标没有任何途径课可以对内网进行探测，此时可以用
    netcat 进行内网ip和端口的扫描。

    命令行：
        - nc -v -n -z -w1 [target_id][start_target_port-stop_target_port]
        - v 表示对错误进行详细输出
        - n 不对目标机器进行 DNS解析，直接使用ip
        - z zero I/O模式，专用与端口扫描。表示对目标IP 发送的数据表中不包含任何payload,这样可以加快扫描速度
        - w1 超时设置为 1s

        .\nc.exe -v -z -w1 -n ip地址 21-22 可以返回 开放的端口

        返回 目标主机 哪些端口开放，并且运行了哪些 服务
        

        

    