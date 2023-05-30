package repositories

import "github.com/yonisaka/solid/internal/entity"

// Userer is a contract
type Userer interface {
	FindOne(id uint) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}
