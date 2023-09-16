package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Config struct {
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

	Config.ConnpassURL = os.Getenv("CONNPASS_URL")
	Config.UserID = os.Getenv("USER_ID")
	Config.ChannelSecret = os.Getenv("CHANNEL_SECRET")
	Config.ChannelAccessToken = os.Getenv("CHANNEL_ACCESS_TOKEN")
	Config.SupabaseURL = os.Getenv("SUPABASE_URL")
	Config.SupabaseKey = os.Getenv("SUPABASE_KEY")

	return nil
}
