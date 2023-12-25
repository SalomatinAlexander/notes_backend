package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"noties/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateNote(c *gin.Context) {
	var noteFromRequest *models.NoteFromCreateRequest
	if err := c.BindJSON(&noteFromRequest); err != nil {
		log.Fatal("Create Note JSON ERROR:" + fmt.Sprint(err))
		c.JSON(http.StatusBadRequest, "JSON PARSE ERROR(Отправила хуйню, перепроверяй json):"+fmt.Sprint(err))
		return

	}
	result, err := h.Service.CreateNewNote(*noteFromRequest)
	if err != nil {
		log.Fatal("Create Note DB ERROR:" + fmt.Sprint(err))
		c.JSON(http.StatusInternalServerError, "CREATE NOTE ERROR(я где то накосячил):"+fmt.Sprint(err))
	}

	response, err := json.Marshal(models.CreateNoteResponse{Id: result})

	c.JSON(http.StatusOK, response)

}

func (h *Handler) GetAllNotes(c *gin.Context) {
	result, err := h.Service.GetALlNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "GET ALL NOTE ERROR(я где то накосячил):"+fmt.Sprint(err))
	}
	decoder := json.NewDecoder(c.Request.Body)
	err = decoder.Decode(&result)
	//response, err := json.Marshal(result)
	c.JSON(http.StatusOK, result)
}
