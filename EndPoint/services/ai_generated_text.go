package services

import (
	"awesomeProject1/EndPoint/models"
	"awesomeProject1/EndPoint/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

// HandleGenerateRequest обрабатывает запрос на генерацию текста
func HandleGenerateRequest(conn *websocket.Conn, db *sqlx.DB, req models.GenerateRequest) {
	// Сгенерированный текст
	generatedText := fmt.Sprintf("Текст для категории %s с данными: %v", req.CategoryID, req.UserData)

	// Сохранение сгенерированного текста в базе данных
	id, err := saveGeneratedText(db, req.CategoryID, generatedText)
	if err != nil {
		logrus.WithError(err).Error("Ошибка сохранения в базу данных")
		return
	}

	// Ответ клиенту
	resp := models.GeneratedText{
		ID:         id,
		CategoryID: req.CategoryID,
		Text:       generatedText,
		Status:     "completed",
		CreateDate: time.Now(),
	}

	// Отправка сгенерированного текста клиенту через WebSocket
	if err := utils.SendJSON(conn, resp); err != nil {
		logrus.WithError(err).Error("Ошибка отправки сообщения через WebSocket")
	}
}

// HandleListRequest обрабатывает запрос на получение списка генераций
func HandleListRequest(conn *websocket.Conn, db *sqlx.DB, req models.ListRequest) {
	// Получение списка генераций из базы данных
	generations, err := getGeneratedTextsByCategory(db, req.CategoryID)
	if err != nil {
		logrus.WithError(err).Error("Ошибка получения списка генераций")
		return
	}

	// Отправка списка генераций клиенту
	if err := utils.SendJSON(conn, generations); err != nil {
		logrus.WithError(err).Error("Ошибка отправки списка генераций через WebSocket")
	}
}

// saveGeneratedText сохраняет сгенерированный текст в базу данных
func saveGeneratedText(db *sqlx.DB, categoryID, text string) (int, error) {
	query := `INSERT INTO generated_texts (category_id, text, status, created_at) 
	          VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := db.Get(&id, query, categoryID, text, "completed", time.Now())
	if err != nil {
		logrus.WithError(err).Error("Ошибка выполнения SQL-запроса")
	}
	return id, err
}

// getGeneratedTextsByCategory получает список генераций по категории
func getGeneratedTextsByCategory(db *sqlx.DB, categoryID string) ([]models.GeneratedText, error) {
	var texts []models.GeneratedText
	query := `SELECT id, category_id, text, status, created_at 
	          FROM generated_texts WHERE category_id = $1`
	err := db.Select(&texts, query, categoryID)
	if err != nil {
		logrus.WithError(err).Error("Ошибка выполнения SQL-запроса для получения списка генераций")
	}
	return texts, err
}
