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
type Employee struct {
        Name string
       
    }
type Person struct {
          Names string
        Employer []Employee
    }

func hello(w http.ResponseWriter, r *http.Request) {
tmple :=template.Must(template.ParseFiles("modify.tmpl"))

       // name := r.FormValue("name")
       
        personname :=Person {
Names: "Here we display name and email of employees",
Employer: []Employee{{"joy"},{"god"},{"thomas"},
},
} 

tmple.Execute(w, personname)

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
http.HandleFunc("/fine", hi) 
   http.HandleFunc("/modify", hello) 
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
