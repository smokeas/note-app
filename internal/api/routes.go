package api

import (
	"net/http"
)

// StartServer запускает HTTP-сервер на указанном порту.
// Оборачиваем обработчики в middleware для логирования/авторизации и т.д.
func StartServer(port string) {
	http.Handle("/register", LoggingMiddleware(http.HandlerFunc(registerHandler)))
	http.Handle("/login", LoggingMiddleware(http.HandlerFunc(loginHandler)))

	// Слушаем на всех интерфейсах (0.0.0.0) и на указанном порту.
	// В реальном приложении лучше обработать ошибку:
	//    if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil { log.Fatal(err) }
	_ = http.ListenAndServe("0.0.0.0:"+port, nil)
}
