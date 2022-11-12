package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		// `id` INT UNSIGNED AUTO_INCREMENT NOT NULL
		field.Uint32("id"),

		// `type` ENUM(`univ`, `corp`) NOT NULL
		field.Enum("type").Values("univ", "corp"),

		// `name` TEXT NOT NULL
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),

		// `name_ja` TEXT NOT NULL
		field.String("name_ja").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),

		// `address` TEXT NOT NULL
		field.String("address").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),

		// `address_ja` TEXT NOT NULL
		field.String("address_ja").SchemaType(map[string]string{
			dialect.MySQL: "text",
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

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return nil
}
