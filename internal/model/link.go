package model

type Link struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}
