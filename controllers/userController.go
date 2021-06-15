package controllers

import (
	"github.com/Chronostasys/wishlist/dto"
	"github.com/Chronostasys/wishlist/models/user"
	"github.com/Chronostasys/wishlist/services"
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type resp map[interface{}]interface{}

type UserController interface {
	Register(ctx *gin.Context)
}
type userController struct {
}

func NewUserCtr() UserController {
	return &userController{}
}
func (uc *userController) Register(ctx *gin.Context) {
	r := &dto.Reg{}

	if err := ctx.BindJSON(r); err != nil {
		panic(err)
	}
	u := &user.User{}
	if err := mapper.Mapper(r, u); err != nil {
		panic(err)
	}
	u.Roles = []*user.Role{{Role: "User"}}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	u.Password = string(hash)
	id, err := services.USvc.AddUser(u)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(200, resp{"id": id})
}
