package app

import (
	"github.com/revel/revel"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"nlpf/app/models"

)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)
var Db *sql.DB
const (
    dbhost = "localhost"
    dbport = "5433"
    dbuser = "postgres"
    dbpass = "postgres"
    dbname = "go"
)

func InitDB() {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        dbhost, dbport,
        dbuser, dbpass, dbname)

	Db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    err = Db.Ping()
    if err != nil {
        panic(err)
    }
	fmt.Println("Successfully connected!")
	createTables()
	fmt.Println("Tables created")	
}

func createTables() {
	sqlStatement := `
	DROP TABLE IF EXISTS users CASCADE;
	DROP TABLE IF EXISTS tags CASCADE;
	CREATE TABLE users (
		id       	SERIAL PRIMARY KEY,
		firstname   varchar(40) NOT NULL,
		lastname    varchar(40) NOT NULL,
		email	    varchar(40) NOT NULL,
		password	varchar(40) NOT NULL,
		admin 		boolean,
		phone		varchar(40) NOT NULL
	);
	CREATE TABLE tags (
		id       	SERIAL PRIMARY KEY,
		userId		integer NOT NULL,
		time  	 	varchar(80) NOT NULL,
		place    	varchar(80) NOT NULL,
		pending     boolean NOT NULL,
		accepted    boolean,
		reason		varchar(80),
		price       int NOT NULL,
		phone		varchar(80),
		motif		varchar(80)
	);`

	_, err := Db.Exec(sqlStatement)
	if err != nil {
  		panic(err)
	}

	eric := models.User{Firstname: "eric", Lastname : "li", Email : "eric@gmail.com", Password : "1234", Phone:"0522398645"}
	tony := models.User{Firstname: "tony", Lastname : "huang", Email : "tony@gmail.com", Password : "1234", Phone:"0522398645"}
	momo := models.User{Firstname: "momo", Lastname : "bennis", Email : "momo@gmail.com", Password : "1234", Phone:"0522398645"}

	tag1 := models.Tag{UserId: 1, Time : "06-0700", Place : "paris", Price : 14, Pending: true}
	tag2 := models.Tag{UserId: 2, Time : "06-0700", Place : "creteil", Price : 15, Pending: true}
	tag3 := models.Tag{UserId: 2, Time : "06-0700", Place : "italie", Price : 16, Pending: false, Accepted: sql.NullBool{true, true}}
	tag4 := models.Tag{UserId: 2, Time : "06-0700", Place : "lol", Price : 54, Pending: false, Accepted: sql.NullBool{false, true}}


	defer createAccount(eric)
	defer createAccount(momo)
	defer createAccount(tony)
	defer createTag(tag1)
	defer createTag(tag2)
	defer createTag(tag3)
	defer createTag(tag4)

	fmt.Println("creation compte")
}

func createAccount(user models.User) {
	sqlStatement := `
INSERT INTO users (firstname, lastname, email, password, admin, phone)
VALUES ($1, $2, $3, $4, false, $5)
RETURNING id`
  id := 0
  err := Db.QueryRow(sqlStatement, user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)
}

func createTag(tag models.Tag) {
	sqlStatement := `
INSERT INTO tags (userId, time, place, pending, price, accepted)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id`
  id := 0
  err := Db.QueryRow(sqlStatement, tag.UserId, tag.Time, tag.Place, tag.Pending, tag.Price, tag.Accepted).Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)
}


func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}



// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
