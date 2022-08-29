package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.String("AccountID"),
		field.Uint64("Balance"),
		field.JSON("Objects", []AccObject{}),
		field.Strings("Transactions"),
		field.Time("Time"),
	}
}

// Edges of the Accounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}
