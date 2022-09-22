package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ID string

func ToID(oid interface{}) ID {
	return ID(oid.(primitive.ObjectID).Hex())
}
