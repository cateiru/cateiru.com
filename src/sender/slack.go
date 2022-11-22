package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SlackPlainTextBlock struct {
	Type  string `json:"type"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}

type SlackHeaderBlock struct {
	Type string              `json:"type"`
	Text SlackPlainTextBlock `json:"text,omitempty"`
}

type SlackSectionBlock struct {
	Type string              `json:"type"`
	Text SlackPlainTextBlock `json:"text,omitempty"`
}

type SlackFiledBlock struct {
	Type  string                `json:"type"`
	Filed []SlackPlainTextBlock `json:"fields"`
}

type SlackDividerBlock struct {
	Type string `json:"type"`
}

type SlackPayload struct {
	Blocks []any `json:"blocks"`
}

func (f *SendForm) SlackSender(webhook string) error {

	userData := []any{
		SlackHeaderBlock{
			Type: "header",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  "User Data",
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Name: %v", f.Name),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Mail: %v", f.Mail),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- IP Address: %v", f.IP),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Language: %v", f.Lang),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Device: %v", f.UserData.Device),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Browser: %v", f.UserData.Browser),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- OS: %v", f.UserData.OS),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Is Mobile: %v", f.UserData.IsMobile),
				Emoji: true,
			},
		},
	}

	formFiled := []any{
		SlackHeaderBlock{
			Type: "header",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  "Inquiry received",
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Subject: %v", f.Subject),
				Emoji: true,
			},
		},
		SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Detail: %v", f.Detail),
				Emoji: true,
			},
		},
	}

	if f.URL != "" {
		formFiled = append(formFiled, SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- URL: %v", f.URL),
				Emoji: true,
			},
		})
	}
	if f.Category != "" {
		formFiled = append(formFiled, SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Category: %v", f.URL),
				Emoji: true,
			},
		})
	}
	if f.CustomTitle != "" {
		formFiled = append(formFiled, SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Custom Subject: %v", f.CustomTitle),
				Emoji: true,
			},
		})
		formFiled = append(formFiled, SlackSectionBlock{
			Type: "section",
			Text: SlackPlainTextBlock{
				Type:  "plain_text",
				Text:  fmt.Sprintf("- Custom Value: %v", f.CustomValue),
				Emoji: true,
			},
		})
	}

	blocks := []any{}
	blocks = append(blocks, formFiled...)
	blocks = append(blocks, SlackDividerBlock{
		Type: "divider",
	})
	blocks = append(blocks, userData...)

	payload := SlackPayload{
		Blocks: blocks,
	}

	body := new(bytes.Buffer)

	err := json.NewEncoder(body).Encode(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhook, "application-json", body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(string(responseBody))
	}
	return nil
}
