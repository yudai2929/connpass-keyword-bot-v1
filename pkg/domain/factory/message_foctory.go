package factory

import "github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"

func CreateMessage(event entity.Event) entity.Message {
	return entity.Message{
		Text: event.Title + "\n" + event.EventURL,
	}
}
