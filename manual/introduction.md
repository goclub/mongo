---
permalink: /mongo/manual/introduction/
---

# MongoDB 简介

欢迎使用 MongoDB 5.0手册！MongoDB 是为便于开发和扩展而设计的 document 数据库。该手册介绍了 MongoDB 中的关键概念，提供了查询语言，并提供了操作和管理方面的考虑因素和过程，以及一个全面的参考部分。

## 安装 MongoDB

使用 docker-compose安装 [goclub/docker](https://github.com/goclub/docker/tree/main/mongo42)

## Document 数据库

MongoDB 中的记录是一个document，它是由字段和值对组成的数据结构。MongoDB  document 类似于 JSON 对象。字段的值可能包括其他 document 和数组还有 document 数组。

![](https://docs.mongodb.com/manual/images/crud-annotated-document.bakedsvg.svg)

使用 document 的好处是:

>  document 指的是 document

使用文件的好处是:

1. 在许多编程语言中(i.e. objects)  document 是原生类型
1. 嵌套(多级)的 document 和数组减少了多表 join 的需求
1. 动态模式能轻易的实现多态


## Collections/Views/On-Demand Materialized Views

MongoDB 将 document 存储在集合中，集合类似于关系数据库中的表。

除了集合之外，MongoDB 还支持:

1. Read-only [Views](/mongo/manual/core/views/) - 只读视图 (从 MongoDB 3.4开始支持)
2. [On-Demand Materialized Views](/mongo/manual/core/materialized-views/) - 按需实现视图 (从 MongoDB 4.2开始支持)

## 主要特点

### 高性能

MongoDB 提供高性能的数据持久性,

- 对嵌入式数据模型的支持减少了数据库系统上的 i/o 活动
- 索引支持更快的查询，可以包括嵌入 document 和 array 中的键

> 这里的嵌入式指的是多级结构,一般在sql中需要附属表来实现多级结构

### 丰富的查询语言

mongoDB 支持丰富的查询语言来支持读写操作(CRUD) ，以及:

- Data Aggregation [数据聚合](/mongo/manual/core/aggregation-pipeline/)
- Text Search [文字搜寻](/mongo/manual/text-search/) 和 Geospatial Queries [地理空间查询](/mongo/manual/tutorial/geospatial-tutorial/).

### 高可用性

MongoDB的 [副本集](/mongo/manual/replication/)提供:

- 自动故障转移
- 数据冗余

副本集是一组 MongoDB 服务器，它们维护相同的数据集，提供冗余并增加数据可用性。

### 水平可扩展性

MongoDB 核心功能的一部分是提供水平可伸缩性:

- 将数据[分布](/mongo/manual/sharding/#std-label-sharding-introduction)在一组机器上
- Starting in 3.4, MongoDB supports creating zones of data based on the shard key. In a balanced cluster, MongoDB directs reads and writes covered by a zone only to those shards inside the zone. See the Zones manual page for more information.

### 支持多种存储引擎

- [WiredTiger Storage Engine](/mongo/manual/core/security-encryption-at-rest/) (including support for Encryption at Rest)
- [In-Memory Storage Engine](/mongo/manual/core/inmemory/).

此外，MongoDB 提供了可插拔的存储引擎 API，允许第三方为 MongoDB 开发存储引擎。










