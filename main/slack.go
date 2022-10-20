package logger

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"time"
)

func sendSlack(msg string, detail string) {
	webhookUrl := "https://hooks.slack.com/services/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"

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Message", Value: msg})
	attachment1.AddField(slack.Field{Title: "Detail", Value: detail})
	payload := slack.Payload{
		Text:        "ERROR time " + fmt.Sprintf("%v", time.Now()),
		Username:    "CONSUMER_BOT",
		Channel:     "consumer_alert",
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
