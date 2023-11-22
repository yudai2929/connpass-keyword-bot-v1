package di

import (
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/infrastructure/external"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/infrastructure/line"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/infrastructure/supabase"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/interface/handler"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/libs"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/usecase"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	c := dig.New()

	c.Provide(libs.NewSupabaseClient)

	c.Provide(supabase.NewNotifiedEventRepository)
	c.Provide(external.NewEventRepository)
	c.Provide(line.NewMessageRepository)
	c.Provide(usecase.NewNotificationUsecase)
	c.Provide(handler.NewNotificationHandler)

	return c
}
