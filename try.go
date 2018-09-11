package main

import (
	"html/template"
	"os"
)

type MyPageData struct {
	Field1 []Foo
	Field2 map[string]Bar
}

type Foo struct {
	Name string
}

type Bar struct {
	Name string
}

func main() {
	t := template.Must(template.New("").Parse(tmpl))
	mpd := MyPageData{
		Field1: []Foo{{"First"}, {"Second"}},
	//	Field2: map[string]Bar{"myKey": {"MyValue"}},
	}
	t.Execute(os.Stdout, mpd)
}

const tmpl = `{{range $idx, $foo := .Field1}} {{$idx}} - {{$foo.Name}}
{{end}}`

