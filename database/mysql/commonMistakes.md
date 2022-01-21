## 工作日常踩坑
### left join 右表筛选条件后移导致的筛选错误
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
