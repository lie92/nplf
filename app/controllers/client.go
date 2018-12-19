package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"io"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
	"os"
	"strconv"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Client struct {
	*revel.Controller
}

//Controller de la page d'accueil du client
func (c Client) Index() revel.Result {

	//Check si le client est authentifier et qu'il n'est pas un admin
	if isAuth() && !isAdmin() {

		var stored int

		_ = cache.Get("id", &stored)

		sqlStatement := `SELECT * FROM tags WHERE userId=$1`


		fmt.Printf("the is is is is : " + strconv.Itoa(stored) + "\n")

		rows, err := app.Db.Query(sqlStatement, stored)
		checkErr(err)
		var total int64 = 0

		var tags []models.Tag
		for rows.Next() {
			var tag models.Tag

			err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
				&tag.Motif, &tag.Orientation)
			checkErr(err)
			total += tag.Price.Int64
			tags = append(tags, tag)
		}

		return c.Render(tags, total)
	} else {
		//Sinon le client (ou admin) est redirigé vers une page 403
		return c.Redirect(routes.App.HTTP403())
	}
}

//Controller de la page facture
func (c Client) Facture() revel.Result {

	if isAuth() && !isAdmin() {
		//On fait la requète SQL
		sqlStatement := `SELECT * FROM tags WHERE userId=$1 and accepted=$2`
		var stored int
		_ = cache.Get("id", &stored)

		rows, err := app.Db.Query(sqlStatement, stored, true)
		checkErr(err)
		var total int64 = 0

		var tags []models.Tag
		//On remplis le tableau de tag
		for rows.Next() {
			var tag models.Tag

			err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
				&tag.Motif, &tag.Orientation)
			checkErr(err)
			total += tag.Price.Int64
			tags = append(tags, tag)
		}

		return c.Render(tags, total)
		} else {
			//Sinon le client (ou admin) est redirigé vers une page 403
			return c.Redirect(routes.App.HTTP403())
		}
	}

//Controller pour modifier la demande de tage
func (c Client) Modify(id int) revel.Result {
		if isAuth() && !isAdmin() {
			var stored int
		_ = cache.Get("id", &stored)
		sqlStatement := `SELECT * FROM tags WHERE userId=$1 AND id=$2`
		rows, err := app.Db.Query(sqlStatement, stored, id)
		checkErr(err)
		var tag models.Tag
		for rows.Next() {
			err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
				&tag.Motif, &tag.Orientation)
			checkErr(err)
		}

		return c.Render(tag)
		} else {
			//Sinon le client (ou admin) est redirigé vers une page 403
			return c.Redirect(routes.App.HTTP403())
		}
}

//Fonction pour modifier la demande de tag
func (c Client) ModifyDemande(address, motif, phone, orientation string, id int) revel.Result {

	sqlStatement := `UPDATE public.tags
	SET place=$1, phone=$2, motif=$3, orientation=$4
	WHERE id = $5`

	_, err := app.Db.Exec(sqlStatement, address, phone, motif, orientation, id)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Client.Index())

}

//Fonction pour créer la demande de tag
func (c Client) ProcessDemande(address, motif, phone, orientation string) revel.Result {

	//On prépare la requète SQL.
	var stored int
	_ = cache.Get("id", &stored)
	sqlStatement := `INSERT INTO tags (userId, place, pending, accepted, motif, phone, time, orientation)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	var id int
	//On l'exécute
	err := app.Db.QueryRow(sqlStatement, stored, address, true, sql.NullBool{false, false}, motif, phone, "01/01/01", orientation).Scan(&id)
	if err != nil {
		panic(err)
	}

	//On récuppère l'image
	file := c.Params.Files["pic"][0]
	//On enregistre l'image dans public/img/IdDuTag.png
	f, err := os.OpenFile("./public/img/"+strconv.Itoa(id)+".png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}

	defer f.Close()

	f2, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/")
	}
	defer f2.Close()

	io.Copy(f, f2)

	defer f2.Close()

	return c.Redirect(routes.Client.Index())
}

//Fonction pour effacer une demande de Tag
func (c Client) DeleteDemande(id int) revel.Result {
	if isAuth() && !isAdmin() {
		sqlStatement := `DELETE FROM tags WHERE id = $1`

		_, err := app.Db.Exec(sqlStatement, id)
		if err != nil {
			panic(err)
		}

		return c.Redirect(routes.Client.Index())
	} else {
		//Sinon le client (ou admin) est redirigé vers une page 403
		return c.Redirect(routes.App.HTTP403())
	}
}

//Controller de la demande de tag
func (c Client) Demande() revel.Result {
	if isAuth() && !isAdmin() {
		today := time.Now()

		y := today.Year()
		var m int = int(today.Month())
		d := today.Day()
		return c.Render(y, m, d)
	} else {
		//Sinon le client (ou admin) est redirigé vers une page 403
		return c.Redirect(routes.App.HTTP403())
	}
}

//Controller de la page profil des tags
func (c Client) Tag(id int) revel.Result {
	if isAuth() && !isAdmin() {
		var stored int

		_ = cache.Get("id", &stored)

		sqlStatement := `SELECT * FROM tags WHERE userId=$1 AND id=$2`

		rows, err := app.Db.Query(sqlStatement, stored, id)
		checkErr(err)

		var tag models.Tag
		for rows.Next() {

			err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
				&tag.Motif, &tag.Orientation)
			checkErr(err)

		}

		return c.Render(tag)
	} else {
		//Sinon le client (ou admin) est redirigé vers une page 403
		return c.Redirect(routes.App.HTTP403())
	}
}
