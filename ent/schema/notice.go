package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Notice holds the schema definition for the Notice entity.
type Notice struct {
	ent.Schema
}

// Fields of the Notice.
func (Notice) Fields() []ent.Field {
	return []ent.Field{
		// `id` INT UNSIGNED AUTO_INCREMENT NOT NULL
		field.Uint32("id"),

		// `user_id` INT UNSIGNED UNIQUE NOT NULL
		field.Uint32("user_id").
			Unique(),

		// `discord_webhook` TEXT
		field.String("discord_webhook").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Optional(),

		// `slack_webhook` TEXT
		field.String("slack_webhook").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Optional(),

		// `mail` TEXT
		field.String("mail").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			Optional(),

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

// Edges of the Notice.
func (Notice) Edges() []ent.Edge {
	return nil
}
