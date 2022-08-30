package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NFTs holds the schema definition for the NFTs entity.
type NFTs struct {
	ent.Schema
}

// Fields of the NFTs.
func (NFTs) Fields() []ent.Field {
	return []ent.Field{
		field.String("ObjectID"),
		field.String("Type"),
		field.JSON("Metadata", map[string]interface{}{}),
		field.Uint64("SequenceID"),
	}
}

// Edges of the NFTs.
func (NFTs) Edges() []ent.Edge {
	return nil
}
