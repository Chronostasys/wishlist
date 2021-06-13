package services

import (
	"log"

	"github.com/Chronostasys/wishlist/models/user"
	"github.com/beego/beego/v2/client/orm"
)

var USvc UserService = NewUserService()

func NewUserService() UserService {
	return &userService{}
}

type UserService interface {
	AddUser(u *user.User) (id int64, err error)
}
type userService struct {
}

func (us *userService) AddUser(u *user.User) (id int64, err error) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(u, "Roles")
	rs := make([]interface{}, len(u.Roles))
	for i, v := range u.Roles {
		rs[i] = v.Role
	}
	q := o.QueryTable(&user.Role{}).Filter("role__in", rs...)
	n, err := q.All(&u.Roles, "id")
	log.Println(n, err)
	for i, v := range u.Roles {
		rs[i] = v
	}
	id, err = o.Insert(u)
	m2m.Add(rs...)
	return
}
