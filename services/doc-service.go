package Services

import (
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
func (s *DocumentService) UpdateDocument(model *models.Document, id uint) error {
	if model == nil {
		return nil
	}

	doc, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	doc.Title = model.Title
	doc.Content = model.Content
	doc.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(&doc); err != nil {
		return err
	}

	return nil
}

func (s *DocumentService) DeleteDocument(id uint) error {
	doc, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(doc.ID)
}
