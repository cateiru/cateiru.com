package schema

import "entgo.io/ent"

// Biography holds the schema definition for the Biography entity.
type Biography struct {
	ent.Schema
}

// Fields of the Biography.
func (Biography) Fields() []ent.Field {
	return nil
}

// Edges of the Biography.
func (Biography) Edges() []ent.Edge {
	return nil
}
