package models

import (
	"database/sql"
	"time"
)

type Tag struct {
	Id       	int
	UserId		int
	Time  	 	time.Time
	Place    	string
	Accepted    sql.NullBool
	Pending     bool
	Reason		sql.NullString
	Price		int
	Phone		sql.NullString
	Motif		sql.NullString
}