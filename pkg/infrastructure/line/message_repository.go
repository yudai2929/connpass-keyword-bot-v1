package line

import (
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/config"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/repository"

	"github.com/line/line-bot-sdk-go/linebot"
)

type MessageRepositoryImpl struct {
	id                 string
	channelSecret      string
	channelAccessToken string
}

func NewMessageRepository() repository.MessageRepository {
	return &MessageRepositoryImpl{
		config.Env.UserID,
		config.Env.ChannelSecret,
		config.Env.ChannelAccessToken,
	}
}

func (repo *MessageRepositoryImpl) Send(messages []entity.Message) error {
	bot, err := linebot.New(
		repo.channelSecret,
		repo.channelAccessToken,
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
		if _, err := bot.PushMessage(repo.id, message).Do(); err != nil {
			return err
		}
	}

	return nil
}
