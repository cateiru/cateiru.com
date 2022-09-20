package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		// `id` INT UNSIGNED AUTO_INCREMENT NOT NULL
		field.Uint32("id"),

		// `to_user_id` INT UNSIGNED NOT NULL
		field.Uint32("to_user_id"),

		// `title` TEXT NOT NULL
		field.String("title").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `detail` TEXT NOT NULL
		field.String("detail").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `mail` TEXT NOT NULL
		field.String("mail").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `category` TEXT NOT NULL
		field.String("category").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `ip` VARCHAR(16) NOT NULL
		field.String("ip").
			Annotations(entsql.Annotation{
				Size: 16,
			}),

		// `device_name` VARCHAR(31) NOT NULL
		field.String("device_name").
			Annotations(entsql.Annotation{
				Size: 31,
			}),

		// `os` VARCHAR(15) NOT NULL
		field.String("os").
			Annotations(entsql.Annotation{
				Size: 15,
			}),

		// `browser_name` VARCHAR(15) NOT NULL
		field.String("browser_name").
			Annotations(entsql.Annotation{
				Size: 15,
			}),

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

func (Contact) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("to_user_id"),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return nil
}
