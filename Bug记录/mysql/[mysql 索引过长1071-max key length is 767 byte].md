# [mysql 索引过长1071-max key length is 767 byte]

### 问题

create table: Specified key was too long; max key length is 767 bytes

 

### 原因

数据库表采用utf8编码，其中varchar(255)的column进行了唯一键索引

而mysql默认情况下单个列的索引不能超过767位(不同版本可能存在差异)

 

于是utf8字符编码下，255*3 byte 超过限制

 

### 解决

1  使用innodb引擎；

2  启用innodb_large_prefix选项，将约束项扩展至3072byte；

3  重新创建数据库；

 

my.cnf配置：

default-storage-engine=INNODB

innodb_large_prefix=on

 

 

一般情况下不建议使用这么长的索引，对性能有一定影响；

 

参考文档：

https://dev.mysql.com/doc/refman/5.5/en/innodb-restrictions.html