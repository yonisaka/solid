package user

import (
	"encoding/json"
	"github.com/yonisaka/solid/internal/dto"
	"github.com/yonisaka/solid/internal/entity"
	"github.com/yonisaka/solid/internal/presentations"
	"github.com/yonisaka/solid/internal/repositories"
	"github.com/yonisaka/solid/internal/service/contract"
)

// userCreate is a struct
type userCreate struct {
	userRepo repositories.Userer
}

// NewUserCreate is a constructor
func NewUserCreate(userRepo repositories.Userer) contract.Service {
	return &userCreate{userRepo: userRepo}
}

// Serve is a method
func (u *userCreate) Serve(req contract.Request) contract.Response {
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

	user := entity.User{
		Name:  param.Name,
		Email: param.Email,
	}

	err = u.userRepo.Create(&user)
	if err != nil {
		return contract.Response{
			Message: err.Error(),
			Data:    nil,
		}
	}

	return contract.Response{
		Message: "success",
		Data:    dto.UserToResponse(user),
	}
}
