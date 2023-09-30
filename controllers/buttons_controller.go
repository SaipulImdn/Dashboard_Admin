package controllers

import (
	"html/template"
	"net/http"
)

func ButtonsController(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/buttons.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Golang MVC Example",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
