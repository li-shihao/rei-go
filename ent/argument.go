// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"rei.io/rei/ent/argument"
)

// Argument is the model entity for the Argument schema.
type Argument struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Type holds the value of the "Type" field.
	Type string `json:"Type,omitempty"`
	// TransactionID holds the value of the "TransactionID" field.
	TransactionID string `json:"TransactionID,omitempty"`
	// Data holds the value of the "Data" field.
	Data string `json:"Data,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Argument) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case argument.FieldID:
			values[i] = new(sql.NullInt64)
		case argument.FieldName, argument.FieldType, argument.FieldTransactionID, argument.FieldData:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Argument", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Argument fields.
func (a *Argument) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case argument.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case argument.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case argument.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Type", values[i])
			} else if value.Valid {
				a.Type = value.String
			}
		case argument.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TransactionID", values[i])
			} else if value.Valid {
				a.TransactionID = value.String
			}
		case argument.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Data", values[i])
			} else if value.Valid {
				a.Data = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Argument.
// Note that you need to call Argument.Unwrap() before calling this method if this Argument
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Argument) Update() *ArgumentUpdateOne {
	return (&ArgumentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Argument entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Argument) Unwrap() *Argument {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Argument is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Argument) String() string {
	var builder strings.Builder
	builder.WriteString("Argument(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("Name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("Type=")
	builder.WriteString(a.Type)
	builder.WriteString(", ")
	builder.WriteString("TransactionID=")
	builder.WriteString(a.TransactionID)
	builder.WriteString(", ")
	builder.WriteString("Data=")
	builder.WriteString(a.Data)
	builder.WriteByte(')')
	return builder.String()
}

// Arguments is a parsable slice of Argument.
type Arguments []*Argument

func (a Arguments) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}