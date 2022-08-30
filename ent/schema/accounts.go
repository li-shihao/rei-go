package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Accounts holds the schema definition for the Accounts entity.
type Accounts struct {
	ent.Schema
}

type AccObject struct {
	ObjectId string
	Type     string
	Metadata map[string]interface{}
}

// Fields of the Accounts.
func (Accounts) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("SequenceID"),
		field.String("AccountID"),
		field.Uint64("Balance"),
		field.JSON("Objects", []AccObject{}),
		field.Strings("Transactions"),
	}
}

// Edges of the Accounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}

func (Accounts) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("SequenceID", "AccountID").Unique(),
	}
}
