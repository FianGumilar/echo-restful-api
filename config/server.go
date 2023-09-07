package config

import "os"

func initServer(conf *AppConfig) {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	conf.Server.Host = host
	conf.Server.Port = port
}
