package main

import (
	"connpass-keyword-bot-v1/config"
	"connpass-keyword-bot-v1/infrastructure/repository"
	"connpass-keyword-bot-v1/interfaces/handler"
	"connpass-keyword-bot-v1/usecase"
	"fmt"
)

func main() {
	config.LoadConfig()

	notifiedEventRepository := repository.NewNotifiedEventRepository(
		config.Config.SupabaseURL,
		config.Config.SupabaseKey,
	)
	eventRepository := repository.NewEventRepository(config.Config.ConnpassURL)
	messageRepository := repository.NewMessageRepository(
		config.Config.UserID,
		config.Config.ChannelSecret,
		config.Config.ChannelAccessToken,
	)

	notificationUsecase := usecase.NewNotificationUsecase(
		eventRepository,
		messageRepository,
		notifiedEventRepository,
	)

	notificationHandler := handler.NewNotificationHandler(notificationUsecase)
	if err := notificationHandler.PostNotification(); err != nil {
		fmt.Println(err)
	}

}
