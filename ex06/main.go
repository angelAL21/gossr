package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl06 := `
	{{- if ge . 18 }}
	Welcome to the bar. Enjoy.
	{{- else if ge . 12 }}
	You are too young.
	{{- else }}
	Sorry.
	{{ end }}
	`
	t, err := template.New("ejercicio06").Parse(tpl06)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	err = t.Execute(os.Stdout, 16)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}
