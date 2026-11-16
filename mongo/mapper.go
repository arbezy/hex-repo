package mongo

import (
	"reflect"
	"time"
)

type MongoMapper[T any] struct {
}

func NewMongoMapper[T any]() (MongoMapper[T], error) {
	return MongoMapper[T]{}, nil
}

// TODO: need to add created at and updated at
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

	// NOTE: this will not preserve created at time, possible todo would be to look for a createdAt field and handle that like ID
	mongoEntity.CreatedAt = time.Now()
	mongoEntity.UpdatedAt = time.Now()

	return mongoEntity, nil
}
