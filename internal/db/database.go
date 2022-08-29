package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"rei.io/rei/ent"
	"rei.io/rei/internal/sui"
)

type EntClient struct {
	client *ent.Client
}

func (c *EntClient) Init(dbType string, dbOption string) {
	cl, err := ent.Open(dbType, dbOption)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := cl.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	fmt.Println("Connected to database successfully")
	c.client = cl
}

func (c *EntClient) CreateTransaction(tx sui.TX) (*ent.Transactions, error) {

	rec := tx.GetRecipient()
	amt := tx.GetTransferAmount()
	pkg := tx.GetContractPackage()
	mod := tx.GetContractModule()
	fn := tx.GetContractFunction()

	txc, err := c.client.Transactions.Create().
		SetType(tx.GetType()).
		SetTime(tx.GetTime()).
		SetTransactionID(tx.GetID()).
		SetStatus(tx.GetStatus()).
		SetSender(tx.GetSender()).
		SetNillableRecipient(rec).
		SetNillableAmount(amt).
		SetNillablePackage(pkg).
		SetNillableModule(mod).
		SetNillableFunction(fn).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating transaction: %w", err)
	}
	log.Println("transaction was created: ", txc)
	return txc, nil
}

func (c *EntClient) CreateEvent(evt sui.Event) (*ent.Events, error) {
	evtc, err := c.client.Events.Create().
		SetObjectID(evt.ObjectId).
		SetRecipient(evt.Recipient).
		SetSender(evt.Sender).
		SetTransactionID(evt.TX).
		SetType(evt.Type).
		SetVersion(evt.Version).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating transaction: %w", err)
	}
	log.Println("transaction was created: ", evtc)
	return evtc, nil
}
