package repository

import (
	"connpass-keyword-bot-v1/domain/entity"
	"connpass-keyword-bot-v1/domain/repository"

	"github.com/line/line-bot-sdk-go/linebot"
)

type MessageRepositoryImpl struct {
	ID                 string
	ChannelSecret      string
	ChannelAccessToken string
}

func NewMessageRepository(
	id string, ChannelSecret string, channelAccessToken string,
) repository.MessageRepository {
	return &MessageRepositoryImpl{
		ID:                 id,
		ChannelSecret:      ChannelSecret,
		ChannelAccessToken: channelAccessToken,
	}
}

func (repo *MessageRepositoryImpl) SendMessage(messages []entity.Message) error {
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
