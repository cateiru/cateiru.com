package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Biography holds the schema definition for the Biography entity.
type Biography struct {
	ent.Schema
}

// Fields of the Biography.
func (Biography) Fields() []ent.Field {
	return []ent.Field{
		// `id` INT UNSIGNED AUTO_INCREMENT NOT NULL
		field.Uint32("id"),

		// `user_id` INT UNSIGNED NOT NULL
		field.Uint32("user_id"),

		// `is_public` BOOLEAN NOT NULL DEFAULT 0
		field.Bool("is_public").
			Default(false),

		// `location_id` INT UNSIGNED NOT NULL
		field.Uint32("location_id"),

		// `position` TEXT NOT NULL
		field.Text("position").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `position_ja` TEXT NOT NULL
		field.Text("position_ja").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `join` DATE NOT NULL
		field.Time("join").
			SchemaType(map[string]string{
				dialect.MySQL: "date",
			}),

		// `leave` DATE NOT NULL
		field.Time("leave").
			SchemaType(map[string]string{
				dialect.MySQL: "date",
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

func (Biography) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),

		index.Fields("user_id", "join"),
	}
}

// Edges of the Biography.
func (Biography) Edges() []ent.Edge {
	return nil
}
