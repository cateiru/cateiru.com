// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru.com/ent/contactdefault"
)

// ContactDefault is the model entity for the ContactDefault schema.
type ContactDefault struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email *string `json:"email,omitempty"`
	// URL holds the value of the "url" field.
	URL *string `json:"url,omitempty"`
	// Category holds the value of the "category" field.
	Category *string `json:"category,omitempty"`
	// CustomTitle holds the value of the "custom_title" field.
	CustomTitle *string `json:"custom_title,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ContactDefault) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contactdefault.FieldID:
			values[i] = new(sql.NullInt64)
		case contactdefault.FieldName, contactdefault.FieldEmail, contactdefault.FieldURL, contactdefault.FieldCategory, contactdefault.FieldCustomTitle, contactdefault.FieldDescription:
			values[i] = new(sql.NullString)
		case contactdefault.FieldCreated, contactdefault.FieldModified:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ContactDefault", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ContactDefault fields.
func (cd *ContactDefault) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contactdefault.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cd.ID = uint32(value.Int64)
		case contactdefault.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cd.Name = new(string)
				*cd.Name = value.String
			}
		case contactdefault.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				cd.Email = new(string)
				*cd.Email = value.String
			}
		case contactdefault.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				cd.URL = new(string)
				*cd.URL = value.String
			}
		case contactdefault.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				cd.Category = new(string)
				*cd.Category = value.String
			}
		case contactdefault.FieldCustomTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field custom_title", values[i])
			} else if value.Valid {
				cd.CustomTitle = new(string)
				*cd.CustomTitle = value.String
			}
		case contactdefault.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				cd.Description = new(string)
				*cd.Description = value.String
			}
		case contactdefault.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				cd.Created = value.Time
			}
		case contactdefault.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				cd.Modified = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ContactDefault.
// Note that you need to call ContactDefault.Unwrap() before calling this method if this ContactDefault
// was returned from a transaction, and the transaction was committed or rolled back.
func (cd *ContactDefault) Update() *ContactDefaultUpdateOne {
	return NewContactDefaultClient(cd.config).UpdateOne(cd)
}

// Unwrap unwraps the ContactDefault entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cd *ContactDefault) Unwrap() *ContactDefault {
	_tx, ok := cd.config.driver.(*txDriver)
	if !ok {
		panic("ent: ContactDefault is not a transactional entity")
	}
	cd.config.driver = _tx.drv
	return cd
}

// String implements the fmt.Stringer.
func (cd *ContactDefault) String() string {
	var builder strings.Builder
	builder.WriteString("ContactDefault(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cd.ID))
	if v := cd.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := cd.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := cd.URL; v != nil {
		builder.WriteString("url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := cd.Category; v != nil {
		builder.WriteString("category=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := cd.CustomTitle; v != nil {
		builder.WriteString("custom_title=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := cd.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(cd.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(cd.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ContactDefaults is a parsable slice of ContactDefault.
type ContactDefaults []*ContactDefault
