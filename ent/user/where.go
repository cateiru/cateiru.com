// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru.com/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GivenName applies equality check predicate on the "given_name" field. It's identical to GivenNameEQ.
func GivenName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGivenName), v))
	})
}

// FamilyName applies equality check predicate on the "family_name" field. It's identical to FamilyNameEQ.
func FamilyName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFamilyName), v))
	})
}

// GivenNameJa applies equality check predicate on the "given_name_ja" field. It's identical to GivenNameJaEQ.
func GivenNameJa(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGivenNameJa), v))
	})
}

// FamilyNameJa applies equality check predicate on the "family_name_ja" field. It's identical to FamilyNameJaEQ.
func FamilyNameJa(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFamilyNameJa), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Mail applies equality check predicate on the "mail" field. It's identical to MailEQ.
func Mail(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMail), v))
	})
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthDate), v))
	})
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// LocationJa applies equality check predicate on the "location_ja" field. It's identical to LocationJaEQ.
func LocationJa(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationJa), v))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Modified applies equality check predicate on the "modified" field. It's identical to ModifiedEQ.
func Modified(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModified), v))
	})
}

// GivenNameEQ applies the EQ predicate on the "given_name" field.
func GivenNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGivenName), v))
	})
}

// GivenNameNEQ applies the NEQ predicate on the "given_name" field.
func GivenNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGivenName), v))
	})
}

// GivenNameIn applies the In predicate on the "given_name" field.
func GivenNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGivenName), v...))
	})
}

// GivenNameNotIn applies the NotIn predicate on the "given_name" field.
func GivenNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGivenName), v...))
	})
}

// GivenNameGT applies the GT predicate on the "given_name" field.
func GivenNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGivenName), v))
	})
}

// GivenNameGTE applies the GTE predicate on the "given_name" field.
func GivenNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGivenName), v))
	})
}

// GivenNameLT applies the LT predicate on the "given_name" field.
func GivenNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGivenName), v))
	})
}

// GivenNameLTE applies the LTE predicate on the "given_name" field.
func GivenNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGivenName), v))
	})
}

// GivenNameContains applies the Contains predicate on the "given_name" field.
func GivenNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGivenName), v))
	})
}

// GivenNameHasPrefix applies the HasPrefix predicate on the "given_name" field.
func GivenNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGivenName), v))
	})
}

// GivenNameHasSuffix applies the HasSuffix predicate on the "given_name" field.
func GivenNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGivenName), v))
	})
}

// GivenNameEqualFold applies the EqualFold predicate on the "given_name" field.
func GivenNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGivenName), v))
	})
}

// GivenNameContainsFold applies the ContainsFold predicate on the "given_name" field.
func GivenNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGivenName), v))
	})
}

// FamilyNameEQ applies the EQ predicate on the "family_name" field.
func FamilyNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFamilyName), v))
	})
}

// FamilyNameNEQ applies the NEQ predicate on the "family_name" field.
func FamilyNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFamilyName), v))
	})
}

// FamilyNameIn applies the In predicate on the "family_name" field.
func FamilyNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFamilyName), v...))
	})
}

// FamilyNameNotIn applies the NotIn predicate on the "family_name" field.
func FamilyNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFamilyName), v...))
	})
}

// FamilyNameGT applies the GT predicate on the "family_name" field.
func FamilyNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFamilyName), v))
	})
}

// FamilyNameGTE applies the GTE predicate on the "family_name" field.
func FamilyNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFamilyName), v))
	})
}

// FamilyNameLT applies the LT predicate on the "family_name" field.
func FamilyNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFamilyName), v))
	})
}

// FamilyNameLTE applies the LTE predicate on the "family_name" field.
func FamilyNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFamilyName), v))
	})
}

// FamilyNameContains applies the Contains predicate on the "family_name" field.
func FamilyNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFamilyName), v))
	})
}

// FamilyNameHasPrefix applies the HasPrefix predicate on the "family_name" field.
func FamilyNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFamilyName), v))
	})
}

// FamilyNameHasSuffix applies the HasSuffix predicate on the "family_name" field.
func FamilyNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFamilyName), v))
	})
}

// FamilyNameEqualFold applies the EqualFold predicate on the "family_name" field.
func FamilyNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFamilyName), v))
	})
}

// FamilyNameContainsFold applies the ContainsFold predicate on the "family_name" field.
func FamilyNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFamilyName), v))
	})
}

// GivenNameJaEQ applies the EQ predicate on the "given_name_ja" field.
func GivenNameJaEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaNEQ applies the NEQ predicate on the "given_name_ja" field.
func GivenNameJaNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaIn applies the In predicate on the "given_name_ja" field.
func GivenNameJaIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGivenNameJa), v...))
	})
}

// GivenNameJaNotIn applies the NotIn predicate on the "given_name_ja" field.
func GivenNameJaNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGivenNameJa), v...))
	})
}

// GivenNameJaGT applies the GT predicate on the "given_name_ja" field.
func GivenNameJaGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaGTE applies the GTE predicate on the "given_name_ja" field.
func GivenNameJaGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaLT applies the LT predicate on the "given_name_ja" field.
func GivenNameJaLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaLTE applies the LTE predicate on the "given_name_ja" field.
func GivenNameJaLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaContains applies the Contains predicate on the "given_name_ja" field.
func GivenNameJaContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaHasPrefix applies the HasPrefix predicate on the "given_name_ja" field.
func GivenNameJaHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaHasSuffix applies the HasSuffix predicate on the "given_name_ja" field.
func GivenNameJaHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaEqualFold applies the EqualFold predicate on the "given_name_ja" field.
func GivenNameJaEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGivenNameJa), v))
	})
}

// GivenNameJaContainsFold applies the ContainsFold predicate on the "given_name_ja" field.
func GivenNameJaContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGivenNameJa), v))
	})
}

// FamilyNameJaEQ applies the EQ predicate on the "family_name_ja" field.
func FamilyNameJaEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaNEQ applies the NEQ predicate on the "family_name_ja" field.
func FamilyNameJaNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaIn applies the In predicate on the "family_name_ja" field.
func FamilyNameJaIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFamilyNameJa), v...))
	})
}

// FamilyNameJaNotIn applies the NotIn predicate on the "family_name_ja" field.
func FamilyNameJaNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFamilyNameJa), v...))
	})
}

// FamilyNameJaGT applies the GT predicate on the "family_name_ja" field.
func FamilyNameJaGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaGTE applies the GTE predicate on the "family_name_ja" field.
func FamilyNameJaGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaLT applies the LT predicate on the "family_name_ja" field.
func FamilyNameJaLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaLTE applies the LTE predicate on the "family_name_ja" field.
func FamilyNameJaLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaContains applies the Contains predicate on the "family_name_ja" field.
func FamilyNameJaContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaHasPrefix applies the HasPrefix predicate on the "family_name_ja" field.
func FamilyNameJaHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaHasSuffix applies the HasSuffix predicate on the "family_name_ja" field.
func FamilyNameJaHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaEqualFold applies the EqualFold predicate on the "family_name_ja" field.
func FamilyNameJaEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFamilyNameJa), v))
	})
}

// FamilyNameJaContainsFold applies the ContainsFold predicate on the "family_name_ja" field.
func FamilyNameJaContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFamilyNameJa), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// MailEQ applies the EQ predicate on the "mail" field.
func MailEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMail), v))
	})
}

// MailNEQ applies the NEQ predicate on the "mail" field.
func MailNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMail), v))
	})
}

// MailIn applies the In predicate on the "mail" field.
func MailIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMail), v...))
	})
}

// MailNotIn applies the NotIn predicate on the "mail" field.
func MailNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMail), v...))
	})
}

// MailGT applies the GT predicate on the "mail" field.
func MailGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMail), v))
	})
}

// MailGTE applies the GTE predicate on the "mail" field.
func MailGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMail), v))
	})
}

// MailLT applies the LT predicate on the "mail" field.
func MailLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMail), v))
	})
}

// MailLTE applies the LTE predicate on the "mail" field.
func MailLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMail), v))
	})
}

// MailContains applies the Contains predicate on the "mail" field.
func MailContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMail), v))
	})
}

// MailHasPrefix applies the HasPrefix predicate on the "mail" field.
func MailHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMail), v))
	})
}

// MailHasSuffix applies the HasSuffix predicate on the "mail" field.
func MailHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMail), v))
	})
}

// MailEqualFold applies the EqualFold predicate on the "mail" field.
func MailEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMail), v))
	})
}

// MailContainsFold applies the ContainsFold predicate on the "mail" field.
func MailContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMail), v))
	})
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthDate), v))
	})
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBirthDate), v))
	})
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBirthDate), v...))
	})
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBirthDate), v...))
	})
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBirthDate), v))
	})
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBirthDate), v))
	})
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBirthDate), v))
	})
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBirthDate), v))
	})
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocation), v))
	})
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLocation), v...))
	})
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLocation), v...))
	})
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocation), v))
	})
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocation), v))
	})
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocation), v))
	})
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocation), v))
	})
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocation), v))
	})
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocation), v))
	})
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocation), v))
	})
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocation), v))
	})
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocation), v))
	})
}

// LocationJaEQ applies the EQ predicate on the "location_ja" field.
func LocationJaEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationJa), v))
	})
}

// LocationJaNEQ applies the NEQ predicate on the "location_ja" field.
func LocationJaNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationJa), v))
	})
}

// LocationJaIn applies the In predicate on the "location_ja" field.
func LocationJaIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLocationJa), v...))
	})
}

// LocationJaNotIn applies the NotIn predicate on the "location_ja" field.
func LocationJaNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLocationJa), v...))
	})
}

// LocationJaGT applies the GT predicate on the "location_ja" field.
func LocationJaGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationJa), v))
	})
}

// LocationJaGTE applies the GTE predicate on the "location_ja" field.
func LocationJaGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationJa), v))
	})
}

// LocationJaLT applies the LT predicate on the "location_ja" field.
func LocationJaLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationJa), v))
	})
}

// LocationJaLTE applies the LTE predicate on the "location_ja" field.
func LocationJaLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationJa), v))
	})
}

// LocationJaContains applies the Contains predicate on the "location_ja" field.
func LocationJaContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationJa), v))
	})
}

// LocationJaHasPrefix applies the HasPrefix predicate on the "location_ja" field.
func LocationJaHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationJa), v))
	})
}

// LocationJaHasSuffix applies the HasSuffix predicate on the "location_ja" field.
func LocationJaHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationJa), v))
	})
}

// LocationJaEqualFold applies the EqualFold predicate on the "location_ja" field.
func LocationJaEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationJa), v))
	})
}

// LocationJaContainsFold applies the ContainsFold predicate on the "location_ja" field.
func LocationJaContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationJa), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// ModifiedEQ applies the EQ predicate on the "modified" field.
func ModifiedEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModified), v))
	})
}

// ModifiedNEQ applies the NEQ predicate on the "modified" field.
func ModifiedNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModified), v))
	})
}

// ModifiedIn applies the In predicate on the "modified" field.
func ModifiedIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldModified), v...))
	})
}

// ModifiedNotIn applies the NotIn predicate on the "modified" field.
func ModifiedNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldModified), v...))
	})
}

// ModifiedGT applies the GT predicate on the "modified" field.
func ModifiedGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModified), v))
	})
}

// ModifiedGTE applies the GTE predicate on the "modified" field.
func ModifiedGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModified), v))
	})
}

// ModifiedLT applies the LT predicate on the "modified" field.
func ModifiedLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModified), v))
	})
}

// ModifiedLTE applies the LTE predicate on the "modified" field.
func ModifiedLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModified), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
