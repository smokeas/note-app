package main

import (
	"log"                   //  Логгер для вывода сообщений
	"note-app/config"       //  Подключаем наш «пульт управления»
	"note-app/internal/api" //  Подключаем HTTP-обработчики
)

func main() {
	cfg := config.Load() //  Загружаем настройки (берём пульт в руки)

	// Запускаем сервер на указанном порту
	log.Printf("Сервер запущен на порту %s", cfg.Server.Port)
	api.StartServer(cfg.Server.Port) //  Нажимаем кнопку ПУСК
}
