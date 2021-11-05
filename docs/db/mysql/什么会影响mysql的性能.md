http://itboyhub.com/2021/07/28/mysql-performance/

1. SQL脚本，索引问题，TPS(事务)，QPS 比较高
2. 数据库连接，事务处理时间长，大表，通过 max_connections 控制
3. 慢sql 导致 CPU 高 -> 考虑多核 CPU,
4. 网卡流量，减少从库数量，避免使用select * 
5. 大表（10G,500~1000w）的update 操作
6. 大事务操作
7. 存储引擎，myisam,innodb
8. 硬盘 固态>机械