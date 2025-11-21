/*
package api

import (

	"net/http"

)

// запускает HTTP-сервер

	func StartServer(port string) {
		// Регистрируем обработчики для путей URL
		http.HandleFunc("/register", registerHandler) // При запросе на /register вызовем registerHandler
		http.HandleFunc("/login", loginHandler)       // При запросе на /login вызовем loginHandler

		// Запускаем сервер
		http.ListenAndServe(":"+port, nil)
	}

//registerHandler обрабатывает запросы регистрации
//w http.ResponseWriter — интерфейс для отправки ответа клиенту
//r *http.Request — указатель на структуру с данными запроса

	func registerHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Форма регистрации")) //  Отправляем ответ клиенту ; ]byte("...") — преобразование строки в байты (HTTP работает с байтами)  ; Write — метод интерфейса ResponseWriter для отправки данных
	}

// loginHandler обрабатывает запросы входа

	func loginHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Форма входа"))
	}
*/
package apis

import (
	"net/http"
)

func StartServer(port string) {
	// Оборачиваем все обработчики в middleware
	http.Handle("/register", LoggingMiddleware(http.HandlerFunc(registerHandler)))
	http.Handle("/login", LoggingMiddleware(http.HandlerFunc(loginHandler)))

	http.ListenAndServe("0.0.0.0:"+port, nil)
}
