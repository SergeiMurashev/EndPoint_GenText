package main

import (
	config "awesomeProject1/EndPoint/configs"
	"awesomeProject1/EndPoint/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Настройка logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	// Инициализация базы данных
	db, err := config.SetupDatabase()
	if err != nil {
		logrus.WithError(err).Fatal("Ошибка подключения к базе данных")
	}
	defer db.Close()

	// Инициализация Gin
	r := gin.Default()

	// WebSocket маршрут для генерации текста
	r.GET("/ws", func(c *gin.Context) {
		handlers.HandleWebSocket(c, db)
	})

	// Запуск сервера на порту 8080
	if err := r.Run(":8080"); err != nil {
		logrus.WithError(err).Fatal("Ошибка запуска сервера")
	}
}
