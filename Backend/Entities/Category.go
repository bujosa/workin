package entities

import (
	"fmt"
)

type Category struct {
	Id int64
	Categorie string
}

type Categories []Category

func (category Category) ToString() string {
	return fmt.Sprintf("Id: %d\nCategorie: %s",
		category.Id ,category.Categorie)
}