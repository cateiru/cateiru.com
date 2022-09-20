// Code generated by ent, DO NOT EDIT.

package contact

import (
	"time"
)

const (
	// Label holds the string label denoting the contact type in the database.
	Label = "contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToUserID holds the string denoting the to_user_id field in the database.
	FieldToUserID = "to_user_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldMail holds the string denoting the mail field in the database.
	FieldMail = "mail"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldDeviceName holds the string denoting the device_name field in the database.
	FieldDeviceName = "device_name"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldBrowserName holds the string denoting the browser_name field in the database.
	FieldBrowserName = "browser_name"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// Table holds the table name of the contact in the database.
	Table = "contacts"
)

// Columns holds all SQL columns for contact fields.
var Columns = []string{
	FieldID,
	FieldToUserID,
	FieldTitle,
	FieldDetail,
	FieldMail,
	FieldCategory,
	FieldIP,
	FieldDeviceName,
	FieldOs,
	FieldBrowserName,
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
