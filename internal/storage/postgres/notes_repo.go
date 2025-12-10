package postgres

import "context"

type NotesRepo struct{}

func NewNotesRepo() *NotesRepo { return &NotesRepo{} }

func (r *NotesRepo) Create(ctx context.Context, userID int, title, body string) error { return nil }
