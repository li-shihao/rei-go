// Code generated by ent, DO NOT EDIT.

package transactions

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"rei.io/rei/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Type applies equality check predicate on the "Type" field. It's identical to TypeEQ.
func Type(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Time applies equality check predicate on the "Time" field. It's identical to TimeEQ.
func Time(v time.Time) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTime), v))
	})
}

// TransactionID applies equality check predicate on the "TransactionID" field. It's identical to TransactionIDEQ.
func TransactionID(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// Status applies equality check predicate on the "Status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Sender applies equality check predicate on the "Sender" field. It's identical to SenderEQ.
func Sender(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSender), v))
	})
}

// Recipient applies equality check predicate on the "Recipient" field. It's identical to RecipientEQ.
func Recipient(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// Amount applies equality check predicate on the "Amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Package applies equality check predicate on the "Package" field. It's identical to PackageEQ.
func Package(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPackage), v))
	})
}

// Module applies equality check predicate on the "Module" field. It's identical to ModuleEQ.
func Module(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModule), v))
	})
}

// Function applies equality check predicate on the "Function" field. It's identical to FunctionEQ.
func Function(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFunction), v))
	})
}

// TypeEQ applies the EQ predicate on the "Type" field.
func TypeEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "Type" field.
func TypeNEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "Type" field.
func TypeIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "Type" field.
func TypeNotIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "Type" field.
func TypeGT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "Type" field.
func TypeGTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "Type" field.
func TypeLT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "Type" field.
func TypeLTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "Type" field.
func TypeContains(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "Type" field.
func TypeHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "Type" field.
func TypeHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "Type" field.
func TypeEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "Type" field.
func TypeContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// TimeEQ applies the EQ predicate on the "Time" field.
func TimeEQ(v time.Time) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTime), v))
	})
}

// TimeNEQ applies the NEQ predicate on the "Time" field.
func TimeNEQ(v time.Time) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTime), v))
	})
}

// TimeIn applies the In predicate on the "Time" field.
func TimeIn(vs ...time.Time) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTime), v...))
	})
}

// TimeNotIn applies the NotIn predicate on the "Time" field.
func TimeNotIn(vs ...time.Time) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTime), v...))
	})
}

// TimeGT applies the GT predicate on the "Time" field.
func TimeGT(v time.Time) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTime), v))
	})
}

// TimeGTE applies the GTE predicate on the "Time" field.
func TimeGTE(v time.Time) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTime), v))
	})
}

// TimeLT applies the LT predicate on the "Time" field.
func TimeLT(v time.Time) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTime), v))
	})
}

// TimeLTE applies the LTE predicate on the "Time" field.
func TimeLTE(v time.Time) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTime), v))
	})
}

// TransactionIDEQ applies the EQ predicate on the "TransactionID" field.
func TransactionIDEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDNEQ applies the NEQ predicate on the "TransactionID" field.
func TransactionIDNEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDIn applies the In predicate on the "TransactionID" field.
func TransactionIDIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDNotIn applies the NotIn predicate on the "TransactionID" field.
func TransactionIDNotIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDGT applies the GT predicate on the "TransactionID" field.
func TransactionIDGT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDGTE applies the GTE predicate on the "TransactionID" field.
func TransactionIDGTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLT applies the LT predicate on the "TransactionID" field.
func TransactionIDLT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLTE applies the LTE predicate on the "TransactionID" field.
func TransactionIDLTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContains applies the Contains predicate on the "TransactionID" field.
func TransactionIDContains(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasPrefix applies the HasPrefix predicate on the "TransactionID" field.
func TransactionIDHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasSuffix applies the HasSuffix predicate on the "TransactionID" field.
func TransactionIDHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDEqualFold applies the EqualFold predicate on the "TransactionID" field.
func TransactionIDEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContainsFold applies the ContainsFold predicate on the "TransactionID" field.
func TransactionIDContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionID), v))
	})
}

// StatusEQ applies the EQ predicate on the "Status" field.
func StatusEQ(v bool) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "Status" field.
func StatusNEQ(v bool) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// SenderEQ applies the EQ predicate on the "Sender" field.
func SenderEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSender), v))
	})
}

// SenderNEQ applies the NEQ predicate on the "Sender" field.
func SenderNEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSender), v))
	})
}

// SenderIn applies the In predicate on the "Sender" field.
func SenderIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSender), v...))
	})
}

// SenderNotIn applies the NotIn predicate on the "Sender" field.
func SenderNotIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSender), v...))
	})
}

// SenderGT applies the GT predicate on the "Sender" field.
func SenderGT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSender), v))
	})
}

// SenderGTE applies the GTE predicate on the "Sender" field.
func SenderGTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSender), v))
	})
}

// SenderLT applies the LT predicate on the "Sender" field.
func SenderLT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSender), v))
	})
}

// SenderLTE applies the LTE predicate on the "Sender" field.
func SenderLTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSender), v))
	})
}

// SenderContains applies the Contains predicate on the "Sender" field.
func SenderContains(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSender), v))
	})
}

// SenderHasPrefix applies the HasPrefix predicate on the "Sender" field.
func SenderHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSender), v))
	})
}

// SenderHasSuffix applies the HasSuffix predicate on the "Sender" field.
func SenderHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSender), v))
	})
}

// SenderEqualFold applies the EqualFold predicate on the "Sender" field.
func SenderEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSender), v))
	})
}

// SenderContainsFold applies the ContainsFold predicate on the "Sender" field.
func SenderContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSender), v))
	})
}

// RecipientEQ applies the EQ predicate on the "Recipient" field.
func RecipientEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// RecipientNEQ applies the NEQ predicate on the "Recipient" field.
func RecipientNEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecipient), v))
	})
}

// RecipientIn applies the In predicate on the "Recipient" field.
func RecipientIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRecipient), v...))
	})
}

// RecipientNotIn applies the NotIn predicate on the "Recipient" field.
func RecipientNotIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRecipient), v...))
	})
}

// RecipientGT applies the GT predicate on the "Recipient" field.
func RecipientGT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecipient), v))
	})
}

// RecipientGTE applies the GTE predicate on the "Recipient" field.
func RecipientGTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecipient), v))
	})
}

// RecipientLT applies the LT predicate on the "Recipient" field.
func RecipientLT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecipient), v))
	})
}

// RecipientLTE applies the LTE predicate on the "Recipient" field.
func RecipientLTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecipient), v))
	})
}

// RecipientContains applies the Contains predicate on the "Recipient" field.
func RecipientContains(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecipient), v))
	})
}

// RecipientHasPrefix applies the HasPrefix predicate on the "Recipient" field.
func RecipientHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecipient), v))
	})
}

// RecipientHasSuffix applies the HasSuffix predicate on the "Recipient" field.
func RecipientHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecipient), v))
	})
}

// RecipientIsNil applies the IsNil predicate on the "Recipient" field.
func RecipientIsNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecipient)))
	})
}

// RecipientNotNil applies the NotNil predicate on the "Recipient" field.
func RecipientNotNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecipient)))
	})
}

// RecipientEqualFold applies the EqualFold predicate on the "Recipient" field.
func RecipientEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecipient), v))
	})
}

// RecipientContainsFold applies the ContainsFold predicate on the "Recipient" field.
func RecipientContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecipient), v))
	})
}

// AmountEQ applies the EQ predicate on the "Amount" field.
func AmountEQ(v float64) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "Amount" field.
func AmountNEQ(v float64) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "Amount" field.
func AmountIn(vs ...float64) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "Amount" field.
func AmountNotIn(vs ...float64) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "Amount" field.
func AmountGT(v float64) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "Amount" field.
func AmountGTE(v float64) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "Amount" field.
func AmountLT(v float64) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "Amount" field.
func AmountLTE(v float64) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountIsNil applies the IsNil predicate on the "Amount" field.
func AmountIsNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmount)))
	})
}

// AmountNotNil applies the NotNil predicate on the "Amount" field.
func AmountNotNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmount)))
	})
}

// PackageEQ applies the EQ predicate on the "Package" field.
func PackageEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPackage), v))
	})
}

// PackageNEQ applies the NEQ predicate on the "Package" field.
func PackageNEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPackage), v))
	})
}

// PackageIn applies the In predicate on the "Package" field.
func PackageIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPackage), v...))
	})
}

// PackageNotIn applies the NotIn predicate on the "Package" field.
func PackageNotIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPackage), v...))
	})
}

// PackageGT applies the GT predicate on the "Package" field.
func PackageGT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPackage), v))
	})
}

// PackageGTE applies the GTE predicate on the "Package" field.
func PackageGTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPackage), v))
	})
}

// PackageLT applies the LT predicate on the "Package" field.
func PackageLT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPackage), v))
	})
}

// PackageLTE applies the LTE predicate on the "Package" field.
func PackageLTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPackage), v))
	})
}

// PackageContains applies the Contains predicate on the "Package" field.
func PackageContains(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPackage), v))
	})
}

// PackageHasPrefix applies the HasPrefix predicate on the "Package" field.
func PackageHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPackage), v))
	})
}

// PackageHasSuffix applies the HasSuffix predicate on the "Package" field.
func PackageHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPackage), v))
	})
}

// PackageIsNil applies the IsNil predicate on the "Package" field.
func PackageIsNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPackage)))
	})
}

// PackageNotNil applies the NotNil predicate on the "Package" field.
func PackageNotNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPackage)))
	})
}

// PackageEqualFold applies the EqualFold predicate on the "Package" field.
func PackageEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPackage), v))
	})
}

// PackageContainsFold applies the ContainsFold predicate on the "Package" field.
func PackageContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPackage), v))
	})
}

// ModuleEQ applies the EQ predicate on the "Module" field.
func ModuleEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModule), v))
	})
}

// ModuleNEQ applies the NEQ predicate on the "Module" field.
func ModuleNEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModule), v))
	})
}

// ModuleIn applies the In predicate on the "Module" field.
func ModuleIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldModule), v...))
	})
}

// ModuleNotIn applies the NotIn predicate on the "Module" field.
func ModuleNotIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldModule), v...))
	})
}

// ModuleGT applies the GT predicate on the "Module" field.
func ModuleGT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModule), v))
	})
}

// ModuleGTE applies the GTE predicate on the "Module" field.
func ModuleGTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModule), v))
	})
}

// ModuleLT applies the LT predicate on the "Module" field.
func ModuleLT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModule), v))
	})
}

// ModuleLTE applies the LTE predicate on the "Module" field.
func ModuleLTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModule), v))
	})
}

// ModuleContains applies the Contains predicate on the "Module" field.
func ModuleContains(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldModule), v))
	})
}

// ModuleHasPrefix applies the HasPrefix predicate on the "Module" field.
func ModuleHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldModule), v))
	})
}

// ModuleHasSuffix applies the HasSuffix predicate on the "Module" field.
func ModuleHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldModule), v))
	})
}

// ModuleIsNil applies the IsNil predicate on the "Module" field.
func ModuleIsNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldModule)))
	})
}

// ModuleNotNil applies the NotNil predicate on the "Module" field.
func ModuleNotNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldModule)))
	})
}

// ModuleEqualFold applies the EqualFold predicate on the "Module" field.
func ModuleEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldModule), v))
	})
}

// ModuleContainsFold applies the ContainsFold predicate on the "Module" field.
func ModuleContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldModule), v))
	})
}

// FunctionEQ applies the EQ predicate on the "Function" field.
func FunctionEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFunction), v))
	})
}

// FunctionNEQ applies the NEQ predicate on the "Function" field.
func FunctionNEQ(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFunction), v))
	})
}

// FunctionIn applies the In predicate on the "Function" field.
func FunctionIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFunction), v...))
	})
}

// FunctionNotIn applies the NotIn predicate on the "Function" field.
func FunctionNotIn(vs ...string) predicate.Transactions {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFunction), v...))
	})
}

// FunctionGT applies the GT predicate on the "Function" field.
func FunctionGT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFunction), v))
	})
}

// FunctionGTE applies the GTE predicate on the "Function" field.
func FunctionGTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFunction), v))
	})
}

// FunctionLT applies the LT predicate on the "Function" field.
func FunctionLT(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFunction), v))
	})
}

// FunctionLTE applies the LTE predicate on the "Function" field.
func FunctionLTE(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFunction), v))
	})
}

// FunctionContains applies the Contains predicate on the "Function" field.
func FunctionContains(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFunction), v))
	})
}

// FunctionHasPrefix applies the HasPrefix predicate on the "Function" field.
func FunctionHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFunction), v))
	})
}

// FunctionHasSuffix applies the HasSuffix predicate on the "Function" field.
func FunctionHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFunction), v))
	})
}

// FunctionIsNil applies the IsNil predicate on the "Function" field.
func FunctionIsNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFunction)))
	})
}

// FunctionNotNil applies the NotNil predicate on the "Function" field.
func FunctionNotNil() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFunction)))
	})
}

// FunctionEqualFold applies the EqualFold predicate on the "Function" field.
func FunctionEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFunction), v))
	})
}

// FunctionContainsFold applies the ContainsFold predicate on the "Function" field.
func FunctionContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFunction), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transactions) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transactions) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transactions) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		p(s.Not())
	})
}
