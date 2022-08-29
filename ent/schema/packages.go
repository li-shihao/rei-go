package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Packages holds the schema definition for the Packages entity.
type Packages struct {
	ent.Schema
}

// Fields of the Packages.
func (Packages) Fields() []ent.Field {
	return []ent.Field{
		field.String("TransactionID"),
		field.String("ObjectID"),
		field.JSON("Bytecode", map[string]interface{}{}),
	}
}

// Edges of the Packages.
func (Packages) Edges() []ent.Edge {
	return nil
}
