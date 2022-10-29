package handler

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/cateiru/cateiru.com/ent/category"
	"github.com/cateiru/cateiru.com/ent/link"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/labstack/echo/v4"
)

// Response all Links
func (h *Handler) LinkHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	links, err := h.DB.Client.Link.
		Query().
		Where(link.UserID(h.User.ID)).
		All(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, links)
}

// Create a new Link
//
// Require form data
// - name
// - name_ja
// - site_url
// - category_id
func (h *Handler) CreateLinkHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	name := e.FormValue("name")
	nameJa := e.FormValue("name_ja")
	siteUrl := e.FormValue("site_url")
	categoryIdStr := e.FormValue("category_id")

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: name")
	}
	if nameJa == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: name_ja")
	}
	if siteUrl == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: site_url")
	}
	if categoryIdStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: category_id")
	}

	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: category_id")
	}

	if err := ValidateURL(siteUrl); err != nil {
		return err
	}
	if err := existCategoryId(ctx, h.DB, uint32(categoryId)); err != nil {
		return err
	}

	favicon, err := GetFavicon(siteUrl)
	if err != nil {
		return err
	}

	link, err := h.DB.Client.Link.
		Create().
		SetName(name).
		SetNameJa(nameJa).
		SetSiteURL(siteUrl).
		SetCategoryID(uint32(categoryId)).
		SetFaviconURL(favicon).
		Save(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, link)
}

func (h *Handler) UpdateLinkHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteLinkHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func existCategoryId(ctx context.Context, db *db.DB, categoryId uint32) error {
	exist, err := db.Client.Category.
		Query().
		Where(category.ID(categoryId)).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: not exist category")
	}
	return nil
}

// Search favicon
//
// - HTML parse
//   - get `link[rel=icon]` elements
//
// - GET `/favicon.ico`
//   - if not found in HTML
func GetFavicon(siteUrl string) (string, error) {
	resp, err := http.Get(siteUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", echo.NewHTTPError(http.StatusBadRequest, "site_url is invalid link")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	query := doc.Find("link[rel=icon]")
	favicons := []string{}
	query.Each(func(i int, s *goquery.Selection) {
		favicon, exist := s.Attr("href")
		if exist {
			favicons = append(favicons, favicon)
		}
	})

	if len(favicons) == 0 {
		// If not found favicon in HTML, request `/favicon.ico` path.
		u, err := url.Parse(siteUrl)
		if err != nil {
			return "", err
		}
		u.Path = "/favicon.ico"

		favicons = append(favicons, u.String())
	}

	for _, f := range favicons {
		reqF, err := http.Get(f)
		if err != nil {
			return "", err
		}
		defer reqF.Body.Close()
		if reqF.StatusCode == 200 {
			return f, nil
		}
	}

	return "", echo.NewHTTPError(http.StatusBadRequest, "favicon is not found")
}
