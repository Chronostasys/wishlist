package primitives

import "time"

type Entity struct {
	IsDeleted  bool      `orm:"null"`
	DeleteTime time.Time `orm:"null"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	Id         int64     `orm:"auto"`
}
