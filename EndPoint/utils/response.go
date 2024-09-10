package utils

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// SendJSON отправляет JSON-ответ через WebSocket
func SendJSON(conn *websocket.Conn, v interface{}) error {
	resp, err := json.Marshal(v)
	if err != nil {
		logrus.WithError(err).Error("Ошибка сериализации JSON")
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, resp); err != nil {
		logrus.WithError(err).Error("Ошибка отправки сообщения через WebSocket")
		return err
	}

	logrus.Info("Сообщение успешно отправлено через WebSocket")
	return nil
}
