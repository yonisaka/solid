package seeder

import (
	"github.com/yonisaka/solid/internal/consts"
	"github.com/yonisaka/solid/internal/entity"
	"github.com/yonisaka/solid/pkg/sqlx"
	"gorm.io/gorm"
)

// User is a seeder
type User struct{}

// Run is a method
func (s *User) Run(db *gorm.DB) error {
	// force delete all data
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}
	// reset auto increment
	if err := db.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error; err != nil {
		return err
	}

	users := []entity.User{
		{
			Name:     "John Doe",
			Email:    "johndoe@mail.com",
			Verified: consts.UserVerified,
		},
		{
			Name:     "Yoni Saka",
			Email:    "yonisaka@mail.com",
			Verified: consts.UserRegistered,
		},
	}

	if err := db.Create(&users).Error; err != nil {
		return err
	}

	return nil
}

var _ sqlx.Seeder = &User{}
