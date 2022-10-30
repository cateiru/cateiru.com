package test

import (
	"context"
	"fmt"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/db"
)

type TestNotice struct {
	UserId         uint32
	DiscordWebhook string
	SlackWebhook   string
	Mail           string

	Notice *ent.Notice
}

func NewTestNotice(userId uint32) (*TestNotice, error) {
	r, err := MakeRandomStr(10)
	if err != nil {
		return nil, err
	}

	return &TestNotice{
		UserId:         userId,
		DiscordWebhook: fmt.Sprintf("https://webhook.discord.com/%s", r),
		SlackWebhook:   fmt.Sprintf("https://webhook.slack.com/%s", r),
		Mail:           fmt.Sprintf("%s@cateiru.com", r),
	}, nil
}

func (c *TestNotice) CreateDB(ctx context.Context, db *db.DB) (*ent.Notice, error) {
	n, err := db.Client.Notice.Create().
		SetUserID(c.UserId).
		SetDiscordWebhook(c.DiscordWebhook).
		SetSlackWebhook(c.SlackWebhook).
		SetMail(c.Mail).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	c.Notice = n

	return n, nil
}
