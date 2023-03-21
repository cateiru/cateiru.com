// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru.com/ent/contact"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// ToUserID holds the value of the "to_user_id" field.
	ToUserID uint32 `json:"to_user_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// Mail holds the value of the "mail" field.
	Mail string `json:"mail,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Lang holds the value of the "lang" field.
	Lang string `json:"lang,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// CustomTitle holds the value of the "custom_title" field.
	CustomTitle string `json:"custom_title,omitempty"`
	// CustomValue holds the value of the "custom_value" field.
	CustomValue string `json:"custom_value,omitempty"`
	// DeviceName holds the value of the "device_name" field.
	DeviceName string `json:"device_name,omitempty"`
	// Os holds the value of the "os" field.
	Os string `json:"os,omitempty"`
	// BrowserName holds the value of the "browser_name" field.
	BrowserName string `json:"browser_name,omitempty"`
	// IsMobile holds the value of the "is_mobile" field.
	IsMobile bool `json:"is_mobile,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldIsMobile:
			values[i] = new(sql.NullBool)
		case contact.FieldID, contact.FieldToUserID:
			values[i] = new(sql.NullInt64)
		case contact.FieldName, contact.FieldTitle, contact.FieldDetail, contact.FieldMail, contact.FieldIP, contact.FieldLang, contact.FieldURL, contact.FieldCategory, contact.FieldCustomTitle, contact.FieldCustomValue, contact.FieldDeviceName, contact.FieldOs, contact.FieldBrowserName:
			values[i] = new(sql.NullString)
		case contact.FieldCreated, contact.FieldModified:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []any) error {
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
			c.ID = uint32(value.Int64)
		case contact.FieldToUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to_user_id", values[i])
			} else if value.Valid {
				c.ToUserID = uint32(value.Int64)
			}
		case contact.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case contact.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case contact.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				c.Detail = value.String
			}
		case contact.FieldMail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mail", values[i])
			} else if value.Valid {
				c.Mail = value.String
			}
		case contact.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				c.IP = value.String
			}
		case contact.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				c.Lang = value.String
			}
		case contact.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				c.URL = value.String
			}
		case contact.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				c.Category = value.String
			}
		case contact.FieldCustomTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field custom_title", values[i])
			} else if value.Valid {
				c.CustomTitle = value.String
			}
		case contact.FieldCustomValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field custom_value", values[i])
			} else if value.Valid {
				c.CustomValue = value.String
			}
		case contact.FieldDeviceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_name", values[i])
			} else if value.Valid {
				c.DeviceName = value.String
			}
		case contact.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				c.Os = value.String
			}
		case contact.FieldBrowserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser_name", values[i])
			} else if value.Valid {
				c.BrowserName = value.String
			}
		case contact.FieldIsMobile:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_mobile", values[i])
			} else if value.Valid {
				c.IsMobile = value.Bool
			}
		case contact.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				c.Created = value.Time
			}
		case contact.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				c.Modified = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return NewContactClient(c.config).UpdateOne(c)
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
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("to_user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ToUserID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(c.Detail)
	builder.WriteString(", ")
	builder.WriteString("mail=")
	builder.WriteString(c.Mail)
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(c.IP)
	builder.WriteString(", ")
	builder.WriteString("lang=")
	builder.WriteString(c.Lang)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(c.URL)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(c.Category)
	builder.WriteString(", ")
	builder.WriteString("custom_title=")
	builder.WriteString(c.CustomTitle)
	builder.WriteString(", ")
	builder.WriteString("custom_value=")
	builder.WriteString(c.CustomValue)
	builder.WriteString(", ")
	builder.WriteString("device_name=")
	builder.WriteString(c.DeviceName)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(c.Os)
	builder.WriteString(", ")
	builder.WriteString("browser_name=")
	builder.WriteString(c.BrowserName)
	builder.WriteString(", ")
	builder.WriteString("is_mobile=")
	builder.WriteString(fmt.Sprintf("%v", c.IsMobile))
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(c.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(c.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact
