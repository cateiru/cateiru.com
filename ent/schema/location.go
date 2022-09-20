package schema

import "entgo.io/ent"

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return nil
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return nil
}
