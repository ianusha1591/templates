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
tmpl := template.Must(template.ParseFiles("extension.tmpl","welcome.php"))
 http.HandleFunc("/extend", func(w http.ResponseWriter, r *http.Request) {
    coverage := Coverage {
       Pagetitle: "interesting",
        NeeoCoverage: []NeoCoverage{{"jasmine"},{"lilly"},{"rose"},
        },
    }
 tmpl.Execute(w, coverage)
})

        http.ListenAndServe(":80", nil)
}
