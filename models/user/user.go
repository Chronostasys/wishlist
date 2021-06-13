package user

import (
	"github.com/Chronostasys/wishlist/models/primitives"
	"github.com/beego/beego/v2/client/orm"
	"github.com/devfeel/mapper"
)

func init() {
	orm.RegisterModel(&User{})
	mapper.Register(&User{})
}

type User struct {
	primitives.Entity
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Roles    []*Role `orm:"rel(m2m)"`
}
type UserRole struct {
	Id     int64
	UserId int64
	RoleId int64
}
