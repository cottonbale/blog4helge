package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"helgeBlog/views"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var homeView *views.View
var aboutView *views.View
var addEntryView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := aboutView.Render(w, nil); err != nil {
		panic(err)
	}
}

func addEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := addEntryView.Render(w, nil); err != nil {
		panic(err)
	}
}

func main() {

	db, err := sql.Open("mysql", "paul:password@/test1")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	var (
		id           int
		blog_element string
	)

	rows, err := db.Query("select id, blog_element from blog_elements where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &blog_element)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, blog_element)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()

	homeView = views.NewView("bootstrap",
		"views/home.gohtml")
	aboutView = views.NewView("bootstrap",
		"views/about.gohtml")
	addEntryView = views.NewView("bootstrap",
		"views/addEntry.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/addEntry", addEntry)
	http.ListenAndServe(":3000", r)
}
