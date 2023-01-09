package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// RenderTemplate function renders a template using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//get all of the files named *.page.gohtml from ./templates
	pages, err := filepath.Glob("./templates/*page.gohtml")
	if err != nil {
		return myCache, err
	}

	// range through pages
	for _, v := range pages {
		name := filepath.Base(v)
		ts, err := template.New(name).ParseFiles(v)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, err
}
