// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateir.com/ent/contact"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return (&ContactClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contact is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact

func (c Contacts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
