package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	cursos := []string{
		"Go desde cero",
		"POO con Go",
		"BD con Go",
		"Testing con Go",
		"APIs con Go",
	}

	// _ = cursos

	// var vacio []string

	tpl07 := `
Cursos de Go dictados en EDteam
{{- range $i, $course := . }}
{{ $i }} {{ $course -}}
{{ else }}
No existen elementos
{{ end }}
`

	t, err := template.New("ejercicio07").Parse(tpl07)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	err = t.Execute(os.Stdout, cursos) //vacio
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}

}
