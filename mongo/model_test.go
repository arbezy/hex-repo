package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateModel(t *testing.T) {
	mongoModel := MongoEntity{
		ID: "12345677890",
	}

	assert.NotNil(t, mongoModel)
}
