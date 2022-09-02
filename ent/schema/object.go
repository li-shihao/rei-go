package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Object holds the schema definition for the Object entity.
type Object struct {
	ent.Schema
}

// Fields of the Object.
func (Object) Fields() []ent.Field {
	return []ent.Field{
		field.String("Status"),
		field.String("DataType"),
		field.String("Type"),
		field.Bool("Has_public_transfer"),
		field.JSON("Fields", map[string]interface{}{}),
		field.String("Owner"),
		field.String("ObjectID"),
		field.Uint64("SequenceID"),
	}
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return nil
}

func (Object) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("SequenceID", "ObjectID").Unique(),
	}
}
