package controllers

import (
	"fmt"
	"lenslocked.com/views"
	"net/http"

	"github.com/gorilla/schema"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *Users) New(w http.ResponseWriter, _ *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type Users struct {
	NewView *views.View
}

// Create is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	form := SignupForm{}
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
