package main

import (
	"log"      // Esta es para capturar errores
	"net/http" // El servidor privado
	"Backend/Router"
)

func main() {

	router := Router.NewRouter()

	Server := http.ListenAndServe(":8080", router)

	log.Fatal(Server)

}

