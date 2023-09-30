package controllers

import (
	"html/template"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
    // Pastikan path file template yang benar
    tmpl, err := template.ParseFiles("views/index.html", "views/components/sidebar.html")
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
