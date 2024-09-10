package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Драйвер для PostgreSQL
	"github.com/sirupsen/logrus"
)

// SetupDatabase подключает к базе данных PostgreSQL через sqlx
func SetupDatabase() (*sqlx.DB, error) {
	connStr := "postgres://postgres:postgres@db:5432/generated_texts_db?sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logrus.WithError(err).Error("Ошибка подключения к базе данных")
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}
	logrus.Info("Подключение к базе данных успешно")
	return db, nil
}
