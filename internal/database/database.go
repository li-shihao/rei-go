package database

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	_ "github.com/lib/pq"
	"rei.io/rei/ent"
	"rei.io/rei/ent/account"
	"rei.io/rei/ent/object"
	"rei.io/rei/ent/schema"
	"rei.io/rei/ent/session"
	"rei.io/rei/ent/user"
	"rei.io/rei/internal/crypto"
	"rei.io/rei/internal/helpers"
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

	var entChanged []schema.Changed
	for _, k := range *tx.Changed {
		var tmp schema.Changed
		tmp.ObjectId = k.ObjectId
		tmp.Version = k.Version
		tmp.Type = k.Type

		entChanged = append(entChanged, tmp)
	}

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
		SetChanged(entChanged).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating transaction: %w", err)
	}
	return txc, nil
}

func (c *EntClient) CreateTransactionFromSocket(tx sui.SocketTX) (*ent.Transaction, error) {

	// Get type
	var _type string
	for k := range tx.Certificate.Kind.Single {
		_type = k
	}

	// Get status
	var _status bool
	switch tx.Effects.Status {
	case "Success":
		_status = true
	default:
		_status = false
	}

	// Get changed
	var changedSet []schema.Changed
	if len(tx.Effects.Created.([]interface{})) > 0 {
		for _, j := range tx.Effects.Created.([]interface{}) {
			var temp schema.Changed
			temp.ObjectId = j.([]interface{})[0].([]interface{})[0].(string)
			temp.Type = "Created"
			temp.Version = int(j.([]interface{})[0].([]interface{})[1].(float64))
			changedSet = append(changedSet, temp)
		}
	}
	if len(tx.Effects.Mutated.([]interface{})) > 0 {
		for _, j := range tx.Effects.Created.([]interface{}) {
			var temp schema.Changed
			temp.ObjectId = j.([]interface{})[0].([]interface{})[0].(string)
			temp.Type = "Mutated"
			temp.Version = int(j.([]interface{})[0].([]interface{})[1].(float64))
			changedSet = append(changedSet, temp)
		}
	}
	if len(tx.Effects.Deleted.([]interface{})) > 0 {
		for _, j := range tx.Effects.Created.([]interface{}) {
			var temp schema.Changed
			temp.ObjectId = j.([]interface{})[0].([]interface{})[0].(string)
			temp.Type = "Deleted"
			temp.Version = int(j.([]interface{})[0].([]interface{})[1].(float64))
			changedSet = append(changedSet, temp)
		}
	}
	if len(tx.Effects.Shared_objects.([]interface{})) > 0 {
		for _, j := range tx.Effects.Created.([]interface{}) {
			var temp schema.Changed
			temp.ObjectId = j.([]interface{})[0].([]interface{})[0].(string)
			temp.Type = "Shared"
			temp.Version = int(j.([]interface{})[0].([]interface{})[1].(float64))
			changedSet = append(changedSet, temp)
		}
	}
	if len(tx.Effects.Wrapped.([]interface{})) > 0 {
		for _, j := range tx.Effects.Created.([]interface{}) {
			var temp schema.Changed
			temp.ObjectId = j.([]interface{})[0].([]interface{})[0].(string)
			temp.Type = "Wrapped"
			temp.Version = int(j.([]interface{})[0].([]interface{})[1].(float64))
			changedSet = append(changedSet, temp)
		}
	}
	if len(tx.Effects.Unwrapped.([]interface{})) > 0 {
		for _, j := range tx.Effects.Created.([]interface{}) {
			var temp schema.Changed
			temp.ObjectId = j.([]interface{})[0].([]interface{})[0].(string)
			temp.Type = "Unwrapped"
			temp.Version = int(j.([]interface{})[0].([]interface{})[1].(float64))
			changedSet = append(changedSet, temp)
		}
	}
	if len(tx.Effects.Gas_object.([]interface{})) > 0 {
		for _, j := range tx.Effects.Created.([]interface{}) {
			var temp schema.Changed
			temp.ObjectId = j.([]interface{})[0].([]interface{})[0].(string)
			temp.Type = "Gas"
			temp.Version = int(j.([]interface{})[0].([]interface{})[1].(float64))
			changedSet = append(changedSet, temp)
		}
	}

	// Get recipient
	var recipient *string
	rec := helpers.RecurseKey(tx.Certificate.Kind.Single, "recipient")
	if rec != nil {
		if reflect.TypeOf(rec) == reflect.TypeOf(map[string]interface{}{}) {
			if v, ok := rec.(map[string]interface{})["ObjectOwner"]; ok {
				ptr := v.(string)
				recipient = &ptr
			} else if v, ok := rec.(map[string]interface{})["AddressOwner"]; ok {
				ptr := v.(string)
				recipient = &ptr
			}
		} else {
			ptr := rec.(string)
			recipient = &ptr
		}
	}

	// Get recipient
	var amount *float64
	amt := helpers.RecurseKey(tx.Certificate.Kind.Single, "amount")
	if amt != nil {
		ptr := amt.(float64)
		amount = &ptr
	}

	// Get pkg
	var _pkg *string
	pkg := helpers.RecurseKey(tx.Certificate.Kind.Single, "package")
	if pkg != nil {
		temp := pkg.([]interface{})[0].(string)
		_pkg = &temp
	}

	// Get module
	var _mod *string
	mod := helpers.RecurseKey(tx.Certificate.Kind.Single, "module")
	if mod != nil {
		temp := mod.(string)
		_mod = &temp
	}

	// Get function
	var _function *string
	function := helpers.RecurseKey(tx.Certificate.Kind.Single, "function")
	if function != nil {
		temp := function.(string)
		_function = &temp
	}

	txc, err := c.client.Transaction.Create().
		SetType(_type).
		SetTime(time.Unix(int64(tx.Time/1000), 0)).
		SetTransactionID(tx.Effects.Transaction_digest).
		SetStatus(_status).
		SetSender(tx.Certificate.Sender).
		SetNillableRecipient(recipient).
		SetNillableAmount(amount).
		SetNillablePackage(_pkg).
		SetNillableModule(_mod).
		SetNillableFunction(_function).
		SetChanged(changedSet).
		SetGas(uint32(tx.Effects.Gas_used.Computation_cost + tx.Effects.Gas_used.Storage_cost)).
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

func (c *EntClient) CreateAccount(acc sui.Acc, sequence int64) (*ent.Account, error) {

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

	log.Println(accc)

	return accc, err
}

func (c *EntClient) CreateDeletedAccount(id string, sequence int64) (*ent.Account, error) {

	// Type conversion from AccObj struct to ent version
	accc, err := c.client.Account.Create().
		SetAccountID(id).
		SetSequenceID(sequence).
		Save(context.Background())

	return accc, err
}

func (c *EntClient) UpsertAccount(acc sui.Acc, sequence int64) error {

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

func (c *EntClient) UpdateAccount(acc string, sequence int64, transaction string) (*ent.Account, error) {

	orig, err := c.client.Account.Query().Where(account.AccountIDEQ(acc)).Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed fetching account to be updated: %w", err)
	}

	accc, err := orig.Update().
		SetTransactions(append(orig.Transactions, transaction)).
		SetSequenceID(sequence).
		Save(context.Background())

	if err != nil {
		helpers.Check(err)
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

func (c *EntClient) CreateNFT(obj sui.AccObject, sequence int64) (*ent.NFT, error) {
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

func (c *EntClient) CreateObject(obj sui.Obj, transactionID string, version int) (*ent.Object, error) {
	objc, err := c.client.Object.Create().
		SetDataType(obj.GetObjectDataType()).
		SetFields(obj.GetObjectMetadata()).
		SetHasPublicTransfer(obj.HasPublicTransfer()).
		SetObjectID(obj.GetObjectID()).
		SetOwner(obj.GetOwner()).
		SetStatus(obj.GetObjectStatus()).
		SetType(obj.GetObjectType()).
		SetTransactionID(transactionID).
		SetVersion(version).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating object: %w", err)
	}
	return objc, nil
}

func (c *EntClient) CreateObjectFromSocket(obj sui.SocketObj) (*ent.Object, error) {

	delete(obj.Fields.SuiMoveStruct, "id")

	objc, err := c.client.Object.Create().
		SetFields(obj.Fields.SuiMoveStruct).
		SetObjectID(obj.Id).
		SetOwner(obj.Owner).
		SetStatus(obj.Status).
		SetType(obj.Type_).
		SetTransactionID(obj.Transaction_id).
		SetVersion(obj.Version).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating object: %w", err)
	}
	return objc, nil
}

func (c *EntClient) CreateDeletedObject(objectID string, transactionID string, version int) (*ent.Object, error) {
	objc, err := c.client.Object.Create().SetFields(map[string]interface{}{}).SetVersion(version).SetStatus("Not Exists").SetObjectID(objectID).SetTransactionID(transactionID).Save(context.Background())
	if err != nil {
		return nil, err
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
