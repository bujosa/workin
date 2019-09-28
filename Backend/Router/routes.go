package Router

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{

	Route{
		"Users",
		"GET",
		"/Users",
		Users,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UserAdd",
		"POST",
		"/UserAdd",
		UserAdd,
	},
	Route{
		"Crimes",
		"GET",
		"/Crimes",
		Crimes,
	},
	Route{
		"CrimeAdd",
		"POST",
		"/CrimeAdd",
		CrimeAdd,
	},
	Route{
		"Locations",
		"GET",
		"/Locations",
		Locations,
	},
	Route{
		"LocationAdd",
		"POST",
		"/LocationAdd",
		LocationAdd,
	},
	Route{
		"Categories",
		"GET",
		"/Categories",
		Categories,
	},
	Route{
		"Unions",
		"GET",
		"/Unions",
		Unions,
	},
}
