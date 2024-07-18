package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rostis232/complicode/internal/handler"
	"github.com/rostis232/complicode/internal/tgbot"
)

type App struct {
	Server  *echo.Echo
	Handler *handler.Handler
	TGBot   *tgbot.TGBot
}

func NewApp(token, key string) (*App, error) {
	bot, err := tgbot.NewBot(token, key)
	if err != nil {
		return nil, err
	}

	a := App{
		Handler: handler.NewHandler(bot),
		TGBot:   bot,
		Server:  echo.New(),
	}

	a.Server.Use(middleware.Logger())
	a.Server.Use(middleware.Recover())
	a.Server.Static("/", "./static")
	a.Server.GET("/", a.Handler.Home)
	a.Server.POST("/send", a.Handler.Send)

	return &a, nil
}

func (a *App) Run(port string) error {
	go a.TGBot.ListenTelegramForKey()

	return a.Server.Start(":" + port)
}
