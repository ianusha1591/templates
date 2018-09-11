package main

import (
	"html/template"
	"net/http"
)

[]Todo string

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("lay.html"))

	http.HandleFunc("/temp", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
                        Todos: []Todo{"hello","happy","bye"}
                        },
			
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
