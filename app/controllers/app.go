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

func isAdmin() bool {

	var isdAmin bool

	if err := cache.Get("admin", &isdAmin); err == nil {
		if isdAmin {
			return true
		} else {
			return false
		}
	}

	return false
}

func isAuth() bool {

	var idStore int

	if err := cache.Get("id", &idStore); err == nil {
		fmt.Printf(strconv.Itoa(idStore) + " is the id")

		if idStore != 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (c App) LogOut() revel.Result {
	go cache.Delete("id")
	go cache.Delete("admin")
	return c.Redirect("/")
}

func (c App) Login(message string) revel.Result {

	if isAuth() {

		if isAdmin() {
			const longForm = "Jan 2, 2006 at 3:04pm (MST)"
			t, _ := time.Parse(longForm, "Dec 29, 2012 at 7:54pm (PST)")
			return c.Redirect(routes.Admin.Administration(t, t))
		} else {
			return c.Redirect(routes.Client.Index())
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
			go cache.Set("admin", true, 30*time.Minute)
			const longForm = "Jan 2, 2006 at 3:04pm (MST)"
			t, _ := time.Parse(longForm, "Dec 29, 2012 at 7:54pm (PST)")
			return c.Redirect(routes.Admin.Administration(t, t))
		} else {
			go cache.Set("admin", false, 30*time.Minute)
			fmt.Printf("not admin ")
			return c.Redirect(routes.Client.Index())
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

	eric := models.User{Firstname: prenom, Lastname: nom, Email: email, Password: password, Phone: phone}
	CreateAccount(eric)
	message := "(Inscription réussie)"

	return c.Redirect(routes.App.Login(message))
}

func CreateAccount(user models.User) {
	sqlStatement := `INSERT INTO users (firstname, lastname, email, password, admin, phone)
VALUES ($1, $2, $3, $4, true, $5) RETURNING id`
	id := 0
	err := app.Db.QueryRow(sqlStatement, user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
