package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("exit.html"))

	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1"},
				{Title: "Task 2"},
				{Title: "Task 3"},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
