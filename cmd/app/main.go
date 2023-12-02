package main

import (
	"fmt"

	"github.com/yudai2929/connpass-keyword-bot-v1/internal/app/di"
	"github.com/yudai2929/connpass-keyword-bot-v1/internal/batch"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/config"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/errors"
)

func main() {

	if err := config.Load(); err != nil {
		panic(err)
	}

	c := di.BuildContainer()

	err := c.Invoke(func(handler *batch.NotificationJob) {
		if err := handler.Send(); err != nil {
			fmt.Printf("%+v", err)
		}
	})

	if err != nil {
		err = errors.Wrap(err, "failed to invoke the function")
		fmt.Printf("%+v", err)
	}
}
