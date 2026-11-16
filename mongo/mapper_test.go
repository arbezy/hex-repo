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

func TestMapper_EntityToMongoEntity_ID(t *testing.T) {
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

func TestMapper_EntityToMongoEntity_Fields(t *testing.T) {
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
	assert.Equal(t, testEntity.Name, mongoEntity.Fields["Name"])
	assert.Equal(t, testEntity.IsActive, mongoEntity.Fields["IsActive"])
}

func TestMapper_EntityToMongoEntity_Ignore(t *testing.T) {
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
	assert.Equal(t, "", mongoEntity.Fields["Internal"])
}
