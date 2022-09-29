package test

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/db"
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

	User    *ent.User
	Session *ent.Session
}

func NewUser() (*TestUser, error) {
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
		Save(ctx)

	if err != nil {
		return err
	}

	c.User = u
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

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
