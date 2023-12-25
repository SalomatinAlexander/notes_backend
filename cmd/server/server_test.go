package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"noties/internal/handlers"
	"noties/internal/models"
	"noties/internal/services"
	"noties/internal/store"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func Test_CREATENOTE(t *testing.T) {

	repo := store.NewRepository(store.New(store.NewConfig()))
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	server := New("8080", handler)

	r := SetUpRouter()
	r.POST("/note/create-note", server.h.CreateNote)

	note := models.NoteFromCreateRequest{
		Title:       "title test1",
		Description: "description test1",
	}

	jsonNote, err := json.Marshal(note)
	if err != nil {
		log.Fatal("PARSE TEST ERROR:" + fmt.Sprint(err))
	}

	request, _ := http.NewRequest(http.MethodPost, "/note/create-note", bytes.NewBuffer(jsonNote))
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, request)

	if rec.Code == http.StatusOK {
		responseData, _ := io.ReadAll(rec.Body)
		fmt.Println("TEST RESPONSE IS:" + fmt.Sprint(responseData))
	} else if rec.Code == http.StatusInternalServerError || rec.Code == http.StatusBadRequest {
		fmt.Println("TEST RESPONSE  ERROR IS:" + fmt.Sprint(rec.Body))
	}

	//assert.
	//assert.Equal(t, rec.Body.String(), "Hello world")
}

func Test_GetAllNotes(t *testing.T) {
	repo := store.NewRepository(store.New(store.NewConfig()))
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	server := New("8080", handler)

	r := SetUpRouter()
	r.GET("/note/get-all", server.h.GetAllNotes)

	request, _ := http.NewRequest(http.MethodGet, "/note/get-all", bytes.NewBuffer(make([]byte, 0)))
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, request)

	if rec.Code == http.StatusOK {
		responseData, _ := io.ReadAll(rec.Body)
		fmt.Println("TEST RESPONSE IS:" + string(responseData))
	} else {
		fmt.Println("TEST RESPONSE  ERROR IS:" + fmt.Sprint(rec.Body))
	}
}
