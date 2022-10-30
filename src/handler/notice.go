package handler

import (
	"context"
	"net/http"
	"strconv"

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
		All(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, n)
}

func (h *Handler) CreateNoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	n := h.DB.Client.Notice.
		Create().
		SetUserID(h.User.ID)

	discordWebhook := e.FormValue("discord_webhook")
	if discordWebhook != "" {
		n = n.SetDiscordWebhook(discordWebhook)
	}
	slackWebhook := e.FormValue("slack_webhook")
	if slackWebhook != "" {
		n = n.SetDiscordWebhook(slackWebhook)
	}
	mail := e.FormValue("mail")
	if mail != "" {
		if err := ValidateMail(mail); err != nil {
			return err
		}

		n = n.SetMail(mail)
	}

	noticeD, err := n.Save(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, noticeD)
}

func (h *Handler) UpdateNoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.FormValue("notice_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: notice_id")
	}

	n := h.DB.Client.Notice.
		Update().
		Where(notice.ID(uint32(id)))

	discordWebhook := e.FormValue("discord_webhook")
	if discordWebhook != "" {
		n = n.SetDiscordWebhook(discordWebhook)
	} else {
		n = n.ClearDiscordWebhook()
	}
	slackWebhook := e.FormValue("slack_webhook")
	if slackWebhook != "" {
		n = n.SetDiscordWebhook(slackWebhook)
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

func (h *Handler) DeleteNoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.QueryParam("notice_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: notice_id")
	}

	_, err = h.DB.Client.Notice.
		Delete().
		Where(notice.ID(uint32(id))).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
