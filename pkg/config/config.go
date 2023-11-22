package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Env struct {
	ConnpassURL        string
	UserID             string
	ChannelSecret      string
	ChannelAccessToken string
	SupabaseURL        string
	SupabaseKey        string
}

func LoadConfig() error {
	err := godotenv.Load()

	if err != nil {
		return fmt.Errorf("failed to load .env file")
	}

	Env.ConnpassURL = os.Getenv("CONNPASS_URL")
	Env.UserID = os.Getenv("USER_ID")
	Env.ChannelSecret = os.Getenv("CHANNEL_SECRET")
	Env.ChannelAccessToken = os.Getenv("CHANNEL_ACCESS_TOKEN")
	Env.SupabaseURL = os.Getenv("SUPABASE_URL")
	Env.SupabaseKey = os.Getenv("SUPABASE_KEY")

	return nil
}
