package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate function renders a template using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
