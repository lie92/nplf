package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
	"strconv"
	"time"
)

type App struct {
	*revel.Controller
}

func isAdmin(id int) bool {

	row := app.Db.QueryRow("Select admin FROM users WHERE id=$1", id)
	var admin bool

	err := row.Scan(&admin)

	if err != nil {
		return true
	}

	return false
}

func (c App) Login(message string) revel.Result {

	var idStore int

	if err := cache.Get("id", &idStore); err == nil {
		fmt.Printf(strconv.Itoa(idStore) + " is the id")

		if idStore != 0 {
			if isAdmin(idStore) {
				return c.Redirect(routes.Admin.Administration())
			} else {
				return c.Redirect(routes.Admin.Administration())
			}
		} else {
			return c.Render(message)
		}
	} else {
		return c.Render(message)
	}
}

func (c App) User(uid string) revel.Result {
	return c.Render(uid)
}

func (c App) Auth(email string, password string) revel.Result {

	row := app.Db.QueryRow("Select email, password, admin, id FROM users WHERE email=$1 AND password=$2", email, password)

	var message string
	var admin bool
	var id int

	switch err := row.Scan(&email, &password, &admin, &id); err {
	case sql.ErrNoRows:
		message = "(Email ou mot de passe introuvable)"
	case nil:
		fmt.Println(email, password)

		go cache.Set("id", id, 30*time.Minute)

		if admin {
			fmt.Printf("is admin")
			return c.Redirect(routes.Admin.Administration())
		} else {
			fmt.Printf("not admin ")
			return c.Redirect(routes.Admin.Administration())
		}
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
