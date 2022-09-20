// Code generated by ent, DO NOT EDIT.

package object

const (
	// Label holds the string label denoting the object type in the database.
	Label = "object"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDataType holds the string denoting the datatype field in the database.
	FieldDataType = "data_type"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldHasPublicTransfer holds the string denoting the has_public_transfer field in the database.
	FieldHasPublicTransfer = "has_public_transfer"
	// FieldFields holds the string denoting the fields field in the database.
	FieldFields = "fields"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldObjectID holds the string denoting the objectid field in the database.
	FieldObjectID = "object_id"
	// FieldTransactionID holds the string denoting the transactionid field in the database.
	FieldTransactionID = "transaction_id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// Table holds the table name of the object in the database.
	Table = "objects"
)

// Columns holds all SQL columns for object fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldDataType,
	FieldType,
	FieldHasPublicTransfer,
	FieldFields,
	FieldOwner,
	FieldObjectID,
	FieldTransactionID,
	FieldVersion,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
