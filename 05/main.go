package main

import (
	"log"
	"os"
	"text/template"
)

type numbers struct {
	X int
	Y int
}

func main() {
	tpl05 := `
 comparators
 eq: equals to. {{ .X  }} equals {{ .Y }}? = {{ eq .X .Y }}
 ne: {{.X}} not equal"!=" {{.Y}}? {{ne .X .Y}}
lt: {{.X}} less than "<" {{.Y}}? {{lt .X .Y}}
le: {{.X}} less or equal to "<=" {{.Y}}? {{le .X .Y}}
gt: {{.X}} greater than ">" {{.Y}}? {{gt .X .Y}}
ge: {{.X}} greater or equal ">=" {{.Y}}? {{ge .X .Y}}
 `

	t, err := template.New("exercise5").Parse(tpl05)
	if err != nil {
		log.Fatalf("error trying to parse template %v ", err)
	}

	n := numbers{5, 6}
	err = t.Execute(os.Stdout, &n)
	if err != nil {
		log.Fatalf("error trying to execute template %v ", err)
	}
}
