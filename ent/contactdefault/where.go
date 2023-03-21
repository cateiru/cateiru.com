// Code generated by ent, DO NOT EDIT.

package contactdefault

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru.com/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldEmail, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldURL, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldCategory, v))
}

// CustomTitle applies equality check predicate on the "custom_title" field. It's identical to CustomTitleEQ.
func CustomTitle(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldCustomTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldDescription, v))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldCreated, v))
}

// Modified applies equality check predicate on the "modified" field. It's identical to ModifiedEQ.
func Modified(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldModified, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContainsFold(FieldEmail, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasSuffix(FieldURL, v))
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIsNull(FieldURL))
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotNull(FieldURL))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContainsFold(FieldURL, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryIsNil applies the IsNil predicate on the "category" field.
func CategoryIsNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIsNull(FieldCategory))
}

// CategoryNotNil applies the NotNil predicate on the "category" field.
func CategoryNotNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotNull(FieldCategory))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContainsFold(FieldCategory, v))
}

// CustomTitleEQ applies the EQ predicate on the "custom_title" field.
func CustomTitleEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldCustomTitle, v))
}

// CustomTitleNEQ applies the NEQ predicate on the "custom_title" field.
func CustomTitleNEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldCustomTitle, v))
}

// CustomTitleIn applies the In predicate on the "custom_title" field.
func CustomTitleIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldCustomTitle, vs...))
}

// CustomTitleNotIn applies the NotIn predicate on the "custom_title" field.
func CustomTitleNotIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldCustomTitle, vs...))
}

// CustomTitleGT applies the GT predicate on the "custom_title" field.
func CustomTitleGT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldCustomTitle, v))
}

// CustomTitleGTE applies the GTE predicate on the "custom_title" field.
func CustomTitleGTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldCustomTitle, v))
}

// CustomTitleLT applies the LT predicate on the "custom_title" field.
func CustomTitleLT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldCustomTitle, v))
}

// CustomTitleLTE applies the LTE predicate on the "custom_title" field.
func CustomTitleLTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldCustomTitle, v))
}

// CustomTitleContains applies the Contains predicate on the "custom_title" field.
func CustomTitleContains(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContains(FieldCustomTitle, v))
}

// CustomTitleHasPrefix applies the HasPrefix predicate on the "custom_title" field.
func CustomTitleHasPrefix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasPrefix(FieldCustomTitle, v))
}

// CustomTitleHasSuffix applies the HasSuffix predicate on the "custom_title" field.
func CustomTitleHasSuffix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasSuffix(FieldCustomTitle, v))
}

// CustomTitleIsNil applies the IsNil predicate on the "custom_title" field.
func CustomTitleIsNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIsNull(FieldCustomTitle))
}

// CustomTitleNotNil applies the NotNil predicate on the "custom_title" field.
func CustomTitleNotNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotNull(FieldCustomTitle))
}

// CustomTitleEqualFold applies the EqualFold predicate on the "custom_title" field.
func CustomTitleEqualFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEqualFold(FieldCustomTitle, v))
}

// CustomTitleContainsFold applies the ContainsFold predicate on the "custom_title" field.
func CustomTitleContainsFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContainsFold(FieldCustomTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldCreated, v))
}

// ModifiedEQ applies the EQ predicate on the "modified" field.
func ModifiedEQ(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldEQ(FieldModified, v))
}

// ModifiedNEQ applies the NEQ predicate on the "modified" field.
func ModifiedNEQ(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNEQ(FieldModified, v))
}

// ModifiedIn applies the In predicate on the "modified" field.
func ModifiedIn(vs ...time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldIn(FieldModified, vs...))
}

// ModifiedNotIn applies the NotIn predicate on the "modified" field.
func ModifiedNotIn(vs ...time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldNotIn(FieldModified, vs...))
}

// ModifiedGT applies the GT predicate on the "modified" field.
func ModifiedGT(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGT(FieldModified, v))
}

// ModifiedGTE applies the GTE predicate on the "modified" field.
func ModifiedGTE(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldGTE(FieldModified, v))
}

// ModifiedLT applies the LT predicate on the "modified" field.
func ModifiedLT(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLT(FieldModified, v))
}

// ModifiedLTE applies the LTE predicate on the "modified" field.
func ModifiedLTE(v time.Time) predicate.ContactDefault {
	return predicate.ContactDefault(sql.FieldLTE(FieldModified, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ContactDefault) predicate.ContactDefault {
	return predicate.ContactDefault(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ContactDefault) predicate.ContactDefault {
	return predicate.ContactDefault(func(s *sql.Selector) {
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
func Not(p predicate.ContactDefault) predicate.ContactDefault {
	return predicate.ContactDefault(func(s *sql.Selector) {
		p(s.Not())
	})
}