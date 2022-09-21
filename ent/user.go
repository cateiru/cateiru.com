// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru.com/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// GivenName holds the value of the "given_name" field.
	GivenName string `json:"given_name,omitempty"`
	// FamilyName holds the value of the "family_name" field.
	FamilyName string `json:"family_name,omitempty"`
	// GivenNameJa holds the value of the "given_name_ja" field.
	GivenNameJa string `json:"given_name_ja,omitempty"`
	// FamilyNameJa holds the value of the "family_name_ja" field.
	FamilyNameJa string `json:"family_name_ja,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Mail holds the value of the "mail" field.
	Mail string `json:"mail,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate time.Time `json:"birth_date,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// LocationJa holds the value of the "location_ja" field.
	LocationJa string `json:"location_ja,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldGivenName, user.FieldFamilyName, user.FieldGivenNameJa, user.FieldFamilyNameJa, user.FieldUserID, user.FieldMail, user.FieldLocation, user.FieldLocationJa:
			values[i] = new(sql.NullString)
		case user.FieldBirthDate, user.FieldCreated, user.FieldModified:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint32(value.Int64)
		case user.FieldGivenName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field given_name", values[i])
			} else if value.Valid {
				u.GivenName = value.String
			}
		case user.FieldFamilyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family_name", values[i])
			} else if value.Valid {
				u.FamilyName = value.String
			}
		case user.FieldGivenNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field given_name_ja", values[i])
			} else if value.Valid {
				u.GivenNameJa = value.String
			}
		case user.FieldFamilyNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family_name_ja", values[i])
			} else if value.Valid {
				u.FamilyNameJa = value.String
			}
		case user.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				u.UserID = value.String
			}
		case user.FieldMail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mail", values[i])
			} else if value.Valid {
				u.Mail = value.String
			}
		case user.FieldBirthDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				u.BirthDate = value.Time
			}
		case user.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				u.Location = value.String
			}
		case user.FieldLocationJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location_ja", values[i])
			} else if value.Valid {
				u.LocationJa = value.String
			}
		case user.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				u.Created = value.Time
			}
		case user.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				u.Modified = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("given_name=")
	builder.WriteString(u.GivenName)
	builder.WriteString(", ")
	builder.WriteString("family_name=")
	builder.WriteString(u.FamilyName)
	builder.WriteString(", ")
	builder.WriteString("given_name_ja=")
	builder.WriteString(u.GivenNameJa)
	builder.WriteString(", ")
	builder.WriteString("family_name_ja=")
	builder.WriteString(u.FamilyNameJa)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(u.UserID)
	builder.WriteString(", ")
	builder.WriteString("mail=")
	builder.WriteString(u.Mail)
	builder.WriteString(", ")
	builder.WriteString("birth_date=")
	builder.WriteString(u.BirthDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(u.Location)
	builder.WriteString(", ")
	builder.WriteString("location_ja=")
	builder.WriteString(u.LocationJa)
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(u.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(u.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}