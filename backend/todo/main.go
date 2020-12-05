package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

type Task struct {
	ID     int
	Name   string
	Status int
}

var db *sql.DB

func dbSetup() {
	var err error
	db, err = sql.Open("mysql", "sample:sample@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error())
	}
}

func load() string {
	var task Task
	var err error

	err = db.QueryRow("SELECT * FROM tasks WHERE id = ?", 1).Scan(&task.ID, &task.Name, &task.Status)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
		return "empty"
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(task.ID, task.Name, task.Status)
		return task.Name
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s!", load())
}

func main() {
	dbSetup()
	// main() 終了時にDBをClseする.
	defer db.Close()

	http.HandleFunc("/", handler)

	fmt.Fprintln(os.Stdout, "Listening...")
	http.ListenAndServe(":8000", nil)
}
