package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackSettings struct {
	Token     string
	ChannelID string
}

func sendSlackMessage(message Message, s SlackSettings) (success bool) {
	api := slack.New(s.Token)
	attachment := slack.Attachment{
		Text: message.TextMessage,
	}
	channelID, timestamp, err := api.PostMessage(
		s.ChannelID,
		slack.MsgOptionText(message.Subject, false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return true
}
