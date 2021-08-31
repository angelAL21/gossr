package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p *person) Saludar() string {
	return "hola, Soy " + p.Name
}

func main() {
	tpl03 := "hola {{ .Name }}, su saludo {{ .Saludar }} "
	t, err := template.New("exercise3").Parse(tpl03)
	if err != nil {
		log.Fatal("error trying to parse template: ", err)
	}

	///err = t.Execute(os.Stdout, tpl02)
	//renderization data. no data now
	p := person{"angel", 23}
	err = t.Execute(os.Stdout, &p)
	if err != nil {
		log.Fatal("error trying to execute template: ", err)
	}
}
