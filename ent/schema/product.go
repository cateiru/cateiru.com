package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		// `id` INT UNSIGNED AUTO_INCREMENT NOT NULL
		field.Uint32("id"),

		// `user_id` INT UNSIGNED NOT NULL
		field.Uint32("user_id"),

		// `name` TEXT NOT NULL
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `name_ja` TEXT NOT NULL`
		field.String("name_ja").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `detail` TEXT NOT NULL
		field.String("detail").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `detail_ja` TEXT NOT NULL
		field.String("detail_ja").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `site_url` TEXT NOT NULL
		field.String("site_url").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `github_url` TEXT
		field.String("github_url").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `dev_time` DATE NOT NULL
		field.Time("dev_time").
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

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),

		index.Fields("user_id", "dev_time"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
