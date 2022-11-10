package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"

	"github.com/cateiru/cateiru.com/ent/notice"
	"github.com/cateiru/cateiru.com/ent/user"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/cateiru/cateiru.com/src/logging"
	"github.com/cateiru/cateiru.com/src/sender"
	goclienthints "github.com/cateiru/go-client-hints/v2"
)

// Contact form API
func (h *Handler) ContactHandler(e echo.Context) error {
	ctx := context.Background()

	name := e.FormValue("name")
	mail := e.FormValue("mail")
	lang := e.FormValue("lang")
	subject := e.FormValue("subject")
	detail := e.FormValue("detail")

	// Optional values
	url := e.FormValue("url")
	category := e.FormValue("category")

	customTitle := e.FormValue("custom_title")
	customValue := e.FormValue("custom_value")

	now := time.Now()

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: name")
	}
	if mail == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: mail")
	}
	if lang == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: lang")
	}
	if subject == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: subject")
	}
	if detail == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: detail")
	}

	// custom filed is pair
	if customTitle != "" {
		if customValue == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid form: custom filed")
		}
	}

	userData, err := GetUserAgent(e)
	if err != nil {
		return err
	}

	ip := e.RealIP()

	forms := &sender.SendForm{
		Name: name,
		Mail: mail,
		Lang: lang,

		Subject: subject,
		Detail:  detail,

		URL:      url,
		Category: category,

		Time: now,
		IP:   ip,

		CustomTitle: customTitle,
		CustomValue: customValue,

		UserData: userData,
	}

	return SwitchPostingService(ctx, h.DB, forms)
}

// Get user data from User-Agent or Client Hints
//
// If useable Client Hints, use it.
// Other, used User-Agent.
func GetUserAgent(e echo.Context) (*sender.UserData, error) {
	req := e.Request()
	if goclienthints.IsSupportClientHints(&req.Header) {
		ch, err := goclienthints.Parse(&req.Header)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return &sender.UserData{
			Browser:  ch.Brand.Brand,
			OS:       string(ch.Platform),
			Device:   "Unknown",
			IsMobile: ch.IsMobile,
		}, nil
	}

	ua := useragent.Parse(req.UserAgent())

	return &sender.UserData{
		Browser:  ua.Name,
		OS:       ua.OS,
		Device:   ua.Device,
		IsMobile: ua.Mobile,
	}, nil
}

func SwitchPostingService(ctx context.Context, db *db.DB, forms *sender.SendForm) error {
	// SELECT * FROM notices
	// WHERE
	// 	notices.user_id in (
	// 		SELECT users.id FORM users
	// 			WHERE users.selected = 1
	// 		);
	notices, err := db.Client.Notice.Query().Where(func(s *sql.Selector) {
		t := sql.Table(user.Table)
		p := sql.EQ(t.C(user.FieldSelected), true)
		s.Where(sql.In(s.C(notice.FieldUserID), sql.Select(t.C(user.FieldID)).From(t).Where(p)))
	}).All(ctx)
	if err != nil {
		return err
	}

	for i, n := range notices {
		logging.Sugar.Infof("Posting form. count: %d, send service: %s", i)

		senderErr := false
		if n.DiscordWebhook != "" {
			err = forms.DiscordSender(n.DiscordWebhook)
			if err != nil {
				senderErr = true
				logging.Sugar.Error(err)
			}
		}
		if n.SlackWebhook != "" {
			err = forms.SlackSender(n.SlackWebhook)
			if err != nil {
				senderErr = true
				logging.Sugar.Error(err)
			}
		}
		if n.Mail != "" {
			err = forms.MailSender(n.Mail)
			if err != nil {
				senderErr = true
				logging.Sugar.Error(err)
			}
		}

		if senderErr {
			return errors.New("dont send forms")
		}
	}

	return nil
}
