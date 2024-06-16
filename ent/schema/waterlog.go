package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// WaterLog holds the schema definition for the WaterLog entity.
type WaterLog struct {
	ent.Schema
}

// Fields of the WaterLog.
func (WaterLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("seconds"),
		field.String("channel"),
		field.Bool("manual"),
		field.Time("time").
			Default(time.Now),
	}
}

// Edges of the WaterLog.
func (WaterLog) Edges() []ent.Edge {
	return nil
}