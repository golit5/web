package repository

import (
	"myapp/internal/model"

	"gorm.io/gorm"
)

type LinkRepository interface {
	Create(link *model.Link) error
	GetAll() ([]model.Link, error)
	GetByID(id uint) (*model.Link, error)
	Update(link *model.Link) error
	Delete(id uint) error
}

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkRepository{db}
}

func (r *linkRepository) Create(link *model.Link) error {
	return r.db.Create(link).Error
}

func (r *linkRepository) GetAll() ([]model.Link, error) {
	var links []model.Link
	err := r.db.Find(&links).Error
	return links, err
}

func (r *linkRepository) GetByID(id uint) (*model.Link, error) {
	var link model.Link
	err := r.db.First(&link, id).Error
	return &link, err
}

func (r *linkRepository) Update(link *model.Link) error {
	return r.db.Save(link).Error
}

func (r *linkRepository) Delete(id uint) error {
	return r.db.Delete(&model.Link{}, id).Error
}
