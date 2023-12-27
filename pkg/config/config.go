package config

import (
	"os"
)

type EnvVars struct {
	ConnpassURL        string
	UserID             string
	ChannelSecret      string
	ChannelAccessToken string
	SupabaseURL        string
	SupabaseKey        string
	YahooClientID      string
}

var Env *EnvVars = &EnvVars{}

func Load() error {

	Env.ConnpassURL = os.Getenv("CONNPASS_URL")
	Env.UserID = os.Getenv("USER_ID")
	Env.ChannelSecret = os.Getenv("CHANNEL_SECRET")
	Env.ChannelAccessToken = os.Getenv("CHANNEL_ACCESS_TOKEN")
	Env.SupabaseURL = os.Getenv("SUPABASE_URL")
	Env.SupabaseKey = os.Getenv("SUPABASE_KEY")
	Env.YahooClientID = os.Getenv("YAHOO_CLIENT_ID")

	return nil
}
