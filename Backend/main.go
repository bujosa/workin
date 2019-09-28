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


/*
func View() {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("------------------------")
			}
		}
	}
}
*/
