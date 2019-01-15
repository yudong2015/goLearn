package main

import (
	"fmt"

	"github.com/yudong2015/goLearn/random"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

// CREATE TABLE IF NOT EXISTS leasing(leasing_id VARCHAR(50)NOT NULL, create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, status VARCHAR(50), PRIMARY KEY (leasing_id));

func main() {
	url := "root:123456@tcp(127.0.0.1:3306)/test"
	conn, err := dbr.Open("mysql", url, nil)
	if err != nil {
		fmt.Printf("Connect to mysql fail: %v", err)
	}
	sess := conn.NewSession(nil)

	sess.Begin()

	result, err := sess.InsertBySql("insert into leasing (leasing_id, status) value(?, ?)", random.RandString(), "running").Exec()
	if err != nil {
		fmt.Printf("failed to insert: %v", err)
	}
	fmt.Println(result.RowsAffected())
}
