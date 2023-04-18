package service_mocks

import (
	"dts-fga-swswg/product-middleware-ut/dto"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
	"dts-fga-swswg/product-middleware-ut/service"
)

var (
	CreateNewUser func(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login         func(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr)
)

type userServiceMock struct{}

func NewUserServiceMock() service.UserService {
	return &userServiceMock{}
}

func (u *userServiceMock) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	return CreateNewUser(payload)
}

func (u *userServiceMock) Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	return Login(payload)
}
