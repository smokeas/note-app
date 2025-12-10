package api

import (
 "net/http"
 "time"

 "note-app/internal/api/handlers"
 "note-app/internal/service"
)

func RegisterRoutes(userService *service.UserService, jwtSecret string, jwtTTL time.Duration) {
 authHandler := handlers.NewAuthHandler(userService, jwtSecret, jwtTTL)

 http.HandleFunc("/register", authHandler.Register)
 http.HandleFunc("/login", authHandler.Login)

 // пример защищённого эндпоинта:
 // http.Handle("/me", middleware.AuthMiddleware(jwtSecret)(http.HandlerFunc(meHandler)))
}


