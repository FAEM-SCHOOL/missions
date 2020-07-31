package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var connStr = "user=zaur password=bober dbname=zaur sslmode=disable"
var db, _= sql.Open("postgres", connStr)
const port = ":5678"
var user User

func main()  {

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/registration.html", registrationPage)
	http.HandleFunc("/personal_account.html", personal_accountPage)
	http.HandleFunc("/exit.html", exitPage)

	fmt.Println("Server start on" + port)
	if err := http.ListenAndServe(port, nil); err != nil{
		log.Fatal("Listen server", err)
	}
}

//Structs
type Task struct {
	ID_user int
	ID_task int
	Date string
	Content string
}

type User struct {
	ID int
	Login string
	Password string
	Port string
	Tasks []Task
}

//Helper functions
func GetUser(login string) User {

	user := User{}
	users := GetUsers()
	for i := 0; i < len(users); i++{
		if users[i].Login == login{
			user = users[i]
			break
		}
	}
	return  user
}

func GetTasksOfUser(id int) []Task {
	var tasks []Task
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		task := Task{}
		if err := rows.Scan(&task.ID_task, &task.ID_user, &task.Date, &task.Content); err != nil{
			log.Println(err)
		}

		if (id == task.ID_user) {
			tasks = append(tasks, task)
		}
	}
	//че как дела уродец
	return tasks
}

func GetUsers() []User {

	var users []User
	rows, err := db.Query("select * from Users")
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		i := User{}
		if err := rows.Scan(&i.ID, &i.Login, &i.Password); err != nil{
			fmt.Println(err)
		}
		i.Port = port

		users = append(users, i)
	}
	return users
}

func UserIsExist(login string) bool {
	users := GetUsers()
	for i := 0; i < len(users); i++{
		if users[i].Login == login{
			return true
		}
	}
	return false
}

func Redirect(page string, w http.ResponseWriter){

	type data struct {
		Port string
		Page string
	}

	new_data := data{Port: port, Page: page}

	templ, _ := template.ParseFiles("scripts/redirect.html")
	templ.Execute(w, new_data)
}

func ParsePage(path string, ) *template.Template {
	var w http.ResponseWriter
	templ, err := template.ParseFiles(path)
	if err != nil{
		http.Error(w, err.Error(), 400)
	}

	return templ
}

//Pages
func indexPage(w http.ResponseWriter, r *http.Request){

	if user.Login != ""{
		Redirect("personal_account.html", w)
	}

	templ := ParsePage("pages/index.html")
	templ.Execute(w, port)

	if r.Method=="POST"{
		if err := r.ParseForm(); err != nil{
			log.Println(err)
		}

		login := r.FormValue("login")
		password := r.FormValue("password")

		if UserIsExist(login) && GetUser(login).Password == password{
			user = GetUser(login)
			user.Port = port
			Redirect("personal_account.html", w)
		} else {
			templ = ParsePage("scripts/faild_log_in.html")
			templ.Execute(w, nil)
		}
	}
}

func registrationPage(w http.ResponseWriter, r *http.Request){
	templ := ParsePage("pages/registration.html")
	templ.Execute(w, port)

	if r.Method == "POST"{
		if err := r.ParseForm(); err != nil{
			log.Println(err)
		}

		Login := r.FormValue("login")
		if UserIsExist(Login){
			templ := ParsePage("scripts/user_exist.html")
			templ.Execute(w, nil)
			return
		}

		login := r.FormValue("login")
		password := r.FormValue("password")
		if len(login) == 0 || len(password) == 0{
			templ := ParsePage("scripts/faild_input.html")
			templ.Execute(w, nil)
			return
		}

		_, err := db.Exec("INSERT INTO users (login, password) VALUES($1, $2)", Login, password)
		if err != nil{
			log.Println(err)
		}

		user = GetUser(login)
		user.Port = port
		Redirect("personal_account.html", w)

	}
}

func personal_accountPage(w http.ResponseWriter, r *http.Request)  {
	if len(user.Login) == 0 {
		Redirect("index.html", w)
		return
	}

	user.Tasks = GetTasksOfUser(user.ID)

	templ := ParsePage("pages/personal_account.html")
	templ.Execute(w, user)

	if r.Method == "POST" && r.FormValue("post") == "add"{
		if err := r.ParseForm(); err != nil{
			log.Println(err)
		}

		request := strings.Split(r.FormValue("task"), " ")

		if len(r.FormValue("task")) == 0{
			templ := ParsePage("scripts/faild_input.html")
			templ.Execute(w, nil)
			return
		}

		date := request[0]
		task := request[1]

		tasks := GetTasksOfUser(user.ID)
		var id int
		if len(tasks) >= 1 {
			id = tasks[len(tasks) - 1].ID_task + 1
		}

		if len(tasks) == 0{
			id = 1
		}


		_, err := db.Exec("INSERT INTO tasks(id, user_id, date, task) VALUES($1, $2, $3, $4)",id, user.ID, date, task)
		if err != nil {
			log.Println(err)
		}


		_, err = db.Exec("UPDATE tasks SET id = $1 WHERE id > $1", id)
		if err != nil {
			log.Println(err)
		}

		Redirect("personal_account.html", w)
	}
	if r.Method == "POST" && r.FormValue("post") == "delete"{
		if err := r.ParseForm(); err != nil{
			log.Println(err)
		}

		id_str := r.FormValue("delete")

		if len(id_str) == 0{
			templ := ParsePage("scripts/faild_input.html")
			templ.Execute(w, nil)
			return
		}

		if id_str == "all"{
			_, err := db.Exec("DELETE FROM tasks WHERE user_id = $1", user.ID)
			if err != nil{
				log.Println(err)
			}
			Redirect("personal_account.html", w)
		}else {
			id, _ := strconv.Atoi(id_str)
			_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
			if err != nil {
				log.Println(err)
			}

			_, err = db.Exec("UPDATE tasks SET id = id - 1  WHERE id > $1", id)
			if err != nil {
				log.Println(err)
			}

			Redirect("personal_account.html", w)
		}

	}

}

func exitPage(w http.ResponseWriter, r *http.Request){
	user = User{}
	user.Port = port
	templ := ParsePage("pages/exit.html")
	templ.Execute(w, user)
}
