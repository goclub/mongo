package mo

import (
	xtype "github.com/goclub/type"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertOneCommand struct {
	ByPassDocumentValidation xtype.OptionBool
}
func (c InsertOneCommand) Options ()(opt []*options.InsertOneOptions) {
	if c.ByPassDocumentValidation.Valid() {
		opt = append(opt, options.InsertOne().SetBypassDocumentValidation(c.ByPassDocumentValidation.Unwrap()))
	}
	return
}
type InsertManyCommand struct {
	ByPassDocumentValidation xtype.OptionBool
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
	Sort interface{}
	Skip xtype.OptionInt64
	AllowPartialResults xtype.OptionBool
	Collation *options.Collation
	Comment *xtype.OptionString
	Hint interface{}
	Max interface{}
	MaxTime xtype.OptionDuration
	MinTime xtype.OptionDuration
	Projection interface{}
	ReturnKey xtype.OptionBool
	ShowRecordID xtype.OptionBool

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
	if c.Comment.Valid() {
		opt = append(opt, options.FindOne().SetComment(c.Comment.Unwrap()))
	}
	if c.Hint != nil {
		opt = append(opt, options.FindOne().SetHint(c.Hint))
	}
	if c.Hint != nil {
		opt = append(opt, options.FindOne().SetHint(c.Hint))
	}
	if c.Max != nil {
		options.FindOne().SetMax(c.Max)
	}
	if c.MaxTime.Valid() {
		options.FindOne().SetMaxTime(c.MaxTime.Unwrap())
	}
	if c.MinTime.Valid() {
		options.FindOne().SetMaxTime(c.MinTime.Unwrap())
	}
	if c.Projection != nil {
		options.FindOne().SetProjection(c.Projection)
	}
	if c.ReturnKey.Valid() {
		options.FindOne().SetReturnKey(c.ReturnKey.Unwrap())
	}
	if c.ShowRecordID.Valid() {
		options.FindOne().SetShowRecordID(c.ShowRecordID.Unwrap())
	}
	return
}