package mo

import (
	xtype "github.com/goclub/type"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertOneCommand struct {
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://docs.mongodb.com/manual/core/schema-validation/ for more information about document
	// validation.
	ByPassDocumentValidation xtype.OptionBool
}

func (c InsertOneCommand) Options() (opt []*options.InsertOneOptions) {
	if c.ByPassDocumentValidation.Valid() {
		opt = append(opt, options.InsertOne().SetBypassDocumentValidation(c.ByPassDocumentValidation.Unwrap()))
	}
	return
}

type InsertManyCommand struct {
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://docs.mongodb.com/manual/core/schema-validation/ for more information about document
	// validation.
	ByPassDocumentValidation xtype.OptionBool

	// If true, no writes will be executed after one fails. The default value is true.
	Ordered xtype.OptionBool
}

func (c InsertManyCommand) Options() (opt []*options.InsertManyOptions) {
	if c.ByPassDocumentValidation.Valid() {
		opt = append(opt, options.InsertMany().SetBypassDocumentValidation(c.ByPassDocumentValidation.Unwrap()))
	}
	if c.Ordered.Valid() {
		opt = append(opt, options.InsertMany().SetOrdered(c.Ordered.Unwrap()))
	}
	return
}

type FindOneCommand struct {
	// If true, an operation on a sharded cluster can return partial results if some shards are down rather than
	// returning an error. The default value is false.
	AllowPartialResults xtype.OptionBool

	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *options.Collation

	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is the empty string, which means that no comment will be included in the logs.
	Comment xtype.OptionString

	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint interface{}

	// A document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max interface{}

	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	MaxTime xtype.OptionDuration

	// A document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min interface{}

	// A document describing which fields will be included in the document returned by the operation. The default value
	// is nil, which means all fields will be included.
	Projection interface{}

	// If true, the document returned by the operation will only contain fields corresponding to the index used. The
	// default value is false.
	ReturnKey xtype.OptionBool

	// If true, a $recordId field with a record identifier will be included in the document returned by the operation.
	// The default value is false.
	ShowRecordID xtype.OptionBool

	// The number of documents to skip before selecting the document to be returned. The default value is 0.
	Sort interface{}

	// A document specifying the sort order to apply to the query. The first document in the sorted order will be
	// returned. The driver will return an error if the sort parameter is a multi-key map.
	Skip xtype.OptionInt64
}

func (c FindOneCommand) Options() (opt []*options.FindOneOptions) {
	if c.Sort != nil {
		opt = append(opt, options.FindOne().SetSort(c.Sort))
	}
	if c.Skip.Valid() {
		options.FindOne().SetSkip(c.Skip.Unwrap())
	}
	if c.AllowPartialResults.Valid() {
		opt = append(opt, options.FindOne().SetAllowPartialResults(c.AllowPartialResults.Unwrap()))
	}
	if c.Collation != nil {
		opt = append(opt, options.FindOne().SetCollation(c.Collation))
	}
	if c.Comment.Valid() {
		opt = append(opt, options.FindOne().SetComment(c.Comment.Unwrap()))
	}
	if c.Hint != nil {
		opt = append(opt, options.FindOne().SetHint(c.Hint))
	}
	if c.Max != nil {
		opt = append(opt, options.FindOne().SetMax(c.Max))
	}
	if c.MaxTime.Valid() {
		opt = append(opt, options.FindOne().SetMaxTime(c.MaxTime.Unwrap()))
	}
	if c.Min != nil {
		opt = append(opt, options.FindOne().SetMin(c.Min))
	}
	if c.Projection != nil {
		opt = append(opt, options.FindOne().SetProjection(c.Projection))
	}
	if c.ReturnKey.Valid() {
		opt = append(opt, options.FindOne().SetReturnKey(c.ReturnKey.Unwrap()))
	}
	if c.ShowRecordID.Valid() {
		opt = append(opt, options.FindOne().SetShowRecordID(c.ShowRecordID.Unwrap()))
	}
	return
}

type FindCommand struct {
	DebugLookupQuery   bool
	DebugLookupResults bool
	// If true, the server can write temporary data to disk while executing the find operation. This option is only
	// valid for MongoDB versions >= 4.4. Server versions >= 3.2 will report an error if this option is specified. For
	// server versions < 3.2, the driver will return a client-side error if this option is specified. The default value
	// is false.
	AllowDiskUse xtype.OptionBool

	// If true, an operation on a sharded cluster can return partial results if some shards are down rather than
	// returning an error. The default value is false.
	AllowPartialResults xtype.OptionBool

	// The maximum number of documents to be included in each batch returned by the server.
	BatchSize xtype.OptionInt32

	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *options.Collation

	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is the empty string, which means that no comment will be included in the logs.
	Comment xtype.OptionString

	// Specifies the type of cursor that should be created for the operation. The default is NonTailable, which means
	// that the cursor will be closed by the server when the last batch of documents is retrieved.
	CursorType *options.CursorType

	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint interface{}

	// The maximum number of documents to return. The default value is 0, which means that all documents matching the
	// filter will be returned. A negative limit specifies that the resulting documents should be returned in a single
	// batch. The default value is 0.
	Limit xtype.OptionUint64

	// A document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max interface{}

	// The maximum amount of time that the server should wait for new documents to satisfy a tailable cursor query.
	// This option is only valid for tailable await cursors (see the CursorType option for more information) and
	// MongoDB versions >= 3.2. For other cursor types or previous server versions, this option is ignored.
	MaxAwaitTime xtype.OptionDuration

	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	MaxTime xtype.OptionDuration

	// A document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min interface{}

	// If true, the cursor created by the operation will not timeout after a period of inactivity. The default value
	// is false.
	NoCursorTimeout xtype.OptionBool

	// This option is for internal replication use only and should not be set.
	//
	// Deprecated: This option has been deprecated in MongoDB version 4.4 and will be ignored by the server if it is
	// set.
	OplogReplay xtype.OptionBool

	// A document describing which fields will be included in the documents returned by the operation. The default value
	// is nil, which means all fields will be included.
	Projection interface{}

	// If true, the documents returned by the operation will only contain fields corresponding to the index used. The
	// default value is false.
	ReturnKey xtype.OptionBool

	// If true, a $recordId field with a record identifier will be included in the documents returned by the operation.
	// The default value is false.
	ShowRecordID xtype.OptionBool

	// The number of documents to skip before adding documents to the result. The default value is 0.
	Skip xtype.OptionUint64

	// If true, the cursor will not return a document more than once because of an intervening write operation. The
	// default value is false.
	//
	// Deprecated: This option has been deprecated in MongoDB version 3.6 and removed in MongoDB version 4.0.
	Snapshot xtype.OptionBool

	// A document specifying the order in which documents should be returned.  The driver will return an error if the
	// sort parameter is a multi-key map.
	Sort interface{}
}

func (c FindCommand) Options() (opt []*options.FindOptions) {
	if c.AllowDiskUse.Valid() {
		opt = append(opt, options.Find().SetAllowDiskUse(c.AllowDiskUse.Unwrap()))
	}
	if c.AllowPartialResults.Valid() {
		opt = append(opt, options.Find().SetAllowPartialResults(c.AllowPartialResults.Unwrap()))
	}
	if c.BatchSize.Valid() {
		opt = append(opt, options.Find().SetBatchSize(c.BatchSize.Unwrap()))
	}
	if c.Collation != nil {
		opt = append(opt, options.Find().SetCollation(c.Collation))
	}
	if c.Comment.Valid() {
		opt = append(opt, options.Find().SetComment(c.Comment.Unwrap()))
	}
	if c.CursorType != nil {
		opt = append(opt, options.Find().SetCursorType(*c.CursorType))
	}
	if c.Hint != nil {
		opt = append(opt, options.Find().SetHint(c.Hint))
	}
	if c.Limit.Valid() {
		opt = append(opt, options.Find().SetLimit(int64(c.Limit.Unwrap())))
	}
	if c.Max != nil {
		opt = append(opt, options.Find().SetMax(c.Max))
	}
	if c.MaxAwaitTime.Valid() {
		opt = append(opt, options.Find().SetMaxAwaitTime(c.MaxAwaitTime.Unwrap()))
	}
	if c.MaxTime.Valid() {
		opt = append(opt, options.Find().SetMaxTime(c.MaxTime.Unwrap()))
	}
	if c.Min != nil {
		opt = append(opt, options.Find().SetMin(c.Min))
	}
	if c.NoCursorTimeout.Valid() {
		opt = append(opt, options.Find().SetNoCursorTimeout(c.NoCursorTimeout.Unwrap()))
	}
	// deprecated OplogReplay
	if c.Projection != nil {
		opt = append(opt, options.Find().SetProjection(c.Projection))
	}
	if c.ReturnKey.Valid() {
		opt = append(opt, options.Find().SetReturnKey(c.ReturnKey.Unwrap()))
	}
	if c.ShowRecordID.Valid() {
		opt = append(opt, options.Find().SetShowRecordID(c.ShowRecordID.Unwrap()))
	}
	if c.Skip.Valid() {
		opt = append(opt, options.Find().SetSkip(int64(c.Skip.Unwrap())))
	}
	// deprecated Snapshot
	if c.Sort != nil {
		opt = append(opt, options.Find().SetSort(c.Sort))
	}
	return
}

type UpdateCommand struct {
	// A set of filters specifying to which array elements an update should apply. This option is only valid for MongoDB
	// versions >= 3.6. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the update will apply to all array elements.
	ArrayFilters *options.ArrayFilters

	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://docs.mongodb.com/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation xtype.OptionBool

	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *options.Collation

	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will return an error
	// if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option
	// is specified. The driver will return an error if this option is specified during an unacknowledged write
	// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint interface{}

	// If true, a new document will be inserted if the filter does not match any documents in the collection. The
	// default value is false.
	Upsert xtype.OptionBool
}

func (c UpdateCommand) Options() (opt []*options.UpdateOptions) {
	if c.ArrayFilters != nil {
		opt = append(opt, options.Update().SetArrayFilters(*c.ArrayFilters))
	}
	if c.BypassDocumentValidation.Valid() {
		opt = append(opt, options.Update().SetBypassDocumentValidation(c.BypassDocumentValidation.Unwrap()))
	}
	if c.Collation != nil {
		opt = append(opt, options.Update().SetCollation(c.Collation))
	}
	if c.Hint != nil {
		opt = append(opt, options.Update().SetHint(c.Hint))
	}
	if c.Upsert.Valid() {
		opt = append(opt, options.Update().SetUpsert(c.Upsert.Unwrap()))
	}
	return
}

type AggregateCommand struct {
	// If true, the operation can write to temporary files in the _tmp subdirectory of the database directory path on
	// the server. The default value is false.
	AllowDiskUse xtype.OptionBool

	// The maximum number of documents to be included in each batch returned by the server.
	BatchSize xtype.OptionInt32

	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://docs.mongodb.com/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation xtype.OptionBool

	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *options.Collation

	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	MaxTime xtype.OptionDuration

	// The maximum amount of time that the server should wait for new documents to satisfy a tailable cursor query.
	// This option is only valid for MongoDB versions >= 3.2 and is ignored for previous server versions.
	MaxAwaitTime xtype.OptionDuration

	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is the empty string, which means that no comment will be included in the logs.
	Comment xtype.OptionString

	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The hint does not apply to $lookup and $graphLookup aggregation stages. The driver will return an
	// error if the hint parameter is a multi-key map. The default value is nil, which means that no hint will be sent.
	Hint interface{}

	// Specifies parameters for the aggregate expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let interface{}
}

func (c AggregateCommand) Options() (opt []*options.AggregateOptions) {
	if c.AllowDiskUse.Valid() {
		opt = append(opt, options.Aggregate().SetAllowDiskUse(c.AllowDiskUse.Unwrap()))
	}
	if c.BatchSize.Valid() {
		opt = append(opt, options.Aggregate().SetBatchSize(c.BatchSize.Unwrap()))
	}
	if c.BypassDocumentValidation.Valid() {
		opt = append(opt, options.Aggregate().SetBypassDocumentValidation(c.BypassDocumentValidation.Unwrap()))
	}
	if c.Collation != nil {
		opt = append(opt, options.Aggregate().SetCollation(c.Collation))
	}
	if c.MaxTime.Valid() {
		opt = append(opt, options.Aggregate().SetMaxTime(c.MaxTime.Unwrap()))
	}
	if c.MaxAwaitTime.Valid() {
		opt = append(opt, options.Aggregate().SetMaxAwaitTime(c.MaxAwaitTime.Unwrap()))
	}
	if c.Comment.Valid() {
		opt = append(opt, options.Aggregate().SetComment(c.Comment.Unwrap()))
	}
	if c.Hint != nil {
		opt = append(opt, options.Aggregate().SetHint(c.Hint))
	}
	if c.Let != nil {
		opt = append(opt, options.Aggregate().SetLet(c.Let))
	}
	return
}

// DeleteCommand represents options that can be used to configure DeleteOne and DeleteMany operations.
type DeleteCommand struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *options.Collation

	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.4 will return an error
	// if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option
	// is specified. The driver will return an error if this option is specified during an unacknowledged write
	// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint interface{}
}

func (c DeleteCommand) Options() (opt []*options.DeleteOptions) {
	if c.Collation != nil {
		opt = append(opt, options.Delete().SetCollation(c.Collation))
	}
	if c.Hint != nil {
		opt = append(opt, options.Delete().SetHint(c.Hint))
	}
	return
}

// CountCommand represents options that can be used to configure a CountDocuments operation.
type CountCommand struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *options.Collation

	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint interface{}

	// The maximum number of documents to count. The default value is 0, which means that there is no limit and all
	// documents matching the filter will be counted.
	Limit xtype.OptionUint64

	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there is
	// no time limit for query execution.
	MaxTime xtype.OptionDuration

	// The number of documents to skip before counting. The default value is 0.
	Skip xtype.OptionUint64
}
func (c CountCommand) Options() (opt []*options.CountOptions) {
	if c.Collation != nil {
		opt = append(opt, options.Count().SetCollation(c.Collation))
	}
	if c.Hint != nil {
		opt = append(opt, options.Count().SetHint(c.Hint))
	}
	if c.Limit.Valid() {
		opt = append(opt, options.Count().SetLimit(int64(c.Limit.Unwrap())))
	}
	if c.MaxTime.Valid() {
		opt = append(opt, options.Count().SetMaxTime(c.MaxTime.Unwrap()))
	}
	if c.Skip.Valid() {
		opt = append(opt, options.Count().SetSkip(int64(c.Skip.Unwrap())))
	}
	return
}