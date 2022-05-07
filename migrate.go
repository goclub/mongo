package mo

import (
	"context"
	xerr "github.com/goclub/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"reflect"
	"strings"
	"time"
)

func Migrate(db *Database, ptr interface{}) {
	log.Print("goclub/mongo: start migrate")
	defer func() {
		log.Print("goclub/mongo: end migrate")
	}()
	coll := NewCollection(db, "goclubMongoMigrateAction")
	rPtrValue := reflect.ValueOf(ptr)
	if rPtrValue.Kind() != reflect.Ptr {
		panic(xerr.New("ExecMigrate(db, ptr) ptr must be pointer"))
	}
	rValue := rPtrValue.Elem()
	rType := rValue.Type()
	methodNames := []string{}
	for i:=0;i<rType.NumMethod();i++ {
		method := rType.Method(i)
		if strings.HasPrefix(method.Name, "Migrate") {
			methodNames = append(methodNames, method.Name)
		}
	}
	for _, methodName := range methodNames {
		document := Action{}
		has, err := coll.FindOne(context.TODO(), bson.M{"name": methodName}, &document, FindOneCommand{}) ; if err != nil {
		    return
		}
		// Migrations are all manually triggered, with no concurrency concerns
		if has { continue }
		log.Print("goclub/mongo: exec: " +methodName)
		outs := rValue.MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(db)})
		errInterface := outs[0].Interface()
		if errInterface != nil {
			execErr := errInterface.(error)
			if execErr != nil {
				log.Printf("goclub/mongo: fail: %+v", execErr)
				break
			}
		}
		_, err = coll.InsertOne(context.TODO(), &Action{
			Name: methodName,
			Time: time.Now(),
		},InsertOneCommand{}) ; if err != nil {
		    return
		}
		log.Printf("goclub/mongo: done: " +methodName)
	}
}
// goclub_mongo_migrate_action
type Action struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Time time.Time `bson:"time"`
	DefaultLifeCycle
}

func (v *Action) AfterInsert(result ResultInsertOne) (err error) {
	if v.ID.IsZero() {
		v.ID, err = result.InsertedObjectID() ; if err != nil {
			return
		}
	}
	return
}