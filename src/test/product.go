package test

import (
	"context"
	"fmt"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/db"
)

type TestProduct struct {
	UserId    uint32
	Name      string
	NameJa    string
	Detail    string
	DetailJa  string
	SiteURL   string
	GithubURL string
	DevTime   time.Time
	Thumbnail string

	Product *ent.Product
}

func NewTestProduct(userId uint32) (*TestProduct, error) {
	name, err := MakeRandomStr(5)
	if err != nil {
		return nil, err
	}
	detail, err := MakeRandomStr(10)
	if err != nil {
		return nil, err
	}

	return &TestProduct{
		UserId:    userId,
		Name:      name,
		NameJa:    name,
		Detail:    detail,
		DetailJa:  detail,
		SiteURL:   fmt.Sprintf("https://cateiru.com/%s", name),
		GithubURL: fmt.Sprintf("https://github.com/cateiru/%s", name),
		DevTime:   time.Now(),
		Thumbnail: fmt.Sprintf("https://cateiru.com/%s", name),

		Product: nil,
	}, nil
}

func (c *TestProduct) CreateDB(ctx context.Context, db *db.DB) (*ent.Product, error) {
	product, err := db.Client.Product.Create().
		SetUserID(c.UserId).
		SetName(c.Name).
		SetNameJa(c.NameJa).
		SetDetail(c.Detail).
		SetDetailJa(c.DetailJa).
		SetSiteURL(c.SiteURL).
		SetGithubURL(c.GithubURL).
		SetDevTime(c.DevTime).
		SetThumbnail(c.Thumbnail).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	c.Product = product

	return product, nil
}
