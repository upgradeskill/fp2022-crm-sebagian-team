package main

import (
	"time"

	"crm-sebagian-team/config"
	"crm-sebagian-team/server"

	"github.com/spf13/viper"
)

func main() {
	cfg := config.NewConfig()
	conn := config.NewDBConn()

	timeoutContext := time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second

	// init repo category repo
	repository := server.NewRepository(conn)
	service := server.NewService(cfg, conn, repository, timeoutContext)
	server.NewHandler(cfg, service)
}
