package main

import (
	_ "fmt"
	"io/ioutil"
	_ "log"
	"net/http"
	_ "strconv"
	"html/template"
	_ "time"
	_ "github.com/clement/apps/polymath/webpolymath/models"
	_ "github.com/clement/apps/polymath/webpolymath/ptypes"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
 	// courses, err := allCourses()
	books, err := allBooks()
	course, err := ptypes.getCourse()
	if err != nil {
		renderErrorPage(w, err)
		return
	}
  //  *** SETUP Template
	buf, err := ioutil.ReadFile("www/polymath/home.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = polyweb.HomePage{AllBooks: books  }
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}


func handleTos(w http.ResponseWriter, r *http.Request) {

	buf, err := ioutil.ReadFile("www/polymath/tos.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	var page = IndexPage{nil}
	tosPage := string(buf)
	t := template.Must(template.New("tosPage").Parse(tosPage))
	t.Execute(w,page)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {

	buf, err := ioutil.ReadFile("www/polymath/about.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	var page = IndexPage{nil}
	aboutPage := string(buf)
	t := template.Must(template.New("aboutPage").Parse(aboutPage))
	t.Execute(w,page)
}

func handleSignPro(w http.ResponseWriter, r *http.Request) {

	buf, err := ioutil.ReadFile("www/polymath/register-pro.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	var page = IndexPage{nil}
	aboutPage := string(buf)
	t := template.Must(template.New("aboutPage").Parse(aboutPage))
	t.Execute(w,page)
}
func handleSignup(w http.ResponseWriter, r *http.Request) {

	buf, err := ioutil.ReadFile("www/polymath/register.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	var page = IndexPage{nil}
	aboutPage := string(buf)
	t := template.Must(template.New("aboutPage").Parse(aboutPage))
	t.Execute(w,page)
}

func handleCourseList(w http.ResponseWriter, r *http.Request) {
 	// courses, err := allCourses()
	books, err := allBooks()
	if err != nil {
		renderErrorPage(w, err)
		return
	}
  //  *** SETUP Template
	buf, err := ioutil.ReadFile("www/polymath/home.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = HomePage{AllBooks: books}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}
func handleLessonsLists(w http.ResponseWriter, r *http.Request) {
 	// courses, err := allCourses()
	books, err := allBooks()
	if err != nil {
		renderErrorPage(w, err)
		return
	}
  //  *** SETUP Template
	buf, err := ioutil.ReadFile("www/polymath/home.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = HomePage{AllBooks: books}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}
