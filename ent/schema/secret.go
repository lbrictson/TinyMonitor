package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Secret holds the schema definition for the Secret entity.
type Secret struct {
	ent.Schema
}

// Fields of the Secret.
func (Secret) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("id").MaxLen(50).NotEmpty().Immutable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("created_by").Default(""),
		field.String("value").MaxLen(3000).NotEmpty(),
	}
}

// Edges of the Secret.
func (Secret) Edges() []ent.Edge {
	return nil
}
