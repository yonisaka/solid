package entity

import (
	"github.com/yonisaka/solid/internal/consts"
	"gorm.io/gorm"
	"time"
)

// UserInterface is an interface
type UserInterface interface {
	EntityInterface
	GetVerifiedStatus() string
}

// User is a struct
type User struct {
	ID        uint       `gorm:"not null;uniqueIndex;primaryKey"`
	Name      string     `gorm:"size:100;not null;"`
	Email     string     `gorm:"size:100;not null;uniqueIndex;"`
	Verified  int        `gorm:"size:1;not null;default:2"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// TableName is a method
func (u User) TableName() string {
	return "users"
}

// GetVerifiedStatus is a method
func (u User) GetVerifiedStatus() string {
	return consts.CodeToStrVerified[u.Verified]
}

var _ UserInterface = &User{}
