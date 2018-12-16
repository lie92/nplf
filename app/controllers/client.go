package controllers

import (
	"github.com/revel/revel"
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
	/*tag1 := models.Tag{UID: 1, UserId: 1, Time: "02/10/18", Place: "lol", Accepted: true, Reason: "Non"}
	tag2 := models.Tag{UID: 2, UserId: 1, Time: "03/10/18", Place: "lol", Accepted: true, Reason: "Non"}
	tag3 := models.Tag{UID: 3, UserId: 1, Time: "04/10/18", Place: "lol", Accepted: true, Reason: "Non"}

	tags := [3]models.Tag{tag1, tag2, tag3}
	price := 0
	for e := range tags {
		price += tags[e].Price
	}*/
	return c.Render(/*tags, price*/)
}

func (c Client) Facture() revel.Result {
	/*sqlStatement := `SELECT * FROM tags WHERE id=$1`
	rows, err := app.Db.Query(sqlStatement, 1)
	checkErr(err)
	for rows.Next() {
		var uid int
		var userId string
		var time time.Time
		var place string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
		
	}

	tags := [3]models.Tag{tag1, tag2, tag3}
	price := 0
	for e := range tags {
		price += tags[e].Price
	}*/
	return c.Render(/*tags, price*/)
}

func (c Client) Demande() revel.Result {

	return c.Render()
}