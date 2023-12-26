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
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	gin.Default().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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
