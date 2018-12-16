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
    dbport = "5432"
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
		UID       	SERIAL PRIMARY KEY,
		userId		integer REFERENCES users(id),
		time  	 	date,
		place    	varchar(80) NOT NULL,
		accepted    boolean,
		reason		varchar(80) NOT NULL,
		price       int NOT NULL
	);`

	_, err := Db.Exec(sqlStatement)
	if err != nil {
  		panic(err)
	}

	eric := models.User{Firstname: "Flavio", Lastname : "Copes", Email : "eric@gmail.com", Password : "1234", Phone:"0522398645"}
	defer createAccount(eric)
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
