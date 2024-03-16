package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type TConfig struct {
	ServerConfig      TServerConfig
	GoogleOAuthConfig TGoogleOAuthConfig
}

type TGoogleOAuthConfig struct {
	CustomConfig TCustomConfig
	OAuthConfig  oauth2.Config
}

type TCustomConfig struct {
	GoogleCallbackRelativeUri string
	GoogleCallbackFullUri     string
	GoogleRandomstateStr      string
	GoogleOathContextDeadline time.Duration
}

type TServerConfig struct {
	Port string
}

var Config TConfig

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading env file, error:%w", err)
	}

	Config.GoogleOAuthConfig.CustomConfig = TCustomConfig{
		GoogleCallbackRelativeUri: os.Getenv("GOOGLE_CALLBACK_PATH"),
		GoogleCallbackFullUri:     os.Getenv("SERVER_DOMAIN") + ":" + os.Getenv("SERVER_PORT") + os.Getenv("GOOGLE_CALLBACK_PATH"),
		GoogleRandomstateStr:      os.Getenv("GOOGLE_RANDOMSTATE_STR"),
		GoogleOathContextDeadline: time.Second * 10,
	}

	Config.ServerConfig = TServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}

	Config.GoogleOAuthConfig.OAuthConfig = oauth2.Config{
		RedirectURL:  Config.GoogleOAuthConfig.CustomConfig.GoogleCallbackFullUri,
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	return nil
}
