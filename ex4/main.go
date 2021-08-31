package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl04 := `
how to create variables
{{- $mynumber := 41  }}
el valor de la variable es:
{{ $mynumber }}
reasign value:
{{- $mynumber  = "hola mundo" }}
new value;
{{ $mynumber }}
	`

	t, err := template.New("exercise4").Parse(tpl04)
	if err != nil {
		log.Fatalf("error trying to parse template %v ", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("error trying to execute template %v ", err)
	}
}
