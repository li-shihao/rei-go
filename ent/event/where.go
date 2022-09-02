// Code generated by ent, DO NOT EDIT.

package event

import (
	"entgo.io/ent/dialect/sql"
	"rei.io/rei/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Type applies equality check predicate on the "Type" field. It's identical to TypeEQ.
func Type(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Sender applies equality check predicate on the "Sender" field. It's identical to SenderEQ.
func Sender(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSender), v))
	})
}

// Recipient applies equality check predicate on the "Recipient" field. It's identical to RecipientEQ.
func Recipient(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// TransactionID applies equality check predicate on the "TransactionID" field. It's identical to TransactionIDEQ.
func TransactionID(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// ObjectID applies equality check predicate on the "ObjectID" field. It's identical to ObjectIDEQ.
func ObjectID(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectID), v))
	})
}

// Version applies equality check predicate on the "Version" field. It's identical to VersionEQ.
func Version(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// TypeEQ applies the EQ predicate on the "Type" field.
func TypeEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "Type" field.
func TypeNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "Type" field.
func TypeIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "Type" field.
func TypeNotIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "Type" field.
func TypeGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "Type" field.
func TypeGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "Type" field.
func TypeLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "Type" field.
func TypeLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "Type" field.
func TypeContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "Type" field.
func TypeHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "Type" field.
func TypeHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "Type" field.
func TypeEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "Type" field.
func TypeContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// SenderEQ applies the EQ predicate on the "Sender" field.
func SenderEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSender), v))
	})
}

// SenderNEQ applies the NEQ predicate on the "Sender" field.
func SenderNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSender), v))
	})
}

// SenderIn applies the In predicate on the "Sender" field.
func SenderIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSender), v...))
	})
}

// SenderNotIn applies the NotIn predicate on the "Sender" field.
func SenderNotIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSender), v...))
	})
}

// SenderGT applies the GT predicate on the "Sender" field.
func SenderGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSender), v))
	})
}

// SenderGTE applies the GTE predicate on the "Sender" field.
func SenderGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSender), v))
	})
}

// SenderLT applies the LT predicate on the "Sender" field.
func SenderLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSender), v))
	})
}

// SenderLTE applies the LTE predicate on the "Sender" field.
func SenderLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSender), v))
	})
}

// SenderContains applies the Contains predicate on the "Sender" field.
func SenderContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSender), v))
	})
}

// SenderHasPrefix applies the HasPrefix predicate on the "Sender" field.
func SenderHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSender), v))
	})
}

// SenderHasSuffix applies the HasSuffix predicate on the "Sender" field.
func SenderHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSender), v))
	})
}

// SenderEqualFold applies the EqualFold predicate on the "Sender" field.
func SenderEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSender), v))
	})
}

// SenderContainsFold applies the ContainsFold predicate on the "Sender" field.
func SenderContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSender), v))
	})
}

// RecipientEQ applies the EQ predicate on the "Recipient" field.
func RecipientEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// RecipientNEQ applies the NEQ predicate on the "Recipient" field.
func RecipientNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecipient), v))
	})
}

// RecipientIn applies the In predicate on the "Recipient" field.
func RecipientIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRecipient), v...))
	})
}

// RecipientNotIn applies the NotIn predicate on the "Recipient" field.
func RecipientNotIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRecipient), v...))
	})
}

// RecipientGT applies the GT predicate on the "Recipient" field.
func RecipientGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecipient), v))
	})
}

// RecipientGTE applies the GTE predicate on the "Recipient" field.
func RecipientGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecipient), v))
	})
}

// RecipientLT applies the LT predicate on the "Recipient" field.
func RecipientLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecipient), v))
	})
}

// RecipientLTE applies the LTE predicate on the "Recipient" field.
func RecipientLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecipient), v))
	})
}

// RecipientContains applies the Contains predicate on the "Recipient" field.
func RecipientContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecipient), v))
	})
}

// RecipientHasPrefix applies the HasPrefix predicate on the "Recipient" field.
func RecipientHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecipient), v))
	})
}

// RecipientHasSuffix applies the HasSuffix predicate on the "Recipient" field.
func RecipientHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecipient), v))
	})
}

// RecipientIsNil applies the IsNil predicate on the "Recipient" field.
func RecipientIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecipient)))
	})
}

// RecipientNotNil applies the NotNil predicate on the "Recipient" field.
func RecipientNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecipient)))
	})
}

// RecipientEqualFold applies the EqualFold predicate on the "Recipient" field.
func RecipientEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecipient), v))
	})
}

// RecipientContainsFold applies the ContainsFold predicate on the "Recipient" field.
func RecipientContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecipient), v))
	})
}

// TransactionIDEQ applies the EQ predicate on the "TransactionID" field.
func TransactionIDEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDNEQ applies the NEQ predicate on the "TransactionID" field.
func TransactionIDNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDIn applies the In predicate on the "TransactionID" field.
func TransactionIDIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDNotIn applies the NotIn predicate on the "TransactionID" field.
func TransactionIDNotIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDGT applies the GT predicate on the "TransactionID" field.
func TransactionIDGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDGTE applies the GTE predicate on the "TransactionID" field.
func TransactionIDGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLT applies the LT predicate on the "TransactionID" field.
func TransactionIDLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLTE applies the LTE predicate on the "TransactionID" field.
func TransactionIDLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContains applies the Contains predicate on the "TransactionID" field.
func TransactionIDContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasPrefix applies the HasPrefix predicate on the "TransactionID" field.
func TransactionIDHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasSuffix applies the HasSuffix predicate on the "TransactionID" field.
func TransactionIDHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDEqualFold applies the EqualFold predicate on the "TransactionID" field.
func TransactionIDEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContainsFold applies the ContainsFold predicate on the "TransactionID" field.
func TransactionIDContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionID), v))
	})
}

// ObjectIDEQ applies the EQ predicate on the "ObjectID" field.
func ObjectIDEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectID), v))
	})
}

// ObjectIDNEQ applies the NEQ predicate on the "ObjectID" field.
func ObjectIDNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObjectID), v))
	})
}

// ObjectIDIn applies the In predicate on the "ObjectID" field.
func ObjectIDIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldObjectID), v...))
	})
}

// ObjectIDNotIn applies the NotIn predicate on the "ObjectID" field.
func ObjectIDNotIn(vs ...string) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldObjectID), v...))
	})
}

// ObjectIDGT applies the GT predicate on the "ObjectID" field.
func ObjectIDGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObjectID), v))
	})
}

// ObjectIDGTE applies the GTE predicate on the "ObjectID" field.
func ObjectIDGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObjectID), v))
	})
}

// ObjectIDLT applies the LT predicate on the "ObjectID" field.
func ObjectIDLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObjectID), v))
	})
}

// ObjectIDLTE applies the LTE predicate on the "ObjectID" field.
func ObjectIDLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObjectID), v))
	})
}

// ObjectIDContains applies the Contains predicate on the "ObjectID" field.
func ObjectIDContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldObjectID), v))
	})
}

// ObjectIDHasPrefix applies the HasPrefix predicate on the "ObjectID" field.
func ObjectIDHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldObjectID), v))
	})
}

// ObjectIDHasSuffix applies the HasSuffix predicate on the "ObjectID" field.
func ObjectIDHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldObjectID), v))
	})
}

// ObjectIDEqualFold applies the EqualFold predicate on the "ObjectID" field.
func ObjectIDEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldObjectID), v))
	})
}

// ObjectIDContainsFold applies the ContainsFold predicate on the "ObjectID" field.
func ObjectIDContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldObjectID), v))
	})
}

// VersionEQ applies the EQ predicate on the "Version" field.
func VersionEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "Version" field.
func VersionNEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "Version" field.
func VersionIn(vs ...uint32) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "Version" field.
func VersionNotIn(vs ...uint32) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "Version" field.
func VersionGT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "Version" field.
func VersionGTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "Version" field.
func VersionLT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "Version" field.
func VersionLTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}