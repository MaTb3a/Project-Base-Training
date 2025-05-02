package Repositories

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id uint) (T, error)
	Create(item *T) error
	Update(item *T) error
	Delete(id uint) error
}
