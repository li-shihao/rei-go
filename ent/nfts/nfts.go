// Code generated by ent, DO NOT EDIT.

package nfts

const (
	// Label holds the string label denoting the nfts type in the database.
	Label = "nf_ts"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldObjectID holds the string denoting the objectid field in the database.
	FieldObjectID = "object_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// Table holds the table name of the nfts in the database.
	Table = "nf_ts"
)

// Columns holds all SQL columns for nfts fields.
var Columns = []string{
	FieldID,
	FieldObjectID,
	FieldType,
	FieldMetadata,
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
