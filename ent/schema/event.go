package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("Type"),
		field.String("Sender"),
		field.String("Recipient").Optional(),
		field.String("TransactionID"),
		field.String("ObjectID"),
		field.Uint32("Version"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("TransactionID", "Type", "ObjectID", "Version").Unique(),
	}
}
