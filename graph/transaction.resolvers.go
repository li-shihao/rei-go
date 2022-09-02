package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"rei.io/rei/ent"
	"rei.io/rei/ent/transaction"
	"rei.io/rei/graph/generated"
)

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context) ([]*ent.Transaction, error) {
	transactions, err := r.client.Transaction.Query().Order(ent.Desc(transaction.FieldTime)).Limit(10).All(ctx)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// Transaction is the resolver for the transaction field.
func (r *queryResolver) Transaction(ctx context.Context, transactionID string) (*ent.Transaction, error) {
	transaction, err := r.client.Transaction.Query().Where(transaction.TransactionIDEQ(transactionID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// TotalTransactions is the resolver for the totalTransactions field.
func (r *queryResolver) TotalTransactions(ctx context.Context) (*int, error) {
	total, err := r.client.Transaction.Query().Count(ctx)
	if err != nil {
		return nil, err
	}
	return &total, nil
}

// Gas is the resolver for the Gas field.
func (r *transactionResolver) Gas(ctx context.Context, obj *ent.Transaction) (int, error) {
	return int(obj.Gas), nil
}

// Transaction returns generated.TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

type transactionResolver struct{ *Resolver }
