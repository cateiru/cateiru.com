package sender

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/logging"
	"github.com/mailgun/mailgun-go"
)

const TEMPLATE_DIR_PATH = "templates"

func (f *SendForm) MailSender(mail string) error {
	m := NewMail(mail, "Inquiry received")

	_, err := m.AddContentsFromTemplate("notice", f)
	if err != nil {
		return err
	}

	return m.Send()
}

type Mail struct {
	EmailAddress string
	Subject      string

	PlainText string
	HTMLText  string
}

func NewMail(email string, subject string) *Mail {
	return &Mail{
		EmailAddress: email,
		Subject:      subject,

		PlainText: "",
		HTMLText:  "",
	}
}

func (c *Mail) AddContents(plain string, html string) *Mail {
	c.PlainText = plain
	c.HTMLText = html

	return c
}

func (c *Mail) AddContentsFromTemplate(tempName string, elements any) (*Mail, error) {
	var err error

	c.HTMLText, err = Template(tempName+".html", elements)
	if err != nil {
		return nil, err
	}
	c.PlainText, err = Template(tempName+".tmpl", elements)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Send mail via Mailgun
func (c *Mail) Send() error {
	domain := config.Config.MailFromDomain
	secret := config.Config.MailgunAPIKey
	sender := config.Config.SenderMailAddress

	mg := mailgun.NewMailgun(domain, secret)
	message := mg.NewMessage(sender, c.Subject, c.PlainText, c.EmailAddress)
	message.SetHtml(c.HTMLText)

	resp, id, err := mg.Send(message)
	if err != nil {
		return err
	}

	logging.Sugar.Infof("Send email. ID: %s, Resp: %s", id, resp)

	return nil
}

func Template(path string, elements any) (string, error) {
	templ, err := template.ParseFiles(filepath.Join(TEMPLATE_DIR_PATH, path))
	if err != nil {
		return "", err
	}

	writer := new(strings.Builder)
	if err := templ.Execute(writer, elements); err != nil {
		return "", err
	}

	return writer.String(), nil
}
