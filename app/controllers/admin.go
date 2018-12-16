package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"nlpf/app"
	"nlpf/app/models"
	"nlpf/app/routes"
	"strconv"
	"time"
)

type Admin struct {
	*revel.Controller
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

	return true
}

func (c Admin) Administration(begin_date_input time.Time, end_date_input time.Time) revel.Result {

	var message string

	if isAuth() {
		sqlStatement := `SELECT * FROM tags` /*WHERE time>$1`*/

		rows, err := app.Db.Query(sqlStatement) //, time.Now)

		var tags []models.Tag

		for rows.Next() {
			var tag models.Tag

			err = rows.Scan(&tag.Id, &tag.UserId, &tag.Time, &tag.Place, &tag.Pending, &tag.Accepted, &tag.Reason, &tag.Price, &tag.Phone,
				&tag.Motif)
			checkErr(err)
			tags = append(tags, tag)
		}

		checkErr(err)

		fmt.Println(err);
		fmt.Println(rows);
		fmt.Println(tags);
		fmt.Println("rows")
		fmt.Println("entering admin")
		fmt.Println(begin_date_input);
		fmt.Println(end_date_input);
		return c.Render()

	} else {
		message = "(Vous n'avez pas les authorisations)"
		return c.Redirect(routes.App.Login(message))
	}
}

func acceptOffer(id int) {
	sqlStatement := `
	UPDATE tags 
	SET accepted = true, pending = false
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	}
	fmt.Println("acceptation demande tag")
}

func refuseOffer(id int, reason string) {
	sqlStatement := `
	UPDATE tags 
	SET pending = false, accepted = false, reason = $2
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id, reason)

	if err != nil {
		panic(err)
	}
	fmt.Println("refus demande tag")
}
