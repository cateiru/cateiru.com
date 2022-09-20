// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BiographiesColumns holds the columns for the "biographies" table.
	BiographiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// BiographiesTable holds the schema information for the "biographies" table.
	BiographiesTable = &schema.Table{
		Name:       "biographies",
		Columns:    BiographiesColumns,
		PrimaryKey: []*schema.Column{BiographiesColumns[0]},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// ContactsColumns holds the columns for the "contacts" table.
	ContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ContactsTable holds the schema information for the "contacts" table.
	ContactsTable = &schema.Table{
		Name:       "contacts",
		Columns:    ContactsColumns,
		PrimaryKey: []*schema.Column{ContactsColumns[0]},
	}
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// LinksTable holds the schema information for the "links" table.
	LinksTable = &schema.Table{
		Name:       "links",
		Columns:    LinksColumns,
		PrimaryKey: []*schema.Column{LinksColumns[0]},
	}
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:       "locations",
		Columns:    LocationsColumns,
		PrimaryKey: []*schema.Column{LocationsColumns[0]},
	}
	// NoticesColumns holds the columns for the "notices" table.
	NoticesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// NoticesTable holds the schema information for the "notices" table.
	NoticesTable = &schema.Table{
		Name:       "notices",
		Columns:    NoticesColumns,
		PrimaryKey: []*schema.Column{NoticesColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "given_name", Type: field.TypeString, Size: 10},
		{Name: "family_name", Type: field.TypeString, Size: 10},
		{Name: "given_name_ja", Type: field.TypeString, Size: 10},
		{Name: "family_name_ja", Type: field.TypeString, Size: 10},
		{Name: "user_id", Type: field.TypeString, Size: 10},
		{Name: "mail", Type: field.TypeString, Size: 254},
		{Name: "birth_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "location", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "location_ja", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BiographiesTable,
		CategoriesTable,
		ContactsTable,
		LinksTable,
		LocationsTable,
		NoticesTable,
		ProductsTable,
		UsersTable,
	}
)

func init() {
}
