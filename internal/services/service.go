package services

import (
	"github.com/SalomatinAlexander/noties/internal/models"
	"github.com/SalomatinAlexander/noties/internal/store"
)

type Note interface {
}

type NoteList interface {
}

type Service struct {
	Note,
	NoteList,
	repo store.Repository
}

func NewService(repo *store.Repository) *Service {
	return &Service{
		repo: *repo,
	}
}

func (s *Service) CreateNewNote(noteFromRequest models.NoteFromCreateRequest) (int, error) {
	noteCreate := models.GetNoteBeforeCreate(&noteFromRequest)
	result, err := s.repo.CreateNewNote(*noteCreate)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *Service) GetALlNotes() ([]models.Note, error) {
	result, err := s.repo.GetAllNotes()
	if err != nil {
		return make([]models.Note, 0), err
	}

	return result, nil
}
