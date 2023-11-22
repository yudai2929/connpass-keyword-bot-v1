package libs

import (
	"github.com/nedpals/supabase-go"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/config"
)

func NewSupabaseClient() *supabase.Client {
	return supabase.CreateClient(config.Env.SupabaseURL,
		config.Env.SupabaseKey)
}
