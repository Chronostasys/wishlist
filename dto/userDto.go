package dto

import "github.com/devfeel/mapper"

func init() {
	mapper.Register(&Reg{})
}

type Reg struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
