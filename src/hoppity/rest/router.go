package rest

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes() {
		handler := route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}


func CreateRoutes() {

	router := NewRouter()

	if err := http.ListenAndServe(":"+"3000", router); err != nil {
		log.Print(err)
	}
}
