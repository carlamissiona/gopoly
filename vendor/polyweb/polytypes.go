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
