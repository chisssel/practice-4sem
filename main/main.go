package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"practice-4sem/config"
	"practice-4sem/routes"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Инициализация БД
	config.ConnectDB()

	// Настройка маршрутов
	r := routes.SetupRoutes()

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "5438"
	}
	r.Run(":" + port)
}
