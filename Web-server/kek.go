package main

import (
	"html/template"
	"log"
	"net/http"
)

var items = []Item{
	{"Sumsung A10", 100, "Phones"},
	{"IPhone SE", 243, "Phones"},
	{"MacBook", 46, "Computers"},
	{"Asus", 189, "Computers"},
	{"Wizadr's diary", 1500, "Books"},
	{"T-shirt", 3589, "Clothes"},
}

type Item struct {

	Name string
	Count int
	Category string
}

func main() {


	http.HandleFunc("/", mainPage)
	http.HandleFunc("/items.html", itemsPage)
	http.HandleFunc("/Phones.html", phonesPage)
	http.HandleFunc("/Computers.html", computersPage)
	http.HandleFunc("/Clothes.html", clothesPage)

	port := ":9090"
	println("Server start on:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil{
		log.Fatal("Listen server", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request)  {

	templ, err := template.ParseFiles("static/index.html")
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}
	templ.Execute(w, nil)

}

func itemsPage(w http.ResponseWriter, r *http.Request){
	templ, err := template.ParseFiles("static/items.html")
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	templ.Execute(w, items)
}

func phonesPage(w http.ResponseWriter, r *http.Request){
	var phones []Item
	for i := 0; i < len(items); i++{
		if(items[i].Category == "Phones"){
			phones = append(phones, items[i])
		}
	}

	templ, err := template.ParseFiles("static/Phones.html")
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	templ.Execute(w, phones)
}

func computersPage(w http.ResponseWriter, r * http.Request)  {
	var computers []Item
	for i := 0; i < len(items); i++{
		if items[i].Category == "Computers"{
			computers = append(computers, items[i])
		}
	}

	templ, err := template.ParseFiles("static/Computers.html")
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	templ.Execute(w, computers)
}

func clothesPage(w http.ResponseWriter, r *http.Request){
	var clothes []Item
	for i := 0; i < len(items); i++{
		if items[i].Category == "Clothes"{
			clothes = append(clothes, items[i])
		}
	}

	templ, err := template.ParseFiles("static/Clothes.html")
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	templ.Execute(w, clothes)
}