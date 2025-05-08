package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/svizcaino26/bookings/internal/config"
	"github.com/svizcaino26/bookings/internal/models"
)

var (
	app             *config.AppConfig
	pathToTemplates = "./templates"
)

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// [AddDefaultData adds data for all templates]
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template
	template, ok := tc[tmpl]
	if !ok {
		log.Fatalf("template %s not found in cache", tmpl)
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := template.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// [CreateTemplateCache creates a template cache as a map]
func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl.html from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl.html
	for _, page := range pages {
		// filepath.Base returns the last element of the filepath
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// this will get matches if the current template uses a layout
		matches, err := filepath.Glob(fmt.Sprintf("%s/templates/*.layout.tmpl.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/templates/*.layout.tmpl.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
