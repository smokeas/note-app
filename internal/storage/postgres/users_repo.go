package postgres

import (
 "context"
 "database/sql"
 "errors"

 "note-app/internal/model"
 "note-app/internal/storage"
)

type pgUserRepo struct {
 db *sql.DB
}

func NewUserRepo(db *sql.DB) storage.UserRepo {
 return &pgUserRepo{db: db}
}

func (r *pgUserRepo) Create(ctx context.Context, u *model.User) error {
 query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, created_at`
 row := r.db.QueryRowContext(ctx, query, u.Email, u.PasswordHash)
 if err := row.Scan(&u.ID, &u.CreatedAt); err != nil {
  return err
 }
 return nil
}

func (r *pgUserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
 var u model.User
 query := `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`
 row := r.db.QueryRowContext(ctx, query, email)
 if err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt); err != nil {
  if errors.Is(err, sql.ErrNoRows) {
   return nil, nil
  }
  return nil, err
 }
 return &u, nil
}

func (r *pgUserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
 var count int
 err := r.db.QueryRowContext(ctx, `SELECT COUNT(1) FROM users WHERE email = $1`, email).Scan(&count)
 if err != nil {
  return false, err
 }
 return count > 0, nil
}


