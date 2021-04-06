package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)
	if err != nil{
		fmt.Println("error getting template", err)
	}
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("| Error | Parsing Template |", err)
	}
}



func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error)  {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil{
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("page is now", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return myCache, err
		}

		matchts, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil{
			return myCache, err
		}

		if len(matchts) > 0{
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil{
				return myCache, err
			}

		}
		myCache[name] = ts

	}


	return myCache, nil
}
