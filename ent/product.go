// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateir.com/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint32 `json:"user_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// NameJa holds the value of the "name_ja" field.
	NameJa string `json:"name_ja,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// DetailJa holds the value of the "detail_ja" field.
	DetailJa string `json:"detail_ja,omitempty"`
	// SiteURL holds the value of the "site_url" field.
	SiteURL string `json:"site_url,omitempty"`
	// GithubURL holds the value of the "github_url" field.
	GithubURL string `json:"github_url,omitempty"`
	// DevTime holds the value of the "dev_time" field.
	DevTime time.Time `json:"dev_time,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldID, product.FieldUserID:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldNameJa, product.FieldDetail, product.FieldDetailJa, product.FieldSiteURL, product.FieldGithubURL:
			values[i] = new(sql.NullString)
		case product.FieldDevTime, product.FieldCreated, product.FieldModified:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint32(value.Int64)
		case product.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pr.UserID = uint32(value.Int64)
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				pr.NameJa = value.String
			}
		case product.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				pr.Detail = value.String
			}
		case product.FieldDetailJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail_ja", values[i])
			} else if value.Valid {
				pr.DetailJa = value.String
			}
		case product.FieldSiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field site_url", values[i])
			} else if value.Valid {
				pr.SiteURL = value.String
			}
		case product.FieldGithubURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field github_url", values[i])
			} else if value.Valid {
				pr.GithubURL = value.String
			}
		case product.FieldDevTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dev_time", values[i])
			} else if value.Valid {
				pr.DevTime = value.Time
			}
		case product.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				pr.Created = value.Time
			}
		case product.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				pr.Modified = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.UserID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(pr.NameJa)
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(pr.Detail)
	builder.WriteString(", ")
	builder.WriteString("detail_ja=")
	builder.WriteString(pr.DetailJa)
	builder.WriteString(", ")
	builder.WriteString("site_url=")
	builder.WriteString(pr.SiteURL)
	builder.WriteString(", ")
	builder.WriteString("github_url=")
	builder.WriteString(pr.GithubURL)
	builder.WriteString(", ")
	builder.WriteString("dev_time=")
	builder.WriteString(pr.DevTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(pr.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(pr.Modified.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
