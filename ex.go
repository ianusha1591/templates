package main
import (
        "html/template"
        "net/http"
)

//func main() {
    type NeoCoverage struct {
        Name string
       
    }

    type Coverage struct {
          Pagetitle string
        NeeoCoverage []NeoCoverage
    }
func main() {
tmpl := template.Must(template.ParseFiles("ex.html"))
 http.HandleFunc("/ex", func(w http.ResponseWriter, r *http.Request) {
    coverage := Coverage {
       Pagetitle: "tough but interesting",
        NeeoCoverage: []NeoCoverage{{"jasmine"},{"lilly"},{"rose"},
        },
    }
 tmpl.Execute(w, coverage)
})

        http.ListenAndServe(":80", nil)
}
