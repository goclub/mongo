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
func (c InsertOneCommand) Options ()(opt []*options.InsertOneOptions) {
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
func (c InsertManyCommand) Options ()(opt []*options.InsertManyOptions) {
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
func (c FindOneCommand) Options ()(opt []*options.FindOneOptions) {
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
func (c UpdateCommand) Options ()(opt []*options.UpdateOptions) {
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