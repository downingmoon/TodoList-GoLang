package util // import "util"

import (
	"database/sql"
	_ "fmt"
	"log"
)

// DB 컬럼
type Todo struct {
	Idx     int    ` json:"idx" form:"idx" query:"idx" `
	Todo    string ` json:"todo" form:"todo" `
	IsDone  string ` json:"isDone" form:"isDone" `
	UserIdx int    ` json:"userIdx" form:"userIdx" query:"userIdx"`
}

type Todos struct {
	Todos []Todo ` json:"todo" `
}

type User struct {
	UserId       string ` json:"userId" form:"userId" query:"userId"`
	UserPassword string ` json:"userPassword" form:"userPassword" query:"userId"`
	UserIdx      int    ` json:"userIdx" form:"userIdx" query:"useridx"`
	Cnt          int
}

// Mysql DB 연결
func DBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:13306)/test")
	ErrorCheck(err)
	return db
}

// 에러 로깅
func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
