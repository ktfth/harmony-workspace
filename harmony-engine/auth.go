package main

import (
	"github.com/ServiceWeaver/weaver"
)

type IAuthResult struct {
	weaver.AutoMarshal
	Token string `json:"token"`
}
