package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"rei.io/rei/ent"
	"rei.io/rei/ent/pkg"
)

// Pkg is the resolver for the pkg field.
func (r *queryResolver) Pkg(ctx context.Context, objectID *string, transactionID *string) (*ent.Pkg, error) {
	if objectID != nil {
		pkg, err := r.client.Pkg.Query().Where(pkg.ObjectIDEQ(*objectID)).Only(ctx)
		if err != nil {
			return nil, err
		}
		return pkg, nil
	}
	if transactionID != nil {
		pkg, err := r.client.Pkg.Query().Where(pkg.TransactionIDEQ(*transactionID)).Only(ctx)
		if err != nil {
			return nil, err
		}
		return pkg, nil
	}
	return nil, errors.New("no parameters fulfilled")
}
