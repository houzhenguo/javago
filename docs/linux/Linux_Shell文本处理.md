

[原文链接](https://mp.weixin.qq.com/s/CtNeak1jbSW1UmyKwRm7Fg)

本文档只记录了常用的，其他的没有使用过的没有记录，可以参考原文。

# Shell处理文本的工具

<!-- TOC -->

- [Shell处理文本的工具](#shell处理文本的工具)
    - [find](#find)
    - [grep](#grep)
    - [wc 统计行与字符的工具](#wc-统计行与字符的工具)
    - [awk 数据流处理工具](#awk-数据流处理工具)

<!-- /TOC -->


## find

1. 查找 txt 和 pdf文件

```bash
find . ( -name "*.txt" -o -name "*.pdf") - print
```
2. 正则表达式查找 txt pdf
```bash
find . -regex ".*(.txt|.pdf)$" # -iregex 忽略大小写
```
3. 否定参数

查找所有非txt文本
```bash
find . ! -name "*.txt" -print
```
4. 按类型搜索：

```bash
find . -type d -print # 只列出目录
# -type f 文件 / l 符号链接
```
按照时间搜索
```bash
-atime 访问时间
-mtime 修改时间
-ctime 创建时间
```
最近7天被访问过的所有文件
```bash
find . -atime 7 -type f -print
```
## grep

```bash
grep match_patten file # 默认访问匹配行
```
常用参数：

-v 输出没有匹配的文本行

-c 统计文件中包含文本的次数
```bash
grep -c "text" filename
```
-n 打印匹配的行号

-i 搜索时候忽略大小写

## wc 统计行与字符的工具

```bash
wc -l file # 统计行数
wc -w file # 统计单词数
wc -c file # 统计字符数
```

## awk 数据流处理工具

https://mp.weixin.qq.com/s/CtNeak1jbSW1UmyKwRm7Fg

以下shell来自【鸟哥的Linux私房菜】 P364

awk 是一个非常棒的数据处理工具，将一行分成数个字段来处理，因此，awk相当适合处理小型的数据。

模式如下

```bash
awk '条件类型1 {动作1} 条件类型2 {动作2} ...' filename

#以"|"分割数据，输出1，2，3，4行数据

awk -F"|" '{print $1,$2,$3,$4}' test.txt
```
