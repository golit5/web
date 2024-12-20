package service

import (
	"myapp/internal/model"
	"myapp/internal/repository"
)

type LinkService interface {
	CreateLink(link *model.Link) error
	GetAllLinks() ([]model.Link, error)
	GetLinkByID(id uint) (*model.Link, error)
	UpdateLink(link *model.Link) error
	DeleteLink(id uint) error
}

type linkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) LinkService {
	return &linkService{repo}
}

func (s *linkService) CreateLink(link *model.Link) error {
	return s.repo.Create(link)
}

func (s *linkService) GetAllLinks() ([]model.Link, error) {
	return s.repo.GetAll()
}

func (s *linkService) GetLinkByID(id uint) (*model.Link, error) {
	return s.repo.GetByID(id)
}

func (s *linkService) UpdateLink(link *model.Link) error {
	return s.repo.Update(link)
}

func (s *linkService) DeleteLink(id uint) error {
	return s.repo.Delete(id)
}
