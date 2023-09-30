package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"qr-code/routes"
)

func main() {
	r := mux.NewRouter()
	routes.SetupRoutes(r)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
