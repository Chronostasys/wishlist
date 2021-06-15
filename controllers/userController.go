package controllers

import (
	"github.com/Chronostasys/wishlist/models/user"
	"github.com/Chronostasys/wishlist/pkg/ginwrapper"
	"github.com/Chronostasys/wishlist/services"
	"github.com/devfeel/mapper"
	"golang.org/x/crypto/bcrypt"
)

type resp map[string]interface{}

type UserController interface {
	Register(b ginwrapper.Req) (response interface{}, code int)
}
type userController struct {
}

func NewUserCtr() UserController {
	return &userController{}
}
func (uc *userController) Register(b ginwrapper.Req) (response interface{}, code int) {
	u := &user.User{}
	if err := mapper.MapperMap(b, u); err != nil {
		return err, 400
	}
	u.Roles = []*user.Role{{Role: "User"}}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err, 500
	}
	u.Password = string(hash)
	id, err := services.USvc.AddUser(u)
	if err != nil {
		return err, 400
	}
	return resp{"id": id}, 200
}
