package controllers

import (
	"github.com/revel/revel"
	"nlpf/app/models"

)

type Client struct {
	*revel.Controller
}

func (c Client) Index() revel.Result {
	net1 := models.NewNettoyage(1, 25, "01/01/11", "lolnet")
	net2 := models.NewNettoyage(1, 78, "01/01/11", "lolnet")
	net3 := models.NewNettoyage(1, 54, "04/12/11", "lolbet")
	price := net1.Prix + net3.Prix + net2.Prix
	nettoyages := [3]*models.Nettoyage{net1, net2, net3}
	return c.Render(nettoyages, price)
}

func (c Client) Facture() revel.Result {
	net1 := models.NewNettoyage(1, 25, "01/01/11", "lolnet")
	net2 := models.NewNettoyage(1, 54, "04/12/11", "lolbet")
	price := net1.Prix + net2.Prix
	nettoyages := [2]*models.Nettoyage{net1, net2}
	return c.Render(nettoyages, price)
}

func (c Client) Demande() revel.Result {

	return c.Render()
}