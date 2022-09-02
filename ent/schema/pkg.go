package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Pkg holds the schema definition for the Pkg entity.
type Pkg struct {
	ent.Schema
}

// Fields of the Pkg.
func (Pkg) Fields() []ent.Field {
	return []ent.Field{
		field.String("TransactionID"),
		field.String("ObjectID"),
		field.JSON("Bytecode", map[string]interface{}{}),
	}
}

// Edges of the Pkg.
func (Pkg) Edges() []ent.Edge {
	return nil
}

func (Pkg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("TransactionID", "ObjectID").Unique(),
	}
}
