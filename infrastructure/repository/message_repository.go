package repository

import (
	"connpass-keyword-bot-v1/domain/message"

	"github.com/line/line-bot-sdk-go/linebot"
)

type MessageRepository struct {
	ID                 string
	ChannelSecret      string
	ChannelAccessToken string
}

func NewMessageRepository(
	id string, ChannelSecret string, channelAccessToken string,
) message.MessageRepository {
	return &MessageRepository{
		ID:                 id,
		ChannelSecret:      ChannelSecret,
		ChannelAccessToken: channelAccessToken,
	}
}

func (repo *MessageRepository) SendMessage(messages []message.Message) error {
	bot, err := linebot.New(
		repo.ChannelSecret,
		repo.ChannelAccessToken,
	)

	if err != nil {
		return err
	}

	var lineMessages []linebot.SendingMessage
	for _, message := range messages {
		lineMessages = append(lineMessages, linebot.NewTextMessage(message.Text))
	}

	if _, err := bot.PushMessage(repo.ID, lineMessages...).Do(); err != nil {
		return err
	}

	return nil
}
