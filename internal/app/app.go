package app

import (
	"billing_service/internal/config"
	"billing_service/internal/db"
	"billing_service/internal/logger"
	"billing_service/internal/server"
	"github.com/spf13/viper"
)

func Run() {
	logger.Init()

	if err := config.InitConfig(); err != nil {
		logger.Logger.Fatal("reading config:", err)
	}

	if err := db.Init(&db.Config{
		Host:        viper.GetString("db.postgres.address.host"),
		Port:        viper.GetString("db.postgres.address.port"),
		DBName:      viper.GetString("db.postgres.name"),
		User:        viper.GetString("db.postgres.user.login"),
		Password:    viper.GetString("db.postgres.user.password"),
		SSLMode:     viper.GetString("db.postgres.ssl_mode"),
		MaxConn:     viper.GetInt("db.postgres.max_conn"),
		MaxIdleConn: viper.GetInt("db.postgres.max_max_idle_connconn"),
	}); err != nil {
		logger.Logger.Fatal("initializing database:", err)
	}

	s := server.NewServer()
	s.InitRouting()

	if err := s.Start(); err != nil {
		logger.Logger.Fatal("starting server:", err)
	}
}
