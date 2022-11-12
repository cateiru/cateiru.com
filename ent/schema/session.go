package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		// `session_token` CHAR(36) UNSIGNED AUTO_INCREMENT NOT NULL
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).StorageKey("session_token"),

		// `user_id` INT UNSIGNED NOT NULL
		field.Uint32("user_id"),

		// `created` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
		field.Time("created").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now).
			Annotations(entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),

		// `period` DATETIME NOT NULL
		field.Time("period").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
