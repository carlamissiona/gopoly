package polyweb

import (
	_ "fmt"
	_ "time"
	_ "github.com/lib/pq"
		"database/sql"
		"os"
		"log"

)






// type Course struct {
// 	ID              int
// 	Name            string
// 	Author          string
// 	Tag             string
// 	Category        string
// 	Partners        string
// 	Difficulty      string
// 	Intro           string
// 	NoOfStudents    int
// 	Price           float32
// 	IsFree          int
// 	PassingScore    int
// 	HoursToComplete int
// 	PublicationDate time.Time
// 	Pages           int
// }
func getCourse() ([]Course, error) {
	//Retrieve
	var db *sql.DB
	tmpDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB

	res := []Course{}


	rows, err := db.Query(`SELECT id, name, author, tag, category,partner_institution,date_created,passing_score,is_free,hours_to_complete,intro,difficulty,no_of_students,price FROM courses`)
	if err != nil {  		return nil, err 	}


	defer rows.Close()
	for rows.Next() {
			var id int
			var name string
			var author string
			var tag string
			var category string
			var partners string
			var passing_score int
			var date_created pq.NullTime
			var pages int
			var is_free int
			var hours_to_complete int
			var difficulty string
			var no_of_students int
			var intro string
			var price float32

			err = rows.Scan( &id, &name, &author, &tag, &category,&partners,&date_created,&passing_score,&is_free,&hours_to_complete,&intro,&difficulty,&no_of_students,&price)
			if err != nil {
				return res, err
			}
			currentCourse := Course{ID: id,	Name:name, 	Author:author,  Tag:tag, Category:category, Partners:partners, 	Difficulty:difficulty,
														  Intro:intro,	NoOfStudents:no_of_students,  Price:price,	IsFree:is_free,  	PassingScore:passing_score, 	HoursToComplete:hours_to_complete, 	PublicationDate:date_created}
			if date_created.Valid != nil {
				currentCourse.PublicationDate = date_created.Time
			}

			res = append(res, currentCourse)
		}


			// err := db.QueryRow(`SELECT id, name, author, tag, category,partner_institution,date_created,passing_score,is_free,hours_to_complete,intro,difficulty,no_of_students,price FROM courses where id = $1`, CourseId)
			//          .Scan(&id,  )
			// if err == nil {
			// 	res = Book{ID: id, Name: name, Author: author, Pages: pages, PublicationDate: publicationDate.Time}
			// }
			//


	return res, err
}
