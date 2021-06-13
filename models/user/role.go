package user

import (
	"github.com/Chronostasys/wishlist/models/primitives"
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(&Role{})
}

type Role struct {
	primitives.Entity
	Role  string  `orm:"unique"`
	Users []*User `orm:"reverse(many)"`
}
