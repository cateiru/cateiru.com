// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/ent/category"
	"github.com/cateiru/cateiru.com/ent/contact"
	"github.com/cateiru/cateiru.com/ent/contactdefault"
	"github.com/cateiru/cateiru.com/ent/link"
	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/cateiru/cateiru.com/ent/notice"
	"github.com/cateiru/cateiru.com/ent/product"
	"github.com/cateiru/cateiru.com/ent/schema"
	"github.com/cateiru/cateiru.com/ent/session"
	"github.com/cateiru/cateiru.com/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	biographyFields := schema.Biography{}.Fields()
	_ = biographyFields
	// biographyDescIsPublic is the schema descriptor for is_public field.
	biographyDescIsPublic := biographyFields[2].Descriptor()
	// biography.DefaultIsPublic holds the default value on creation for the is_public field.
	biography.DefaultIsPublic = biographyDescIsPublic.Default.(bool)
	// biographyDescCreated is the schema descriptor for created field.
	biographyDescCreated := biographyFields[8].Descriptor()
	// biography.DefaultCreated holds the default value on creation for the created field.
	biography.DefaultCreated = biographyDescCreated.Default.(func() time.Time)
	// biographyDescModified is the schema descriptor for modified field.
	biographyDescModified := biographyFields[9].Descriptor()
	// biography.DefaultModified holds the default value on creation for the modified field.
	biography.DefaultModified = biographyDescModified.Default.(func() time.Time)
	// biography.UpdateDefaultModified holds the default value on update for the modified field.
	biography.UpdateDefaultModified = biographyDescModified.UpdateDefault.(func() time.Time)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreated is the schema descriptor for created field.
	categoryDescCreated := categoryFields[4].Descriptor()
	// category.DefaultCreated holds the default value on creation for the created field.
	category.DefaultCreated = categoryDescCreated.Default.(func() time.Time)
	// categoryDescModified is the schema descriptor for modified field.
	categoryDescModified := categoryFields[5].Descriptor()
	// category.DefaultModified holds the default value on creation for the modified field.
	category.DefaultModified = categoryDescModified.Default.(func() time.Time)
	// category.UpdateDefaultModified holds the default value on update for the modified field.
	category.UpdateDefaultModified = categoryDescModified.UpdateDefault.(func() time.Time)
	contactFields := schema.Contact{}.Fields()
	_ = contactFields
	// contactDescCreated is the schema descriptor for created field.
	contactDescCreated := contactFields[16].Descriptor()
	// contact.DefaultCreated holds the default value on creation for the created field.
	contact.DefaultCreated = contactDescCreated.Default.(func() time.Time)
	// contactDescModified is the schema descriptor for modified field.
	contactDescModified := contactFields[17].Descriptor()
	// contact.DefaultModified holds the default value on creation for the modified field.
	contact.DefaultModified = contactDescModified.Default.(func() time.Time)
	// contact.UpdateDefaultModified holds the default value on update for the modified field.
	contact.UpdateDefaultModified = contactDescModified.UpdateDefault.(func() time.Time)
	contactdefaultFields := schema.ContactDefault{}.Fields()
	_ = contactdefaultFields
	// contactdefaultDescCreated is the schema descriptor for created field.
	contactdefaultDescCreated := contactdefaultFields[7].Descriptor()
	// contactdefault.DefaultCreated holds the default value on creation for the created field.
	contactdefault.DefaultCreated = contactdefaultDescCreated.Default.(func() time.Time)
	// contactdefaultDescModified is the schema descriptor for modified field.
	contactdefaultDescModified := contactdefaultFields[8].Descriptor()
	// contactdefault.DefaultModified holds the default value on creation for the modified field.
	contactdefault.DefaultModified = contactdefaultDescModified.Default.(func() time.Time)
	// contactdefault.UpdateDefaultModified holds the default value on update for the modified field.
	contactdefault.UpdateDefaultModified = contactdefaultDescModified.UpdateDefault.(func() time.Time)
	linkFields := schema.Link{}.Fields()
	_ = linkFields
	// linkDescCreated is the schema descriptor for created field.
	linkDescCreated := linkFields[7].Descriptor()
	// link.DefaultCreated holds the default value on creation for the created field.
	link.DefaultCreated = linkDescCreated.Default.(func() time.Time)
	// linkDescModified is the schema descriptor for modified field.
	linkDescModified := linkFields[8].Descriptor()
	// link.DefaultModified holds the default value on creation for the modified field.
	link.DefaultModified = linkDescModified.Default.(func() time.Time)
	// link.UpdateDefaultModified holds the default value on update for the modified field.
	link.UpdateDefaultModified = linkDescModified.UpdateDefault.(func() time.Time)
	locationFields := schema.Location{}.Fields()
	_ = locationFields
	// locationDescCreated is the schema descriptor for created field.
	locationDescCreated := locationFields[6].Descriptor()
	// location.DefaultCreated holds the default value on creation for the created field.
	location.DefaultCreated = locationDescCreated.Default.(func() time.Time)
	// locationDescModified is the schema descriptor for modified field.
	locationDescModified := locationFields[7].Descriptor()
	// location.DefaultModified holds the default value on creation for the modified field.
	location.DefaultModified = locationDescModified.Default.(func() time.Time)
	// location.UpdateDefaultModified holds the default value on update for the modified field.
	location.UpdateDefaultModified = locationDescModified.UpdateDefault.(func() time.Time)
	noticeFields := schema.Notice{}.Fields()
	_ = noticeFields
	// noticeDescCreated is the schema descriptor for created field.
	noticeDescCreated := noticeFields[5].Descriptor()
	// notice.DefaultCreated holds the default value on creation for the created field.
	notice.DefaultCreated = noticeDescCreated.Default.(func() time.Time)
	// noticeDescModified is the schema descriptor for modified field.
	noticeDescModified := noticeFields[6].Descriptor()
	// notice.DefaultModified holds the default value on creation for the modified field.
	notice.DefaultModified = noticeDescModified.Default.(func() time.Time)
	// notice.UpdateDefaultModified holds the default value on update for the modified field.
	notice.UpdateDefaultModified = noticeDescModified.UpdateDefault.(func() time.Time)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreated is the schema descriptor for created field.
	productDescCreated := productFields[10].Descriptor()
	// product.DefaultCreated holds the default value on creation for the created field.
	product.DefaultCreated = productDescCreated.Default.(func() time.Time)
	// productDescModified is the schema descriptor for modified field.
	productDescModified := productFields[11].Descriptor()
	// product.DefaultModified holds the default value on creation for the modified field.
	product.DefaultModified = productDescModified.Default.(func() time.Time)
	// product.UpdateDefaultModified holds the default value on update for the modified field.
	product.UpdateDefaultModified = productDescModified.UpdateDefault.(func() time.Time)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreated is the schema descriptor for created field.
	sessionDescCreated := sessionFields[2].Descriptor()
	// session.DefaultCreated holds the default value on creation for the created field.
	session.DefaultCreated = sessionDescCreated.Default.(func() time.Time)
	// sessionDescID is the schema descriptor for id field.
	sessionDescID := sessionFields[0].Descriptor()
	// session.DefaultID holds the default value on creation for the id field.
	session.DefaultID = sessionDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescSelected is the schema descriptor for selected field.
	userDescSelected := userFields[12].Descriptor()
	// user.DefaultSelected holds the default value on creation for the selected field.
	user.DefaultSelected = userDescSelected.Default.(bool)
	// userDescCreated is the schema descriptor for created field.
	userDescCreated := userFields[13].Descriptor()
	// user.DefaultCreated holds the default value on creation for the created field.
	user.DefaultCreated = userDescCreated.Default.(func() time.Time)
	// userDescModified is the schema descriptor for modified field.
	userDescModified := userFields[14].Descriptor()
	// user.DefaultModified holds the default value on creation for the modified field.
	user.DefaultModified = userDescModified.Default.(func() time.Time)
	// user.UpdateDefaultModified holds the default value on update for the modified field.
	user.UpdateDefaultModified = userDescModified.UpdateDefault.(func() time.Time)
}
