package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	return &Server{
		Echo: echo.New(),
	}
}

func (s *Server) InitMiddlewares() {
	s.Echo.Use(newLoggerMiddleware())
}

func (s *Server) Start() error {
	return s.Echo.Start(fmt.Sprintf(
		"%s:%d",
		viper.GetString("app.address.host"),
		viper.GetInt("app.address.port"),
	))
}
