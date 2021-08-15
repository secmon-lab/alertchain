package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/m-mizutani/alertchain/types"
)

// Alert holds the schema definition for the Alert entity.
type Alert struct {
	ent.Schema
}

// Fields of the Alert.
func (Alert) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(types.AlertID("")).Immutable(),
		field.String("title").Optional(),
		field.String("description").Optional(),
		field.String("detector").Optional(),
		field.String("status").GoType(types.AlertStatus("")).Default(string(types.StatusNew)),
		field.Time("created_at").Immutable().Default(func() time.Time {
			return time.Now().UTC()
		}),
		field.Time("closed_at").Optional().Nillable(),
	}
}

// Edges of the Alert.
func (Alert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attributes", Attribute.Type),
	}
}
