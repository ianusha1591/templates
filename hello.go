package main
import (
        "html/template"
"encoding/json"
        "net/http"
"os"
"fmt"
"log"
)
type Data struct {
    Name string
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("hello.tmpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
name:= r.FormValue("username")
data  := &Data{name}
b, err :=json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
f, err := os.OpenFile("loginoutput.go", os.O_APPEND|os.O_WRONLY, 0644)
f.Write(b)
}
}
func main() {
//    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":80", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
