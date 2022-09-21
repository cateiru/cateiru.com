// Code generated by ent, DO NOT EDIT.

package biography

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru.com/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// IsPublic applies equality check predicate on the "is_public" field. It's identical to IsPublicEQ.
func IsPublic(v bool) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPublic), v))
	})
}

// LocationID applies equality check predicate on the "location_id" field. It's identical to LocationIDEQ.
func LocationID(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationID), v))
	})
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// Join applies equality check predicate on the "join" field. It's identical to JoinEQ.
func Join(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJoin), v))
	})
}

// Leave applies equality check predicate on the "leave" field. It's identical to LeaveEQ.
func Leave(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeave), v))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Modified applies equality check predicate on the "modified" field. It's identical to ModifiedEQ.
func Modified(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModified), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint32) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint32) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// IsPublicEQ applies the EQ predicate on the "is_public" field.
func IsPublicEQ(v bool) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPublic), v))
	})
}

// IsPublicNEQ applies the NEQ predicate on the "is_public" field.
func IsPublicNEQ(v bool) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsPublic), v))
	})
}

// LocationIDEQ applies the EQ predicate on the "location_id" field.
func LocationIDEQ(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationID), v))
	})
}

// LocationIDNEQ applies the NEQ predicate on the "location_id" field.
func LocationIDNEQ(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationID), v))
	})
}

// LocationIDIn applies the In predicate on the "location_id" field.
func LocationIDIn(vs ...uint32) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLocationID), v...))
	})
}

// LocationIDNotIn applies the NotIn predicate on the "location_id" field.
func LocationIDNotIn(vs ...uint32) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLocationID), v...))
	})
}

// LocationIDGT applies the GT predicate on the "location_id" field.
func LocationIDGT(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationID), v))
	})
}

// LocationIDGTE applies the GTE predicate on the "location_id" field.
func LocationIDGTE(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationID), v))
	})
}

// LocationIDLT applies the LT predicate on the "location_id" field.
func LocationIDLT(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationID), v))
	})
}

// LocationIDLTE applies the LTE predicate on the "location_id" field.
func LocationIDLTE(v uint32) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationID), v))
	})
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPosition), v))
	})
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPosition), v...))
	})
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPosition), v...))
	})
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPosition), v))
	})
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPosition), v))
	})
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPosition), v))
	})
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPosition), v))
	})
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPosition), v))
	})
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPosition), v))
	})
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPosition), v))
	})
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPosition), v))
	})
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPosition), v))
	})
}

// JoinEQ applies the EQ predicate on the "join" field.
func JoinEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJoin), v))
	})
}

// JoinNEQ applies the NEQ predicate on the "join" field.
func JoinNEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJoin), v))
	})
}

// JoinIn applies the In predicate on the "join" field.
func JoinIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldJoin), v...))
	})
}

// JoinNotIn applies the NotIn predicate on the "join" field.
func JoinNotIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldJoin), v...))
	})
}

// JoinGT applies the GT predicate on the "join" field.
func JoinGT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJoin), v))
	})
}

// JoinGTE applies the GTE predicate on the "join" field.
func JoinGTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJoin), v))
	})
}

// JoinLT applies the LT predicate on the "join" field.
func JoinLT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJoin), v))
	})
}

// JoinLTE applies the LTE predicate on the "join" field.
func JoinLTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJoin), v))
	})
}

// LeaveEQ applies the EQ predicate on the "leave" field.
func LeaveEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeave), v))
	})
}

// LeaveNEQ applies the NEQ predicate on the "leave" field.
func LeaveNEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLeave), v))
	})
}

// LeaveIn applies the In predicate on the "leave" field.
func LeaveIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLeave), v...))
	})
}

// LeaveNotIn applies the NotIn predicate on the "leave" field.
func LeaveNotIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLeave), v...))
	})
}

// LeaveGT applies the GT predicate on the "leave" field.
func LeaveGT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLeave), v))
	})
}

// LeaveGTE applies the GTE predicate on the "leave" field.
func LeaveGTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLeave), v))
	})
}

// LeaveLT applies the LT predicate on the "leave" field.
func LeaveLT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLeave), v))
	})
}

// LeaveLTE applies the LTE predicate on the "leave" field.
func LeaveLTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLeave), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// ModifiedEQ applies the EQ predicate on the "modified" field.
func ModifiedEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModified), v))
	})
}

// ModifiedNEQ applies the NEQ predicate on the "modified" field.
func ModifiedNEQ(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModified), v))
	})
}

// ModifiedIn applies the In predicate on the "modified" field.
func ModifiedIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldModified), v...))
	})
}

// ModifiedNotIn applies the NotIn predicate on the "modified" field.
func ModifiedNotIn(vs ...time.Time) predicate.Biography {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldModified), v...))
	})
}

// ModifiedGT applies the GT predicate on the "modified" field.
func ModifiedGT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModified), v))
	})
}

// ModifiedGTE applies the GTE predicate on the "modified" field.
func ModifiedGTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModified), v))
	})
}

// ModifiedLT applies the LT predicate on the "modified" field.
func ModifiedLT(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModified), v))
	})
}

// ModifiedLTE applies the LTE predicate on the "modified" field.
func ModifiedLTE(v time.Time) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModified), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Biography) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Biography) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
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
func Not(p predicate.Biography) predicate.Biography {
	return predicate.Biography(func(s *sql.Selector) {
		p(s.Not())
	})
}
