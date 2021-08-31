package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func main() {
	tpl02 := "hola {{ .Name }}, tienes {{ .Age }} años"
	t, err := template.New("exercise2").Parse(tpl02)
	if err != nil {
		log.Fatal("error trying to parse template: ", err)
	}

	///err = t.Execute(os.Stdout, tpl02)
	//renderization data. no data now
	p := person{"angel", 23}
	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal("error trying to execute template: ", err)
	}
}
