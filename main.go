package main // import "zzz"

import (
	"net/http"
	"fmt"
	_"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"
	"util"
	"log"
)

func main() {
	fmt.Print("start")
	// echo 초기화 & CORS 설정
	e := echo.New()
	e.Use(middleware.CORS())

	/*** Controllers ***/
	// main
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Home")
	})
	// e.GET("/user/:id", getUser)

	// TODO List
	e.GET("/todos", selectTodoList)

	// TODO Complete
	// e.PUT("/todos/:idx", todoComplete)

	// TODO Delete
	// e.DELETE("/todos/:idx", deleteTodo)

	// TODO Insert
	// e.POST("/todos", insertTodo)

	// 시작
	e.Logger.Fatal(e.Start(":4000"))
}

func selectTodoList(c echo.Context) error {
	rows := util.SelectTodoList()
	result := util.Todos{}
	defer rows.Close()

	for rows.Next(){
		todo := util.Todo{}
		err2 := rows.Scan(&todo.Idx, &todo.Todo, &todo.IsDone)
		if err2 != nil { log.Fatal(err2) }
		result.Todos = append(result.Todos, todo)
	}
	return c.JSON(http.StatusOK, result)
}