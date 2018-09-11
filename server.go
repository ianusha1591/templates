package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/users", users)
	http.ListenAndServe(":80", nil)
}

// user profile
type Profile struct {
	Name    string
	Hobbies []string
}

// user profiles
type Profiles []Profile

// context for templates
type Context struct {
    Title string
    Profiles Profiles
}

// handler func for `/users`
func users(w http.ResponseWriter, r *http.Request) {

	profiles := Profiles{
		Profile{"Jack", []string{"snowboarding", "croquet"}},
		Profile{"Jill", []string{"knitting", "minecraft"}},
	}

    context := Context{"User Profiles", profiles}

	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"layouts.html",
		"profiles.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
