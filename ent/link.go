// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateir.com/ent/link"
)

// Link is the model entity for the Link schema.
type Link struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint32 `json:"user_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// NameJa holds the value of the "name_ja" field.
	NameJa string `json:"name_ja,omitempty"`
	// SiteURL holds the value of the "site_url" field.
	SiteURL string `json:"site_url,omitempty"`
	// FaviconURL holds the value of the "favicon_url" field.
	FaviconURL string `json:"favicon_url,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID uint32 `json:"category_id,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Link) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case link.FieldID, link.FieldUserID, link.FieldCategoryID:
			values[i] = new(sql.NullInt64)
		case link.FieldName, link.FieldNameJa, link.FieldSiteURL, link.FieldFaviconURL:
			values[i] = new(sql.NullString)
		case link.FieldCreated, link.FieldModified:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Link", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Link fields.
func (l *Link) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case link.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = uint32(value.Int64)
		case link.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				l.UserID = uint32(value.Int64)
			}
		case link.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case link.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				l.NameJa = value.String
			}
		case link.FieldSiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field site_url", values[i])
			} else if value.Valid {
				l.SiteURL = value.String
			}
		case link.FieldFaviconURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field favicon_url", values[i])
			} else if value.Valid {
				l.FaviconURL = value.String
			}
		case link.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				l.CategoryID = uint32(value.Int64)
			}
		case link.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				l.Created = value.Time
			}
		case link.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				l.Modified = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Link.
// Note that you need to call Link.Unwrap() before calling this method if this Link
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Link) Update() *LinkUpdateOne {
	return (&LinkClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Link entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Link) Unwrap() *Link {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Link is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Link) String() string {
	var builder strings.Builder
	builder.WriteString("Link(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", l.UserID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(l.NameJa)
	builder.WriteString(", ")
	builder.WriteString("site_url=")
	builder.WriteString(l.SiteURL)
	builder.WriteString(", ")
	builder.WriteString("favicon_url=")
	builder.WriteString(l.FaviconURL)
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", l.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(l.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(l.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Links is a parsable slice of Link.
type Links []*Link

func (l Links) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
