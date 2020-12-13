package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
type StoreRequest struct {
	Name string `json:"name"`
}

var db *sql.DB
var err error

func dbSetup() (err error) {
	db, err = sql.Open("mysql", "sample:sample@tcp(db:3306)/sample")
	return
}

func getTasks(limit int) (tasks []Task, err error) {
	rows, err := db.Query("SELECT id, name, status FROM tasks LIMIT ?", limit)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		task := Task{}
		if err = rows.Scan(&task.Id, &task.Name, &task.Status); err != nil {
			return
		}
		tasks = append(tasks, task)
	}
	return
}

func handleGetAll(c *gin.Context) {
	tasks, err := getTasks(30)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, tasks)
}

func storeTask(name string) (err error) {
	fmt.Println(name)
	_, err = db.Exec("INSERT INTO tasks(name, status) VALUES (?, 0)", name)
	return
}

func handleStore(c *gin.Context) {
	var sr StoreRequest
	if c.ShouldBindJSON(&sr) == nil {
		err = storeTask(sr.Name)
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusInternalServerError, "")
	}
	return
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Next()
	}
}

func runServer() {
	r := gin.Default()
	r.Use(corsMiddleware())

	r.GET("/tasks", handleGetAll)
	r.POST("/tasks", handleStore)

	r.Run(":8000")
}

func main() {
	_ = dbSetup()
	// main() 終了時にDBをClseする.
	defer db.Close()

	runServer()
}
