package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/ent/link"
	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/cateiru/cateiru.com/ent/product"
	"github.com/cateiru/cateiru.com/ent/user"
	"github.com/labstack/echo/v4"
)

type Public struct {
	// User data
	GivenName    string    `json:"given_name,omitempty"`
	FamilyName   string    `json:"family_name,omitempty"`
	GivenNameJa  string    `json:"given_name_ja,omitempty"`
	FamilyNameJa string    `json:"family_name_ja,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
	BirthDate    time.Time `json:"birth_date,omitempty"`
	Location     string    `json:"location,omitempty"`
	LocationJa   string    `json:"location_ja,omitempty"`
	AvatarURL    string    `json:"avatar_url,omitempty"`
	Created      time.Time `json:"created,omitempty"`
	Modified     time.Time `json:"modified,omitempty"`

	Biographies []PublicBioGraphy    `json:"biographies"`
	Products    []PublicShortProduct `json:"products"`
	Links       []PublicLinkCategory `json:"links"`
}

type PublicBioGraphy struct {
	Position   string    `json:"position,omitempty"`
	PositionJa string    `json:"position_ja,omitempty"`
	Join       time.Time `json:"join,omitempty"`
	Leave      time.Time `json:"leave,omitempty"`

	PublicLocation
}

type PublicLocation struct {
	Type      location.Type `json:"type,omitempty"`
	Name      string        `json:"name,omitempty"`
	NameJa    string        `json:"name_ja,omitempty"`
	Address   string        `json:"address,omitempty"`
	AddressJa string        `json:"address_ja,omitempty"`
}

type PublicShortProduct struct {
	ID        uint32    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NameJa    string    `json:"name_ja,omitempty"`
	Detail    string    `json:"detail,omitempty"`
	DetailJa  string    `json:"detail_ja,omitempty"`
	DevTime   time.Time `json:"dev_time,omitempty"`
	Thumbnail string    `json:"thumbnail,omitempty"`
	GithubURL string    `json:"github_url,omitempty"`
	SiteURL   string    `json:"site_url"`
}

type PublicLinkCategory struct {
	CategoryId     uint32 `json:"category_id"`
	CategoryName   string `json:"category_name,omitempty"`
	CategoryNameJa string `json:"category_name_ja,omitempty"`
	Emoji          string `json:"emoji,omitempty"`

	Links []PublicLink `json:"links"`
}

type PublicLink struct {
	Name       string `json:"name,omitempty"`
	NameJa     string `json:"name_ja,omitempty"`
	SiteURL    string `json:"site_url,omitempty"`
	FaviconURL string `json:"favicon_url,omitempty"`
}

type PublicProduct struct {
	ID        uint32    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NameJa    string    `json:"name_ja,omitempty"`
	Detail    string    `json:"detail,omitempty"`
	DetailJa  string    `json:"detail_ja,omitempty"`
	SiteURL   string    `json:"site_url,omitempty"`
	GithubURL string    `json:"github_url,omitempty"`
	DevTime   time.Time `json:"dev_time,omitempty"`
	Thumbnail string    `json:"thumbnail,omitempty"`
}

// Response public profiles
func (h *Handler) PublicProfileHandler(e echo.Context) error {
	ctx := e.Request().Context()

	u, err := h.DB.Client.User.Query().Where(user.Selected(true)).First(ctx)
	if _, ok := err.(*ent.NotFoundError); ok {
		return echo.ErrNotFound
	}
	if err != nil {
		return err
	}

	// Get Biography
	bios, err := h.DB.Client.Biography.
		Query().
		Where(biography.And(biography.UserID(u.ID), biography.IsPublic(true))).
		Order(ent.OrderFunc(ent.Asc(biography.FieldJoin))).
		All(ctx)
	if err != nil {
		return err
	}
	var locationMap map[uint32]*ent.Location
	publicBios := make([]PublicBioGraphy, len(bios))
	for i, bio := range bios {
		var loc *ent.Location
		if locationMap[bio.LocationID] != nil {
			loc = locationMap[bio.LocationID]
		} else {
			loc, err = h.DB.Client.Location.Get(ctx, bio.LocationID)
			if err != nil {
				return err
			}
		}
		publicBios[i] = PublicBioGraphy{
			Position:   bio.Position,
			PositionJa: bio.PositionJa,
			Join:       bio.Join,
			Leave:      bio.Leave,

			PublicLocation: PublicLocation{
				Type:      loc.Type,
				Name:      loc.Name,
				NameJa:    loc.NameJa,
				Address:   loc.Address,
				AddressJa: loc.AddressJa,
			},
		}
	}

	// Get Products
	prods, err := h.DB.Client.Product.
		Query().
		Where(product.UserID(u.ID)).
		Order(ent.Desc(product.FieldDevTime)).
		All(ctx)
	if err != nil {
		return err
	}

	publicProds := make([]PublicShortProduct, len(prods))
	for i, prod := range prods {
		publicProds[i] = PublicShortProduct{
			ID:        prod.ID,
			Name:      prod.Name,
			NameJa:    prod.NameJa,
			Detail:    prod.Detail,
			DetailJa:  prod.DetailJa,
			DevTime:   prod.DevTime,
			Thumbnail: prod.Thumbnail,
			GithubURL: prod.GithubURL,
			SiteURL:   prod.SiteURL,
		}
	}

	// Get Links
	links, err := h.DB.Client.Link.
		Query().
		Where(link.UserID(u.ID)).
		Order(ent.Asc(link.FieldCategoryID)).
		Order(ent.Asc(link.FieldName)).
		All(ctx)
	if err != nil {
		return err
	}
	// var categoryMap map[uint32]int{}
	publicLinkCategories := []PublicLinkCategory{}
	for _, l := range links {

		var categoriesIndex int = 0
		var isFind bool = false
		for i, c := range publicLinkCategories {
			if c.CategoryId == l.CategoryID {
				isFind = true
				categoriesIndex = i
			}
		}

		if isFind {
			publicLinkCategories[categoriesIndex].Links = append(publicLinkCategories[categoriesIndex].Links, PublicLink{
				Name:       l.Name,
				NameJa:     l.NameJa,
				SiteURL:    l.SiteURL,
				FaviconURL: l.FaviconURL,
			})
		} else {
			category, err := h.DB.Client.Category.Get(ctx, l.CategoryID)
			if err != nil {
				return err
			}

			publicLinkCategories = append(publicLinkCategories, PublicLinkCategory{
				CategoryId:     category.ID,
				CategoryName:   category.Name,
				CategoryNameJa: category.NameJa,
				Emoji:          category.Emoji,

				Links: []PublicLink{{
					Name:       l.Name,
					NameJa:     l.NameJa,
					SiteURL:    l.SiteURL,
					FaviconURL: l.FaviconURL,
				}},
			})
		}
	}

	publicUser := Public{
		GivenName:    u.GivenName,
		FamilyName:   u.FamilyName,
		GivenNameJa:  u.GivenNameJa,
		FamilyNameJa: u.FamilyNameJa,
		UserID:       u.UserID,
		BirthDate:    u.BirthDate,
		Location:     u.Location,
		LocationJa:   u.LocationJa,
		AvatarURL:    u.AvatarURL,
		Created:      u.Created,
		Modified:     u.Modified,

		Biographies: publicBios,

		Products: publicProds,

		Links: publicLinkCategories,
	}

	return e.JSON(http.StatusOK, publicUser)
}

func (h *Handler) PublicProductsHandler(e echo.Context) error {
	ctx := e.Request().Context()

	productIdStr := e.QueryParam("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid query: product_id")
	}

	p, err := h.DB.Client.Product.
		Query().
		Where(product.ID(uint32(productId))).
		First(ctx)
	if _, ok := err.(*ent.NotFoundError); ok {
		return echo.ErrNotFound
	}
	if err != nil {
		return err
	}

	publicProduct := PublicProduct{
		ID:        p.ID,
		Name:      p.Name,
		NameJa:    p.NameJa,
		Detail:    p.Detail,
		DetailJa:  p.DetailJa,
		SiteURL:   p.SiteURL,
		GithubURL: p.GithubURL,
		DevTime:   p.DevTime,
		Thumbnail: p.Thumbnail,
	}

	return e.JSON(http.StatusOK, publicProduct)
}
