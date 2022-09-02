package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
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
		field.Uint32("Gas"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return nil
}
