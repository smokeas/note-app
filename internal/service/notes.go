package service

import "context"

// NoteService — заглушка для сервиса заметок
type NoteService struct {
	// добавь сюда зависимости (repo и т.д.)
}

func NewNoteService() *NoteService {
	return &NoteService{}
}

// Примеры методов — заглушки
func (s *NoteService) Create(ctx context.Context, userID int, title, body string) error {
	// TODO: реализовать
	return nil
}

func (s *NoteService) List(ctx context.Context, userID int) ([]any, error) {
	// TODO: реализовать
	return nil, nil
}
