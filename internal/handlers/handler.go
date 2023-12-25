package handlers

import (
	"noties/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) InitRout() *gin.Engine {
	router := gin.New()
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
