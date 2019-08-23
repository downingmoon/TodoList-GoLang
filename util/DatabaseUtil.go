package util // import "util"

import (
	_"fmt"
	"database/sql"
	"log"
)

// DB 컬럼
type Todo struct {
	Idx int ` json:"idx" `
	Todo string ` json:"todo" `
	IsDone string ` json:"isDone"`
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

// 모든 TODO Rows 조회
// @param sql string
// @return Todos{}
func SelectTodoList() *sql.Rows {
	db := DBConnection()
	defer db.Close()
	rows, err := db.Query("SELECT IDX, TODO, ISDONE FROM TODOLIST ORDER BY IDX ASC")
	if err != nil { log.Fatal(err) }
	return rows
}