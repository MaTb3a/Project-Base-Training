package Repositories

import (
	"errors"

	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) GetAll() ([]T, error) {
	var items []T
	result := r.db.Find(&items)
	return items, result.Error
}

func (r *GormRepository[T]) GetByID(id uint) (T, error) {
	var item T
	result := r.db.First(&item, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		var empty T
		return empty, result.Error
	}
	return item, nil
}

func (r *GormRepository[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r *GormRepository[T]) Update(item *T) error {
	return r.db.Save(item).Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	var item T
	result := r.db.Delete(&item, id)
	return result.Error
}
