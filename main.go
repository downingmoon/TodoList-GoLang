package main // import "zzz"

import (
	_"encoding/json"
	_"fmt"
	"log"
	"net/http"
	"util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// echo 초기화 & CORS 설정
	e := echo.New()
	e.Use(middleware.CORS())

	/*** Controllers ***/
	// main
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Home")
	})

	// TODO List
	e.GET("/todos", selectTodoList)

	// TODO Complete
	e.PUT("/todos/:idx", todoComplete)

	// TODO Delete
	e.DELETE("/todos/:idx", deleteTodo)

	// TODO Insert
	e.POST("/todos", insertTodo)

	// 시작
	e.Logger.Fatal(e.Start(":4000"))
}

func selectTodoList(c echo.Context) error {
	db := util.DBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT IDX, TODO, ISDONE FROM TODOLIST ORDER BY IDX ASC")
	if err != nil { log.Fatal(err) }
	result := util.Todos{}
	defer rows.Close()

	for rows.Next() {
		todo := util.Todo{}
		err2 := rows.Scan(&todo.Idx, &todo.Todo, &todo.IsDone)
		if err2 != nil {
			log.Fatal(err2)
		}
		result.Todos = append(result.Todos, todo)
	}
	return c.JSON(http.StatusOK, result)
}

func todoComplete(c echo.Context) error {
	idx := c.Param("idx")
	db := util.DBConnection()
	defer db.Close()
	res, err := db.Query("UPDATE TODOLIST SET ISDONE = '1' WHERE IDX = ?", idx)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, res)
}

func insertTodo(c echo.Context) error {
	todo := new(util.Todo)
	if err := c.Bind(todo); err != nil { log.Fatal(err) }

	db := util.DBConnection()
	defer db.Close()

	res, err := db.Query("INSERT INTO TODOLIST (TODO) VALUES (?)", todo.Todo)
	if err != nil { log.Fatal(err) }
	return c.JSON(http.StatusOK, res)
}

func deleteTodo(c echo.Context) error {
	idx := c.Param("idx")
	db := util.DBConnection()
	defer db.Close()

	res, err := db.Query("DELETE FROM TODOLIST WHERE IDX = ?", idx)
	if err != nil { log.Fatal(err) }
	return c.JSON(http.StatusOK, res)
}