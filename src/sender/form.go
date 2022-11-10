package sender

import "time"

type UserData struct {
	Browser  string
	OS       string
	Device   string
	IsMobile bool
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
