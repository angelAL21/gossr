package main

import (
	"html/template"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	myFunctions := template.FuncMap{
		"addOne": func(a int) int {
			a++
			return a
		},
		"lower": func(a string) string {
			return strings.ToLower(a)
		},
		"hours": func(a, b int) string {
			return "las horas son: " + strconv.Itoa((a+1)*b)
		},
	}

	cursos := []string{
		"Go desde cero",
		"POO con Go",
		"BD con Go",
		"Testing con Go",
		"APIs con Go",
	}

	tpl08 := `
Cursos de Go dictados en EDteam
{{- range $i, $course := . }}
{{ addOne $i }} {{ lower $course }} {{ hours $i 3 -}}
{{ else }}
No existen elementos
{{ end }}
`

	t, err := template.New("ejercicio08").Funcs(myFunctions).Parse(tpl08)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	err = t.Execute(os.Stdout, cursos)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}
