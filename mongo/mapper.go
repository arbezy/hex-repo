package mongo

import (
	"reflect"
)

type MongoMapper[T any] struct {
}

func NewMongoMapper[T any]() (MongoMapper[T], error) {
	return MongoMapper[T]{}, nil
}

func (mm *MongoMapper[T]) ToMongo(entity T) (MongoEntity, error) {
	var mongoEntity MongoEntity
	mongoEntity.Fields = make(map[string]any)

	// iterate through fields
	v := reflect.ValueOf(entity)
	for i := 0; i < v.NumField(); i++ {
		// match 'repo' tags
		repoTag := v.Type().Field(i).Tag.Get("repo")
		if repoTag != "" {
			if repoTag == "_id" {
				// TODO: think I need to pad this since mongo ids have a set length
				mongoEntity.ID = v.Field(i).String()
			}
			// skip fields with dashes
			if repoTag == "-" {
				continue
			}
			// add to mongo entity fields
			mongoEntity.Fields[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	}

	return mongoEntity, nil
}
