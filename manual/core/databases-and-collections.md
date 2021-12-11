---
permalink: /manual/core/databases-and-collections/
---

# 数据库和集合

## 概览

MongoDB 将数据记录存储为文档 ,这些文档收集在集合中。一个数据库存储一个或多个集合。

## 数据库

在 MongoDB 中，数据库保存一个或多个集合。要选择要使用的数据库,请使用:

```go
package main
import (
    "context"
    mo "github.com/goclub/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
func main() {
    ctx := context.Background()
    uri := "mongodb://goclub:goclub@localhost:27017/goclub?authSource=goclub"
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri)) ; if err != nil {
        return
    }
    err = client.Ping(ctx, readpref.Primary()) ; if err != nil {
        return
    }
    db := mo.NewDatabase(client, "goclub")
}
```

### 创建数据库

如果数据库不存在，MongoDB 会在您第一次为该数据库存储数据时创建该数据库。因此，您可以直接插入数据:

```shell
use myNewDB
db.myNewCollection1.insertOne( { x: 1 } )
```

如果 `insertOne()` 数据库 `myNewDB` 和集合 `myNewCollection1` 尚不存在，则该操作会同时创建它们。确保数据库和集合名称都遵循 MongoDB[命名限制](/manual/reference/limits/#std-label-restrictions-on-db-names)。

## 集合

MongoDB 将文档存储在集合中。集合类似于关系数据库中的表。


![](https://docs.mongodb.com/manual/images/crud-annotated-collection.bakedsvg.svg)

### 创建一个集合

如果集合不存在，MongoDB 会在您第一次存储该集合的数据时创建该集合。

```shell
db.myNewCollection2.insertOne( { x: 1 } )
db.myNewCollection3.createIndex( { y: 1 } )
```

如果它们不存在，则insertOne()和 createIndex()操作都会创建它们各自的集合。确保集合名称遵循 MongoDB[命名限制](https://docs.mongodb.com/manual/reference/limits/#std-label-restrictions-on-db-names)。

### 显式创建

MongoDB 提供了db.createCollection()使用各种选项显式创建集合的方法，例如设置最大大小或文档验证规则。如果您未指定这些选项，则无需显式创建集合，因为 MongoDB 在您首次存储集合数据时会创建新集合。

要修改这些集合选项，请参阅[collMod](/manual/reference/command/collMod/#mongodb-dbcommand-dbcmd.collMod)。

### 文件验证


默认情况下，集合不要求其文档具有相同的架构；即单个集合中的文档不需要具有相同的字段集，并且一个字段的数据类型可以在集合内的文档之间不同。

但是，从 MongoDB 3.2 开始，您可以在更新和插入操作期间为集合强制执行文档验证规则。有关详细信息，请参阅[架构验证](https://docs.mongodb.com/manual/core/schema-validation/)。

### 修改文档结构

要更改集合中文档的结构，例如添加新字段、删除现有字段或将字段值更改为新类型，请将文档更新为新结构。


### Unique Identifiers 唯一标识符

> 3.6版中的新功能。

集合被分配一个不可变的UUID。集合 UUID 在副本集的所有成员和分片集群中的分片中保持不变。

要检索集合的 UUID，请运行 [listCollections](/manual/reference/command/listCollections/)命令或[db.getCollectionInfos](/manual/reference/method/db.getCollectionInfos/#mongodb-method-db.getCollectionInfos)方法。








