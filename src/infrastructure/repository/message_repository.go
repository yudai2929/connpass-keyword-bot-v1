package repository

import (
	"connpass-keyword-bot-v1/src/domain/entity"
	"connpass-keyword-bot-v1/src/domain/repository"

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

	const maxConcurrentMessages = 5
	concurrentMessages := make(chan linebot.SendingMessage, maxConcurrentMessages)

	go func() {
		for _, message := range messages {
			concurrentMessages <- linebot.NewTextMessage(message.Text)
		}
		close(concurrentMessages)
	}()

	for message := range concurrentMessages {
		if _, err := bot.PushMessage(repo.ID, message).Do(); err != nil {
			return err
		}
	}

	return nil
}
