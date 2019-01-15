package main

import (
	"fmt"
	"time"

	"github.com/yudong2015/goLearn/random"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

//Leasing for test
type Leasing struct {
	Id               string
	ResouceId        string
	ResouceVersionId string
	UserId           string
	CreateTime       time.Time
	Status           string //updating / updated / overtime
	Duration         int32
}

func main() {
	url := "root:123456@tcp(127.0.0.1:3306)/test"
	conn, err := dbr.Open("mysql", url, nil)
	if err != nil {
		fmt.Printf("Connect to mysql fail: %v", err)
	}
	sess := conn.NewSession(nil)

	sess.Begin()

	result, err := sess.InsertBySql("insert into leasing (test_id, status) value(?, ?)", random.RandString(), "running").Exec()
	if err != nil {
		fmt.Printf("failed to insert: %v", err)
	}
	fmt.Println(result.RowsAffected())
}
