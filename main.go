package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/view/"):]
	// if err != nil {
	// 	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	// 	return
	// }
	renderTemplate(w, "message")
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
