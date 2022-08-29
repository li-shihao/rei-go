package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Events holds the schema definition for the Events entity.
type Events struct {
	ent.Schema
}

// Fields of the Events.
func (Events) Fields() []ent.Field {
	return []ent.Field{
		field.String("Type"),
		field.String("Sender"),
		field.String("Recipient"),
		field.String("TransactionID"),
		field.String("ObjectID"),
		field.Uint32("Version"),
	}
}

// Edges of the Events.
func (Events) Edges() []ent.Edge {
	return nil
}
