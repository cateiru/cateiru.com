// Code generated by ent, DO NOT EDIT.

package link

import (
	"time"
)

const (
	// Label holds the string label denoting the link type in the database.
	Label = "link"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNameJa holds the string denoting the name_ja field in the database.
	FieldNameJa = "name_ja"
	// FieldSiteURL holds the string denoting the site_url field in the database.
	FieldSiteURL = "site_url"
	// FieldFaviconURL holds the string denoting the favicon_url field in the database.
	FieldFaviconURL = "favicon_url"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// Table holds the table name of the link in the database.
	Table = "links"
)

// Columns holds all SQL columns for link fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldName,
	FieldNameJa,
	FieldSiteURL,
	FieldFaviconURL,
	FieldCategoryID,
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