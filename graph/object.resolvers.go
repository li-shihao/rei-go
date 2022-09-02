package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"rei.io/rei/ent"
	"rei.io/rei/ent/object"
	"rei.io/rei/graph/generated"
)

// SequenceID is the resolver for the SequenceID field.
func (r *objectResolver) SequenceID(ctx context.Context, obj *ent.Object) (int, error) {
	return int(obj.SequenceID), nil
}

// Object is the resolver for the object field.
func (r *queryResolver) Object(ctx context.Context, objectID string) (*ent.Object, error) {
	object, err := r.client.Object.Query().Where(object.ObjectIDEQ(objectID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// Objects is the resolver for the Objects field.
func (r *queryResolver) Objects(ctx context.Context, owner string) ([]*ent.Object, error) {
	objects, err := r.client.Object.Query().Where(object.OwnerEQ(owner)).All(ctx)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

// Object returns generated.ObjectResolver implementation.
func (r *Resolver) Object() generated.ObjectResolver { return &objectResolver{r} }

type objectResolver struct{ *Resolver }
