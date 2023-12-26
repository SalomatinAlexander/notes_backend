package handlers

import (
	"github.com/SalomatinAlexander/noties/internal/services"

	_ "github.com/SalomatinAlexander/noties/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) InitRout() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	notiesGroup := router.Group("/note")
	{
		notiesGroup.POST("/create-note", h.CreateNote)
		notiesGroup.GET("/get-all", h.GetAllNotes)
		//	notiesGroup.DELETE("/remove_note")
	}

	notiesListGroup := router.Group("/list")
	{
		notiesListGroup.POST("/create-list", h.createList)
		//	notiesListGroup.PATCH("/update_list")
		//	notiesGroup.DELETE("/delete_list")
		//	notiesGroup.GET("/:id")

	}
	return router
}
