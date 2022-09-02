package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"rei.io/rei/ent"
	"rei.io/rei/graph/generated"
)

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context) (*ent.Transactions, error) {
	return r.client.Transactions.Query().Only(ctx)
}

// Time is the resolver for the Time field.
func (r *transactionsResolver) Time(ctx context.Context, obj *ent.Transactions) (string, error) {
	return obj.Time.Format(time.RFC3339), nil
}

// Gas is the resolver for the Gas field.
func (r *transactionsResolver) Gas(ctx context.Context, obj *ent.Transactions) (int, error) {
	return int(obj.Gas), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Transactions returns generated.TransactionsResolver implementation.
func (r *Resolver) Transactions() generated.TransactionsResolver { return &transactionsResolver{r} }

type queryResolver struct{ *Resolver }
type transactionsResolver struct{ *Resolver }
