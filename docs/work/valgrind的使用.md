## valgrind的安装

1. yum install valgrind

2. 输入valgrind–h显示valgrind的参数及提示，说明安装成功

注意对应版本的安装，否则检查会有问题。

## C Demo

```c
#include <stdlib.h>

void fun()
{
	int *p=(int*)malloc(10*sizeof(int));
	p[10]=0;
}

int main(int argc, char* argv[])
{
	fun();
	return 0;
}

```
编译：
```bash
gcc test.c -o test
```
```bash
valgrind --log-file=./report.log ./test
```

https://blog.csdn.net/erlang_hell/article/details/51360149

https://www.ibm.com/developerworks/cn/linux/l-cn-valgrind/


## Linux 分析

```bash
cat /proc/self/status
```