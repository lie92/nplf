package controllers

import (
	"github.com/revel/revel"
	"myapp/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Login(message string) revel.Result {
	return c.Render(message)
}

func (c App) Inscription() revel.Result {
	return c.Render()
}


func (c App) SignIn(nom string) revel.Result {

	c.Validation.MinSize(nom, 8).Message("Your name is not long enough!")

	c.Flash.Success("Welcome, " + nom)


	message := "(Inscription réussie)"

	return c.Redirect(routes.App.Login(message))
}
