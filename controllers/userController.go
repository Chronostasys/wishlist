package controllers

import (
	"github.com/Chronostasys/wishlist/dto"
	"github.com/Chronostasys/wishlist/models/user"
	"github.com/Chronostasys/wishlist/services"
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
)

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
	u.Roles = []*user.Role{{Role: "User"}, {Role: "Test"}}
	id, err := services.USvc.AddUser(u)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, map[string]int64{"id": id})
}
