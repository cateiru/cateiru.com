package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// ContactDefault holds the schema definition for the ContactDefault entity.
type ContactDefault struct {
	ent.Schema
}

// Fields of the ContactDefault.
func (ContactDefault) Fields() []ent.Field {
	return []ent.Field{
		// `id` INT UNSIGNED AUTO_INCREMENT NOT NULL
		field.Uint32("id"),

		// `name` TEXT
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Nillable(),

		// `email` TEXT
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Nillable(),

		// `category` TEXT
		field.String("url").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Nillable(),

		// `category` TEXT
		field.String("category").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Nillable(),

		// `custom_title` TEXT
		field.String("custom_title").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Nillable(),

		// `description` TEXT
		field.String("description").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Nillable(),

		// `created` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
		field.Time("created").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now).
			Annotations(entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),

		// `modified` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
		field.Time("modified").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entsql.Annotation{
				// FIXME: If DDL is available, use it. ref.https://github.com/ent/ent/issues/558
				Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the ContactDefault.
func (ContactDefault) Edges() []ent.Edge {
	return nil
}
