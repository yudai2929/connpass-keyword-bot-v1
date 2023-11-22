package main

import (
	"fmt"

	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/config"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/infrastructure/external"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/infrastructure/line"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/infrastructure/supabase"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/interface/handler"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/libs"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/usecase"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	libs.InitSupabase()

	notifiedEventRepository := supabase.NewNotifiedEventRepository(
		libs.SupabaseClient,
	)

	eventRepository := external.NewEventRepository(config.Env.ConnpassURL)
	messageRepository := line.NewMessageRepository(
		config.Env.UserID,
		config.Env.ChannelSecret,
		config.Env.ChannelAccessToken,
	)

	notificationUsecase := usecase.NewNotificationUsecase(
		eventRepository,
		messageRepository,
		notifiedEventRepository,
	)

	notificationHandler := handler.NewNotificationHandler(notificationUsecase)

	if err := notificationHandler.Send(); err != nil {
		fmt.Printf("%+v", err)
	}

}
