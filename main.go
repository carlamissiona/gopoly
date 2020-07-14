package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "dbname=d9fbdugs64n2l user=fcirxmuauuolmn password=505e9b87e24f6c1a647682111d98337e4f2f01a46e75b3cbf09cd3123e244ea8 host=ec2-3-215-83-17.compute-1.amazonaws.com sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB

	sqlCreate :=`
		CREATE DATABASE books_database
		  WITH OWNER = postgres
		       ENCODING = 'UTF8'
		       TABLESPACE = pg_default
		       CONNECTION LIMIT = -1;

		-- Table: books

		-- DROP TABLE books;

		CREATE TABLE books
		(
		 id serial NOT NULL,
		 name character varying NOT NULL,
		 author character varying,
		 pages integer,
		 publication_date date,
		 CONSTRAINT pk_books PRIMARY KEY (id )
		)
		WITH (
		 OIDS=FALSE
		);
		ALTER TABLE books
		 OWNER TO postgres;
		 `

	_, err = db.Exec(sqlCreate)
	if err == nil {
		panic(err)
	}

}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListBooks)
	http.HandleFunc("/book.html", handleViewBook)
	http.HandleFunc("/save", handleSaveBook)
	http.HandleFunc("/delete", handleDeleteBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
