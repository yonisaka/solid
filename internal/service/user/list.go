package user

import (
	"github.com/yonisaka/solid/internal/dto"
	"github.com/yonisaka/solid/internal/repositories"
	"github.com/yonisaka/solid/internal/service/contract"
)

// userList is a struct
type userList struct {
	userRepo repositories.Userer
}

// NewUserList is a constructor
func NewUserList(userRepo repositories.Userer) contract.Service {
	return &userList{userRepo: userRepo}
}

// Serve is a method
func (u *userList) Serve(req contract.Request) contract.Response {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return contract.Response{
			Message: err.Error(),
			Data:    nil,
		}
	}

	return contract.Response{
		Message: "success",
		Data:    dto.UsersToResponse(users),
	}
}
