package sender

import "time"

type UserData struct {
	Browser  string `json:"browser"`
	OS       string `json:"os"`
	Device   string `json:"device"`
	IsMobile bool   `json:"is_mobile"`
}

type SendForm struct {
	Name string
	Mail string
	Lang string

	Subject string
	Detail  string

	URL      string
	Category string

	Time time.Time
	IP   string

	CustomTitle string
	CustomValue string

	UserData *UserData
}
