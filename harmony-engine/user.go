package main

import (
	"github.com/ServiceWeaver/weaver"
)

type User struct {
	weaver.AutoMarshal
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type IUserResult struct {
	weaver.AutoMarshal
	Id int64 `json:"id"`
}
