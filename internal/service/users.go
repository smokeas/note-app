package service

import (
    "context"
    "errors"

    "note-app/internal/model"
    "note-app/internal/storage"

    "golang.org/x/crypto/bcrypt"
)

var ErrEmailTaken = errors.New("email already taken")
var ErrInvalidCredentials = errors.New("invalid email or password")

type UserService struct {
    repo storage.UserRepo
}

func NewUserService(repo storage.UserRepo) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, email, password string) (*model.User, error) {
    ok, err := s.repo.ExistsByEmail(ctx, email)
    if err != nil {
        return nil, err
    }
    if ok {
        return nil, ErrEmailTaken
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        Email:        email,
        PasswordHash: string(hash),
    }

    if err := s.repo.Create(ctx, user); err != nil {
        return nil, err
    }

    return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
    u, err := s.repo.FindByEmail(ctx, email)
    if err != nil {
        return nil, err
    }
    if u == nil {
        return nil, ErrInvalidCredentials
    }

    if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
        return nil, ErrInvalidCredentials
    }

    u.PasswordHash = ""
    return u, nil
}
