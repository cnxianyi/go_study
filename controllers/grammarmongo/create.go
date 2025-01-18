package grammarmongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

// omitempty 自动生成_id
type accountInfo struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	Account string        `bson:"account"`
}

type PhoneInfo struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	AccountId bson.ObjectID `bson:"accountId"`
	Phone     string        `bson:"phone"`
}
