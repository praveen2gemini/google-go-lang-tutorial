package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func handlePostProcessor(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml", d)

}

func handleGetProcessor(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml", d)

}

func main() {
	fmt.Printf("Hello GOOGLE GO!")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/process_details_post", handlePostProcessor)
	http.HandleFunc("/process_details_get", handleGetProcessor)
	http.ListenAndServe(":8000", nil)
}
