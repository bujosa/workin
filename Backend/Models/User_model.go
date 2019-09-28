package models

import (
	"Backend/Entities"
	"database/sql"
)

type UserModel struct {
	Db *sql.DB
}

var users []entities.User

// Para agregar Usuarios
func (usermodel UserModel) CreateUser(UserN entities.User) (int64, error) {


	count,err4 := usermodel.Db.Query("select count(*) from Usuarios where Email = ?", UserN.Email)
	
	if err4 != nil{
		return 0, err4
	} else{
		for count.Next() {

			var x int64

			err3 := count.Scan(&x)

			if err3 != nil{
				return 0, err3
			}

			if x == 1{
			return 0,err3
			}
		}
	}

	result, err := usermodel.Db.Exec("INSERT Usuarios (Nombre,Apellido,Contrasena, Email,Telefono, Genero,Direccion) values (?, ?, ?, ?, ?, ?, ?)", 
	UserN.Name, UserN.Lastname, UserN.Password, UserN.Email, UserN.Tel, UserN.Gender, UserN.Address)

	if err != nil {
		return 0, err
	} else {
		UserN.Id,_ = result.LastInsertId() 
		users = append(users, UserN)
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
		}
		
}

// Esto es para listar los usuarios
func (usermodel UserModel) FindAll() ([]entities.User, error) {
	rows, err := usermodel.Db.Query("select * from Usuarios")

	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		for rows.Next() {
			var id             int64 
			var	name 		   string
			var lastname       string
			var password       string
			var email          string
			var tel		       string
			var gender		   string
			var address        string 

			err2 := rows.Scan(&id, &name, &lastname, &password, &email, &tel, &gender, &address)
			if err2 != nil {
				return nil, err2
			} else {
				user := entities.User{
					Id:             id,
					Name: 			name,
					Lastname:       lastname,
					Password:       password,
					Email:          email,
					Tel:         	tel,
					Gender:         gender,
					Address:		address,
				}

				users = append(users, user)
			}
		}

		return users, nil
	}

}
