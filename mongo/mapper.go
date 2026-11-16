package mongo

import (
	"fmt"
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

	// NOTE: need to ignore the - field!

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
			// add to mongo entity fields
			mongoEntity.Fields[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	}

	fmt.Println(mongoEntity)
	fmt.Println("FIELDS:")
	fmt.Println(mongoEntity.Fields)

	return mongoEntity, nil
}
