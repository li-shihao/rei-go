package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sort"
	"time"

	"rei.io/rei/ent"
	"rei.io/rei/ent/account"
	"rei.io/rei/ent/event"
	"rei.io/rei/ent/transaction"
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
		var object model.AccObject
		object.ObjectID = j.ObjectId
		object.Type = j.Type
		object.Metadata = j.Metadata
		objects = append(objects, &object)
	}
	return objects, nil
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, accountID string) (*ent.Account, error) {
	account, err := r.client.Account.Query().Where(account.AccountIDEQ(accountID)).Order(ent.Desc(account.FieldSequenceID)).Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	return account[0], nil
}

// AccountHistory is the resolver for the accountHistory field.
func (r *queryResolver) AccountHistory(ctx context.Context, accountID string) (*model.AccHistory, error) {
	var acchistory model.AccHistory
	var objects []*model.AccObject

	accountList, err := r.client.Account.Query().Where(account.AccountIDEQ(accountID)).Order(ent.Desc(account.FieldSequenceID)).Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}

	account := accountList[0]
	acchistory.AccountID = accountID
	acchistory.Balance = int(account.Balance)

	for _, j := range account.Objects {
		var object model.AccObject
		object.ObjectID = j.ObjectId
		object.Type = j.Type
		object.Metadata = j.Metadata
		objects = append(objects, &object)
	}
	acchistory.Objects = objects

	for _, k := range account.Transactions {
		transaction, err := r.client.Transaction.Query().Where(transaction.TransactionIDEQ(k)).Only(context.Background())
		if err != nil {
			return nil, err
		}

		event, err := r.client.Event.Query().Where(event.TransactionIDEQ(k)).All(context.Background())
		if err != nil {
			return nil, err
		}
		acchistory.Transactions = append(acchistory.Transactions, transaction)
		acchistory.Events = append(acchistory.Events, event...)
	}
	sort.Slice(acchistory.Transactions, func(i, j int) bool {
		return acchistory.Transactions[i].Time.Unix() > acchistory.Transactions[j].Time.Unix()
	})
	return &acchistory, nil
}

// Accounts is the resolver for the accounts field.
func (r *subscriptionResolver) Accounts(ctx context.Context) (<-chan []*model.HomeAcc, error) {
	c := make(chan []*model.HomeAcc)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			test, _ := r.client.QueryContext(context.Background(), "SELECT balance, account_id FROM accounts u1 WHERE sequence_id = (SELECT MAX(sequence_id) FROM accounts u2 WHERE u1.account_id = u2.account_id) ORDER BY balance DESC LIMIT 10;")

			var ns []*model.HomeAcc

			for test.Next() {
				var r model.HomeAcc
				test.Scan(&r.Balance, &r.AccountID)
				ns = append(ns, &r)
			}

			c <- ns
		}
	}()
	return c, nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type accountResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
