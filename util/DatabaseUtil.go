package util // import "util"

import (
	_"fmt"
	"database/sql"
	"log"
)

// DB 컬럼
type Todo struct {
	Idx int ` json:"idx" form:"idx" query:"idx" `
	Todo string ` json:"todo" form:"todo" `
	IsDone string ` json:"isDone" form:"isDone" `
}

type Todos struct {
	Todos []Todo ` json:"todo" `
}

// Mysql DB 연결
func DBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:13306)/test")
	if err != nil { log.Fatal(err) }
	return db
}