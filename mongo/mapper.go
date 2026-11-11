package mongo

import ()

type MongoMapper[T any] struct {
}

func NewMongoMapper[T any]() (MongoMapper[T], error) {
	return MongoMapper[T]{}, nil
}

func (mm *MongoMapper[T]) ToMongo(T) (MongoEntity, error) {

	return MongoEntity{}, nil
}
