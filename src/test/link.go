package test

import (
	"context"
	"fmt"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/db"
)

type TestLink struct {
	UserId     uint32
	Name       string
	NameJa     string
	SiteURL    string
	FaviconURL string

	CategoryId uint32

	CategoryName   string
	CategoryNameJa string
	Emoji          string

	Link     *ent.Link
	Category *ent.Category
}

const UNDEFINED_CATEGORY_ID uint32 = 1000000000

func NewTestLink(userId uint32) (*TestLink, error) {
	name, err := MakeRandomStr(5)
	if err != nil {
		return nil, err
	}
	categoryName, err := MakeRandomStr(5)
	if err != nil {
		return nil, err
	}

	return &TestLink{
		UserId:     userId,
		Name:       name,
		NameJa:     name,
		SiteURL:    fmt.Sprintf("https://cateiru.com/%s", name),
		FaviconURL: fmt.Sprintf("https://cateiru.com/img/%s", name),

		CategoryId: UNDEFINED_CATEGORY_ID,

		CategoryName:   categoryName,
		CategoryNameJa: categoryName,
		Emoji:          "🍣",
	}, nil
}

func (c *TestLink) CreateDB(ctx context.Context, db *db.DB) (*ent.Link, error) {
	if c.Category != nil {
		c.CategoryId = c.Category.ID
	} else if c.CategoryId == UNDEFINED_CATEGORY_ID {
		if _, err := c.CreateCategoryDB(ctx, db); err != nil {
			return nil, err
		}
	}

	link, err := db.Client.Link.
		Create().
		SetUserID(c.UserId).
		SetName(c.Name).
		SetNameJa(c.NameJa).
		SetSiteURL(c.SiteURL).
		SetFaviconURL(c.FaviconURL).
		SetCategoryID(c.CategoryId).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	c.Link = link
	return link, nil
}

func (c *TestLink) CreateCategoryDB(ctx context.Context, db *db.DB) (*ent.Category, error) {
	category, err := db.Client.Category.
		Create().
		SetName(c.CategoryName).
		SetNameJa(c.CategoryNameJa).
		SetEmoji(c.Emoji).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	c.CategoryId = category.ID
	c.Category = category

	return category, nil
}
