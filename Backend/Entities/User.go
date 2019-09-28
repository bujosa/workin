package entities

import (
	"fmt"
)

type User struct {
	Id             int64 
	Name 		   string
	Lastname       string
	Password       string
	Email          string
	Tel		       string
	Gender		   string
	Address        string 
}

type Users []User

func (user User) ToString() string {
	return fmt.Sprintf("id: %d\n name: %s\n lastname: %s\n password: %s\n email: %s\n tel: %s\ngender: %s\n adress: %s",
		user.Id, user.Name, user.Lastname, user.Password, user.Email, user.Tel, user.Gender, user.Address)
}