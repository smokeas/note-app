package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"note-app/config"
	"note-app/internal/api"
	"note-app/internal/service"
	pg "note-app/internal/storage/postgres"
)

func main() {
	cfg := config.Load()
	log.Printf("DEBUG: config DSN = %q", cfg.DB.DSN)

	// 1) подключаемся к БД
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer db.Close()

	// 2) проверяем соединение
	if err := db.Ping(); err != nil {
		log.Fatalf("db ping: %v", err)
	}

	// 3) создаём репозиторий и сервисы
	userRepo := pg.NewUserRepo(db)
	userService := service.NewUserService(userRepo)

	// 4) регистрируем роуты (передаём TTL как time.Duration)
	api.RegisterRoutes(userService, cfg.JWT.Secret, cfg.JWT.TTL)

	// 5) запускаем сервер
	addr := ":" + cfg.Server.Port
	log.Printf("Сервер запущен на порту %s", cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
