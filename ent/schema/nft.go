package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NFT holds the schema definition for the NFT entity.
type NFT struct {
	ent.Schema
}

// Fields of the NFT.
func (NFT) Fields() []ent.Field {
	return []ent.Field{
		field.String("ObjectID"),
		field.String("Type"),
		field.JSON("Metadata", map[string]interface{}{}),
		field.Int64("SequenceID"),
	}
}

// Edges of the NFT.
func (NFT) Edges() []ent.Edge {
	return nil
}

func (NFT) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ObjectID", "SequenceID").Unique(),
	}
}
