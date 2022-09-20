// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateir.com/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameJa applies equality check predicate on the "name_ja" field. It's identical to NameJaEQ.
func NameJa(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameJa), v))
	})
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailJa applies equality check predicate on the "detail_ja" field. It's identical to DetailJaEQ.
func DetailJa(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetailJa), v))
	})
}

// SiteURL applies equality check predicate on the "site_url" field. It's identical to SiteURLEQ.
func SiteURL(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSiteURL), v))
	})
}

// GithubURL applies equality check predicate on the "github_url" field. It's identical to GithubURLEQ.
func GithubURL(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubURL), v))
	})
}

// DevTime applies equality check predicate on the "dev_time" field. It's identical to DevTimeEQ.
func DevTime(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDevTime), v))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Modified applies equality check predicate on the "modified" field. It's identical to ModifiedEQ.
func Modified(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModified), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint32) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint32) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint32) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// NameJaEQ applies the EQ predicate on the "name_ja" field.
func NameJaEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameJa), v))
	})
}

// NameJaNEQ applies the NEQ predicate on the "name_ja" field.
func NameJaNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameJa), v))
	})
}

// NameJaIn applies the In predicate on the "name_ja" field.
func NameJaIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNameJa), v...))
	})
}

// NameJaNotIn applies the NotIn predicate on the "name_ja" field.
func NameJaNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNameJa), v...))
	})
}

// NameJaGT applies the GT predicate on the "name_ja" field.
func NameJaGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameJa), v))
	})
}

// NameJaGTE applies the GTE predicate on the "name_ja" field.
func NameJaGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameJa), v))
	})
}

// NameJaLT applies the LT predicate on the "name_ja" field.
func NameJaLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameJa), v))
	})
}

// NameJaLTE applies the LTE predicate on the "name_ja" field.
func NameJaLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameJa), v))
	})
}

// NameJaContains applies the Contains predicate on the "name_ja" field.
func NameJaContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameJa), v))
	})
}

// NameJaHasPrefix applies the HasPrefix predicate on the "name_ja" field.
func NameJaHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameJa), v))
	})
}

// NameJaHasSuffix applies the HasSuffix predicate on the "name_ja" field.
func NameJaHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameJa), v))
	})
}

// NameJaEqualFold applies the EqualFold predicate on the "name_ja" field.
func NameJaEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameJa), v))
	})
}

// NameJaContainsFold applies the ContainsFold predicate on the "name_ja" field.
func NameJaContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameJa), v))
	})
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetail), v))
	})
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDetail), v...))
	})
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDetail), v...))
	})
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetail), v))
	})
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetail), v))
	})
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetail), v))
	})
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetail), v))
	})
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetail), v))
	})
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetail), v))
	})
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetail), v))
	})
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetail), v))
	})
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetail), v))
	})
}

// DetailJaEQ applies the EQ predicate on the "detail_ja" field.
func DetailJaEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetailJa), v))
	})
}

// DetailJaNEQ applies the NEQ predicate on the "detail_ja" field.
func DetailJaNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetailJa), v))
	})
}

// DetailJaIn applies the In predicate on the "detail_ja" field.
func DetailJaIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDetailJa), v...))
	})
}

// DetailJaNotIn applies the NotIn predicate on the "detail_ja" field.
func DetailJaNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDetailJa), v...))
	})
}

// DetailJaGT applies the GT predicate on the "detail_ja" field.
func DetailJaGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetailJa), v))
	})
}

// DetailJaGTE applies the GTE predicate on the "detail_ja" field.
func DetailJaGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetailJa), v))
	})
}

// DetailJaLT applies the LT predicate on the "detail_ja" field.
func DetailJaLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetailJa), v))
	})
}

// DetailJaLTE applies the LTE predicate on the "detail_ja" field.
func DetailJaLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetailJa), v))
	})
}

// DetailJaContains applies the Contains predicate on the "detail_ja" field.
func DetailJaContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetailJa), v))
	})
}

// DetailJaHasPrefix applies the HasPrefix predicate on the "detail_ja" field.
func DetailJaHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetailJa), v))
	})
}

// DetailJaHasSuffix applies the HasSuffix predicate on the "detail_ja" field.
func DetailJaHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetailJa), v))
	})
}

// DetailJaEqualFold applies the EqualFold predicate on the "detail_ja" field.
func DetailJaEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetailJa), v))
	})
}

// DetailJaContainsFold applies the ContainsFold predicate on the "detail_ja" field.
func DetailJaContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetailJa), v))
	})
}

// SiteURLEQ applies the EQ predicate on the "site_url" field.
func SiteURLEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSiteURL), v))
	})
}

// SiteURLNEQ applies the NEQ predicate on the "site_url" field.
func SiteURLNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSiteURL), v))
	})
}

// SiteURLIn applies the In predicate on the "site_url" field.
func SiteURLIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSiteURL), v...))
	})
}

// SiteURLNotIn applies the NotIn predicate on the "site_url" field.
func SiteURLNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSiteURL), v...))
	})
}

// SiteURLGT applies the GT predicate on the "site_url" field.
func SiteURLGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSiteURL), v))
	})
}

// SiteURLGTE applies the GTE predicate on the "site_url" field.
func SiteURLGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSiteURL), v))
	})
}

// SiteURLLT applies the LT predicate on the "site_url" field.
func SiteURLLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSiteURL), v))
	})
}

// SiteURLLTE applies the LTE predicate on the "site_url" field.
func SiteURLLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSiteURL), v))
	})
}

// SiteURLContains applies the Contains predicate on the "site_url" field.
func SiteURLContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSiteURL), v))
	})
}

// SiteURLHasPrefix applies the HasPrefix predicate on the "site_url" field.
func SiteURLHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSiteURL), v))
	})
}

// SiteURLHasSuffix applies the HasSuffix predicate on the "site_url" field.
func SiteURLHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSiteURL), v))
	})
}

// SiteURLEqualFold applies the EqualFold predicate on the "site_url" field.
func SiteURLEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSiteURL), v))
	})
}

// SiteURLContainsFold applies the ContainsFold predicate on the "site_url" field.
func SiteURLContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSiteURL), v))
	})
}

// GithubURLEQ applies the EQ predicate on the "github_url" field.
func GithubURLEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubURL), v))
	})
}

// GithubURLNEQ applies the NEQ predicate on the "github_url" field.
func GithubURLNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGithubURL), v))
	})
}

// GithubURLIn applies the In predicate on the "github_url" field.
func GithubURLIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGithubURL), v...))
	})
}

// GithubURLNotIn applies the NotIn predicate on the "github_url" field.
func GithubURLNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGithubURL), v...))
	})
}

// GithubURLGT applies the GT predicate on the "github_url" field.
func GithubURLGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGithubURL), v))
	})
}

// GithubURLGTE applies the GTE predicate on the "github_url" field.
func GithubURLGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGithubURL), v))
	})
}

// GithubURLLT applies the LT predicate on the "github_url" field.
func GithubURLLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGithubURL), v))
	})
}

// GithubURLLTE applies the LTE predicate on the "github_url" field.
func GithubURLLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGithubURL), v))
	})
}

// GithubURLContains applies the Contains predicate on the "github_url" field.
func GithubURLContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGithubURL), v))
	})
}

// GithubURLHasPrefix applies the HasPrefix predicate on the "github_url" field.
func GithubURLHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGithubURL), v))
	})
}

// GithubURLHasSuffix applies the HasSuffix predicate on the "github_url" field.
func GithubURLHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGithubURL), v))
	})
}

// GithubURLIsNil applies the IsNil predicate on the "github_url" field.
func GithubURLIsNil() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGithubURL)))
	})
}

// GithubURLNotNil applies the NotNil predicate on the "github_url" field.
func GithubURLNotNil() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGithubURL)))
	})
}

// GithubURLEqualFold applies the EqualFold predicate on the "github_url" field.
func GithubURLEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGithubURL), v))
	})
}

// GithubURLContainsFold applies the ContainsFold predicate on the "github_url" field.
func GithubURLContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGithubURL), v))
	})
}

// DevTimeEQ applies the EQ predicate on the "dev_time" field.
func DevTimeEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDevTime), v))
	})
}

// DevTimeNEQ applies the NEQ predicate on the "dev_time" field.
func DevTimeNEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDevTime), v))
	})
}

// DevTimeIn applies the In predicate on the "dev_time" field.
func DevTimeIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDevTime), v...))
	})
}

// DevTimeNotIn applies the NotIn predicate on the "dev_time" field.
func DevTimeNotIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDevTime), v...))
	})
}

// DevTimeGT applies the GT predicate on the "dev_time" field.
func DevTimeGT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDevTime), v))
	})
}

// DevTimeGTE applies the GTE predicate on the "dev_time" field.
func DevTimeGTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDevTime), v))
	})
}

// DevTimeLT applies the LT predicate on the "dev_time" field.
func DevTimeLT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDevTime), v))
	})
}

// DevTimeLTE applies the LTE predicate on the "dev_time" field.
func DevTimeLTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDevTime), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// ModifiedEQ applies the EQ predicate on the "modified" field.
func ModifiedEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModified), v))
	})
}

// ModifiedNEQ applies the NEQ predicate on the "modified" field.
func ModifiedNEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModified), v))
	})
}

// ModifiedIn applies the In predicate on the "modified" field.
func ModifiedIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldModified), v...))
	})
}

// ModifiedNotIn applies the NotIn predicate on the "modified" field.
func ModifiedNotIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldModified), v...))
	})
}

// ModifiedGT applies the GT predicate on the "modified" field.
func ModifiedGT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModified), v))
	})
}

// ModifiedGTE applies the GTE predicate on the "modified" field.
func ModifiedGTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModified), v))
	})
}

// ModifiedLT applies the LT predicate on the "modified" field.
func ModifiedLT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModified), v))
	})
}

// ModifiedLTE applies the LTE predicate on the "modified" field.
func ModifiedLTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModified), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
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
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
