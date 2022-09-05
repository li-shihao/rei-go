package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"rei.io/rei/ent"
	"rei.io/rei/ent/nft"
	"rei.io/rei/graph/generated"
	"rei.io/rei/graph/model"
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

// Nfts is the resolver for the nfts field.
func (r *subscriptionResolver) Nfts(ctx context.Context) (<-chan []*model.NFTCount, error) {
	c := make(chan []*model.NFTCount)
	go func() {

		var v []struct {
			Type  string `json:"type"`
			Count int    `json:"count"`
		}

		var temp []*model.NFTCount

		time.Sleep(1 * time.Second)
		_ = r.client.NFT.Query().
			Where(
				nft.Not(
					nft.TypeHasPrefix("0x2::coin"),
				),
				nft.Not(
					func(s *sql.Selector) {
						s.Where(sqljson.ValueEQ(nft.FieldMetadata, "{}"))
					},
				),
			).
			GroupBy(nft.FieldType).
			Aggregate(ent.Count()).
			Scan(ctx, &v)

		sort.Slice(v, func(i, j int) bool {
			return v[i].Count > v[j].Count
		})

		for i, d := range v {
			var inner model.NFTCount
			temp = append(temp, &inner)
			temp[i].Count = d.Count
			temp[i].Type = d.Type
		}

		log.Println(temp)

		select {
		case c <- temp:

		default:
			return
		}
	}()
	return c, nil
}

// NFT returns generated.NFTResolver implementation.
func (r *Resolver) NFT() generated.NFTResolver { return &nFTResolver{r} }

type nFTResolver struct{ *Resolver }
