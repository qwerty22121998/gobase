package test

import "github.com/qwerty22121998/gobase/base_model"

type ModelA struct {
	base_model.Model
	A string
}

func (ModelA) TableName() string {
	return "model_a"
}

type ModelB struct {
	base_model.Model
	B int
}

func (ModelB) TableName() string {
	return "model_b"
}
