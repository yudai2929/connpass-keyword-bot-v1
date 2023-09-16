package main

import (
	"connpass-keyword-bot-v1/pkg/config"
	"connpass-keyword-bot-v1/pkg/handler"
	"connpass-keyword-bot-v1/pkg/infrastructure/repository"
	"connpass-keyword-bot-v1/pkg/usecase"
	"fmt"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

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
