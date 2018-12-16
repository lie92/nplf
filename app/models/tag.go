package models

import "database/sql"

type Tag struct {
	Id       	int
	UserId		int
	Time  	 	string
	Place    	string
	Accepted    sql.NullBool
	Pending     bool
	Reason		sql.NullString
	Price		int
	Phone		sql.NullString
	Motif		sql.NullString
}