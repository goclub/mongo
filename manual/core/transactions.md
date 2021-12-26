---
permalink: /manual/core/transactions/
---

# Transactions  事务

在 MongoDB 中，对单个文档的操作是原子操作。因为您可以使用嵌入的文档和数组来管理单个文档结构中的数据之间的关系，而不是跨多个文档和集合进行规范化，所以这种单文档原子性避免了许多实际用例对多文档事务的需求。

> 译者:你应该尽量在单个集合中通过嵌入文档(多级结构)来避免使用事务

在单个或多个集合中需要进行原子性操作时,MongoDB支持多文档事务.对于分布式事务，可以跨集合、数据库、文档和分片使用事务。

## Transactions API


此示例主要展示事务API

这个示例使用 callback API处理事务,开启事务执行指定操作并提交事务(or abours on error).

The new callback API incorporates retry logic for "[TransientTransactionError](/manual/core/transactions-in-applications/#std-label-transient-transaction-error)" or "[UnknownTransactionCommitResult](https://docs.mongodb.com/manual/core/transactions-in-applications/#std-label-unknown-transaction-commit-result)" commit errors.

> **重要** 
> - Recommended. Use the MongoDB driver updated for the version of your MongoDB deployment. For transactions on MongoDB 4.2 deployments (replica sets and sharded clusters), clients must use MongoDB drivers updated for MongoDB 4.2. 
> - 当使用驱动程序时，事务中的每个操作都必须与会话关联(将`mongo.SessionContext`传递给每个操作函数)。
> - 事务中操作使用 [transaction-level read concern - 事务级别读](#std-label-transactions-read-concern) [transaction-level write concern - 事务级别写](#std-label-transactions-write-concern) [transaction-level read preference - 读优先](#std-label-transactions-read-preference)
> - 在MongoDB 4.2和更早的版本中，不能在事务中创建集合。如果在事务内部运行，导致文档插入的写操作(例如使用upsert: true进行插入或更新操作)必须针对现有的集合。
> - 从MongoDB 4.4开始，您可以隐式或显式地在事务中创建集合。请参见在事务中创建集合和索引。


[示例代码](./transactions_test.go?blob)





