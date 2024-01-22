package main

import (
	"github.com/ServiceWeaver/weaver"
)

type Prompt struct {
	weaver.AutoMarshal
	Id    int    `json:"id"`
	Text  string `json:"text"`
	Model string `json:"model"`
	Tags  string `json:"tags"`
}

type IPromptResult struct {
	weaver.AutoMarshal
	Id int64 `json:"id"`
}

type IPromptListResult struct {
	weaver.AutoMarshal
	Prompts []Prompt `json:"prompts"`
}
