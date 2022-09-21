// Code generated by ent, DO NOT EDIT.

package notice

import (
	"time"
)

const (
	// Label holds the string label denoting the notice type in the database.
	Label = "notice"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDiscordWebhook holds the string denoting the discord_webhook field in the database.
	FieldDiscordWebhook = "discord_webhook"
	// FieldSlackWebhook holds the string denoting the slack_webhook field in the database.
	FieldSlackWebhook = "slack_webhook"
	// FieldMail holds the string denoting the mail field in the database.
	FieldMail = "mail"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// Table holds the table name of the notice in the database.
	Table = "notices"
)

// Columns holds all SQL columns for notice fields.
var Columns = []string{
	FieldID,
	FieldDiscordWebhook,
	FieldSlackWebhook,
	FieldMail,
	FieldCreated,
	FieldModified,
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

var (
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// DefaultModified holds the default value on creation for the "modified" field.
	DefaultModified func() time.Time
	// UpdateDefaultModified holds the default value on update for the "modified" field.
	UpdateDefaultModified func() time.Time
)