package models

import (
	"time"
)

type Note struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	ListId      int    `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type NoteCreate struct {
	UserId      int    `json:"user_id"`
	ListId      int    `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type NoteFromCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateNoteResponse struct {
	Id int `json:"id" binding:"required"`
}

func GetNoteBeforeCreate(noteRequest *NoteFromCreateRequest) *NoteCreate {
	t := time.Now()
	timeFormat := t.UTC().Format("2006-01-02T15:04:05")
	return &NoteCreate{
		UserId:      0,
		ListId:      0,
		Title:       noteRequest.Title,
		Description: noteRequest.Description,
		CreateAt:    timeFormat,
		UpdateAt:    timeFormat,
	}
}
