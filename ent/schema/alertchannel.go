package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// AlertChannel holds the schema definition for the AlertChannel entity.
type AlertChannel struct {
	ent.Schema
}

// Fields of the AlertChannel.
func (AlertChannel) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("id").MaxLen(50).NotEmpty().Immutable(),
		field.String("alert_channel_type").Immutable(),
		field.JSON("config", map[string]interface{}{}),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AlertChannel.
func (AlertChannel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("monitors", Monitor.Type),
	}
}
