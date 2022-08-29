package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Objects holds the schema definition for the Objects entity.
type Objects struct {
	ent.Schema
}

// Fields of the Objects.
func (Objects) Fields() []ent.Field {
	return []ent.Field{
		field.String("Status"),
		field.String("DataType"),
		field.String("Type"),
		field.Bool("Has_public_transfer"),
		field.JSON("Fields", map[string]interface{}{}),
		field.String("Owner"),
		field.String("ObjectID"),
	}
}

// Edges of the Objects.
func (Objects) Edges() []ent.Edge {
	return nil
}
