package user

import (
	"encoding/json"
	"github.com/yonisaka/solid/internal/dto"
	"github.com/yonisaka/solid/internal/presentations"
	"github.com/yonisaka/solid/internal/repositories"
	"github.com/yonisaka/solid/internal/service/contract"
)

// userDetail is a struct
type userDetail struct {
	userRepo repositories.Userer
}

// NewUserDetail is a constructor
func NewUserDetail(userRepo repositories.Userer) contract.Service {
	return &userDetail{userRepo: userRepo}
}

// Serve is a method
func (u *userDetail) Serve(req contract.Request) contract.Response {
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

	return contract.Response{
		Message: "success",
		Data:    dto.UserToResponse(*user),
	}
}
