package database

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"rei.io/rei/ent"
	"rei.io/rei/ent/account"
	"rei.io/rei/ent/object"
	"rei.io/rei/ent/schema"
	"rei.io/rei/ent/session"
	"rei.io/rei/ent/user"
	"rei.io/rei/internal/crypto"
	"rei.io/rei/internal/sui"
)

type EntClient struct {
	client *ent.Client
}

func (c *EntClient) GetClient() *ent.Client {
	return c.client
}

func (c *EntClient) FirstRun(dbType string, dbOption string) {
	cl, err := ent.Open(dbType, dbOption)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := cl.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	cl.User.Create().SetUsername("arthur").SetHash("ENZNT+7rT+h9XRHUB1DUCQx6LZKqX/o1y5irrbckJIzbcvqpGEUhEuEaau7InLxJjucV/WcRgiyO").Save(context.Background())

	c.client = cl
}

func (c *EntClient) Init(dbType string, dbOption string) {
	cl, err := ent.Open(dbType, dbOption)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	c.client = cl
}

func (c *EntClient) CreateTransaction(tx sui.TX) (*ent.Transaction, error) {

	// Nullable fields
	rec := tx.GetRecipient()
	amt := tx.GetTransferAmount()
	pkg := tx.GetContractPackage()
	mod := tx.GetContractModule()
	fn := tx.GetContractFunction()

	txc, err := c.client.Transaction.Create().
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

func (c *EntClient) CreateEvent(evt sui.Event) (*ent.Event, error) {
	evtc, err := c.client.Event.Create().
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

func (c *EntClient) CreateAccount(acc sui.Acc, sequence uint64) (*ent.Account, error) {

	// Type conversion from AccObj struct to ent version
	var obj []schema.AccObject
	for _, v := range acc.Objects {
		temp := schema.AccObject{}
		temp.Type = v.Type
		temp.Metadata = v.Metadata
		temp.ObjectId = v.ObjectId
		obj = append(obj, temp)
	}

	accc, err := c.client.Account.Create().
		SetAccountID(acc.ID).
		SetBalance(acc.Balance).
		SetTransactions(acc.Transactions).
		SetObjects(obj).
		SetSequenceID(sequence).
		Save(context.Background())

	return accc, err
}

func (c *EntClient) UpsertAccount(acc sui.Acc, sequence uint64) error {

	// Type conversion from AccObj struct to ent version
	var obj []schema.AccObject
	for _, v := range acc.Objects {
		temp := schema.AccObject{}
		temp.Type = v.Type
		temp.Metadata = v.Metadata
		temp.ObjectId = v.ObjectId
		obj = append(obj, temp)
	}

	err := c.client.Account.Create().
		SetAccountID(acc.ID).
		SetBalance(acc.Balance).
		SetObjects(obj).
		SetSequenceID(sequence).
		OnConflictColumns(account.FieldAccountID).
		UpdateNewValues().
		Exec(context.Background())

	return err
}

func (c *EntClient) UpdateAccount(acc string, sequence uint64, transaction string) (*ent.Account, error) {

	orig, err := c.client.Account.Query().Where(account.AccountIDEQ(acc)).Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed fetching account to be updated: %w", err)
	}

	accc, err := orig.Update().
		SetTransactions(append(orig.Transactions, transaction)).
		SetSequenceID(sequence).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed updating account: %w", err)
	}
	return accc, nil
}

func (c *EntClient) CreateArgument(arg sui.Arg) (*ent.Argument, error) {
	argc, err := c.client.Argument.Create().
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

func (c *EntClient) CreateNFT(obj sui.AccObject, sequence uint64) (*ent.NFT, error) {
	nftc, err := c.client.NFT.Create().
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

func (c *EntClient) CreateObject(obj sui.Obj, sequence uint64) (*ent.Object, error) {
	objc, err := c.client.Object.Create().
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

func (c *EntClient) CreatePackage(pkg sui.Package) (*ent.Pkg, error) {
	pkgc, err := c.client.Pkg.Create().
		SetBytecode(pkg.Bytecode).
		SetObjectID(pkg.ID).
		SetTransactionID(pkg.DeployTX).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating package: %w", err)
	}
	return pkgc, nil
}

func (c *EntClient) QueryAccountFirstLoad(accId string) bool {
	accc, err := c.client.Account.
		Query().
		Where(account.AccountIDEQ(accId)).
		Exist(context.Background())

	if err != nil {
		return false
	}

	return accc
}

func (c *EntClient) QueryObjectFirstLoad(objId string) bool {
	objc, err := c.client.Object.
		Query().
		Where(object.ObjectID(objId)).
		Exist(context.Background())

	if err != nil || !objc {
		return false
	} else {
		return true
	}
}

func (c *EntClient) QueryTotalTransactionCount() (*int, error) {
	count, err := c.client.Transaction.Query().Count(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed getting transaction count %w", err)
	}
	return &count, nil
}

func (c *EntClient) QueryUserExist(username string) (*bool, error) {
	exist, err := c.client.User.Query().Where(user.UsernameEQ(username)).Exist(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed querying for user %w", err)
	}
	return &exist, nil
}

func (c *EntClient) QueryUserCredentials(username string, password string) (*bool, error) {
	user, err := c.client.User.Query().Where(user.UsernameEQ(username)).Only(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed getting user %w", err)
	}

	canDecrypt := crypto.CheckEncryptPwdMatch(password, user.Hash)
	return &canDecrypt, nil
}

func (c *EntClient) CreateUser(username string, password string) (*ent.User, error) {
	hashString, err := crypto.EncryptPwd(password)

	if err != nil {
		return nil, fmt.Errorf("failed creating hash string %w", err)
	}

	user, err := c.client.User.Create().SetUsername(username).SetHash(*hashString).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating user %w", err)
	}
	return user, nil
}

func (c *EntClient) CreateSession(username string, ip string) (*ent.Session, error) {
	user, err := c.client.User.Query().Where(user.UsernameEQ(username)).Exist(context.Background())
	if err != nil || !user {
		return nil, fmt.Errorf("user does not exist %w", err)
	}

	session, err := c.client.Session.Create().SetUsername(username).SetLoginIP(ip).SetLoginTime(time.Now()).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating user session %w", err)
	}
	return session, nil
}

func (c *EntClient) DeleteSession(username string) error {
	_, err := c.client.Session.Delete().Where(session.UsernameEQ(username)).Exec(context.Background())

	if err != nil {
		return fmt.Errorf("failed deleting user session %w", err)
	}
	return nil
}

func (c *EntClient) QuerySession(username string) (*bool, *string, error) {
	session, err := c.client.Session.Query().Where(session.UsernameEQ(username)).All(context.Background())

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
