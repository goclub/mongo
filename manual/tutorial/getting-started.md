---
permalink: /manual/tutorial/getting-started/
---

# 开始

本教程指导您将测试数据插入到 MongoDB 数据库中，
并使用 [goclub/mongo](https://github.com/goclub/mongo)查询数据。

## 插入

将 document 存储在 [collections](/mongo/manual/core/databases-and-collections/) 中。集合类似于关系数据库中的表。如果一个集合不存在，MongoDB 将在您第一次为该集合存储数据时创建该集合。

下面的示例使用 [Collection.InsertMany](https://pkg.go.dev/github.com/goclub/mongo#Collection.InsertMany)db.collection.insertMany ()方法将新文档插入到电影集合中。

将存储在集合中。集合类似于关系数据库中的表。如果一个集合不存在，MongoDB 将在您第一次为该集合存储数据时创建该集合。

下面的示例使用 `Collection.InsertMany` 方法将新文档插入到电影集合中。

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

## 查询



## 其他例子

### 查询文档示例

- [查询文档](/mongo/manual/tutorial/query-documents/)
- [嵌套式文档的查询](/mongo/manual/tutorial/query-embedded-documents/))
- [查询数组](/mongo/manual/tutorial/query-arrays/)
- [查询嵌入式文档数组](/mongo/manual/tutorial/query-array-of-documents/)
- [从查询返回的项目字段](/mongo/manual/tutorial/project-fields-from-query-results/)
- [查询空字段或缺失字段](/mongo/manual/tutorial/query-for-null-fields/)

### 更新文档示例

- [更新文档](/mongo/manual/tutorial/update-documents/)

### 删除文档示例
- [删除文档](/mongo/manual/tutorial/remove-documents/)

