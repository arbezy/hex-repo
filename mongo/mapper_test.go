package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestMapperEntity struct {
	ID       string `repo:"_id"`
	Name     string `repo:"name"`
	Email    string `repo:"email"`
	IsActive bool   `repo:"is_active"`
	Internal string `repo:"-"`
}

func TestMapper_MongoMapperGetIdFromEntity(t *testing.T) {
	mapper := NewMongoMapper[TestMapperEntity]{}

	assert.Equal(t, "ID", mapper.idField)
}

func TestMapper_EntityToMongoEntity(t *testing.T) {
	mapper := NewMongoMapper[TestMapperEntity]{}

	testEntity := TestMapperEntity{
		ID:       "123def",
		Name:     "John Mongo",
		Email:    "johnmongo@mongo.com",
		IsActive: true,
		Internal: "ignoreme",
	}

	mongoEntity := mapper.ToMongo(testEntity)

	assert.NotNil(t, mongoEntity)
	assert.Equal(t, testEntity.ID, mongoEntity.ID)
}
