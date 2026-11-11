package mongo

import ()

type MongoMapper[T any] struct {
}

func NewMongoMapper[T any]() (MongoMapper[T], error) {
	return MongoMapper[T]{}, nil
}
