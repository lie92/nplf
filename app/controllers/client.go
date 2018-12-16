package controllers

import (
	"github.com/revel/revel"
	"nlpf/app"
	"nlpf/app/models"
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

func (c Client) Demande() revel.Result {

	return c.Render()
}