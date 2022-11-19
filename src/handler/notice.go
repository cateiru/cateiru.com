package handler

import (
	"context"
	"net/http"

	"github.com/cateiru/cateiru.com/ent/notice"
	"github.com/labstack/echo/v4"
)

func (h *Handler) NoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	n, err := h.DB.Client.Notice.
		Query().
		Where(notice.UserID(h.User.ID)).
		First(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, n)
}

func (h *Handler) UpdateNoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	n := h.DB.Client.Notice.
		Update().
		Where(notice.UserID(h.User.ID))

	discordWebhook := e.FormValue("discord_webhook")
	if discordWebhook != "" {
		if err := ValidateURL(discordWebhook); err != nil {
			return err
		}
		n = n.SetDiscordWebhook(discordWebhook)
	} else {
		n = n.ClearDiscordWebhook()
	}
	slackWebhook := e.FormValue("slack_webhook")
	if slackWebhook != "" {
		if err := ValidateURL(slackWebhook); err != nil {
			return err
		}
		n = n.SetSlackWebhook(slackWebhook)
	} else {
		n = n.ClearSlackWebhook()
	}
	mail := e.FormValue("mail")
	if mail != "" {
		if err := ValidateMail(mail); err != nil {
			return err
		}

		n = n.SetMail(mail)
	} else {
		n = n.ClearMail()
	}

	noticeD, err := n.Save(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, noticeD)
}
