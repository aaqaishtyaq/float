package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

func serveHTTP() {
	t, err := template.New("index").Parse(pageTmpl)
	if err != nil {
		log.Fatal(err)
	}

	staticFS := http.FS(assets)
	fs := http.FileServer(staticFS)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, err := pageContent()
		if err != nil {
			log.Fatal(err)
		}

		var indexPage bytes.Buffer

		err = t.Execute(&indexPage, page)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(indexPage.Bytes())
	})

	http.Handle("/static/", fs)

	log.Println("Starting http server at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
