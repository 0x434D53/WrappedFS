package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/0x434D53/WrappedFS"
	_ "github.com/0x434d53/WrappedFS/example/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	file, err := statikFS.Open("/templates/template.html")

	if err != nil {
		log.Fatalf("Templates could not be read: %s", err)
	}

	tmpl, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Tempalte could not be read from: %s", err)
	}
	indextemplate, err := template.New("template.html").Parse(string(tmpl))

	if err != nil {
		log.Fatalf("Tempalte could ne be initiliazed: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indextemplate.Execute(w, nil)
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(wrappedfs.New(statikFS, "/public/"))))

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
