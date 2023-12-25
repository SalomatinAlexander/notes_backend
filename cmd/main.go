package main

import (
	"flag"
	"fmt"
	"log"
	"noties/cmd/server"
	"noties/internal/handlers"
	"noties/internal/services"
	"noties/internal/store"

	"github.com/gin-gonic/gin"
)

var (
	configPath string
)

func main() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	repo := store.NewRepository(store.New(store.NewConfig()))
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	server := server.New("8080", handler)
	if err := server.Run(); err != nil {
		log.Fatal("ERROR:" + fmt.Sprint(err))
	}

}
