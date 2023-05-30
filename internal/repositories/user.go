package repositories

import (
	"github.com/yonisaka/solid/internal/entity"
	"gorm.io/gorm"
)

// userRepo is a struct
type userRepo struct {
	db *gorm.DB
}

// NewUserRepo is a constructor
func NewUserRepo(db *gorm.DB) Userer {
	return &userRepo{db: db}
}

// FindOne is a method to find one user
func (u *userRepo) FindOne(id uint) (*entity.User, error) {
	var user entity.User

	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll is a method to find all user
func (u *userRepo) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Create is a method to create user
func (u *userRepo) Create(user *entity.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// Update is a method to update user
func (u *userRepo) Update(user *entity.User) error {
	err := u.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete is a method to delete user
func (u *userRepo) Delete(id uint) error {
	err := u.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
