package main

import (
	"fmt"
	"github.com/kr/pretty"
	"log"
	"net/http"
	"secondProject/db"
	"secondProject/pkg/settings"
	"secondProject/routes"
)

func init() {
	settings.Setup("./config/config.json")
	db.Setup()
}

func main() {
	defer deferFunc()

	endPoint := fmt.Sprintf(":%d", settings.Config.App.Port)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routes.Init(),
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen:", err)
	}
}

func deferFunc() {
	pretty.Logln("[MAIN] Work has stopped!")
	db.CloseDB()
}
