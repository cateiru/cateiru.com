package sender

import (
	"github.com/gtuk/discordwebhook"
)

func (f *SendForm) DiscordSender(webhook string) error {

	isMobileStr := "yes"
	if !f.UserData.IsMobile {
		isMobileStr = "no"
	}

	userData := discordwebhook.Embed{
		Title: toPtr("User Data"),
		Fields: &[]discordwebhook.Field{
			{
				Name:  toPtr("Name"),
				Value: &f.Name,
			},
			{
				Name:  toPtr("Mail"),
				Value: &f.Mail,
			},
			{
				Name:  toPtr("IP Address"),
				Value: &f.IP,
			},
			{
				Name:  toPtr("Language"),
				Value: &f.Lang,
			},
			{
				Name:  toPtr("Device"),
				Value: &f.UserData.Device,
			},
			{
				Name:  toPtr("Browser"),
				Value: &f.UserData.Browser,
			},
			{
				Name:  toPtr("OS"),
				Value: &f.UserData.OS,
			},
			{
				Name:  toPtr("Is Mobile"),
				Value: &isMobileStr,
			},
		},
	}

	formFiled := []discordwebhook.Field{}

	if f.URL != "" {
		formFiled = append(formFiled, discordwebhook.Field{
			Name:  toPtr("URL"),
			Value: &f.URL,
		})
	}
	if f.Category != "" {
		formFiled = append(formFiled, discordwebhook.Field{
			Name:  toPtr("Category"),
			Value: &f.Category,
		})
	}
	if f.CustomTitle != "" {
		formFiled = append(formFiled, discordwebhook.Field{
			Name:  &f.CustomTitle,
			Value: &f.CustomValue,
		})
	}

	form := discordwebhook.Embed{
		Title:       &f.Subject,
		Description: &f.Detail,
		Fields:      &formFiled,
	}

	if f.URL != "" {
		form.Url = &f.URL
	}

	message := discordwebhook.Message{
		Content: toPtr("Inquiry received"),
		Embeds:  &[]discordwebhook.Embed{form, userData},
	}

	return discordwebhook.SendMessage(webhook, message)
}

func toPtr(s string) *string {
	return &s
}
