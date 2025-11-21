package api

import (
	"net/http"
)

// Простой middleware для логирования запросов
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логируем метод и URL запроса
		println("Запрос:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
