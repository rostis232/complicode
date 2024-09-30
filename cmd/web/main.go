package main

import (
	"fmt"
	"os"

	"github.com/rostis232/complicode/internal/pkg/app"

	"github.com/labstack/gommon/log"
)

func main() {
	fmt.Println("Starting app...")
	token := os.Getenv("TOKEN")
	key := os.Getenv("KEY")
	port := os.Getenv("PORT")
	if token == "" || key == "" || port == "" {
		log.Panicf("Some env are not set:\nTOKEN: %s\nKEY: %s\nPORT: %s", token, key, port)
	}
	log.Infof("PORT: %s", port)
	
	a, err := app.NewApp(token, key)
	if err != nil {
		log.Panic(err)
	}

	if err := a.Run(port); err != nil {
		log.Panic(err)
	}
}
