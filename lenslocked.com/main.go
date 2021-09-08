package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"lenslocked.com/views"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	_ = http.ListenAndServe(":3000", r)
}
