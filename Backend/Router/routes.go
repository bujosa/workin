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
		"UserDelete",
		"POST",
		"/UserDelete",
		DeleteUser,
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
		"CrimeDelete",
		"POST",
		"/CrimeDelete",
		DeleteCrime,
	},
	Route{
		"CrimeUpdate",
		"POST",
		"/CrimeUpdate",
		UpdateCrime,
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
		"CategoryAdd",
		"POST",
		"/CategoryAdd",
		CategoryAdd,
	},
	Route{
		"CategoryUpdate",
		"POST",
		"/CategoryUpdate",
		UpdateCategory,
	},
	Route{
		"CategoryDelete",
		"POST",
		"/CategoryDelete",
		DeleteCategory,
	},
	Route{
		"Unions",
		"GET",
		"/Unions",
		Unions,
	},
}
