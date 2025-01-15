package model

// Link представляет ссылку
// @Description Модель для работы с ссылками
type Link struct {
	ID        uint   `json:"id" gorm:"primary_key"` // @example 1
	URL       string `json:"url"`                   // @example "https://example.com"
	Title     string `json:"title"`                 // @example "Example"
	CreatedAt string `json:"created_at"`            // @example "2025-01-15T00:00:00Z"
}
