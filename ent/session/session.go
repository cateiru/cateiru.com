// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "session_token"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldPeriod holds the string denoting the period field in the database.
	FieldPeriod = "period"
	// Table holds the table name of the session in the database.
	Table = "sessions"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldCreated,
	FieldPeriod,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
