package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"rei.io/rei/ent"
	"rei.io/rei/ent/event"
	"rei.io/rei/ent/object"
	"rei.io/rei/graph/generated"
)

// Object is the resolver for the Object field.
func (r *eventResolver) Object(ctx context.Context, obj *ent.Event) (*ent.Object, error) {
	objc, err := r.client.Object.Query().Where(object.And(object.TransactionIDEQ(obj.TransactionID), object.ObjectIDEQ(obj.ObjectID))).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return objc, nil
}

// Version is the resolver for the Version field.
func (r *eventResolver) Version(ctx context.Context, obj *ent.Event) (int, error) {
	return int(obj.Version), nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context, transactionID string) ([]*ent.Event, error) {
	events, err := r.client.Event.Query().Where(event.TransactionIDEQ(transactionID)).All(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
