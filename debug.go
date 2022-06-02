package mo

import (
	"context"
	"encoding/json"
	"fmt"
	xerr "github.com/goclub/error"
	xtype "github.com/goclub/type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"regexp"
	"strconv"
	"strings"
)

var reDebugOIDToObjectID = regexp.MustCompile(`\{\"\$oid\":\"(.*?)\"\}`)
var reDebugNumberIntToInt = regexp.MustCompile(`\{\"\$numberInt\":\"(.*?)\"\}`)
var reDebugNumberIntToLong = regexp.MustCompile(`\{\"\$numberLong\":\"(.*?)\"\}`)
func debugMarshalBSON(v interface{}) []byte {
	bsonBytes, marshalErr := bson.Marshal(v) ; if marshalErr  != nil {
		// bson marshal 失败就用json marshal
		bsonBytes, marshalErr = json.Marshal(v) ; if marshalErr != nil {
			return []byte(fmt.Sprintf("goclub/mongo:lookup marshal bson error:\n%+v", marshalErr))
		} else {
			return bsonBytes
		}
	}
	return debugBsonRawToQuery(bson.Raw(bsonBytes))
}
func debugBsonRawToQuery(raw bson.Raw) []byte {
	bsonBytes := []byte(raw.String())
	bsonBytes = reDebugOIDToObjectID.ReplaceAll(bsonBytes, []byte(`ObjectId('$1')`))
	bsonBytes = reDebugNumberIntToInt.ReplaceAll(bsonBytes, []byte(`$1`))
	bsonBytes = reDebugNumberIntToLong.ReplaceAll(bsonBytes, []byte(`$1`))
	return bsonBytes
}
func (cmd FindOneCommand) checkAndRunLookupQuery(filter interface{}) {
	lookupQueryCmd{
		LookupQuery:  cmd.LookupQuery,
		Collation:    cmd.Collation,
		// Limit:        cmd.Limit,
		MaxTime:      cmd.MaxTime,
		Projection:   cmd.Projection,
		Skip:         cmd.Skip,
		Sort:         cmd.Sort,
	}.checkAndRunLookupQuery(filter)
}
func (cmd FindOneCommand) checkAndRunLookupResults(res *mongo.SingleResult) (err error) {
	if cmd.LookupResult {
		var raw bson.Raw
		raw, err = res.DecodeBytes() ; if err != nil {
			return
		}
		Logger.Printf("goclub/mongo:lookup result:\n%s", string(debugBsonRawToQuery(raw)))
	}
	return
}

func (cmd FindCommand) checkAndRunLookupQuery(filter interface{}) {
	lookupQueryCmd{
		LookupQuery:  cmd.LookupQuery,
		Collation:    cmd.Collation,
		Limit:        cmd.Limit,
		MaxTime:      cmd.MaxTime,
		Projection:   cmd.Projection,
		Skip:         cmd.Skip,
		Sort:         cmd.Sort,
	}.checkAndRunLookupQuery(filter)
}
func (cmd FindCommand) checkAndRunLookupResult(ctx context.Context, cursor *mongo.Cursor) (err error) {
	return checkAndRunLookupCursor(ctx, cmd.LookupResult, cursor)
}
func (cmd AggregateCommand) checkAndRunLookupQuery(query interface{}) {
	lookupQueryCmd{
		LookupQuery:  cmd.LookupQuery,
		Collation:    cmd.Collation,
		MaxTime:      cmd.MaxTime,
	}.checkAndRunLookupQuery(query)
}
func (cmd AggregateCommand) checkAndRunLookupResult(ctx context.Context, cursor *mongo.Cursor) (err error) {
	return checkAndRunLookupCursor(ctx, cmd.LookupResult, cursor)
}
func checkAndRunLookupCursor (ctx context.Context, lookupResult bool, cursor *mongo.Cursor) (err error) {
	if lookupResult == false {
		return
	}
	results := []map[string]interface{}{}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		debugItem := map[string]interface{}{}
		debugErr := cursor.Decode(&debugItem) ; if debugErr  != nil {
			Logger.Printf("goclub/mongo:\n%+v", debugErr)
		}
		results = append(results, debugItem)
	}
	debugResultsBytes, jsonMarshalErr := json.Marshal(results) ; if jsonMarshalErr != nil {
		Logger.Printf("goclub/mongo:\n%+v", jsonMarshalErr)
	} else {
		Logger.Printf("goclub/mongo:lookup result:\nlen:%d\nresults:\n%s", len(results), string(debugResultsBytes))
	}
	// 如果将开启了 LookupResult 的代码提交到线上,线上会以为 cursor.Next() 已经读完数据导致结果为空
	err = xerr.New("goclub/mongo:Because cmd.LookupResult is true, so return err")
	return
}

type lookupQueryCmd struct {
	LookupQuery  bool
	Collation *options.Collation
	Limit xtype.OptionUint64
	MaxTime xtype.OptionDuration
	Projection interface{}
	Skip xtype.OptionUint64
	Sort interface{}
}
func (cmd lookupQueryCmd) checkAndRunLookupQuery(query interface{}) {
	if cmd.LookupQuery  == false{
		return
	}
	message := []string{"goclub/mongo:lookup query:"}
	nilMessage := []string{}
	message = append(message, "filter:")
	message = append(message, string(debugMarshalBSON(query)))
	if cmd.Projection == nil {
		nilMessage = append(nilMessage, "projection")
	} else {
		message = append(message, "projection:" + string(debugMarshalBSON(cmd.Projection)))
	}
	if cmd.Sort == nil {
		nilMessage = append(nilMessage, "sort")
	} else {
		message = append(message, "sort:" + string(debugMarshalBSON(cmd.Sort)))
	}
	if cmd.MaxTime.Valid() == false {
		nilMessage = append(nilMessage, "maxTime")
	} else {
		message = append(message, "maxTime:" + cmd.MaxTime.Unwrap().String())
	}
	if cmd.Collation == nil {
		nilMessage = append(nilMessage, "collation")
	} else {
		message = append(message, "collation:" + string(debugMarshalBSON(cmd.Collation)))
	}
	if cmd.Skip.Valid() == false {
		nilMessage = append(nilMessage, "skip")
	} else {
		message = append(message, "skip:", strconv.FormatUint(cmd.Skip.Unwrap(), 10))
	}
	if cmd.Limit.Valid() == false {
		nilMessage = append(nilMessage, "limit")
	} else {
		message = append(message, "limit:", strconv.FormatUint(cmd.Limit.Unwrap(), 10))
	}
	message = append(message, "nil options:", strings.Join(nilMessage, " "))
	Logger.Print(strings.Join(message, "\n"))
}
