package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

type AccObject struct {
	ObjectId string
	Type     string
	Metadata map[string]interface{}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("SequenceID"),
		field.String("AccountID"),
		field.Uint64("Balance"),
		field.JSON("Objects", []AccObject{}),
		field.Strings("Transactions").Optional(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("SequenceID", "AccountID").Unique(),
	}
}
