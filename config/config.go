package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Env string
	}
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed load .env file: ", err)
	}

	if appConfig == nil {
		appConfig = &AppConfig{}
		initApp(appConfig)
		initDB(appConfig)
		initServer(appConfig)
	}
	return appConfig
}
