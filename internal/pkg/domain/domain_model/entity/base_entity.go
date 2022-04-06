package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseEntity interface {
	SetID(primitive.ObjectID)
	CollectionName() string
}
