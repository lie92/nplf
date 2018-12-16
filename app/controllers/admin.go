package controllers

import (
	"github.com/revel/revel"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) Administration() revel.Result {
	return c.Render()
}