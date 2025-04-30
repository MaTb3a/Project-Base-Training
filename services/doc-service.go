package Services

import (
	"errors"
	"time"

	"github.com/MaTb3aa/Project-Base-Training/models"
	Repositories "github.com/MaTb3aa/Project-Base-Training/repository"
)

type DocumentService struct {
	repo Repositories.Repository[models.Document]
}

func NewDocumentService(repo Repositories.Repository[models.Document]) *DocumentService {
	return &DocumentService{
		repo: repo,
	}
}
func (s *DocumentService) CreateDoc(model *models.Document) error {
	if model == nil {
		return nil
	}
	model.ID = 0
	model.CreatedAt = time.Now()
	model.Author = "katreen"
	s.repo.Create(model)
	return nil
}
func (s *DocumentService) GetAllDocuments() ([]models.Document, error) {
	docs, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return docs, nil
}
func (s *DocumentService) GetDocumentByID(id uint) (models.Document, error) {
	doc, err := s.repo.GetByID(id)
	if err != nil {
		return models.Document{}, err
	}
	return doc, nil
}
func (s *DocumentService) UpdateDocument(model *models.Document) error {
	if model == nil {
		return nil
	}

	doc, err := s.repo.GetByID(model.ID)
	if err != nil {
		return err
	}
	if doc.ID == 0 {
		return errors.New("Document not found")
	}

	model.UpdatedAt = time.Now()
	err = s.repo.Update(model)
	if err != nil {
		return err
	}

	return nil
}
