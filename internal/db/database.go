package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"rei.io/rei/ent"
	"rei.io/rei/ent/accounts"
	"rei.io/rei/ent/objects"
	"rei.io/rei/ent/schema"
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

	// Nullable fields
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
		SetGas(tx.GetGas()).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating transaction: %w", err)
	}
	return txc, nil
}

func (c *EntClient) CreateEvent(evt sui.Event) (*ent.Events, error) {
	evtc, err := c.client.Events.Create().
		SetObjectID(evt.ObjectId).
		SetNillableRecipient(evt.Recipient).
		SetSender(evt.Sender).
		SetTransactionID(evt.TX).
		SetType(evt.Type).
		SetVersion(evt.Version).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating event: %w", err)
	}
	return evtc, nil
}

func (c *EntClient) CreateAccount(acc sui.Acc, sequence uint64) (*ent.Accounts, error) {

	// Type conversion from AccObj struct to ent version
	var obj []schema.AccObject
	for _, v := range acc.Objects {
		temp := schema.AccObject{}
		temp.Type = v.Type
		temp.Metadata = v.Metadata
		temp.ObjectId = v.ObjectId
		obj = append(obj, temp)
	}

	accc, err := c.client.Accounts.Create().
		SetAccountID(acc.ID).
		SetBalance(acc.Balance).
		SetObjects(obj).
		SetTransactions(acc.Transactions).
		SetSequenceID(sequence).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating account: %w", err)
	}
	return accc, nil
}

func (c *EntClient) CreateArgument(arg sui.Arg) (*ent.Arguments, error) {
	argc, err := c.client.Arguments.Create().
		SetData(arg.Data).
		SetName(arg.Name).
		SetTransactionID(arg.ID).
		SetType(arg.Type).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating argument: %w", err)
	}
	return argc, nil
}

func (c *EntClient) CreateNFT(obj sui.AccObject, sequence uint64) (*ent.NFTs, error) {
	nftc, err := c.client.NFTs.Create().
		SetType(obj.Type).
		SetMetadata(obj.Metadata).
		SetSequenceID(sequence).
		SetObjectID(obj.ObjectId).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating nft: %w", err)
	}
	return nftc, nil
}

func (c *EntClient) CreateObject(obj sui.Obj, sequence uint64) (*ent.Objects, error) {
	objc, err := c.client.Objects.Create().
		SetDataType(obj.GetObjectDataType()).
		SetFields(obj.GetObjectMetadata()).
		SetHasPublicTransfer(obj.HasPublicTransfer()).
		SetObjectID(obj.GetObjectID()).
		SetOwner(obj.GetOwner()).
		SetStatus(obj.GetObjectStatus()).
		SetType(obj.GetObjectType()).
		SetSequenceID(sequence).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating object: %w", err)
	}
	return objc, nil
}

func (c *EntClient) CreatePackage(pkg sui.Package) (*ent.Packages, error) {
	pkgc, err := c.client.Packages.Create().
		SetBytecode(pkg.Bytecode).
		SetObjectID(pkg.ID).
		SetTransactionID(pkg.DeployTX).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating package: %w", err)
	}
	return pkgc, nil
}

func (c *EntClient) QueryAccountFirstLoad(accId string, until uint64) bool {
	accc, err := c.client.Accounts.
		Query().
		Where(
			accounts.And(
				accounts.AccountID(accId),
				accounts.SequenceIDLTE(until),
			),
		).
		Exist(context.Background())

	if err != nil || !accc {
		return false
	} else {
		return true
	}
}

func (c *EntClient) QueryObjectFirstLoad(objId string, until uint64) bool {
	objc, err := c.client.Objects.
		Query().
		Where(
			objects.And(
				objects.ObjectID(objId),
				objects.SequenceIDLTE(until),
			),
		).
		Exist(context.Background())

	if err != nil || !objc {
		return false
	} else {
		return true
	}
}