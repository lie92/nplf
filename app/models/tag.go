package models

import "time"

type Tag struct {
	UID       	int
	UserId		int
	Time  	 	time.Time
	Place    	string
	Accepted    bool
	Reason		string
	Price		int
}