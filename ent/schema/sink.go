package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Sink holds the schema definition for the Sink entity.
type Sink struct {
	ent.Schema
}

// Fields of the Sink.
func (Sink) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("id").MaxLen(50).NotEmpty().Immutable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("sink_type").MaxLen(150).NotEmpty(),
		field.JSON("config", map[string]interface{}{}),
	}
}

// Edges of the Sink.
func (Sink) Edges() []ent.Edge {
	return nil
}
