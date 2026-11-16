package mongo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestEntity struct {
	ID       string `repo:"_id"`
	Name     string `repo:"name"`
	Email    string `repo:"email"`
	IsActive bool   `repo:"is_active"`
	Internal string `repo:"-"`
}

var testEntity = TestEntity{
	ID:       "123def",
	Name:     "John Mongo",
	Email:    "johnmongo@mongo.com",
	IsActive: true,
	Internal: "ignoreme",
}

func TestMapper_InitMapper(t *testing.T) {
	mapper := MongoMapper[TestEntity]{}

	assert.NotNil(t, mapper)
}

func TestMapper_ToMongoEntity_ID(t *testing.T) {
	mapper, _ := NewMongoMapper[TestEntity]()

	mongoEntity, err := mapper.ToMongo(testEntity)
	assert.Nil(t, err)

	assert.NotNil(t, mongoEntity)
	assert.Equal(t, testEntity.ID, mongoEntity.ID)
}

func TestMapper_ToMongoEntity_Fields(t *testing.T) {
	mapper, _ := NewMongoMapper[TestEntity]()

	mongoEntity, err := mapper.ToMongo(testEntity)
	assert.Nil(t, err)

	assert.NotNil(t, mongoEntity)
	assert.Equal(t, testEntity.Name, mongoEntity.Fields["Name"])
	assert.Equal(t, testEntity.IsActive, mongoEntity.Fields["IsActive"])
}

func TestMapper_ToMongoEntity_Ignore(t *testing.T) {
	mapper, _ := NewMongoMapper[TestEntity]()

	mongoEntity, err := mapper.ToMongo(testEntity)
	assert.Nil(t, err)

	assert.NotNil(t, mongoEntity)
	assert.Nil(t, mongoEntity.Fields["Internal"])
}

func TestMapper_ToMongoEntity_CreatedandUpdated(t *testing.T) {
	mapper, _ := NewMongoMapper[TestEntity]()

	mongoEntity, err := mapper.ToMongo(testEntity)
	assert.Nil(t, err)

	assert.NotNil(t, mongoEntity)
	assert.NotNil(t, mongoEntity.CreatedAt)
	assert.NotNil(t, mongoEntity.UpdatedAt)
}

var testMongoEntity = MongoEntity{
	ID:        "1234",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Fields: map[string]any{
		"Name":     "Steve Mingo",
		"Email":    "stevemingo@hexrepo.com",
		"IsActive": false,
		"Age":      4,
	},
}

func TestMapper_FromMongoEntity_ID(t *testing.T) {
	mapper, _ := NewMongoMapper[TestEntity]()

	testEntity, err := mapper.FromMongo(testMongoEntity)
	assert.Nil(t, err)

	assert.NotNil(t, testEntity)
	assert.Equal(t, testMongoEntity.ID, testEntity.ID)
}
