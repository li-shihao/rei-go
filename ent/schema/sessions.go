package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Sessions holds the schema definition for the Sessions entity.
type Sessions struct {
	ent.Schema
}

// Fields of the Sessions.
func (Sessions) Fields() []ent.Field {
	return []ent.Field{
		field.String("Username"),
		field.Time("LoginTime"),
		field.String("LoginIP"),
	}
}

// Edges of the Sessions.
func (Sessions) Edges() []ent.Edge {
	return nil
}
