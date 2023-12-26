package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SalomatinAlexander/noties/cmd/server"
	"github.com/SalomatinAlexander/noties/internal/handlers"
	"github.com/SalomatinAlexander/noties/internal/services"
	"github.com/SalomatinAlexander/noties/internal/store"

	"github.com/gin-gonic/gin"
)

var (
	configPath string
)

// @title           Note Application
// @version         1.0
// @description     Pet Project
// @host       		8080
// @BasePath  		v2250198.hosted-by-vdsina.ru/note

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
