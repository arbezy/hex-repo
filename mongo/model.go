package mongo

import (
	"time"
)

type MongoEntity struct {
	ID        string              `bson:"_id,omitempty"`
	CreatedAt time.Time           `bson:"created_at,omitempty"`
	UpdatedAt time.Time           `bson:"updated_at,omitempty"`
	Fields    map[string]struct{} `bson:",inline"`
}
