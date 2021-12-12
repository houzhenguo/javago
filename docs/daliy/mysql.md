## 读写分离
1. 读多写少
2. 准实时，有延迟，避免立刻读取（业务规避）
3. MyCat -> 独立部署的，避免中心节点-> 相当于代理
4. sharding-jdbc 轻量无侵入

## 主从同步
1. slave -> 创建一个IO线程 向master 请求更新binlog
2. master 创建一个binlog dump 线程来发送binlog 
3. slave 接收binlog写到 relay log (中继日志)
4. slave 读取relay log同步到本地，sql 线程在本地执行里面内容
binlog 可以做主从同步和数据恢复。

## 高性能高可用
1. mycat 集群，去中心化（心跳，zk 始终维护一个leader）
2. 做读写分离
3. 做水平拆分和垂直拆分
4. sql 语句优化 -> 慢查询日志 
高并发 -> 限流 -> mq -> 缓存 -> mysql