package Router

import (
	"Backend/Config" // Mi conexion a base de datos
	"Backend/Entities"
	"Backend/Models" // Mi modelo
	"encoding/json"
	"fmt"
	"log"
	"net/http" // El servidor privado
	_ "github.com/gorilla/mux"
)


func Users(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		users, err2 := userModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(users)
		}
	}
}

func Unions(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		unionModel := models.UnionModel{
			Db: db,
		}
		unions, err2 := unionModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(unions)
		}
	}
}


func Crimes(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		crimemodel := models.CrimeModel{
			Db: db,
		}
		crimes, err2 := crimemodel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(crimes)
		}
	}
}

func Locations(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		locationmodel := models.LocationModel{
			Db: db,
		}
		locations, err2 := locationmodel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(locations)
		}
	}
}

func Categories(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
	} else {
		categorymodel := models.CategoryModel{
			Db: db,
		}
		categories, err2 := categorymodel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(categories)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde mi servidor privado klk wawawa")
}

func UserAdd(w http.ResponseWriter, r *http.Request) {
	
	db, err := config.GetDB()
	
	if err != nil {
		fmt.Println(err)
	} else {
	decoder := json.NewDecoder(r.Body)

	var User entities.User

	err2 := decoder.Decode(&User)

	if err2 != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(User)

	usermodel := models.UserModel{
		Db: db,
	}

	result,err3 := usermodel.CreateUser(User)
	if err3 != nil {
		panic(err3)
	}
	log.Println(result) // Con el result se puede confirmar 
	// si añadio el usuario correctamente
  }
}

func LocationAdd(w http.ResponseWriter, r *http.Request) {
	
	db, err := config.GetDB()
	
	if err != nil {
		fmt.Println(err)
	} else {
	decoder := json.NewDecoder(r.Body)

	var Location entities.Location

	err2 := decoder.Decode(&Location)

	if err2 != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(Location)

	locationmodel := models.LocationModel{
		Db: db,
	}

	result,err3 := locationmodel.CreateLocation(Location)
	if err3 != nil {
		panic(err3)
	}
	log.Println(result) // Con el result se puede confirmar 
	// si añadio el usuario correctamente
  }
}

func CrimeAdd(w http.ResponseWriter, r *http.Request) {
	
	db, err := config.GetDB()
	
	if err != nil {
		fmt.Println(err)
	} else {
	decoder := json.NewDecoder(r.Body)

	var Crime entities.Crime

	err2 := decoder.Decode(&Crime)

	if err2 != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(Crime)

	crimemodel := models.CrimeModel{
		Db: db,
	}

	result,err3 := crimemodel.CreateCrime(Crime)
	if err3 != nil {
		panic(err3)
	}
	log.Println(result) // Con el result se puede confirmar 
	// si añadio el usuario correctamente
  }
}

/*
func ProductsAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var Producto entities.Product

	err := decoder.Decode(&Producto)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(Producto)

	products = append(products, Producto)
}*/
