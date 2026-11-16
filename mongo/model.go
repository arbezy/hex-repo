package mongo

import (
	"time"
)

// NOTE: maybe I don't even need an ID field at all, since this will get created automatically when interacting with mongodb?
type MongoEntity struct {
	ID        string         `bson:"_id,omitempty"`
	CreatedAt time.Time      `bson:"created_at,omitempty"`
	UpdatedAt time.Time      `bson:"updated_at,omitempty"`
	Fields    map[string]any `bson:",inline"`
}
