package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplate function renders a template using html template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// template cache variable
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have a template in cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("creating new template")
		err = createTemplateCache(t)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gohtml",
	}

	// parsing template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache
	tc[t] = tmpl
	return nil
}
