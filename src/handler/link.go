package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/category"
	"github.com/cateiru/cateiru.com/ent/link"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/labstack/echo/v4"
)

type LinkResponse struct {
	Link     ent.Link     `json:"link"`
	Category ent.Category `json:"category"`
}

// Response all Links
func (h *Handler) LinkHandler(e echo.Context) error {
	ctx := e.Request().Context()

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

	resp := []LinkResponse{}
	categories := map[uint32]*ent.Category{}
	for _, l := range links {
		c, ok := categories[l.CategoryID]
		if ok {
			resp = append(resp, LinkResponse{
				Link:     *l,
				Category: *c,
			})
		} else {
			newC, err := h.DB.Client.Category.
				Query().
				Where(category.ID(l.CategoryID)).
				First(ctx)
			if err != nil {
				return err
			}
			categories[newC.ID] = newC

			resp = append(resp, LinkResponse{
				Link:     *l,
				Category: *newC,
			})
		}
	}

	return e.JSON(http.StatusOK, resp)
}

// Create a new Link
//
// Require form data
// - name
// - name_ja
// - site_url
// - category_id
func (h *Handler) CreateLinkHandler(e echo.Context) error {
	ctx := e.Request().Context()

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

	linkSch := h.DB.Client.Link.
		Create().
		SetUserID(h.User.ID).
		SetName(name).
		SetNameJa(nameJa).
		SetSiteURL(siteUrl).
		SetCategoryID(uint32(categoryId))

	if favicon != "" {
		linkSch = linkSch.SetFaviconURL(favicon)
	}

	link, err := linkSch.Save(ctx)
	if err != nil {
		return err
	}

	c, err := h.DB.Client.Category.Get(ctx, uint32(categoryId))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, LinkResponse{
		Link:     *link,
		Category: *c,
	})
}

func (h *Handler) UpdateLinkHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.FormValue("link_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: link_id")
	}

	l := h.DB.Client.Link.
		Update().
		Where(link.And(
			link.UserID(h.User.ID),
			link.ID(uint32(id)),
		))
	changed := false

	name := e.FormValue("name")
	if name != "" {
		l = l.SetName(name)
		changed = true
	}
	nameJa := e.FormValue("name_ja")
	if nameJa != "" {
		l = l.SetNameJa(nameJa)
		changed = true
	}
	siteUrl := e.FormValue("site_url")
	if siteUrl != "" {
		if err := ValidateURL(siteUrl); err != nil {
			return err
		}
		favicon, err := GetFavicon(siteUrl)
		if err != nil {
			return err
		}
		if favicon != "" {
			l = l.SetSiteURL(siteUrl).SetFaviconURL(favicon)
		}
		changed = true
	}
	categoryIdStr := e.FormValue("category_id")
	if categoryIdStr != "" {
		categoryId, err := strconv.Atoi(categoryIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: category_id")
		}

		if err := existCategoryId(ctx, h.DB, uint32(categoryId)); err != nil {
			return err
		}

		l = l.SetCategoryID(uint32(categoryId))
		changed = true
	}

	if !changed {
		return echo.NewHTTPError(http.StatusBadRequest, "no changes")
	}

	if err := l.Exec(ctx); err != nil {
		return err
	}

	linkDB, err := h.DB.Client.Link.
		Query().
		Where(link.And(
			link.UserID(h.User.ID),
			link.ID(uint32(id)),
		)).
		First(ctx)
	if err != nil {
		return err
	}
	category, err := h.DB.Client.Category.
		Query().
		Where(category.ID(linkDB.CategoryID)).
		First(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, LinkResponse{
		Link:     *linkDB,
		Category: *category,
	})
}

func (h *Handler) DeleteLinkHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.QueryParam("link_id")
	if idStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invlaid form: link_id")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invlaid form: link_id")
	}

	_, err = h.DB.Client.Link.
		Delete().
		Where(link.And(
			link.UserID(h.User.ID),
			link.ID(uint32(id)),
		)).
		Exec(ctx)
	if err != nil {
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

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	query := doc.Find(`link[rel="shortcut icon"], link[rel="icon"]`)
	favicons := []string{}
	query.Each(func(i int, s *goquery.Selection) {
		favicon, exist := s.Attr("href")
		if exist {
			favicons = append(favicons, favicon)
		}
	})

	u, err := url.Parse(siteUrl)
	if err != nil {
		return "", err
	}

	if len(favicons) == 0 {
		// If not found favicon in HTML, request `/favicon.ico` path.
		u.Path = "/favicon.ico"

		favicons = append(favicons, u.String())
	}

	for _, f := range favicons {
		favUrl := ""
		if err := ValidateURL(f); err != nil {
			favUrl = fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, f)
		} else {
			favUrl = f
		}

		reqF, err := http.Get(favUrl)
		if err != nil {
			return "", err
		}
		defer reqF.Body.Close()
		if reqF.StatusCode == 200 {
			return favUrl, nil
		}
	}

	return "", nil
}
