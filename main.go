package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {
  tmpDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB

}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	// http.HandleFunc("/", handleListBooks)
	http.HandleFunc("/book.html", handleViewBook)
	http.HandleFunc("/save", handleSaveBook)
	http.HandleFunc("/delete", handleDeleteBook)


	http.HandleFunc("/", handleHome)
	http.HandleFunc("/signup", handleRegistration)
	http.HandleFunc("/login", handleLogout)
	http.HandleFunc("/courses", handleCourseList)
	http.HandleFunc("/lessons", handleLessonsLists)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
