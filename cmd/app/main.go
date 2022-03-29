package main

import (
	"first-go-app/internal/app/adapter"

	"github.com/spf13/viper"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "user"
)

func main() {
	r := adapter.Router()

	// for local development
	viper.SetDefault("PGHOST", "postgres")
	viper.SetDefault("PGUSER", "user")
	viper.SetDefault("PGPASSWORD", "password")

	viper.BindEnv("PGHOST", host)
	viper.BindEnv("PGUSER", user)
	viper.BindEnv("PGPASSWORD", user)
	r.Run(":8080")
}
