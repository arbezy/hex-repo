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

	// NOTE: ID is a special case! dont want to add this to fields!
	// NOTE: need to ignore the - field!

	// iterate through fields
	v := reflect.ValueOf(entity)
	for i := 0; i < v.NumField(); i++ {
		// match 'repo' tags
		repoTag := v.Type().Field(i).Tag.Get("repo")
		if repoTag != "" {
			// add to mongo entity fields
			mongoEntity.Fields[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	}

	// save the field type and map this to the new entity
	// need to create bson fields on the new entry

	fmt.Println(mongoEntity)
	fmt.Println("FIELDS:")
	fmt.Println(mongoEntity.Fields)

	return mongoEntity, nil
}
