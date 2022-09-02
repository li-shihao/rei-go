package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Arguments holds the schema definition for the Arguments entity.
type Argument struct {
	ent.Schema
}

// Fields of the Arguments.
func (Argument) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.String("Type"),
		field.String("TransactionID"),
		field.String("Data"),
	}
}

// Edges of the Arguments.
func (Argument) Edges() []ent.Edge {
	return nil
}

func (Argument) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Name", "Type", "TransactionID").Unique(),
	}
}
