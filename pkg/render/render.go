package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil{
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf,nil)
	_, err = buf.WriteTo(w)
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
