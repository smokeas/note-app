package api

import (
 "net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "text/plain; charset=utf-8")
 w.Write([]byte("Форма регистрации. Отправь POST запрос с email и password"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "text/plain; charset=utf-8")
 w.Write([]byte("Форма входа. Отправь POST запрос с email и password"))
}

