# 工作日常踩坑
## left join 右表筛选条件后移导致的筛选错误
表 A:客户信息主表

|id|deleted_at|
|----|----|

表 B:客户补充信息表，一个客户可能没有补充信息。

|id|name|a_id|deleted_at|
|----|----|----|----|

目的：查询出所有客户信息，若该客户有补充信息需要带有补充信息

错误 sql：
```mariadb
select a.id,b.name from a left join b on a.id = b.a_id where b.deleted_at = '1970-01-01 00:00:00'
```
分析：

1、由于表 b 可删除，表 b 中可能包含多条属于同一条 a 表某一列的数据，因此需要筛选 deleted_at 字段。

2、由于一个客户可能没有补充信息，因此 left join 后可能有右表为空的情况，而在 where 中进行筛选会在联表后再进行筛选，导致丢失没有补充信息的客户信息。

修复 sql：
```mariadb
select a.id,b.name from a left join b on a.id = b.a_id and b.deleted_at = '1970-01-01 00:00:00'
```

在联表时就对删除信息进行筛选。


## 分页排序列包含相同数据导致数据返回错误
### 问题场景
一条分页查询 sql 语句中，有同一条数据出现在相隔的两页中。
```sql
select rl from #{#entityName} rl where rl.xxx=xxx order by rl.createdAt desc limit 0,1
```

### 原因分析
查询语句以create_time进行倒序排序，通过limit进行分页，在正常情况下不会出现问题。

但当业务并发量比较大，导致create_time存在大量相同值时，再基于limit进行分页，就会出现乱序问题。

查看了Mysql 5.7和8.0的官方文档，描述如下：

If multiple rows have identical values in the ORDER BY columns, the server is free to return those rows in any order, and may do so differently depending on the overall execution plan. In other words, the sort order of those rows is nondeterministic with respect to the nonordered columns.

上述内容概述：在使用ORDER BY对列进行排序时，如果对应(ORDER BY的列)列存在多行相同数据，(Mysql)服务器会按照任意顺序返回这些行，并且可能会根据整体执行计划以不同的方式返回。

简单来说就是：ORDER BY查询的数据，如果ORDER BY列存在多行相同数据，Mysql会随机返回。这就会导致虽然使用了排序，但也会发生乱序的状况。

### 解决方案
可增加筛选条件
```sql
select rl from #{#entityName} rl where rl.xxx=xxx order by rl.createdAt,rl.id desc limit 0,1
```
