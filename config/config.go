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
}

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	Config.ConnpassURL = os.Getenv("CONNPASS_URL")
	Config.UserID = os.Getenv("USER_ID")
	Config.ChannelSecret = os.Getenv("CHANNEL_SECRET")
	Config.ChannelAccessToken = os.Getenv("CHANNEL_ACCESS_TOKEN")

}
