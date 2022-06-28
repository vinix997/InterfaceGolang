package service

import (
	"fmt"
	"interface/entity"
)

type UserServiceIface interface {
	Register(user *entity.User)
}

type UserSvc struct{}

func NewUserService() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity.User) {
	fmt.Printf("%v", user)
}
