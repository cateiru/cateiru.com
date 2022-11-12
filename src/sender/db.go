package sender

import (
	"context"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/db"
)

func (f *SendForm) InsertDB(ctx context.Context, db *db.DB, userId uint32) (*ent.Contact, error) {
	contact := db.Client.Contact.Create().
		SetToUserID(userId).
		SetName(f.Name).
		SetTitle(f.Subject).
		SetDetail(f.Detail).
		SetMail(f.Mail).
		SetIP(f.IP).
		SetLang(f.Lang).
		SetIsMobile(f.UserData.IsMobile)

	if f.URL != "" {
		contact = contact.SetURL(f.URL)
	}
	if f.Category != "" {
		contact = contact.SetCategory(f.Category)
	}
	if f.CustomTitle != "" && f.CustomValue != "" {
		contact = contact.
			SetCustomTitle(f.CustomTitle).
			SetCustomValue(f.CustomValue)
	}
	if f.UserData.Device != "" {
		contact = contact.SetDeviceName(f.UserData.Device)
	}
	if f.UserData.Browser != "" {
		contact = contact.SetBrowserName(f.UserData.Browser)
	}
	if f.UserData.OS != "" {
		contact = contact.SetOs(f.UserData.OS)
	}

	return contact.Save(ctx)
}
