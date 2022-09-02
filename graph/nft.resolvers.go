package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"rei.io/rei/ent"
	"rei.io/rei/ent/nft"
	"rei.io/rei/graph/generated"
)

// SequenceID is the resolver for the SequenceID field.
func (r *nFTResolver) SequenceID(ctx context.Context, obj *ent.NFT) (int, error) {
	return int(obj.SequenceID), nil
}

// Nft is the resolver for the nft field.
func (r *queryResolver) Nft(ctx context.Context, objectID string) (*ent.NFT, error) {
	nft, err := r.client.NFT.Query().Where(nft.ObjectIDEQ(objectID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return nft, nil
}

// NFT returns generated.NFTResolver implementation.
func (r *Resolver) NFT() generated.NFTResolver { return &nFTResolver{r} }

type nFTResolver struct{ *Resolver }
