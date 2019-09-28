package entities

import (
	"fmt"
)

type Category struct {
	Categorie string
}

type Categories []Category

func (category Category) ToString() string {
	return fmt.Sprintf("Categorie: %s",
		category.Categorie)
}