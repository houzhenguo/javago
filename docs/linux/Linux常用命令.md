
1. du : 本文件夹的磁盘使用情况
2.  free : 显示内存和交换空间的使用
3. ps -ef | grep java
4. top / top -Hp
5. netstat -nalp t 查看网络端口号的情况。（htop命令）
```
 - n: 列出的带ip地址而不是域名
 - t: tcp/ u: udp
 - p: process 进程的名称
```
6. lsof -i :6379 监控6979端口 / -iTCP 监控TCP端口

7. less 命令，可以前后翻看 pageDown pageUp,
```
按 / 可以进行查询。 
可以按大写 F，就会有类似 tail -f 的效果，读取写入文件的最新内容， 按 ctrl+C 停止
n: next
N: pre
```
8. tail -n -f / head
9. Ctrl + R 搜索历史命令
10. nohup
```
nohup command > myout.file 2>&1 &   
```

## 参考 
[UNIX / LINUX CHEAT SHEET](http://cheatsheetworld.com/programming/unix-linux-cheat-sheet/)