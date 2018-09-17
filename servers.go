package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/users", users)
	http.ListenAndServe(":80", nil)
}

type Act struct {
	Name    string
	Hab []string
}

type Acts []Act


type Co struct {
    Title string
    Acts Acts
}


func users(w http.ResponseWriter, r *http.Request) {

	cp := Acts{
         	Act{"u1", []string{"travel", "music"}},
		Act{"u2", []string{"read", "play"}},
	}

    context := Co{"Users and their habits", cp}


	templates, err := template.ParseFiles(
		"ly.html",
		"pg.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
