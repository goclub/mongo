package mo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AliasObjectID struct {
	primitive.ObjectID
}

func (id AliasObjectID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.ObjectID, []byte(id.Hex()), nil
}
func (id *AliasObjectID) UnmarshalBSONValue(t bsontype.Type, b[]byte) error {
	var err error
	id.ObjectID, err = primitive.ObjectIDFromHex(fmt.Sprintf("%x", b)) ; if err != nil {
		return err
	}
	return nil
}