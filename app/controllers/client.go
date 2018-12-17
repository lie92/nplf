package controllers

import (
	"database/sql"
	"github.com/revel/revel"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
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

func (c Client) Index() revel.Result {

	var message string

	if isAuth() && !isAdmin() {

	sqlStatement := `SELECT * FROM tags WHERE userId=$1`

	rows, err := app.Db.Query(sqlStatement, 2)
	checkErr(err)
	var total = 0

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif)
		checkErr(err)
		total += tag.Price
		tags = append(tags, tag)
	}

	return c.Render(tags, total)} else {
		message = "(Vous n'avez pas les authorisations)"
		return c.Redirect(routes.App.Login(message))
	}
}

func (c Client) Facture() revel.Result {
	sqlStatement := `SELECT * FROM tags WHERE userId=$1`

	rows, err := app.Db.Query(sqlStatement, 2)
	checkErr(err)
	var total = 0

	
	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag

		err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
			&tag.Motif)
		checkErr(err)
		total += tag.Price
		tags = append(tags, tag)
	}

	return c.Render(tags, total)
}

func (c Client) ProcessDemande(address, date, hour, motif, phone string) revel.Result {
	booking := date + " " + hour

	sqlStatement := `INSERT INTO tags (userId, time, place, pending, price, accepted, motif, phone)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	var id int
	err := app.Db.QueryRow(sqlStatement, 2, booking, address, true, 20, sql.NullBool{false, false},
		phone, motif).Scan(&id)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Client.Index())
	}

func (c Client) DeleteDemande(id int) revel.Result {

	sqlStatement := `DELETE FROM tags WHERE id = $1`

	_, err := app.Db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Client.Index())
}

func (c Client) Demande() revel.Result {
	today := time.Now()

	y := today.Year()
	var m int = int (today.Month())
	d := today.Day()
	return c.Render(y, m, d)
}