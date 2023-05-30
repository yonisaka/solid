package dto

import (
	"github.com/yonisaka/solid/internal/entity"
	"github.com/yonisaka/solid/internal/presentations"
)

// UserToResponse is a function to convert entity.User to presentations.ResponseUser
func UserToResponse(src entity.User) presentations.ResponseUser {
	result := presentations.ResponseUser{
		ID:       src.ID,
		Name:     src.Name,
		Email:    src.Email,
		Verified: src.GetVerifiedStatus(),
	}

	return result
}

// UsersToResponse is a function to convert []entity.User to []presentations.ResponseUser
func UsersToResponse(inputs []entity.User) []presentations.ResponseUser {
	var result []presentations.ResponseUser

	for i := 0; i < len(inputs); i++ {
		result = append(result, UserToResponse(inputs[i]))
	}

	return result
}
