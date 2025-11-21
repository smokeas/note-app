package api

import (
	"net/http"
)

// запускает HTTP-сервер
func StartServer(port string) {
	// Регистрируем обработчики для путей
	http.HandleFunc("/register", registerHandler) // При запросе на /register вызовем registerHandler
	http.HandleFunc("/login", loginHandler)       // При запросе на /login вызовем loginHandler

	// Запускаем сервер
	http.ListenAndServe(":"+port, nil)
}

// registerHandler обрабатывает запросы регистрации
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма регистрации")) //  Отправляем ответ клиенту
}

// loginHandler обрабатывает запросы входа
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма входа"))
}
