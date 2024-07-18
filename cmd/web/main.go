package main

import (
	"github.com/rostis232/complicode/internal/pkg/app"
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	token := os.Getenv("TOKEN")
	key := os.Getenv("KEY")
	a, err := app.NewApp(token, key)
	if err != nil {
		log.Panic(err)
	}
	port := os.Getenv("PORT")
	log.Infof("PORT: %s", port)
	if err := a.Run(port); err != nil {
		log.Panic(err)
	}
}
