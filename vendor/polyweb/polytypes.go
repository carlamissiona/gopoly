package polyweb

import "time"

type Book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}
type HomePage struct {
	AllBooks []Book
	// AllCourses []Course
}


type Course struct {
	ID              int
	Name            string
	Author          string
	Tag             string
	Category        string
	Partners        string
	Difficulty      string
	Intro           string
	NoOfStudents    int
	Price           float32
	IsFree          int
	PassingScore    int
	HoursToComplete int
	PublicationDate time.Time
	Pages           int
}
