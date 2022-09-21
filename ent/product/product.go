// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNameJa holds the string denoting the name_ja field in the database.
	FieldNameJa = "name_ja"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldDetailJa holds the string denoting the detail_ja field in the database.
	FieldDetailJa = "detail_ja"
	// FieldSiteURL holds the string denoting the site_url field in the database.
	FieldSiteURL = "site_url"
	// FieldGithubURL holds the string denoting the github_url field in the database.
	FieldGithubURL = "github_url"
	// FieldDevTime holds the string denoting the dev_time field in the database.
	FieldDevTime = "dev_time"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// Table holds the table name of the product in the database.
	Table = "products"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldName,
	FieldNameJa,
	FieldDetail,
	FieldDetailJa,
	FieldSiteURL,
	FieldGithubURL,
	FieldDevTime,
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