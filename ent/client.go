// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"rei.io/rei/ent/migrate"

	"rei.io/rei/ent/accounts"
	"rei.io/rei/ent/arguments"
	"rei.io/rei/ent/events"
	"rei.io/rei/ent/nfts"
	"rei.io/rei/ent/objects"
	"rei.io/rei/ent/packages"
	"rei.io/rei/ent/transactions"
	"rei.io/rei/ent/users"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Accounts is the client for interacting with the Accounts builders.
	Accounts *AccountsClient
	// Arguments is the client for interacting with the Arguments builders.
	Arguments *ArgumentsClient
	// Events is the client for interacting with the Events builders.
	Events *EventsClient
	// NFTs is the client for interacting with the NFTs builders.
	NFTs *NFTsClient
	// Objects is the client for interacting with the Objects builders.
	Objects *ObjectsClient
	// Packages is the client for interacting with the Packages builders.
	Packages *PackagesClient
	// Transactions is the client for interacting with the Transactions builders.
	Transactions *TransactionsClient
	// Users is the client for interacting with the Users builders.
	Users *UsersClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Accounts = NewAccountsClient(c.config)
	c.Arguments = NewArgumentsClient(c.config)
	c.Events = NewEventsClient(c.config)
	c.NFTs = NewNFTsClient(c.config)
	c.Objects = NewObjectsClient(c.config)
	c.Packages = NewPackagesClient(c.config)
	c.Transactions = NewTransactionsClient(c.config)
	c.Users = NewUsersClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Accounts:     NewAccountsClient(cfg),
		Arguments:    NewArgumentsClient(cfg),
		Events:       NewEventsClient(cfg),
		NFTs:         NewNFTsClient(cfg),
		Objects:      NewObjectsClient(cfg),
		Packages:     NewPackagesClient(cfg),
		Transactions: NewTransactionsClient(cfg),
		Users:        NewUsersClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Accounts:     NewAccountsClient(cfg),
		Arguments:    NewArgumentsClient(cfg),
		Events:       NewEventsClient(cfg),
		NFTs:         NewNFTsClient(cfg),
		Objects:      NewObjectsClient(cfg),
		Packages:     NewPackagesClient(cfg),
		Transactions: NewTransactionsClient(cfg),
		Users:        NewUsersClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Accounts.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Accounts.Use(hooks...)
	c.Arguments.Use(hooks...)
	c.Events.Use(hooks...)
	c.NFTs.Use(hooks...)
	c.Objects.Use(hooks...)
	c.Packages.Use(hooks...)
	c.Transactions.Use(hooks...)
	c.Users.Use(hooks...)
}

// AccountsClient is a client for the Accounts schema.
type AccountsClient struct {
	config
}

// NewAccountsClient returns a client for the Accounts from the given config.
func NewAccountsClient(c config) *AccountsClient {
	return &AccountsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accounts.Hooks(f(g(h())))`.
func (c *AccountsClient) Use(hooks ...Hook) {
	c.hooks.Accounts = append(c.hooks.Accounts, hooks...)
}

// Create returns a builder for creating a Accounts entity.
func (c *AccountsClient) Create() *AccountsCreate {
	mutation := newAccountsMutation(c.config, OpCreate)
	return &AccountsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Accounts entities.
func (c *AccountsClient) CreateBulk(builders ...*AccountsCreate) *AccountsCreateBulk {
	return &AccountsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Accounts.
func (c *AccountsClient) Update() *AccountsUpdate {
	mutation := newAccountsMutation(c.config, OpUpdate)
	return &AccountsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountsClient) UpdateOne(a *Accounts) *AccountsUpdateOne {
	mutation := newAccountsMutation(c.config, OpUpdateOne, withAccounts(a))
	return &AccountsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountsClient) UpdateOneID(id int) *AccountsUpdateOne {
	mutation := newAccountsMutation(c.config, OpUpdateOne, withAccountsID(id))
	return &AccountsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Accounts.
func (c *AccountsClient) Delete() *AccountsDelete {
	mutation := newAccountsMutation(c.config, OpDelete)
	return &AccountsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountsClient) DeleteOne(a *Accounts) *AccountsDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AccountsClient) DeleteOneID(id int) *AccountsDeleteOne {
	builder := c.Delete().Where(accounts.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountsDeleteOne{builder}
}

// Query returns a query builder for Accounts.
func (c *AccountsClient) Query() *AccountsQuery {
	return &AccountsQuery{
		config: c.config,
	}
}

// Get returns a Accounts entity by its id.
func (c *AccountsClient) Get(ctx context.Context, id int) (*Accounts, error) {
	return c.Query().Where(accounts.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountsClient) GetX(ctx context.Context, id int) *Accounts {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountsClient) Hooks() []Hook {
	return c.hooks.Accounts
}

// ArgumentsClient is a client for the Arguments schema.
type ArgumentsClient struct {
	config
}

// NewArgumentsClient returns a client for the Arguments from the given config.
func NewArgumentsClient(c config) *ArgumentsClient {
	return &ArgumentsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `arguments.Hooks(f(g(h())))`.
func (c *ArgumentsClient) Use(hooks ...Hook) {
	c.hooks.Arguments = append(c.hooks.Arguments, hooks...)
}

// Create returns a builder for creating a Arguments entity.
func (c *ArgumentsClient) Create() *ArgumentsCreate {
	mutation := newArgumentsMutation(c.config, OpCreate)
	return &ArgumentsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Arguments entities.
func (c *ArgumentsClient) CreateBulk(builders ...*ArgumentsCreate) *ArgumentsCreateBulk {
	return &ArgumentsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Arguments.
func (c *ArgumentsClient) Update() *ArgumentsUpdate {
	mutation := newArgumentsMutation(c.config, OpUpdate)
	return &ArgumentsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArgumentsClient) UpdateOne(a *Arguments) *ArgumentsUpdateOne {
	mutation := newArgumentsMutation(c.config, OpUpdateOne, withArguments(a))
	return &ArgumentsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArgumentsClient) UpdateOneID(id int) *ArgumentsUpdateOne {
	mutation := newArgumentsMutation(c.config, OpUpdateOne, withArgumentsID(id))
	return &ArgumentsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Arguments.
func (c *ArgumentsClient) Delete() *ArgumentsDelete {
	mutation := newArgumentsMutation(c.config, OpDelete)
	return &ArgumentsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArgumentsClient) DeleteOne(a *Arguments) *ArgumentsDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ArgumentsClient) DeleteOneID(id int) *ArgumentsDeleteOne {
	builder := c.Delete().Where(arguments.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArgumentsDeleteOne{builder}
}

// Query returns a query builder for Arguments.
func (c *ArgumentsClient) Query() *ArgumentsQuery {
	return &ArgumentsQuery{
		config: c.config,
	}
}

// Get returns a Arguments entity by its id.
func (c *ArgumentsClient) Get(ctx context.Context, id int) (*Arguments, error) {
	return c.Query().Where(arguments.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArgumentsClient) GetX(ctx context.Context, id int) *Arguments {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ArgumentsClient) Hooks() []Hook {
	return c.hooks.Arguments
}

// EventsClient is a client for the Events schema.
type EventsClient struct {
	config
}

// NewEventsClient returns a client for the Events from the given config.
func NewEventsClient(c config) *EventsClient {
	return &EventsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `events.Hooks(f(g(h())))`.
func (c *EventsClient) Use(hooks ...Hook) {
	c.hooks.Events = append(c.hooks.Events, hooks...)
}

// Create returns a builder for creating a Events entity.
func (c *EventsClient) Create() *EventsCreate {
	mutation := newEventsMutation(c.config, OpCreate)
	return &EventsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Events entities.
func (c *EventsClient) CreateBulk(builders ...*EventsCreate) *EventsCreateBulk {
	return &EventsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Events.
func (c *EventsClient) Update() *EventsUpdate {
	mutation := newEventsMutation(c.config, OpUpdate)
	return &EventsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EventsClient) UpdateOne(e *Events) *EventsUpdateOne {
	mutation := newEventsMutation(c.config, OpUpdateOne, withEvents(e))
	return &EventsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EventsClient) UpdateOneID(id int) *EventsUpdateOne {
	mutation := newEventsMutation(c.config, OpUpdateOne, withEventsID(id))
	return &EventsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Events.
func (c *EventsClient) Delete() *EventsDelete {
	mutation := newEventsMutation(c.config, OpDelete)
	return &EventsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EventsClient) DeleteOne(e *Events) *EventsDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EventsClient) DeleteOneID(id int) *EventsDeleteOne {
	builder := c.Delete().Where(events.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EventsDeleteOne{builder}
}

// Query returns a query builder for Events.
func (c *EventsClient) Query() *EventsQuery {
	return &EventsQuery{
		config: c.config,
	}
}

// Get returns a Events entity by its id.
func (c *EventsClient) Get(ctx context.Context, id int) (*Events, error) {
	return c.Query().Where(events.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EventsClient) GetX(ctx context.Context, id int) *Events {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EventsClient) Hooks() []Hook {
	return c.hooks.Events
}

// NFTsClient is a client for the NFTs schema.
type NFTsClient struct {
	config
}

// NewNFTsClient returns a client for the NFTs from the given config.
func NewNFTsClient(c config) *NFTsClient {
	return &NFTsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `nfts.Hooks(f(g(h())))`.
func (c *NFTsClient) Use(hooks ...Hook) {
	c.hooks.NFTs = append(c.hooks.NFTs, hooks...)
}

// Create returns a builder for creating a NFTs entity.
func (c *NFTsClient) Create() *NFTsCreate {
	mutation := newNFTsMutation(c.config, OpCreate)
	return &NFTsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NFTs entities.
func (c *NFTsClient) CreateBulk(builders ...*NFTsCreate) *NFTsCreateBulk {
	return &NFTsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NFTs.
func (c *NFTsClient) Update() *NFTsUpdate {
	mutation := newNFTsMutation(c.config, OpUpdate)
	return &NFTsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NFTsClient) UpdateOne(nt *NFTs) *NFTsUpdateOne {
	mutation := newNFTsMutation(c.config, OpUpdateOne, withNFTs(nt))
	return &NFTsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NFTsClient) UpdateOneID(id int) *NFTsUpdateOne {
	mutation := newNFTsMutation(c.config, OpUpdateOne, withNFTsID(id))
	return &NFTsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NFTs.
func (c *NFTsClient) Delete() *NFTsDelete {
	mutation := newNFTsMutation(c.config, OpDelete)
	return &NFTsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NFTsClient) DeleteOne(nt *NFTs) *NFTsDeleteOne {
	return c.DeleteOneID(nt.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *NFTsClient) DeleteOneID(id int) *NFTsDeleteOne {
	builder := c.Delete().Where(nfts.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NFTsDeleteOne{builder}
}

// Query returns a query builder for NFTs.
func (c *NFTsClient) Query() *NFTsQuery {
	return &NFTsQuery{
		config: c.config,
	}
}

// Get returns a NFTs entity by its id.
func (c *NFTsClient) Get(ctx context.Context, id int) (*NFTs, error) {
	return c.Query().Where(nfts.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NFTsClient) GetX(ctx context.Context, id int) *NFTs {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NFTsClient) Hooks() []Hook {
	return c.hooks.NFTs
}

// ObjectsClient is a client for the Objects schema.
type ObjectsClient struct {
	config
}

// NewObjectsClient returns a client for the Objects from the given config.
func NewObjectsClient(c config) *ObjectsClient {
	return &ObjectsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `objects.Hooks(f(g(h())))`.
func (c *ObjectsClient) Use(hooks ...Hook) {
	c.hooks.Objects = append(c.hooks.Objects, hooks...)
}

// Create returns a builder for creating a Objects entity.
func (c *ObjectsClient) Create() *ObjectsCreate {
	mutation := newObjectsMutation(c.config, OpCreate)
	return &ObjectsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Objects entities.
func (c *ObjectsClient) CreateBulk(builders ...*ObjectsCreate) *ObjectsCreateBulk {
	return &ObjectsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Objects.
func (c *ObjectsClient) Update() *ObjectsUpdate {
	mutation := newObjectsMutation(c.config, OpUpdate)
	return &ObjectsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ObjectsClient) UpdateOne(o *Objects) *ObjectsUpdateOne {
	mutation := newObjectsMutation(c.config, OpUpdateOne, withObjects(o))
	return &ObjectsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ObjectsClient) UpdateOneID(id int) *ObjectsUpdateOne {
	mutation := newObjectsMutation(c.config, OpUpdateOne, withObjectsID(id))
	return &ObjectsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Objects.
func (c *ObjectsClient) Delete() *ObjectsDelete {
	mutation := newObjectsMutation(c.config, OpDelete)
	return &ObjectsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ObjectsClient) DeleteOne(o *Objects) *ObjectsDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ObjectsClient) DeleteOneID(id int) *ObjectsDeleteOne {
	builder := c.Delete().Where(objects.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ObjectsDeleteOne{builder}
}

// Query returns a query builder for Objects.
func (c *ObjectsClient) Query() *ObjectsQuery {
	return &ObjectsQuery{
		config: c.config,
	}
}

// Get returns a Objects entity by its id.
func (c *ObjectsClient) Get(ctx context.Context, id int) (*Objects, error) {
	return c.Query().Where(objects.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ObjectsClient) GetX(ctx context.Context, id int) *Objects {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ObjectsClient) Hooks() []Hook {
	return c.hooks.Objects
}

// PackagesClient is a client for the Packages schema.
type PackagesClient struct {
	config
}

// NewPackagesClient returns a client for the Packages from the given config.
func NewPackagesClient(c config) *PackagesClient {
	return &PackagesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `packages.Hooks(f(g(h())))`.
func (c *PackagesClient) Use(hooks ...Hook) {
	c.hooks.Packages = append(c.hooks.Packages, hooks...)
}

// Create returns a builder for creating a Packages entity.
func (c *PackagesClient) Create() *PackagesCreate {
	mutation := newPackagesMutation(c.config, OpCreate)
	return &PackagesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Packages entities.
func (c *PackagesClient) CreateBulk(builders ...*PackagesCreate) *PackagesCreateBulk {
	return &PackagesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Packages.
func (c *PackagesClient) Update() *PackagesUpdate {
	mutation := newPackagesMutation(c.config, OpUpdate)
	return &PackagesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PackagesClient) UpdateOne(pa *Packages) *PackagesUpdateOne {
	mutation := newPackagesMutation(c.config, OpUpdateOne, withPackages(pa))
	return &PackagesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PackagesClient) UpdateOneID(id int) *PackagesUpdateOne {
	mutation := newPackagesMutation(c.config, OpUpdateOne, withPackagesID(id))
	return &PackagesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Packages.
func (c *PackagesClient) Delete() *PackagesDelete {
	mutation := newPackagesMutation(c.config, OpDelete)
	return &PackagesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PackagesClient) DeleteOne(pa *Packages) *PackagesDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PackagesClient) DeleteOneID(id int) *PackagesDeleteOne {
	builder := c.Delete().Where(packages.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PackagesDeleteOne{builder}
}

// Query returns a query builder for Packages.
func (c *PackagesClient) Query() *PackagesQuery {
	return &PackagesQuery{
		config: c.config,
	}
}

// Get returns a Packages entity by its id.
func (c *PackagesClient) Get(ctx context.Context, id int) (*Packages, error) {
	return c.Query().Where(packages.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PackagesClient) GetX(ctx context.Context, id int) *Packages {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PackagesClient) Hooks() []Hook {
	return c.hooks.Packages
}

// TransactionsClient is a client for the Transactions schema.
type TransactionsClient struct {
	config
}

// NewTransactionsClient returns a client for the Transactions from the given config.
func NewTransactionsClient(c config) *TransactionsClient {
	return &TransactionsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transactions.Hooks(f(g(h())))`.
func (c *TransactionsClient) Use(hooks ...Hook) {
	c.hooks.Transactions = append(c.hooks.Transactions, hooks...)
}

// Create returns a builder for creating a Transactions entity.
func (c *TransactionsClient) Create() *TransactionsCreate {
	mutation := newTransactionsMutation(c.config, OpCreate)
	return &TransactionsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transactions entities.
func (c *TransactionsClient) CreateBulk(builders ...*TransactionsCreate) *TransactionsCreateBulk {
	return &TransactionsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transactions.
func (c *TransactionsClient) Update() *TransactionsUpdate {
	mutation := newTransactionsMutation(c.config, OpUpdate)
	return &TransactionsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionsClient) UpdateOne(t *Transactions) *TransactionsUpdateOne {
	mutation := newTransactionsMutation(c.config, OpUpdateOne, withTransactions(t))
	return &TransactionsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionsClient) UpdateOneID(id int) *TransactionsUpdateOne {
	mutation := newTransactionsMutation(c.config, OpUpdateOne, withTransactionsID(id))
	return &TransactionsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transactions.
func (c *TransactionsClient) Delete() *TransactionsDelete {
	mutation := newTransactionsMutation(c.config, OpDelete)
	return &TransactionsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransactionsClient) DeleteOne(t *Transactions) *TransactionsDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TransactionsClient) DeleteOneID(id int) *TransactionsDeleteOne {
	builder := c.Delete().Where(transactions.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionsDeleteOne{builder}
}

// Query returns a query builder for Transactions.
func (c *TransactionsClient) Query() *TransactionsQuery {
	return &TransactionsQuery{
		config: c.config,
	}
}

// Get returns a Transactions entity by its id.
func (c *TransactionsClient) Get(ctx context.Context, id int) (*Transactions, error) {
	return c.Query().Where(transactions.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionsClient) GetX(ctx context.Context, id int) *Transactions {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TransactionsClient) Hooks() []Hook {
	return c.hooks.Transactions
}

// UsersClient is a client for the Users schema.
type UsersClient struct {
	config
}

// NewUsersClient returns a client for the Users from the given config.
func NewUsersClient(c config) *UsersClient {
	return &UsersClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `users.Hooks(f(g(h())))`.
func (c *UsersClient) Use(hooks ...Hook) {
	c.hooks.Users = append(c.hooks.Users, hooks...)
}

// Create returns a builder for creating a Users entity.
func (c *UsersClient) Create() *UsersCreate {
	mutation := newUsersMutation(c.config, OpCreate)
	return &UsersCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Users entities.
func (c *UsersClient) CreateBulk(builders ...*UsersCreate) *UsersCreateBulk {
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Users.
func (c *UsersClient) Update() *UsersUpdate {
	mutation := newUsersMutation(c.config, OpUpdate)
	return &UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersClient) UpdateOne(u *Users) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsers(u))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersClient) UpdateOneID(id int) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsersID(id))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Users.
func (c *UsersClient) Delete() *UsersDelete {
	mutation := newUsersMutation(c.config, OpDelete)
	return &UsersDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersClient) DeleteOne(u *Users) *UsersDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UsersClient) DeleteOneID(id int) *UsersDeleteOne {
	builder := c.Delete().Where(users.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersDeleteOne{builder}
}

// Query returns a query builder for Users.
func (c *UsersClient) Query() *UsersQuery {
	return &UsersQuery{
		config: c.config,
	}
}

// Get returns a Users entity by its id.
func (c *UsersClient) Get(ctx context.Context, id int) (*Users, error) {
	return c.Query().Where(users.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersClient) GetX(ctx context.Context, id int) *Users {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UsersClient) Hooks() []Hook {
	return c.hooks.Users
}
