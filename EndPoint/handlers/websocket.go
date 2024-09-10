package handlers

import (
	"awesomeProject1/EndPoint/models"
	"awesomeProject1/EndPoint/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Upgrader для WebSocket соединений
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleWebSocket обрабатывает WebSocket соединения
func HandleWebSocket(c *gin.Context, db *sqlx.DB) {
	if db == nil {
		logrus.Error("Подключение к базе данных отсутствует")
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.WithError(err).Error("Ошибка обновления до WebSocket")
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logrus.WithError(err).Error("Ошибка закрытия WebSocket")
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.WithError(err).Error("Неожиданное закрытие WebSocket")
			}
			break
		}

		var generateReq models.GenerateRequest
		if err := json.Unmarshal(msg, &generateReq); err == nil {
			services.HandleGenerateRequest(conn, db, generateReq)
		} else {
			var listReq models.ListRequest
			if err := json.Unmarshal(msg, &listReq); err == nil {
				services.HandleListRequest(conn, db, listReq)
			} else {
				logrus.WithError(err).Warn("Неправильный формат сообщения")
			}
		}
	}
}
