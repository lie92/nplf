package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"nlpf/app"
	"nlpf/app/models"
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

	return c.Render()
}

func (c Client) Facture() revel.Result {
	sqlStatement := `SELECT * FROM tags WHERE userId=$1`
	rows, err := app.Db.Query(sqlStatement, 1)
	checkErr(err)
	var total = 0

	
	var tags []models.Tag
	for rows.Next() {

		var uid int
		var userId int
		var time time.Time
		var place string
		var accepter bool
		var reason string
		var price int

		total += price
		err = rows.Scan(&uid, &userId, &time, &place, &accepter, &reason, &price)
		checkErr(err)
		fmt.Println("test")
		tags = append(tags, models.Tag{uid, userId, time, place, accepter,
			reason, price})
	}

	return c.Render(tags, total)
}

func (c Client) Demande() revel.Result {

	return c.Render()
}