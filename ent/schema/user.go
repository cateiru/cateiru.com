package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// `id` INT UNSIGNED AUTO_INCREMENT NOT NULL
		field.Uint32("id"),

		// `given_name` VARCHAR(10) NOT NULL
		field.String("given_name").
			Annotations(entsql.Annotation{
				Size: 10,
			}),

		// `family_name` VARCHAR(10) NOT NULL
		field.String("family_name").
			Annotations(entsql.Annotation{
				Size: 10,
			}),

		// `given_name_ja` VARCHAR(10) NOT NULL
		field.String("given_name_ja").
			Annotations(entsql.Annotation{
				Size: 10,
			}),

		// `family_name_ja` VARCHAR(10) NOT NULL
		field.String("family_name_ja").
			Annotations(entsql.Annotation{
				Size: 10,
			}),

		// `user_id` VARCHAR(10) NOT NULL
		field.String("user_id").
			Annotations(entsql.Annotation{
				Size: 10,
			}),

		// `mail` VARCHAR(256) NOT NULL
		field.String("mail").
			Annotations(entsql.Annotation{
				Size: 254,
			}),

		// `birth_date` DATE NOT NULL
		field.Time("birth_date").
			SchemaType(map[string]string{
				dialect.MySQL: "date",
			}),

		// `location` TEXT NOT NULL
		field.Text("location").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),

		// `location_ja` TEXT NOT NULL
		field.Text("location_ja").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),

		// `sso_token` TEXT NOT NULL
		field.Text("sso_token").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `avatar_url` TEXT
		field.Text("avatar_url").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		// `selected` BOOL NOT NULL DEFAULT 0
		field.Bool("selected").
			Default(false),

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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
