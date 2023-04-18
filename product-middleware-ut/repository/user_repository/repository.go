package user_repository

import (
	"dts-fga-swswg/product-middleware-ut/entity"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(user entity.User) errs.MessageErr
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
}
