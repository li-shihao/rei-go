package database

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"rei.io/rei/ent"
	"rei.io/rei/ent/accounts"
	"rei.io/rei/ent/objects"
	"rei.io/rei/ent/schema"
	"rei.io/rei/ent/sessions"
	"rei.io/rei/ent/users"
	"rei.io/rei/internal/crypto"
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

func (c *EntClient) QueryTotalTransactionCount() (*int, error) {
	count, err := c.client.Transactions.Query().Count(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed getting transaction count %w", err)
	}
	return &count, nil
}

func (c *EntClient) QueryUserExist(username string) (*bool, error) {
	exist, err := c.client.Users.Query().Where(users.UsernameEQ(username)).Exist(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed querying for user %w", err)
	}
	return &exist, nil
}

func (c *EntClient) QueryUserCredentials(username string, password string) (*bool, error) {
	user, err := c.client.Users.Query().Where(users.UsernameEQ(username)).Only(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed getting user %w", err)
	}

	canDecrypt := crypto.CheckEncryptPwdMatch(password, user.Hash)
	return &canDecrypt, nil
}

func (c *EntClient) CreateUser(username string, password string) (*ent.Users, error) {
	hashString, err := crypto.EncryptPwd(password)

	if err != nil {
		return nil, fmt.Errorf("failed creating hash string %w", err)
	}

	user, err := c.client.Users.Create().SetUsername(username).SetHash(*hashString).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating user %w", err)
	}
	return user, nil
}

func (c *EntClient) CreateSession(username string, ip string) (*ent.Sessions, error) {
	user, err := c.client.Users.Query().Where(users.UsernameEQ(username)).Exist(context.Background())
	if err != nil || !user {
		return nil, fmt.Errorf("user does not exist %w", err)
	}

	session, err := c.client.Sessions.Create().SetUsername(username).SetLoginIP(ip).SetLoginTime(time.Now()).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating user session %w", err)
	}
	return session, nil
}

func (c *EntClient) DeleteSession(username string) error {
	_, err := c.client.Sessions.Delete().Where(sessions.UsernameEQ(username)).Exec(context.Background())

	if err != nil {
		return fmt.Errorf("failed deleting user session %w", err)
	}
	return nil
}

func (c *EntClient) QuerySession(username string) (*bool, *string, error) {
	session, err := c.client.Sessions.Query().Where(sessions.UsernameEQ(username)).All(context.Background())

	if err != nil {
		return nil, nil, fmt.Errorf("failed querying user session %w", err)
	}
	if len(session) == 0 {
		loggedIn := false
		return &loggedIn, nil, nil
	}
	loggedIn := true
	return &loggedIn, &session[0].LoginIP, nil
}
