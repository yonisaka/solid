package user

import (
	"encoding/json"
	"github.com/yonisaka/solid/internal/dto"
	"github.com/yonisaka/solid/internal/presentations"
	"github.com/yonisaka/solid/internal/repositories"
	"github.com/yonisaka/solid/internal/service/contract"
)

// userDelete is a struct
type userDelete struct {
	userRepo repositories.Userer
}

// NewUserDelete is a constructor
func NewUserDelete(userRepo repositories.Userer) contract.Service {
	return &userDelete{userRepo: userRepo}
}

// Serve is a method
func (u *userDelete) Serve(req contract.Request) contract.Response {
	var (
		param presentations.RequestUser
	)

	data, err := json.Marshal(req.Data)
	if err != nil {
		return contract.Response{
			Message: err.Error(),
			Data:    nil,
		}
	}
	err = json.Unmarshal(data, &param)
	if err != nil {
		return contract.Response{
			Message: err.Error(),
			Data:    nil,
		}
	}

	user, err := u.userRepo.FindOne(param.ID)
	if err != nil {
		return contract.Response{
			Message: err.Error(),
			Data:    nil,
		}
	}

	if user == nil {
		return contract.Response{
			Message: "record not found",
			Data:    nil,
		}
	}

	err = u.userRepo.Delete(param.ID)
	if err != nil {
		return contract.Response{
			Message: err.Error(),
			Data:    nil,
		}
	}

	return contract.Response{
		Message: "success",
		Data:    dto.UserToResponse(*user),
	}
}
