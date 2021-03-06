package main
 
import (
    "fmt"
    "log"
    "net/http"
"html/template"
)
type NeoCoverage struct {
        Name string
       
    }
type Coverage struct {
          Pagetitle string
        NeeoCoverage []NeoCoverage
    }


func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/postgethello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "postget.html")
    case "POST":

        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

        name := r.FormValue("name")
        email := r.FormValue("email")
        fmt.Fprintf(w, "Name = %s\n", name)
        fmt.Fprintf(w, "Email = %s\n", email)  


    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }

}
func hi(w http.ResponseWriter, r *http.Request) {
tmpl := template.Must(template.ParseFiles("mytrial.tmpl"))
 coverage := Coverage {
       Pagetitle: "tough but interesting",
        NeeoCoverage: []NeoCoverage{{"jasmine"},{"lilly"},{"rose"},
        },
    }
tmpl.Execute(w, coverage)
}
func main() {
http.HandleFunc("/postgetnew", hi) 
   http.HandleFunc("/postgethello", hello) 
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
