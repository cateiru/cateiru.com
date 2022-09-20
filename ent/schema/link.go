package schema

import "entgo.io/ent"

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return nil
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return nil
}
