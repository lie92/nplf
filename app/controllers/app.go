package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Login(message string) revel.Result {
	return c.Render(message)
}

func (c App) Admin(uid string) revel.Result {
	return c.Render(uid)
}

func (c App) User(uid string) revel.Result {
	return c.Render(uid)
}

func (c App) Auth(email string, password string) revel.Result {

	row := app.Db.QueryRow("Select email, password, admin, uid FROM users WHERE email=$1 AND password=$2", email, password)

	var message string
	var admin bool
	var uid string

	switch err := row.Scan(&email, &password, &admin, &uid); err {
	case sql.ErrNoRows:
		message = "(Email ou mot de passe introuvable)"
	case nil:
		fmt.Println(email, password)
		if (admin) {
			fmt.Printf("is admin")
			c.Redirect(routes.App.Admin(uid))
		} else {
			fmt.Printf("not admin")
			c.Redirect(routes.App.User(uid))
		}

		c.Redirect("/")
	default:
		message = "(Connexion impossible)"
		panic(err)
	}



	return c.Redirect(routes.App.Login(message))
}

func (c App) Inscription() revel.Result {
	return c.Render()
}

func (c App) SignIn(nom string, prenom string, email string, password string, phone string) revel.Result {

	c.Validation.MinSize(nom, 8).Message("Your name is not long enough!")

	c.Flash.Success("Welcome, " + nom)

	eric := models.User{Firstname: prenom, Lastname: nom, Email: password, Password: "1234", Phone: phone}
	CreateAccount(eric)

	message := "(Inscription réussie)"

	return c.Redirect(routes.App.Login(message))
}

func CreateAccount(user models.User) {
	sqlStatement := `
		INSERT INTO users (firstname, lastname, email, password, admin, phone)
		VALUES ($1, $2, $3, $4, false, $5)
		RETURNING id`
	id := 0
	err := app.Db.QueryRow(sqlStatement, user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
