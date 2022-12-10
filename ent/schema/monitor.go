package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Monitor holds the schema definition for the Monitor entity.
type Monitor struct {
	ent.Schema
}

// Fields of the Monitor.
func (Monitor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Immutable(),
		field.String("description").Default(""),
		field.String("status"),
		field.Time("last_checked_at").Nillable().Optional(),
		field.Time("status_last_changed_at").Default(time.Now),
		field.String("monitor_type").Immutable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.JSON("config", map[string]interface{}{}),
		field.Int("interval_seconds").Default(60),
		field.Bool("paused").Default(false),
	}
}

// Edges of the Monitor.
func (Monitor) Edges() []ent.Edge {
	return nil
}
