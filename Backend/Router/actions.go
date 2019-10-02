package Router

import (
	config "Backend/Config" // Mi conexion a base de datos
	entities "Backend/Entities"
	models "Backend/Models" // Mi modelo
	"encoding/json"
	"fmt"
	"log"
	"net/http" // El servidor privado

	_ "github.com/gorilla/mux"
	_ "github.com/rs/cors"
)

/*
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
    // process the request...
}
*/

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods"," GET, POST, PUT, DELETE")
}

func Users(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

    enableCors(&w)

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

	enableCors(&w)

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

	enableCors(&w)

	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()

	if err != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		crimemodel := models.CrimeModel{
			Db: db,
		}
		crimes, err2 := crimemodel.FindAll()
		if err2 != nil || err3 != nil {
			json.NewEncoder(w).Encode(messages)
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(crimes)
		}
	}
}

func Locations(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()

	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		locationmodel := models.LocationModel{
			Db: db,
		}
		locations, err2 := locationmodel.FindAll()
		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(locations)
		}
	}
}

func Categories(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()
	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		categorymodel := models.CategoryModel{
			Db: db,
		}
		categories, err2 := categorymodel.FindAll()
		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			fmt.Println(err2)
		} else {
			json.NewEncoder(w).Encode(categories)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor Privado corriendo, las Api se encuentran en funcionamiento")
}

func UserAdd(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()
	db, err := config.GetDB()

	if err != nil || err3!=nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var User entities.User

		err2 := decoder.Decode(&User)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(User)

		usermodel := models.UserModel{
			Db: db,
		}

		result, err3 := usermodel.CreateUser(User)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}

		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)
		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func LocationAdd(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()

	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var Location entities.Location

		err2 := decoder.Decode(&Location)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(Location)

		locationmodel := models.LocationModel{
			Db: db,
		}

		result, err3 := locationmodel.CreateLocation(Location)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}

		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)
		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func CrimeAdd(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()

	if err != nil || err3 != nil{
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var Crime entities.Crime

		err2 := decoder.Decode(&Crime)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(Crime)

		crimemodel := models.CrimeModel{
			Db: db,
		}

		result, err3 := crimemodel.CreateCrime(Crime)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}
		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)
		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func CategoryAdd(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()

	if err != nil || err3 != nil{
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var Category entities.Category

		err2 := decoder.Decode(&Category)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(Category)

		categorymodel := models.CategoryModel{
			Db: db,
		}

		result, err3 := categorymodel.CreateCategory(Category)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}
		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)
		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func DeleteCrime(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()
	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var Crime entities.Crime

		err2 := decoder.Decode(&Crime)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(Crime)

		crimemodel := models.CrimeModel{
			Db: db,
		}

		result, err3 := crimemodel.DeleteCrime(Crime)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}

		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)

		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()
	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var Category entities.Category

		err2 := decoder.Decode(&Category)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(Category)

		categorymodel := models.CategoryModel{
			Db: db,
		}

		result, err3 := categorymodel.DeleteCategory(Category)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}

		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)

		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)
	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()
	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var User entities.User

		err2 := decoder.Decode(&User)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(User)

		usermodel := models.UserModel{
			Db: db,
		}

		result, err3 := usermodel.DeleteUser(User)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}

		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)

		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func UpdateCrime(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)

	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()

	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var Crime entities.Crime

		err2 := decoder.Decode(&Crime)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(Crime)

		crimemodel := models.CrimeModel{
			Db: db,
		}

		result, err3 := crimemodel.UpdateCrime(Crime)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}

		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)

		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	db, err := config.GetDB()

	enableCors(&w)

	messagemodel := models.MessageModel{}
	messages, err3 := messagemodel.Denied()

	if err != nil || err3 != nil {
		json.NewEncoder(w).Encode(messages)
		fmt.Println(err)
	} else {
		decoder := json.NewDecoder(r.Body)

		var Category entities.Category

		err2 := decoder.Decode(&Category)

		if err2 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err)
		}

		defer r.Body.Close()

		log.Println(Category)

		categorymodel := models.CategoryModel{
			Db: db,
		}

		result, err3 := categorymodel.UpdateCategory(Category)
		if err3 != nil {
			json.NewEncoder(w).Encode(messages)
			panic(err3)
		}

		messages, err3 = messagemodel.Success()
		json.NewEncoder(w).Encode(messages)

		log.Println(result) // Con el result se puede confirmar
		// si añadio el usuario correctamente
	}
}