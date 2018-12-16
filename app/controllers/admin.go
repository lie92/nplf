package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"nlpf/app"
	"nlpf/app/models"
	"time"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) Administration(begin_date_input time.Time, end_date_input time.Time) revel.Result {
	sqlStatement := `SELECT * FROM tags` /*WHERE time>$1`*/

	rows, err := app.Db.Query(sqlStatement)//, time.Now)

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
}


func acceptOffer(id int) {
	sqlStatement := `
	UPDATE tags 
	SET accepted = true
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
	SET accepted = false, reason = $2
	WHERE id = $1; `

	_, err := app.Db.Exec(sqlStatement, id, reason)
	
	if err != nil {
  		panic(err)
	}
	fmt.Println("refus demande tag")
}