package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {

	types := []string{
		"string",
		"int",
		"int64",
		"float64",
	}

	var tpl *template.Template
	var headerTpl *template.Template

	dir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}

	fileName := path.Join(dir, "generated_funcs.go")
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	tpl = template.Must(template.New("sliceContainsTemplate").Parse(sliceContainsTemplate()))
	headerTpl = template.Must(template.New("headerTemplate").Parse(headerTemplate()))

	headerTpl.Execute(w, nil)

	for _, v := range types {
		tpl.Execute(
			w,
			templateData{
				CapitalType: strings.Title(v),
				LowerType:   v,
			},
		)
	}

	w.Flush()
}

type templateData struct {
	CapitalType string
	LowerType   string
}

func headerTemplate() string {
	return `package main`
}

func sliceContainsTemplate() string {
	return `
		func SliceContains{{ .CapitalType }}(slice1 []{{.LowerType}}, val {{.LowerType}}) bool {
			for _, v := range slice1 {
				if v == val {
					return true
				}
			}
			return false
		}
	`
}
