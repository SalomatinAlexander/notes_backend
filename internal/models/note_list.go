package models

import "time"

type NoteList struct {
	Id       int       `json:"id"`
	UserId   int       `json:"user_id"`
	Title    string    `json:"title"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
