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
		field.String("DataType").Optional(),
		field.String("Type").Optional(),
		field.Bool("Has_public_transfer").Optional(),
		field.JSON("Fields", map[string]interface{}{}).Optional(),
		field.String("Owner").Optional(),
		field.String("ObjectID"),
		field.String("TransactionID"),
	}
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return nil
}

func (Object) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("TransactionID", "ObjectID").Unique(),
	}
}
