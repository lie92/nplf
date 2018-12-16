package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"nlpf/app"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) Administration() revel.Result {
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
	fmt.Println("refus demande tag")
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