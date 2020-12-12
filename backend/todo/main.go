package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"path"
	"strconv"
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

func getTask(id int) (task Task, err error) {
	err = db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(&task.Id, &task.Name, &task.Status)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
		return
	case err != nil:
		return
	default:
		fmt.Println(task.Id, task.Name, task.Status)
	}
	return
}

func handleGetDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	task, err := getTask(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	json, _ := json.Marshal(task)
	w.Write(json)
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

func handleGetAll(w http.ResponseWriter, r *http.Request) (err error) {
	tasks, err := getTasks(30)
	json, _ := json.Marshal(tasks)
	w.Write(json)
	return
}

func storeTask(name string) (err error) {
	_, err = db.Exec("INSERT INTO tasks(name, status) VALUES (?, 0)", name)
	return
}

func handleStore(w http.ResponseWriter, r *http.Request) (err error) {
	var sr StoreRequest
	json.NewDecoder(r.Body).Decode(&sr)
	err = storeTask(sr.Name)
	return
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	var err error
	switch r.Method {
	case "GET":
		err = handleGetAll(w, r)
	case "POST":
		err = handleStore(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	_ = dbSetup()
	// main() 終了時にDBをClseする.
	defer db.Close()

	http.HandleFunc("/tasks", handleRequest)

	fmt.Fprintln(os.Stdout, "Listening...")
	http.ListenAndServe(":8000", nil)
}
