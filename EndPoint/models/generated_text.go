package models

import "time"

// GenerateRequest описывает запрос на генерацию текста
type GenerateRequest struct {
	CategoryID string                 `json:"category_id"`
	UserData   map[string]interface{} `json:"user_data"`
}

// ListRequest описывает запрос на список генераций
type ListRequest struct {
	CategoryID string `json:"category_id"`
}

// GeneratedText описывает структуру сгенерированного текста
type GeneratedText struct {
	ID         int       `db:"id" json:"id"`
	CategoryID string    `db:"category_id" json:"category_id"`
	Text       string    `db:"text" json:"text"`
	Status     string    `db:"status" json:"status"`
	CreateDate time.Time `json:"createDate" db:"createDate"`
}
