// Code generated by ent, DO NOT EDIT.

package nfts

import (
	"entgo.io/ent/dialect/sql"
	"rei.io/rei/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ObjectID applies equality check predicate on the "ObjectID" field. It's identical to ObjectIDEQ.
func ObjectID(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectID), v))
	})
}

// Type applies equality check predicate on the "Type" field. It's identical to TypeEQ.
func Type(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// ObjectIDEQ applies the EQ predicate on the "ObjectID" field.
func ObjectIDEQ(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectID), v))
	})
}

// ObjectIDNEQ applies the NEQ predicate on the "ObjectID" field.
func ObjectIDNEQ(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObjectID), v))
	})
}

// ObjectIDIn applies the In predicate on the "ObjectID" field.
func ObjectIDIn(vs ...string) predicate.NFTs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldObjectID), v...))
	})
}

// ObjectIDNotIn applies the NotIn predicate on the "ObjectID" field.
func ObjectIDNotIn(vs ...string) predicate.NFTs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldObjectID), v...))
	})
}

// ObjectIDGT applies the GT predicate on the "ObjectID" field.
func ObjectIDGT(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObjectID), v))
	})
}

// ObjectIDGTE applies the GTE predicate on the "ObjectID" field.
func ObjectIDGTE(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObjectID), v))
	})
}

// ObjectIDLT applies the LT predicate on the "ObjectID" field.
func ObjectIDLT(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObjectID), v))
	})
}

// ObjectIDLTE applies the LTE predicate on the "ObjectID" field.
func ObjectIDLTE(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObjectID), v))
	})
}

// ObjectIDContains applies the Contains predicate on the "ObjectID" field.
func ObjectIDContains(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldObjectID), v))
	})
}

// ObjectIDHasPrefix applies the HasPrefix predicate on the "ObjectID" field.
func ObjectIDHasPrefix(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldObjectID), v))
	})
}

// ObjectIDHasSuffix applies the HasSuffix predicate on the "ObjectID" field.
func ObjectIDHasSuffix(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldObjectID), v))
	})
}

// ObjectIDEqualFold applies the EqualFold predicate on the "ObjectID" field.
func ObjectIDEqualFold(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldObjectID), v))
	})
}

// ObjectIDContainsFold applies the ContainsFold predicate on the "ObjectID" field.
func ObjectIDContainsFold(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldObjectID), v))
	})
}

// TypeEQ applies the EQ predicate on the "Type" field.
func TypeEQ(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "Type" field.
func TypeNEQ(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "Type" field.
func TypeIn(vs ...string) predicate.NFTs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "Type" field.
func TypeNotIn(vs ...string) predicate.NFTs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "Type" field.
func TypeGT(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "Type" field.
func TypeGTE(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "Type" field.
func TypeLT(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "Type" field.
func TypeLTE(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "Type" field.
func TypeContains(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "Type" field.
func TypeHasPrefix(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "Type" field.
func TypeHasSuffix(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "Type" field.
func TypeEqualFold(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "Type" field.
func TypeContainsFold(v string) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NFTs) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NFTs) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
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
func Not(p predicate.NFTs) predicate.NFTs {
	return predicate.NFTs(func(s *sql.Selector) {
		p(s.Not())
	})
}
