package test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/google/uuid"
)

type TestUser struct {
	GivenName    string
	FamilyName   string
	UserId       string
	Mail         string
	BirthDate    time.Time
	GivenNameJa  string
	FamilyNameJa string
	Location     string
	LocationJa   string
	AvatarURL    string
	SSOToken     string

	User    *ent.User
	Session *ent.Session
}

func NewUser(d ...ent.User) (*TestUser, error) {
	if len(d) == 1 {
		u := d[0]
		return &TestUser{
			GivenName:    u.GivenName,
			FamilyName:   u.FamilyName,
			UserId:       u.UserID,
			Mail:         u.Mail,
			BirthDate:    u.BirthDate,
			GivenNameJa:  u.GivenNameJa,
			FamilyNameJa: u.FamilyNameJa,
			Location:     u.Location,
			LocationJa:   u.LocationJa,
			AvatarURL:    u.AvatarURL,
			SSOToken:     u.SSOToken,

			User:    nil,
			Session: nil,
		}, nil
	}

	userId, err := MakeRandomStr(10)
	if err != nil {
		return nil, err
	}

	return &TestUser{
		GivenName:    "iii",
		FamilyName:   "aaa",
		UserId:       userId,
		Mail:         fmt.Sprintf("%s@cateiru.com", userId),
		BirthDate:    time.Now(),
		GivenNameJa:  "いいい",
		FamilyNameJa: "あああ",
		Location:     "hogehoge",
		LocationJa:   "ほげほげ",
		AvatarURL:    fmt.Sprintf("https://%s.cateiru.com", userId),
		SSOToken:     userId,

		User:    nil,
		Session: nil,
	}, nil
}

// Insert from User Databases
func (c *TestUser) CreateDB(ctx context.Context, db *db.DB) error {
	u, err := db.Client.User.Create().
		SetGivenName(c.GivenName).
		SetFamilyName(c.FamilyName).
		SetUserID(c.UserId).
		SetMail(c.Mail).
		SetBirthDate(c.BirthDate).
		SetGivenNameJa(c.GivenNameJa).
		SetFamilyNameJa(c.FamilyNameJa).
		SetLocation(c.Location).
		SetLocationJa(c.LocationJa).
		SetSSOToken(c.SSOToken).
		Save(ctx)

	if err != nil {
		return err
	}

	c.User = u
	return nil
}

func (c *TestUser) HandlerSession(ctx context.Context, db *db.DB, m *mock.MockHandler) error {
	sessionToken, err := c.CreateSession(ctx, db)
	if err != nil {
		return err
	}

	sessionCookie := &http.Cookie{
		Name:  config.Config.SessionCookieName,
		Value: sessionToken.String(),
	}
	checkCookie := &http.Cookie{
		Name:  config.Config.SessionConfirmationCookieName,
		Value: "true",
	}
	m.Cookie([]*http.Cookie{sessionCookie, checkCookie})

	return nil
}

// Create user session
// returns uuid is session token
func (c *TestUser) CreateSession(ctx context.Context, db *db.DB) (uuid.UUID, error) {
	if c.User == nil {
		return uuid.UUID{}, errors.New("user is not inserted in db")
	}

	s, err := db.Client.Session.Create().
		SetUserID(c.User.ID).
		SetPeriod(time.Now().Add(config.Config.DBSessionTime)).
		Save(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	c.Session = s

	return s.ID, nil
}

func (c *TestUser) SelectStatus(ctx context.Context, s bool) error {
	return c.User.Update().SetSelected(s).Exec(ctx)
}

func (c *TestUser) CreateBio() (*TestBio, error) {
	if c.User == nil {
		return nil, errors.New("user is not inserted in db")
	}

	return NewTestBio(c.User.ID)
}

func (c *TestUser) CreateProduct() (*TestProduct, error) {
	if c.User == nil {
		return nil, errors.New("user is not inserted in db")
	}

	return NewTestProduct(c.User.ID)
}

func (c *TestUser) CreateLink() (*TestLink, error) {
	if c.User == nil {
		return nil, errors.New("user is not inserted in db")
	}

	return NewTestLink(c.User.ID)
}

func (c *TestUser) CreateNotice() (*TestNotice, error) {
	if c.User == nil {
		return nil, errors.New("user is not inserted in db")
	}

	return NewTestNotice(c.User.ID)
}
