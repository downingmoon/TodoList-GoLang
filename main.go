package main // import "zzz"

import (
	_ "encoding/json"
	_ "fmt"
	_ "log"
	"net/http"
	"time"
	"util"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
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
	e.GET("/todos/:idx", selectTodoList)

	// TODO Complete
	e.PUT("/todos/:idx", todoComplete)

	// TODO Delete
	e.DELETE("/todos/:idx", deleteTodo)

	// TODO Insert
	e.POST("/todos", insertTodo)

	// user signup
	e.POST("/user", insertUser)

	// login
	e.POST("/login", login)

	// Restricted Group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	// 시작
	e.Logger.Fatal(e.Start(":4000"))
}

// todo 리스트 조회
func selectTodoList(c echo.Context) error {
	idx := c.Param("idx")
	db := util.DBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT IDX, TODO, ISDONE FROM TODOLIST WHERE USER_IDX = ? ORDER BY IDX ASC", idx)
	util.ErrorCheck(err)
	result := util.Todos{}
	defer rows.Close()

	for rows.Next() {
		todo := util.Todo{}
		err2 := rows.Scan(&todo.Idx, &todo.Todo, &todo.IsDone)
		util.ErrorCheck(err2)
		result.Todos = append(result.Todos, todo)
	}

	return c.JSON(http.StatusOK, result)
}

// todo 완료
func todoComplete(c echo.Context) error {
	idx := c.Param("idx")
	db := util.DBConnection()
	defer db.Close()

	res, err := db.Query("UPDATE TODOLIST SET ISDONE = '1' WHERE IDX = ?", idx)
	util.ErrorCheck(err)

	return c.JSON(http.StatusOK, res)
}

// new todo
func insertTodo(c echo.Context) error {
	todo := new(util.Todo)
	err := c.Bind(todo)
	util.ErrorCheck(err)

	db := util.DBConnection()
	defer db.Close()
	res, err := db.Query("INSERT INTO TODOLIST (TODO, USER_IDX) VALUES (?, ?)", todo.Todo, todo.UserIdx)
	util.ErrorCheck(err)

	return c.JSON(http.StatusOK, res)
}

// todo삭제
func deleteTodo(c echo.Context) error {
	idx := c.Param("idx")
	db := util.DBConnection()
	defer db.Close()

	res, err := db.Query("DELETE FROM TODOLIST WHERE IDX = ?", idx)
	util.ErrorCheck(err)

	return c.JSON(http.StatusOK, res)
}

// 회원가입
func insertUser(c echo.Context) error {
	user := new(util.User)
	err := c.Bind(user)
	util.ErrorCheck(err)

	db := util.DBConnection()
	defer db.Close()

	res, err := db.Query("INSERT INTO TODOUSER (USER_ID, USER_PWD) VALUES (?, ?)", user.UserId, user.UserPassword)
	util.ErrorCheck(err)

	return c.JSON(http.StatusOK, res)
}

func login(c echo.Context) error {
	user := new(util.User)
	err := c.Bind(user)
	util.ErrorCheck(err)

	db := util.DBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(USER_ID), USER_ID, USER_IDX FROM TODOUSER WHERE USER_ID = ? AND USER_PWD = ?", user.UserId, user.UserPassword)
	err = row.Scan(&user.Cnt, &user.UserId, &user.UserIdx)
	util.ErrorCheck(err)

	if user.Cnt == 0 {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.UserId
	claims["userIdx"] = user.UserIdx
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	util.ErrorCheck(err)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":   t,
		"userId":  user.UserId,
		"userIdx": user.UserIdx,
	})
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// userId := claims["userId"].(string)
	return c.JSON(http.StatusOK, claims)
}
