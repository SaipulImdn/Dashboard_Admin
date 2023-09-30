package routes

import (
	"net/http"
	"qr-code/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.HomeController).Methods("GET")
	r.HandleFunc("/buttons", controllers.ButtonsController).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
