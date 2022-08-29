package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Transactions holds the schema definition for the Transactions entity.
type Transactions struct {
	ent.Schema
}

// Fields of the Transactions.
func (Transactions) Fields() []ent.Field {
	return []ent.Field{
		field.String("Type"),
		field.Time("Time"),
		field.String("TransactionID").Unique(),
		field.Bool("Status"),
		field.String("Sender"),
		field.String("Recipient").Optional(),
		field.Float("Amount").Optional(),
		field.String("Package").Optional(),
		field.String("Module").Optional(),
		field.String("Function").Optional(),
	}
}

// Edges of the Transactions.
func (Transactions) Edges() []ent.Edge {
	return nil
}
