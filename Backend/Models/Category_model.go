package models

import (
	"Backend/Entities"
	"database/sql"
)

type CategoryModel struct {
	Db *sql.DB
}

func (categorymodel CategoryModel) FindAll() ([]entities.Category, error) {
	rows, err := categorymodel.Db.Query("select Categoria from Categorias")

	if err != nil {
		return nil, err
	} else {
		var categories []entities.Category
		for rows.Next() {
			var	categorie string

			err2 := rows.Scan(&categorie)
			if err2 != nil {
				return nil, err2
			} else {
				category := entities.Category{
					Categorie: categorie,
				}

				categories = append(categories, category)
			}
		}

		return categories, nil
	}

}
