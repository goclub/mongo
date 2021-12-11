---
permalink: /manual/tutorial/getting-started/
---

# 开始

本教程指导您将测试数据插入到 MongoDB 数据库中，
并使用 [goclub/mongo](https://github.com/goclub/mongo)查询数据。

## 插入

将 document 存储在 [collections](/manual/core/databases-and-collections/) 中。集合类似于关系数据库中的表。如果一个集合不存在，MongoDB 将在您第一次为该集合存储数据时创建该集合。

下面的示例使用 [Collection.InsertMany](https://pkg.go.dev/github.com/goclub/mongo#Collection.InsertMany) 方法将新文档插入到电影集合中。

将存储在集合中。集合类似于关系数据库中的表。如果一个集合不存在，MongoDB 将在您第一次为该集合存储数据时创建该集合。

下面的示例使用 `Collection.InsertMany` 方法将新文档插入到电影集合中。

> goclub/mongo 使用了 [Document](https://pkg.go.dev/github.com/goclub/mongo#Document) 和 [ManyDocument](https://pkg.go.dev/github.com/goclub/mongo#ManyDocument)
> 作为一些方法的入参,定义一些满足 Document ManyDocument 接口的结构体能让减少很多重复的工作量.让代码更健壮

[点击查看示例代码](./getting-started-insert_test.go)


插入操作会返回包含每个成功插入的文档的 _id 的数组。 Collection.InsertMany 会自动将这些 _id 赋值给 mo.ManyExampleMovie,
因为 mo.ManyExampleMovie 实现了

```go
func (many ManyExampleMovie) BeforeInsertMany(data BeforeInsertManyData) (err error) {
	IDs := data.ObjectIDs()
	for i,_ := range many {
		many[i].ID = IDs[i]
	}
	return
}
```

要验证插入，可以查询集合

## 查询全部文档

要从集合中查询文档，可以使用  [Collection.Find](https://pkg.go.dev/github.com/goclub/mongo#Collection.Find)  方法。
若要选择集合中的所有文档，请将一个空的 `bson.M{}` 作为查询筛选器文档传递给该方法。

[点击查看示例代码](./getting-started-find_test.go)

## 按条件查询数据

> 你可能看到 bson.M bson.D bson.A bson.E 会有点懵
> 如果基于 json 去理解: `bson.A` 表示 `array`, `bson.D`和 `bson.M` 对应 `object`
> 因为 go 的map遍历时是无序的所以一般情况下操作符(filter)类型的参数用 `bson.M`,插入数据类型时用 `bson.D`
> `bson.E` 是用来实现 `bson.D`的,自己去看下他们的源码就能很快理解了

在 filter 中设置查询条件传递给 [Collection.Find](https://pkg.go.dev/github.com/goclub/mongo#Collection.Find) 方法。

查询2000年之前发布的电影:

```go
filter := bson.M{
    "released": bson.M{
        "$lt": time.Date(2000,1,1,0,0,0,0, time.UTC), // 中国时区用 time.FixedZone("CST", 8*3600) 代替 time.UTC
    },
}
```
查询获得100多个奖项的电影:

```go
filter := bson.M{
    "awards.wins": bson.M{
        "$gt": 100,
    },
}
```

查询 `languages` 包含`Japanese`或`Mandarin`的电影:

```go
filter := bson.M{
    "languages": bson.M{
        "$in": []string{"Japanese", "Mandarin"},
    },
}
```

[点击查看示例代码](./getting-started-filter-data_test.go)

## 指定返回字段

要指定要返回的字段，请将一个`mo.FindCommand{ Projection: bson.M{...} }`传递给 [Collection.Find](https://pkg.go.dev/github.com/goclub/mongo#Collection.Find) 方法。

- `<field>: 1` 在返回的文档中包含字段
- `<field>: 0` 在返回的文档中排除字段

在 Go 中，运行以下查询，返回电影集合中所有文档的 `id`、`title`、`directors`和`year`字段:

```go
cursor, err := moviesColl.Find(ctx, filter, mo.FindCommand{
        Projection: bson.M{
            "title": 1,
            "directors": 1,
            "year": 1,
        },
    }) ; if err != nil {
    return
}
    list := []bson.M{}
    err = cursor.All(ctx, &list) ; if err != nil {
    return
}
```

您不必指定 `_id` 字段来返回该字段。默认情况下它会返回。若要排除该字段，请在 `projection` 中将其设置为`0`。例如，运行以下查询只返回标题和匹配文档中的类型字段:

```go
cursor, err := moviesColl.Find(ctx, filter, mo.FindCommand{
    Projection: bson.M{
        "_id": 0,
        "title": 1,
        "genres": 1,
    },
}) ; if err != nil {
    return
}
list := []bson.M{}
err = cursor.All(ctx, &list) ; if err != nil {
    return
}
```

[点击查看示例代码](./getting-started-projection_test.go)

## 聚合

可以使用聚合对多个文档中的值进行分组，并返回单个结果。MongoDB 中的聚合是通过[聚合管道](/manual/aggregation/#std-label-aggregation-framework)执行的。

虽然 find() 操作也可以查询数据，但是与简单的 CRUD 操作相比，聚合管道允许您操作数据、执行计算和编写表达能力更强的查询。


在 shell 中，运行以下聚合管道来计算每种类型值的出现次数:

```shell
db.movies.aggregate( [
   { $unwind: "$genres" },
   {
     $group: {
       _id: "$genres",
       genreCount: { $sum: 1 }
     }
   },
   { $sort: { "genreCount": -1 } }
])
```

> 因为 5.0 版本之前不支持 `$group: {_id: "$genres", genreCount: { $count: { } } }` 所以本文档使用 `$sum: 1` 代替 $group 中的 $count.

管道用途:

- [`$unwind`](/manual/reference/operator/aggregation/unwind/#mongodb-pipeline-pipe.-unwind) 将指定数组(`$genres`)中的字段打平
- [`$group`](/manual/reference/operator/aggregation/group/#mongodb-pipeline-pipe.-group) 分组处理
- [`$sort`](/manual/reference/operator/aggregation/sort/#mongodb-pipeline-pipe.-sort) 排序


[点击查看示例代码](./getting-started-aggregate_test.go)


## 其他例子

### 查询文档示例

- [查询文档](/manual/tutorial/query-documents/)
- [嵌套式文档的查询](/manual/tutorial/query-embedded-documents/))
- [查询数组](/manual/tutorial/query-arrays/)
- [查询嵌入式文档数组](/manual/tutorial/query-array-of-documents/)
- [从查询返回的项目字段](/manual/tutorial/project-fields-from-query-results/)
- [查询空字段或缺失字段](/manual/tutorial/query-for-null-fields/)

### 更新文档示例

- [更新文档](/manual/tutorial/update-documents/)

### 删除文档示例
- [删除文档](/manual/tutorial/remove-documents/)

