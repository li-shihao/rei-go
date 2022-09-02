package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"rei.io/rei/ent"
	"rei.io/rei/ent/account"
	"rei.io/rei/graph/generated"
	"rei.io/rei/graph/model"
)

// SequenceID is the resolver for the SequenceID field.
func (r *accountResolver) SequenceID(ctx context.Context, obj *ent.Account) (int, error) {
	return int(obj.SequenceID), nil
}

// Balance is the resolver for the Balance field.
func (r *accountResolver) Balance(ctx context.Context, obj *ent.Account) (int, error) {
	return int(obj.Balance), nil
}

// Objects is the resolver for the Objects field.
func (r *accountResolver) Objects(ctx context.Context, obj *ent.Account) ([]*model.AccObject, error) {
	var objects []*model.AccObject
	for _, j := range obj.Objects {
		var object *model.AccObject
		object.ObjectID = j.ObjectId
		object.Type = j.Type
		object.Metadata = j.Metadata
		objects = append(objects, object)
	}
	return objects, nil
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, accountID string) (*ent.Account, error) {
	account, err := r.client.Account.Query().Where(account.AccountIDEQ(accountID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type accountResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }