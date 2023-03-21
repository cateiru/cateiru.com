// Code generated by ent, DO NOT EDIT.

package notice

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru.com/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Notice {
	return predicate.Notice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Notice {
	return predicate.Notice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Notice {
	return predicate.Notice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Notice {
	return predicate.Notice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Notice {
	return predicate.Notice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Notice {
	return predicate.Notice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Notice {
	return predicate.Notice(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint32) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldUserID, v))
}

// DiscordWebhook applies equality check predicate on the "discord_webhook" field. It's identical to DiscordWebhookEQ.
func DiscordWebhook(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldDiscordWebhook, v))
}

// SlackWebhook applies equality check predicate on the "slack_webhook" field. It's identical to SlackWebhookEQ.
func SlackWebhook(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldSlackWebhook, v))
}

// Mail applies equality check predicate on the "mail" field. It's identical to MailEQ.
func Mail(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldMail, v))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldCreated, v))
}

// Modified applies equality check predicate on the "modified" field. It's identical to ModifiedEQ.
func Modified(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldModified, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint32) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint32) predicate.Notice {
	return predicate.Notice(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint32) predicate.Notice {
	return predicate.Notice(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint32) predicate.Notice {
	return predicate.Notice(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint32) predicate.Notice {
	return predicate.Notice(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint32) predicate.Notice {
	return predicate.Notice(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint32) predicate.Notice {
	return predicate.Notice(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint32) predicate.Notice {
	return predicate.Notice(sql.FieldLTE(FieldUserID, v))
}

// DiscordWebhookEQ applies the EQ predicate on the "discord_webhook" field.
func DiscordWebhookEQ(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldDiscordWebhook, v))
}

// DiscordWebhookNEQ applies the NEQ predicate on the "discord_webhook" field.
func DiscordWebhookNEQ(v string) predicate.Notice {
	return predicate.Notice(sql.FieldNEQ(FieldDiscordWebhook, v))
}

// DiscordWebhookIn applies the In predicate on the "discord_webhook" field.
func DiscordWebhookIn(vs ...string) predicate.Notice {
	return predicate.Notice(sql.FieldIn(FieldDiscordWebhook, vs...))
}

// DiscordWebhookNotIn applies the NotIn predicate on the "discord_webhook" field.
func DiscordWebhookNotIn(vs ...string) predicate.Notice {
	return predicate.Notice(sql.FieldNotIn(FieldDiscordWebhook, vs...))
}

// DiscordWebhookGT applies the GT predicate on the "discord_webhook" field.
func DiscordWebhookGT(v string) predicate.Notice {
	return predicate.Notice(sql.FieldGT(FieldDiscordWebhook, v))
}

// DiscordWebhookGTE applies the GTE predicate on the "discord_webhook" field.
func DiscordWebhookGTE(v string) predicate.Notice {
	return predicate.Notice(sql.FieldGTE(FieldDiscordWebhook, v))
}

// DiscordWebhookLT applies the LT predicate on the "discord_webhook" field.
func DiscordWebhookLT(v string) predicate.Notice {
	return predicate.Notice(sql.FieldLT(FieldDiscordWebhook, v))
}

// DiscordWebhookLTE applies the LTE predicate on the "discord_webhook" field.
func DiscordWebhookLTE(v string) predicate.Notice {
	return predicate.Notice(sql.FieldLTE(FieldDiscordWebhook, v))
}

// DiscordWebhookContains applies the Contains predicate on the "discord_webhook" field.
func DiscordWebhookContains(v string) predicate.Notice {
	return predicate.Notice(sql.FieldContains(FieldDiscordWebhook, v))
}

// DiscordWebhookHasPrefix applies the HasPrefix predicate on the "discord_webhook" field.
func DiscordWebhookHasPrefix(v string) predicate.Notice {
	return predicate.Notice(sql.FieldHasPrefix(FieldDiscordWebhook, v))
}

// DiscordWebhookHasSuffix applies the HasSuffix predicate on the "discord_webhook" field.
func DiscordWebhookHasSuffix(v string) predicate.Notice {
	return predicate.Notice(sql.FieldHasSuffix(FieldDiscordWebhook, v))
}

// DiscordWebhookIsNil applies the IsNil predicate on the "discord_webhook" field.
func DiscordWebhookIsNil() predicate.Notice {
	return predicate.Notice(sql.FieldIsNull(FieldDiscordWebhook))
}

// DiscordWebhookNotNil applies the NotNil predicate on the "discord_webhook" field.
func DiscordWebhookNotNil() predicate.Notice {
	return predicate.Notice(sql.FieldNotNull(FieldDiscordWebhook))
}

// DiscordWebhookEqualFold applies the EqualFold predicate on the "discord_webhook" field.
func DiscordWebhookEqualFold(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEqualFold(FieldDiscordWebhook, v))
}

// DiscordWebhookContainsFold applies the ContainsFold predicate on the "discord_webhook" field.
func DiscordWebhookContainsFold(v string) predicate.Notice {
	return predicate.Notice(sql.FieldContainsFold(FieldDiscordWebhook, v))
}

// SlackWebhookEQ applies the EQ predicate on the "slack_webhook" field.
func SlackWebhookEQ(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldSlackWebhook, v))
}

// SlackWebhookNEQ applies the NEQ predicate on the "slack_webhook" field.
func SlackWebhookNEQ(v string) predicate.Notice {
	return predicate.Notice(sql.FieldNEQ(FieldSlackWebhook, v))
}

// SlackWebhookIn applies the In predicate on the "slack_webhook" field.
func SlackWebhookIn(vs ...string) predicate.Notice {
	return predicate.Notice(sql.FieldIn(FieldSlackWebhook, vs...))
}

// SlackWebhookNotIn applies the NotIn predicate on the "slack_webhook" field.
func SlackWebhookNotIn(vs ...string) predicate.Notice {
	return predicate.Notice(sql.FieldNotIn(FieldSlackWebhook, vs...))
}

// SlackWebhookGT applies the GT predicate on the "slack_webhook" field.
func SlackWebhookGT(v string) predicate.Notice {
	return predicate.Notice(sql.FieldGT(FieldSlackWebhook, v))
}

// SlackWebhookGTE applies the GTE predicate on the "slack_webhook" field.
func SlackWebhookGTE(v string) predicate.Notice {
	return predicate.Notice(sql.FieldGTE(FieldSlackWebhook, v))
}

// SlackWebhookLT applies the LT predicate on the "slack_webhook" field.
func SlackWebhookLT(v string) predicate.Notice {
	return predicate.Notice(sql.FieldLT(FieldSlackWebhook, v))
}

// SlackWebhookLTE applies the LTE predicate on the "slack_webhook" field.
func SlackWebhookLTE(v string) predicate.Notice {
	return predicate.Notice(sql.FieldLTE(FieldSlackWebhook, v))
}

// SlackWebhookContains applies the Contains predicate on the "slack_webhook" field.
func SlackWebhookContains(v string) predicate.Notice {
	return predicate.Notice(sql.FieldContains(FieldSlackWebhook, v))
}

// SlackWebhookHasPrefix applies the HasPrefix predicate on the "slack_webhook" field.
func SlackWebhookHasPrefix(v string) predicate.Notice {
	return predicate.Notice(sql.FieldHasPrefix(FieldSlackWebhook, v))
}

// SlackWebhookHasSuffix applies the HasSuffix predicate on the "slack_webhook" field.
func SlackWebhookHasSuffix(v string) predicate.Notice {
	return predicate.Notice(sql.FieldHasSuffix(FieldSlackWebhook, v))
}

// SlackWebhookIsNil applies the IsNil predicate on the "slack_webhook" field.
func SlackWebhookIsNil() predicate.Notice {
	return predicate.Notice(sql.FieldIsNull(FieldSlackWebhook))
}

// SlackWebhookNotNil applies the NotNil predicate on the "slack_webhook" field.
func SlackWebhookNotNil() predicate.Notice {
	return predicate.Notice(sql.FieldNotNull(FieldSlackWebhook))
}

// SlackWebhookEqualFold applies the EqualFold predicate on the "slack_webhook" field.
func SlackWebhookEqualFold(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEqualFold(FieldSlackWebhook, v))
}

// SlackWebhookContainsFold applies the ContainsFold predicate on the "slack_webhook" field.
func SlackWebhookContainsFold(v string) predicate.Notice {
	return predicate.Notice(sql.FieldContainsFold(FieldSlackWebhook, v))
}

// MailEQ applies the EQ predicate on the "mail" field.
func MailEQ(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldMail, v))
}

// MailNEQ applies the NEQ predicate on the "mail" field.
func MailNEQ(v string) predicate.Notice {
	return predicate.Notice(sql.FieldNEQ(FieldMail, v))
}

// MailIn applies the In predicate on the "mail" field.
func MailIn(vs ...string) predicate.Notice {
	return predicate.Notice(sql.FieldIn(FieldMail, vs...))
}

// MailNotIn applies the NotIn predicate on the "mail" field.
func MailNotIn(vs ...string) predicate.Notice {
	return predicate.Notice(sql.FieldNotIn(FieldMail, vs...))
}

// MailGT applies the GT predicate on the "mail" field.
func MailGT(v string) predicate.Notice {
	return predicate.Notice(sql.FieldGT(FieldMail, v))
}

// MailGTE applies the GTE predicate on the "mail" field.
func MailGTE(v string) predicate.Notice {
	return predicate.Notice(sql.FieldGTE(FieldMail, v))
}

// MailLT applies the LT predicate on the "mail" field.
func MailLT(v string) predicate.Notice {
	return predicate.Notice(sql.FieldLT(FieldMail, v))
}

// MailLTE applies the LTE predicate on the "mail" field.
func MailLTE(v string) predicate.Notice {
	return predicate.Notice(sql.FieldLTE(FieldMail, v))
}

// MailContains applies the Contains predicate on the "mail" field.
func MailContains(v string) predicate.Notice {
	return predicate.Notice(sql.FieldContains(FieldMail, v))
}

// MailHasPrefix applies the HasPrefix predicate on the "mail" field.
func MailHasPrefix(v string) predicate.Notice {
	return predicate.Notice(sql.FieldHasPrefix(FieldMail, v))
}

// MailHasSuffix applies the HasSuffix predicate on the "mail" field.
func MailHasSuffix(v string) predicate.Notice {
	return predicate.Notice(sql.FieldHasSuffix(FieldMail, v))
}

// MailIsNil applies the IsNil predicate on the "mail" field.
func MailIsNil() predicate.Notice {
	return predicate.Notice(sql.FieldIsNull(FieldMail))
}

// MailNotNil applies the NotNil predicate on the "mail" field.
func MailNotNil() predicate.Notice {
	return predicate.Notice(sql.FieldNotNull(FieldMail))
}

// MailEqualFold applies the EqualFold predicate on the "mail" field.
func MailEqualFold(v string) predicate.Notice {
	return predicate.Notice(sql.FieldEqualFold(FieldMail, v))
}

// MailContainsFold applies the ContainsFold predicate on the "mail" field.
func MailContainsFold(v string) predicate.Notice {
	return predicate.Notice(sql.FieldContainsFold(FieldMail, v))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldLTE(FieldCreated, v))
}

// ModifiedEQ applies the EQ predicate on the "modified" field.
func ModifiedEQ(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldEQ(FieldModified, v))
}

// ModifiedNEQ applies the NEQ predicate on the "modified" field.
func ModifiedNEQ(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldNEQ(FieldModified, v))
}

// ModifiedIn applies the In predicate on the "modified" field.
func ModifiedIn(vs ...time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldIn(FieldModified, vs...))
}

// ModifiedNotIn applies the NotIn predicate on the "modified" field.
func ModifiedNotIn(vs ...time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldNotIn(FieldModified, vs...))
}

// ModifiedGT applies the GT predicate on the "modified" field.
func ModifiedGT(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldGT(FieldModified, v))
}

// ModifiedGTE applies the GTE predicate on the "modified" field.
func ModifiedGTE(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldGTE(FieldModified, v))
}

// ModifiedLT applies the LT predicate on the "modified" field.
func ModifiedLT(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldLT(FieldModified, v))
}

// ModifiedLTE applies the LTE predicate on the "modified" field.
func ModifiedLTE(v time.Time) predicate.Notice {
	return predicate.Notice(sql.FieldLTE(FieldModified, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Notice) predicate.Notice {
	return predicate.Notice(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Notice) predicate.Notice {
	return predicate.Notice(func(s *sql.Selector) {
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
func Not(p predicate.Notice) predicate.Notice {
	return predicate.Notice(func(s *sql.Selector) {
		p(s.Not())
	})
}
