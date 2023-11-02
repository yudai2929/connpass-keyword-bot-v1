package main

func main() {

	// if err := config.LoadConfig(); err != nil {
	// 	panic(err)
	// }

	// notifiedEventRepository := repository.NewNotifiedEventRepository(
	// 	config.Config.SupabaseURL,
	// 	config.Config.SupabaseKey,
	// )
	// eventRepository := repository.NewEventRepository(config.Config.ConnpassURL)
	// messageRepository := repository.NewMessageRepository(
	// 	config.Config.UserID,
	// 	config.Config.ChannelSecret,
	// 	config.Config.ChannelAccessToken,
	// )

	// notificationUsecase := usecase.NewNotificationUsecase(
	// 	eventRepository,
	// 	messageRepository,
	// 	notifiedEventRepository,
	// )

	// notificationHandler := handler.NewNotificationHandler(notificationUsecase)
	// if err := notificationHandler.PostNotification(); err != nil {
	// 	fmt.Println(err)
	// }

}
