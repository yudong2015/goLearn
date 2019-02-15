package testing01

import (
	"testing"
	"os"
	"flag"
	"fmt"
)

var db struct{
	Url string
}

//test main
func TestMain(m *testing.M){
	// Pretend to open DB Connection
	db.Url = os.Getenv("DATABASE_URL")
	if db.Url == "" {
		db.Url = "localhost:3306"
	}

	flag.Parse()
	exitCode := m.Run()

	//Pretend to close DB connection
	db.Url = ""
	//Exit
	os.Exit(exitCode)
}

func TestDBWithMain(t *testing.T) {
	//Pretend to use DB
	fmt.Println(db.Url)
}