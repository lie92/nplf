package models

type Nettoyage struct {
	ClientId          int
	Date, Addresse     string
	Prix              int
}

func NewNettoyage(clientId, prix int, date, addresse string) *Nettoyage {
	u := &Nettoyage{
		ClientId: clientId,
		Date: date,
		Addresse:    addresse,
		Prix: prix,
	}
	return u
}