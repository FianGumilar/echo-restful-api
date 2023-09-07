package config

import "os"

func initDB(conf *AppConfig) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	conf.Database.Host = host
	conf.Database.Port = port
	conf.Database.User = user
	conf.Database.Pass = pass
	conf.Database.Name = name
}
