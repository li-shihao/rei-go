package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"rei.io/rei/ent"
	"rei.io/rei/ent/argument"
)

// Argument is the resolver for the argument field.
func (r *queryResolver) Argument(ctx context.Context, transactionID string) (*ent.Argument, error) {
	argument, err := r.client.Argument.Query().Where(argument.TransactionIDEQ(transactionID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return argument, nil
}
