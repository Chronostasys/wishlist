package services

import (
	"errors"

	"github.com/Chronostasys/wishlist/models/user"
	"github.com/beego/beego/v2/client/orm"
)

var USvc UserService = NewUserService()
var badRolesErr error = errors.New("need roles not empty")

func NewUserService() UserService {
	return &userService{}
}

type UserService interface {
	AddUser(u *user.User) (id int64, err error)
}
type userService struct {
}

func (us *userService) AddUser(u *user.User) (id int64, err error) {
	if len(u.Roles) < 1 {
		err = badRolesErr
		return
	}
	o := orm.NewOrm()
	tx, err := o.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	m2m := o.QueryM2M(u, "Roles")
	rs := make([]interface{}, len(u.Roles))
	for i, v := range u.Roles {
		rs[i] = v.Role
	}
	q := o.QueryTable(&user.Role{}).Filter("role__in", rs...)
	_, err = q.All(&u.Roles, "id")
	if err != nil {
		return
	}
	for i, v := range u.Roles {
		rs[i] = v
	}
	id, err = o.Insert(u)
	if err != nil {
		return
	}
	_, err = m2m.Add(rs...)
	if err != nil {
		return
	}
	err = tx.Commit()
	return
}
