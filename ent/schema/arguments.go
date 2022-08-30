package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Argumentss holds the schema definition for the Argumentss entity.
type Arguments struct {
	ent.Schema
}

// Fields of the Argumentss.
func (Arguments) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.String("Type"),
		field.String("TransactionID"),
		field.String("Data"),
	}
}

// Edges of the Argumentss.
func (Arguments) Edges() []ent.Edge {
	return nil
}

func (Arguments) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Name", "Type", "TransactionID").Unique(),
	}
}
