package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/errors"
)

type EnvVars struct {
	ConnpassURL        string
	UserID             string
	ChannelSecret      string
	ChannelAccessToken string
	SupabaseURL        string
	SupabaseKey        string
}

var Env *EnvVars = &EnvVars{}

func Load() error {
	err := godotenv.Load()

	if err != nil {
		return errors.Wrap(err, "failed to load .env file")
	}

	Env.ConnpassURL = os.Getenv("CONNPASS_URL")
	Env.UserID = os.Getenv("USER_ID")
	Env.ChannelSecret = os.Getenv("CHANNEL_SECRET")
	Env.ChannelAccessToken = os.Getenv("CHANNEL_ACCESS_TOKEN")
	Env.SupabaseURL = os.Getenv("SUPABASE_URL")
	Env.SupabaseKey = os.Getenv("SUPABASE_KEY")

	return nil
}
