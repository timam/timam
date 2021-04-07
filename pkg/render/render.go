package render

import (
	"bytes"
	"fmt"
	"github.com/timam/timam/pkg/config"
	"github.com/timam/timam/pkg/modles"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig
var functions = template.FuncMap{

}

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *modles.TemplateData) *modles.TemplateData  {
	return td
}

//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *modles.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache{
		//get the template cache form app config
		tc = app.TemplateCache
	}else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template form template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf,td)

	_, err := buf.WriteTo(w)
	if err != nil{
		fmt.Println("Error writing template to browser", err)
	}

}


//CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error)  {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil{
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

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
