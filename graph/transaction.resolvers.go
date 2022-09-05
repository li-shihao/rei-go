package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"rei.io/rei/ent"
	"rei.io/rei/ent/transaction"
	"rei.io/rei/graph/generated"
)

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

// Tps is the resolver for the TPS field.
func (r *subscriptionResolver) Tps(ctx context.Context) (<-chan *float64, error) {
	c := make(chan *float64)
	go func() {
		time.Sleep(1 * time.Second)
		tps, _ := r.client.Transaction.Query().Where(transaction.And(transaction.TimeGT(time.Now().Add(-1 * time.Minute)))).Count(ctx)
		tps_ptr := float64(tps) / 60
		select {
		case c <- &tps_ptr:
		default:
			return
		}
	}()
	return c, nil
}

// Transactions is the resolver for the transactions field.
func (r *subscriptionResolver) Transactions(ctx context.Context) (<-chan []*ent.Transaction, error) {
	c := make(chan []*ent.Transaction)
	go func() {
		time.Sleep(5 * time.Second)
		transactions, _ := r.client.Transaction.Query().Order(ent.Desc(transaction.FieldTime)).Limit(10).All(ctx)
		log.Println(transactions)
		select {
		case c <- transactions:
		default:
			return
		}
	}()
	return c, nil
}

// Gas is the resolver for the Gas field.
func (r *transactionResolver) Gas(ctx context.Context, obj *ent.Transaction) (int, error) {
	return int(obj.Gas), nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Transaction returns generated.TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

type subscriptionResolver struct{ *Resolver }
type transactionResolver struct{ *Resolver }
