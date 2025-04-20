package render

//
// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// )
//
// func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template: ", err)
// 	}
// }
//
// // tc is an implementation of a template cache using a map
// var tc = make(map[string]*template.Template)
//
// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error
//
// 	// search for the template in cache map
// 	_, inMap := tc[t]
//
// 	if !inMap {
// 		// add the parsed template to cache map
// 		log.Println("creating the template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// template is in cache map
// 		log.Println("using cached template")
// 	}
//
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
//
// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}
//
// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
//
// 	// add template to map cache
// 	tc[t] = tmpl
// 	return nil
// }
