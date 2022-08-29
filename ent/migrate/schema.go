// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "account_id", Type: field.TypeString},
		{Name: "balance", Type: field.TypeUint64},
		{Name: "objects", Type: field.TypeJSON},
		{Name: "transactions", Type: field.TypeJSON},
		{Name: "time", Type: field.TypeTime},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// ArgumentsColumns holds the columns for the "arguments" table.
	ArgumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "transaction_id", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
	}
	// ArgumentsTable holds the schema information for the "arguments" table.
	ArgumentsTable = &schema.Table{
		Name:       "arguments",
		Columns:    ArgumentsColumns,
		PrimaryKey: []*schema.Column{ArgumentsColumns[0]},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "sender", Type: field.TypeString},
		{Name: "recipient", Type: field.TypeString, Nullable: true},
		{Name: "transaction_id", Type: field.TypeString},
		{Name: "object_id", Type: field.TypeString},
		{Name: "version", Type: field.TypeUint32},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// NfTsColumns holds the columns for the "nf_ts" table.
	NfTsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "object_id", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "time", Type: field.TypeTime},
	}
	// NfTsTable holds the schema information for the "nf_ts" table.
	NfTsTable = &schema.Table{
		Name:       "nf_ts",
		Columns:    NfTsColumns,
		PrimaryKey: []*schema.Column{NfTsColumns[0]},
	}
	// ObjectsColumns holds the columns for the "objects" table.
	ObjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString},
		{Name: "data_type", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "has_public_transfer", Type: field.TypeBool},
		{Name: "fields", Type: field.TypeJSON},
		{Name: "owner", Type: field.TypeString},
		{Name: "object_id", Type: field.TypeString},
	}
	// ObjectsTable holds the schema information for the "objects" table.
	ObjectsTable = &schema.Table{
		Name:       "objects",
		Columns:    ObjectsColumns,
		PrimaryKey: []*schema.Column{ObjectsColumns[0]},
	}
	// PackagesColumns holds the columns for the "packages" table.
	PackagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "transaction_id", Type: field.TypeString},
		{Name: "object_id", Type: field.TypeString},
		{Name: "bytecode", Type: field.TypeJSON},
	}
	// PackagesTable holds the schema information for the "packages" table.
	PackagesTable = &schema.Table{
		Name:       "packages",
		Columns:    PackagesColumns,
		PrimaryKey: []*schema.Column{PackagesColumns[0]},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "time", Type: field.TypeTime},
		{Name: "transaction_id", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeBool},
		{Name: "sender", Type: field.TypeString},
		{Name: "recipient", Type: field.TypeString, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "package", Type: field.TypeString, Nullable: true},
		{Name: "module", Type: field.TypeString, Nullable: true},
		{Name: "function", Type: field.TypeString, Nullable: true},
		{Name: "gas", Type: field.TypeUint32},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		ArgumentsTable,
		EventsTable,
		NfTsTable,
		ObjectsTable,
		PackagesTable,
		TransactionsTable,
	}
)

func init() {
}
