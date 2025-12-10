package storage

import (
    "context"

    "note-app/internal/model"
)

type UserRepo interface {
    Create(ctx context.Context, u *model.User) error
    FindByEmail(ctx context.Context, email string) (*model.User, error)
    ExistsByEmail(ctx context.Context, email string) (bool, error)
}
