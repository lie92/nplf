package models

type User struct {
	Firstname   string
	Lastname    string
	Email	    string
	Password	string
	Admin 		bool
	UID       	int
	Phone		string
	Blacklist	bool
	//sql.nullstring
}
