
视频地址： https://www.bilibili.com/video/av32081351

官网地址： http://hadoop.apache.org/

## 简介

大数据主要解决，海量数据的`存储` 和 海量数据的`分析计算`问题，是一个分布式基础架构，后期学习的框架都是依赖 hadoop。它通常指的是一个 广泛的生态圈，例如 Hbase,Hive,Zk,...


数据存储单位： bit Byte kb mb gb tb pb eb 1024

当前，个人计算机硬盘的容量是 TB级别。

## 特点
- 大数据量
- 高速
- 多样，数据结构化 非结构化数据

## 论文

- GFS -> HDFS hadoop distribute file system
- Map-Reduce ->MR
- BigTable -> HBase

## 三大发行版本
- Apache 最基础版本，适合入门
- Cloudera 在互联网企业中使用
- Hortonworks 文档最好

## 优势
- 高可靠 底层维护多个副本
- 高扩展 可以方便的扩展数以千计的节点，动态
- 高效性 在 MapReduce 思想下，并行工作，加速任务处理速度
- 高容错性 自动将失败的任务重新分配

## 组成
- common(辅助工具)
- HDFS(数据存储)
- Yarn(资源调度)
- MapReduce(计算))

## HDFS 架构概述

1. NameNode(nn) ； 存储 文件的元数据，如文件名，文件的目录结构，文件属性（生成时间，副本数，文件权限），以及每个文件的块列表和块所在的DataNode等。（相当于目录）

2. DataNode(dn): 在本地文件系统存储文件块数据，以及块数据的校验和。

3. Secondary NameNode(2nn):用来监控 HDFS状态的辅助后台程序，每隔一段时间获取 HDFS元数据的快照。

