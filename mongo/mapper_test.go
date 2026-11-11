package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEntity struct {
	ID       string `repo:"_id"`
	Name     string `repo:"name"`
	Email    string `repo:"email"`
	IsActive bool   `repo:"is_active"`
	Internal string `repo:"-"`
}

func TestMapper_InitMapper(t *testing.T) {
	mapper := MongoMapper[TestEntity]{}

	assert.NotNil(t, mapper)
}

func TestMapper_EntityToMongoEntity(t *testing.T) {
	mapper, _ := NewMongoMapper[TestEntity]()

	testEntity := TestEntity{
		ID:       "123def",
		Name:     "John Mongo",
		Email:    "johnmongo@mongo.com",
		IsActive: true,
		Internal: "ignoreme",
	}

	mongoEntity, err := mapper.ToMongo(testEntity)
	assert.Nil(t, err)

	assert.NotNil(t, mongoEntity)
	assert.Equal(t, testEntity.ID, mongoEntity.ID)
}
