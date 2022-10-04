// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BiographiesColumns holds the columns for the "biographies" table.
	BiographiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "is_public", Type: field.TypeBool, Default: false},
		{Name: "location_id", Type: field.TypeUint32},
		{Name: "position", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "join", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "leave", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// BiographiesTable holds the schema information for the "biographies" table.
	BiographiesTable = &schema.Table{
		Name:       "biographies",
		Columns:    BiographiesColumns,
		PrimaryKey: []*schema.Column{BiographiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "biography_user_id",
				Unique:  false,
				Columns: []*schema.Column{BiographiesColumns[1]},
			},
			{
				Name:    "biography_user_id_join",
				Unique:  false,
				Columns: []*schema.Column{BiographiesColumns[1], BiographiesColumns[5]},
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "name_ja", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "emoji", Type: field.TypeString, Size: 1, SchemaType: map[string]string{"mysql": "char(1)"}},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// ContactsColumns holds the columns for the "contacts" table.
	ContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "to_user_id", Type: field.TypeUint32},
		{Name: "title", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "detail", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "mail", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "category", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "ip", Type: field.TypeString, Size: 16},
		{Name: "device_name", Type: field.TypeString, Size: 31},
		{Name: "os", Type: field.TypeString, Size: 15},
		{Name: "browser_name", Type: field.TypeString, Size: 15},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// ContactsTable holds the schema information for the "contacts" table.
	ContactsTable = &schema.Table{
		Name:       "contacts",
		Columns:    ContactsColumns,
		PrimaryKey: []*schema.Column{ContactsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "contact_to_user_id",
				Unique:  false,
				Columns: []*schema.Column{ContactsColumns[1]},
			},
		},
	}
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "name_ja", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "site_url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "favicon_url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "category_id", Type: field.TypeUint32},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// LinksTable holds the schema information for the "links" table.
	LinksTable = &schema.Table{
		Name:       "links",
		Columns:    LinksColumns,
		PrimaryKey: []*schema.Column{LinksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "link_user_id",
				Unique:  false,
				Columns: []*schema.Column{LinksColumns[1]},
			},
		},
	}
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"univ", "corp"}},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "name_ja", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "address", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "address_ja", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:       "locations",
		Columns:    LocationsColumns,
		PrimaryKey: []*schema.Column{LocationsColumns[0]},
	}
	// NoticesColumns holds the columns for the "notices" table.
	NoticesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "discord_webhook", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "slack_webhook", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "mail", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// NoticesTable holds the schema information for the "notices" table.
	NoticesTable = &schema.Table{
		Name:       "notices",
		Columns:    NoticesColumns,
		PrimaryKey: []*schema.Column{NoticesColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "name_ja", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "detail", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "detail_ja", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "site_url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "github_url", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "dev_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "modified", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "product_user_id",
				Unique:  false,
				Columns: []*schema.Column{ProductsColumns[1]},
			},
			{
				Name:    "product_user_id_dev_time",
				Unique:  false,
				Columns: []*schema.Column{ProductsColumns[1], ProductsColumns[8]},
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "session_token", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "period", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "session_user_id",
				Unique:  false,
				Columns: []*schema.Column{SessionsColumns[1]},
			},
		},
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
		{Name: "sso_token", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "text"}},
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
		SessionsTable,
		UsersTable,
	}
)

func init() {
}
