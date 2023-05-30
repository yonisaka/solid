package user

import (
	"encoding/json"
	"github.com/yonisaka/solid/internal/dto"
	"github.com/yonisaka/solid/internal/entity"
	"github.com/yonisaka/solid/internal/presentations"
	"github.com/yonisaka/solid/internal/repositories"
	"github.com/yonisaka/solid/internal/service/contract"
)

// userUpdate is a struct
type userUpdate struct {
	userRepo repositories.Userer
}

// NewUserUpdate is a constructor
func NewUserUpdate(userRepo repositories.Userer) contract.Service {
	return &userUpdate{userRepo: userRepo}
}

// Serve is a method
func (u *userUpdate) Serve(req contract.Request) contract.Response {
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
		ID:       param.ID,
		Name:     param.Name,
		Email:    param.Email,
		Verified: param.Verified,
	}

	err = u.userRepo.Update(&user)
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
