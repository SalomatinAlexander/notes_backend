package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SalomatinAlexander/noties/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary      Create Note
// @Description  Создание новой заметки
// @Tags         Notes
// @Accept       json
// @Produce      json
// @Param        title      body string  true  "Название"
// @Param        description     body string true  "Описание"
// @Success      200  {object}  models.CreateNoteResponse
// @Failure      400  {string}  Отправила хуйню, перепроверяй json
// @Failure      500  {string}  я где то накосячил
// @Router       /note/create-note [post]
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

// @Summary      Get All
// @Description  Метод для получения всех заметок
// @Tags         Notes
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Note
// @Failure      400  {string}  Отправила хуйню, перепроверяй json
// @Failure      500  {string}  я где то накосячил
// @Router       /note/get-all [get]
func (h *Handler) GetAllNotes(c *gin.Context) {
	result, err := h.Service.GetALlNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "GET ALL NOTE ERROR(я где то накосячил):"+fmt.Sprint(err))
	}
	decoder := json.NewDecoder(c.Request.Body)
	err = decoder.Decode(&result)
	//response, err := json.Marshal(result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
