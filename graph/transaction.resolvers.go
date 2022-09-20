package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"rei.io/rei/ent"
	"rei.io/rei/ent/argument"
	"rei.io/rei/ent/object"
	"rei.io/rei/ent/transaction"
	"rei.io/rei/graph/generated"
	"rei.io/rei/graph/model"
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
		for {
			time.Sleep(time.Second)
			tps, _ := r.client.Transaction.Query().Where(transaction.And(transaction.TimeGT(time.Now().Add(-1 * time.Minute)))).Count(ctx)
			tps_ptr := float64(tps) / 60
			c <- &tps_ptr
		}
	}()
	return c, nil
}

// Transactions is the resolver for the transactions field.
func (r *subscriptionResolver) Transactions(ctx context.Context) (<-chan []*ent.Transaction, error) {
	c := make(chan []*ent.Transaction)
	go func() {
		for {
			time.Sleep(time.Second)
			transactions, _ := r.client.Transaction.Query().Order(ent.Desc(transaction.FieldTime)).Limit(10).All(ctx)
			c <- transactions
		}
	}()
	return c, nil
}

// LastDay is the resolver for the lastDay field.
func (r *subscriptionResolver) LastDay(ctx context.Context) (<-chan []*model.Frequency, error) {
	c := make(chan []*model.Frequency)
	go func() {
		for {
			test, _ := r.client.QueryContext(context.Background(), "SELECT dd, count('time') FROM generate_series( date_bin('5 minutes', now(), timestamp '2001-01-01') - interval '1 days' + interval '1 hour', date_bin('5 minutes', now(), timestamp '2001-01-01') , interval '5 minutes') dd LEFT JOIN transactions ON dd = date_bin('5 minutes', time, timestamp '2001-01-01') GROUP BY dd")
			var ns []*model.Frequency

			for test.Next() {
				var r model.Frequency
				test.Scan(&r.Time, &r.Count)
				ns = append(ns, &r)
			}
			c <- ns
			time.Sleep(time.Minute)
		}
	}()
	return c, nil
}

// HotModule is the resolver for the hotModule field.
func (r *subscriptionResolver) HotModule(ctx context.Context) (<-chan []*model.HotModule, error) {
	c := make(chan []*model.HotModule)

	go func() {
		for {
			test, _ := r.client.QueryContext(context.Background(), "SELECT package, MODULE, COUNT(*) FROM transactions WHERE package IS NOT NULL AND TIME + (1 * INTERVAL '1 day') > NOW() GROUP BY package, MODULE ORDER BY COUNT DESC LIMIT 10")
			var ns []*model.HotModule

			for test.Next() {
				var r model.HotModule
				test.Scan(&r.Package, &r.Module, &r.Count)
				ns = append(ns, &r)
			}
			c <- ns
		}
	}()
	return c, nil
}

// Args is the resolver for the Args field.
func (r *transactionResolver) Args(ctx context.Context, obj *ent.Transaction) ([]*ent.Argument, error) {
	argc, err := r.client.Argument.Query().Where(argument.TransactionIDEQ(obj.TransactionID)).All(context.Background())
	if err != nil {
		return nil, err
	}
	return argc, nil
}

// Gas is the resolver for the Gas field.
func (r *transactionResolver) Gas(ctx context.Context, obj *ent.Transaction) (int, error) {
	return int(obj.Gas), nil
}

// Change is the resolver for the Change field.
func (r *transactionResolver) Change(ctx context.Context, obj *ent.Transaction) ([]*model.Change, error) {
	var tmp []*model.Change
	for _, k := range obj.Changed {
		var inner model.Change
		inner.Type = k.Type
		inner.Version = k.Version
		if k.Version > 1 {
			bef, err := r.client.Object.Query().Where(object.And(object.ObjectIDEQ(k.ObjectId), object.VersionEQ(k.Version-1))).Only(context.Background())
			if err != nil {
				return nil, err
			}
			inner.Before = bef
		}

		aft, err := r.client.Object.Query().Where(object.And(object.ObjectIDEQ(k.ObjectId), object.VersionEQ(k.Version))).Only(context.Background())
		if err != nil {
			return nil, err
		}
		inner.After = aft

		tmp = append(tmp, &inner)
	}
	return tmp, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Transaction returns generated.TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

type subscriptionResolver struct{ *Resolver }
type transactionResolver struct{ *Resolver }
